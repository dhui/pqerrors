// Used to generate the errors

// +build ignore

//
// Usage:
//
// go run generate.go

package main

import (
	"bufio"
	"bytes"
	"flag"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

import (
	"github.com/lib/pq"
)

const (
	errClassPrefix = "Class "
	filename       = "consts.go"
)

const (
	fileWriteMode = 0644
)

func varName(s string) string {
	toks := strings.Split(s, "_")
	for i := range toks {
		toks[i] = strings.Replace(strings.Replace(strings.Title(toks[i]),
			"Sql", "SQL", -1), "Xml", "XML", -1)
	}
	return strings.Join(toks, "")
}

// ErrorClass wraps pq's ErrorClass
type ErrorClass struct{ pq.ErrorClass }

// VarName converts the ErrorClass name to one that can be used as a Go variable name
func (ec ErrorClass) VarName() string {
	return varName(ec.Name())
}

// SafeVarName is guaranteed to return a unique VarName
func (ec ErrorClass) SafeVarName() string {
	return ec.VarName()
}

// String gets the string backing the pq.ErrorClass
func (ec ErrorClass) String() string {
	return string(ec.ErrorClass)
}

// NewErrorClass creates a new ErrorClass
func NewErrorClass(s string) ErrorClass {
	return ErrorClass{pq.ErrorClass(s)}
}

// ErrorCode wraps pq's ErrorCode
type ErrorCode struct{ pq.ErrorCode }

// VarName converts the ErrorCode name to one that can be used as a Go variable name
func (ec ErrorCode) VarName() string {
	return varName(ec.Name())
}

// SafeVarName is guaranteed to return a unique VarName
func (ec ErrorCode) SafeVarName() string {
	return ErrorClass{ec.Class()}.VarName() + ec.VarName()
}

// String gets the string backing the pq.ErrorCode
func (ec ErrorCode) String() string {
	return string(ec.ErrorCode)
}

// NewErrorCode creates a new ErrorCode
func NewErrorCode(s string) ErrorCode {
	return ErrorCode{pq.ErrorCode(s)}
}

// GetErrorClasses extracts ErrorClasses from a table
func GetErrorClasses(lines []string) []ErrorClass {
	var errClasses []ErrorClass
	for _, line := range lines {
		if strings.HasPrefix(line, errClassPrefix) {
			if errClass := NewErrorClass(line[len(errClassPrefix) : len(errClassPrefix)+2]); errClass.Name() != "" {
				errClasses = append(errClasses, errClass)
			} else {
				log.Printf("WARNING: Postgres error class %q (%s) not supported by pq\n",
					errClass, errClass.Name())
			}
		}
	}
	return errClasses
}

// GetErrorCodes extracts ErrorClasses from a table
func GetErrorCodes(lines []string) []ErrorCode {
	var errCodes []ErrorCode
	for _, line := range lines {
		if !strings.HasPrefix(line, errClassPrefix) {
			if errCode := NewErrorCode(line[:5]); errCode.Name() != "" {
				errCodes = append(errCodes, errCode)
			} else {
				log.Printf("WARNING: Postgres error code %q (%s) not supported by pq\n",
					errCode, errCode.Name())
			}
		}
	}
	return errCodes
}

// TmplCtx is passed to template.Execute()
type TmplCtx struct {
	PkgName       string
	PqConstructor string
	Consts        interface{}
}

// GenSource generates formatted Go source code using the given template and context
func GenSource(tmpl *template.Template, tmplCtx TmplCtx) []byte {
	if tmpl == nil {
		return nil
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, tmplCtx); err != nil {
		log.Fatal("Failed to execute template:", err)
	}
	// return buf.Bytes()
	srcFormatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal("Failed to format generated go source. tmplCtx:", tmplCtx,
			"error:", err)
	}
	return srcFormatted
}

func main() {
	var tableFilename, classesPkg, codesPkg string
	flag.StringVar(&tableFilename, "errorsTable", "errors_table.txt",
		"Filename of Postgres errors table, sourced from: https://www.postgresql.org/docs/10/static/errcodes-appendix.html")
	flag.StringVar(&classesPkg, "classesPkg", "pqerrcls", "Name of package containing generated Go code for pq ErrorClass constants")
	flag.StringVar(&codesPkg, "codesPkg", "pqerrcode", "Name of package containing generated Go code for pq ErrorCodeConstants")
	flag.Parse()

	tableFile, err := os.Open(tableFilename)
	if err != nil {
		log.Fatal("Unable to open errors table file:", err)
	}
	defer tableFile.Close()

	var tableLines []string
	tableScanner := bufio.NewScanner(tableFile)
	for tableScanner.Scan() {
		tableLines = append(tableLines, tableScanner.Text())
	}
	if err := tableScanner.Err(); err != nil {
		log.Fatal("Failed to parse errors table file", err)
	}

	tmpl, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		log.Fatal("Failed to parse Go source code template:", err)
	}

	errClassCtx := TmplCtx{
		PkgName:       classesPkg,
		PqConstructor: "pq.ErrorClass",
		Consts:        GetErrorClasses(tableLines)}
	errClassSrc := GenSource(tmpl, errClassCtx)
	errCodeCtx := TmplCtx{
		PkgName:       codesPkg,
		PqConstructor: "pq.ErrorCode",
		Consts:        GetErrorCodes(tableLines)}
	errCodeSrc := GenSource(tmpl, errCodeCtx)

	var classesFilename = filepath.Join(classesPkg, filename)
	if err := ioutil.WriteFile(classesFilename, errClassSrc, fileWriteMode); err != nil {
		log.Fatal("Failed to write", classesFilename, "error:", err)
	}
	var codesFilename = filepath.Join(codesPkg, filename)
	if err := ioutil.WriteFile(codesFilename, errCodeSrc, fileWriteMode); err != nil {
		log.Fatal("Failed to write", codesFilename, "error:", err)
	}
	log.Println("Successfully generated pq error constants!")
}

const tmpl = `
// Package {{.PkgName}} provides consts for {{.PqConstructor}}
//
// DO NOT EDIT
//
// Code is auto-generated

package {{.PkgName}}

import (
"github.com/lib/pq"
)

const (
{{range .Consts -}}
{{.SafeVarName}} = {{$.PqConstructor}}("{{.String}}")
{{end}}
)
`
