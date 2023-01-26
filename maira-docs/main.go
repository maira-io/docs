// Copyright 2021 Maira, Inc. All rights reserved.
// Copying, sharing, publishing prohibited without written
// permission from Maira, Inc.

package main

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/maira-io/playbook/cmd/maira-docs/docgen"
	"github.com/maira-io/playbook/lang/cmd_schema"
)

func generate(tmplt, outDir string, cmdName []string, cmd *cmd_schema.Command) {
	if len(cmd.Subcommands) == 0 {
		outfilePath := outDir + "/" + filepath.Join(cmdName...)
		dir, _ := filepath.Split(outfilePath)
		if err := os.MkdirAll(dir, 0700); err != nil {
			log.Fatalf("Error creating %s directory: %v", dir, err)
		}
		outfile, err := os.Create(outfilePath + ".md")
		if err != nil {
			log.Fatalf("Error creating output file %s: %v", outfilePath, err)
		}
		defer outfile.Close()
		err = docgen.Execute(outfile, strings.Join(cmdName, " "), cmd, tmplt)
		if err != nil {
			log.Fatalf("Error executing template: %v", err)
		}
		return
	}
	cmd.Subcommands.Range(func(name string, c *cmd_schema.Command) {
		generate(tmplt, outDir, append(cmdName, name), c)
	})
}

//go:embed docs.tmplt
var tmplt string

func main() {
	var outDir = flag.StringP("output", "o", "docs", "output directory")
	var schemaDir = flag.StringP("schema", "s", "", "schema directory")
	var examplesDir = flag.StringP("examples", "e", "", "examples directory")
	flag.Parse()
	schema, err := cmd_schema.ParseSchemaDir(*schemaDir, cmd_schema.WithNoExtHandlers, cmd_schema.WithNoHttpHandlers)
	if err != nil {
		log.Fatalf("Parse schema failed with error " + err.Error())
	}
	examples, err := cmd_schema.ParseExamplesDir(*examplesDir, cmd_schema.WithNoExtHandlers, cmd_schema.WithNoHttpHandlers)
	if err != nil {
		log.Fatalf("Parse examples failed with error " + err.Error())
	}
	cmd_schema.MergeExamples(schema.Commands, examples.Commands)
	schema.Commands.Range(func(cmdName string, command *cmd_schema.Command) {
		generate(tmplt, *outDir, []string{cmdName}, command)
	})
}
