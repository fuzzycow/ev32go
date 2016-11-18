package codegen

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"text/template"
	"strings"
	"go/format"
	"bytes"

	"path/filepath"

	"os"
)


type Properties struct {
	SpecFile     string
	GoSrcFile    string
	TemplateFile string
	IncludeDir   string
	PackageName  string
	DeviceName   string
}

func SpecToGo(prop *Properties) {
	var spec Spec
	contents, err := ioutil.ReadFile(prop.SpecFile)
	if err != nil {
		log.Fatalf("failed to read spec file '%v': %v", prop.SpecFile, err)
	}
	err = json.Unmarshal(contents, &spec)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("%+v",data)
	var packageFuncs = template.FuncMap{
		"packageName": func() string { return prop.PackageName },
	}

	mainTmplName := filepath.Base(prop.TemplateFile)

	tmpl, err := template.New(mainTmplName).
	Delims("<<", ">>").
	Funcs(TemplateFuncs).
	Funcs(packageFuncs).
	ParseGlob(filepath.Join(prop.IncludeDir, "*.tmpl"))

	if err != nil {
		log.Fatalf("error parsing template files: %v", err)
	}
	log.Printf("loaded templates%v", tmpl.DefinedTemplates())


	log.Printf("processing template %v", mainTmplName)
	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	var specDevice *Device
	seen := make([]string, 0, len(spec.Classes))
	for name, dev := range spec.Classes {
		seen = append(seen, name)
		if strings.ToLower(prop.DeviceName) == strings.ToLower(name) {
			specDevice = dev
			break
		}
	}
	if specDevice == nil {
		log.Fatalf("Failed to find device %v in spec file (known: %+v)", prop.DeviceName, seen)
	}

	err = tmpl.Execute(buf, specDevice)
	if err != nil {
		log.Fatalf("error executing template %v: %v", prop.TemplateFile, err)
	}

	goSrcDir := filepath.Dir(prop.GoSrcFile)
	log.Printf("verifying that go src dir exists: %v", goSrcDir)
	if err := os.MkdirAll(goSrcDir, 0755); err != nil {
		log.Fatalf("Failed to create go src dir '%s': %v", goSrcDir, err)
	}

	out := prop.GoSrcFile
	log.Printf("writing generated go code into %v", out)
	err = ioutil.WriteFile(out, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("failed to write unformatted generated code into %v: %v", out, err)
	}
	log.Printf("wrote unfortmatted go code into %v", out)
}

func GoFormat(prop *Properties) {
	log.Printf("loading go src file: %v", prop.GoSrcFile)
	b, err := ioutil.ReadFile(prop.GoSrcFile)
	if err != nil {
		log.Fatalf("failed to read unformatted generated code from %v: %v", prop.GoSrcFile, err)
	}

	formated, err := format.Source(b)
	if err != nil {
		log.Fatalf("formatting failed %v: %v", prop.GoSrcFile, err)
	}

	err = ioutil.WriteFile(prop.GoSrcFile, formated, 0644)
	if err != nil {
		log.Fatalf("failed to save formatted code into %v: %v", prop.GoSrcFile, err)
	}
	log.Printf("wrote re-formatted go src file: %v", prop.GoSrcFile)
}
