// +build tools

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/template"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalln("exactly 4 arguments must be provided")
	}
	argsWithoutProg := os.Args[1:]

	pkg := argsWithoutProg[0]
	srcPath, err := filepath.Abs(argsWithoutProg[1])
	log.Println("srcPath: ", srcPath)
	if err != nil {
		log.Fatalln("failed to load absolute path for the source")
	}
	dstPath, err := filepath.Abs(argsWithoutProg[2])
	log.Println("dstPath: ", dstPath)
	if err != nil {
		log.Fatalln("failed to load absolute path for the source")
	}

	srcData, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Fatalln(err)
	}

	var m metadata
	if err := json.Unmarshal(srcData, &m); err != nil {
		log.Fatalln(fmt.Errorf("failed to unmarshal source: %v", err))
	}

	regions := make([]string, 0, len(m.AMIs))
	for region := range m.AMIs {
		regions = append(regions, region)
	}
	sort.Strings(regions)

	tinput := struct {
		Pkg     string
		Regions []string
	}{Pkg: pkg, Regions: regions}

	t := template.Must(template.New("ami_regions").Parse(tmpl))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, tinput); err != nil {
		log.Fatalln(fmt.Errorf("failed to execute the template: %v", err))
	}

	if err := ioutil.WriteFile(dstPath, buf.Bytes(), 0664); err != nil {
		log.Fatalln(err)
	}
}

type metadata struct {
	AMIs map[string]struct {
		HVM string `json:"hvm"`
	} `json:"amis"`
}

var tmpl = `// Code generated by ami_regions_generate.go; DO NOT EDIT.

package {{ .Pkg }}

// AMIRegoins is a list of regions where the RHEL CoreOS is published.
var AMIRegions = []string{
{{- range $region := .Regions}}
    "{{ $region }}",
{{- end}}
}
`
