package main

import (
	"flag"
	"github.com/fuzzycow/ev32go/codegen"
	"path/filepath"
	"os"
	"log"

)

var (
	specFile = flag.String("spec","../docs/spec/spec.json", "specfile")
	templateDir = flag.String("td","../codegen/templates", "templates directory")
	mainTemplate = flag.String("t","", "main template file (relative to templates directory)")
	outFile = flag.String("o","", "output go src file")
	deviceName = flag.String("d","", "device name (matched against spec file keys)")
	packageName = flag.String("p","", "go package name for generated code")
)

func main() {
	flag.Parse()
	gofile := os.Getenv("GOFILE")
	if gofile == "" {
		log.Fatal("GOFILE not set. GOFILE is used for workdir detection.")
	}

	ev3dir,err := filepath.Abs(filepath.Dir(os.Getenv("GOFILE")))
	if err != nil {
		log.Fatalf("Failed to detect working directory absolute path: %v", err)
	}
	log.Printf("Work dir: %s",ev3dir)
	if *specFile == "" || *templateDir == "" || *mainTemplate == "" || *outFile == "" || *deviceName == "" {
		log.Fatalf("missing required parameters")
	}
	tmpl := filepath.Join(*templateDir,*mainTemplate)

	log.Printf("- specfile: %s", *specFile)
	log.Printf("- main template: %s (%s)", *mainTemplate, tmpl)
	log.Printf("- template include path: %s", *templateDir)
	log.Printf("- go src file: %s", *outFile)

	if *packageName == "" {
		*packageName = filepath.Base(filepath.Dir(*outFile))
	}
	log.Printf("- go package name: %v", *packageName)

	prop := &codegen.Properties{
		SpecFile: *specFile,
		GoSrcFile: *outFile,
		TemplateFile: tmpl,
		IncludeDir: *templateDir,
		PackageName: *packageName,
		DeviceName: *deviceName,
	}

	codegen.SpecToGo(prop)
	codegen.GoFormat(prop)
}
