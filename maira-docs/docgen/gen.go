// Copyright 2021, 2022 Maira, Inc. All rights reserved.
// Copying, sharing, publishing prohibited without written
// permission from Maira, Inc.

package docgen

import (
	"io"
	"log"
	"strings"
	"text/template"

	"github.com/maira-io/playbook/lang/cmd_schema"
	"github.com/maira-io/playbook/lang/exec"
)

type Command struct {
	*cmd_schema.Command
	Name string
}

func (c *Command) Synopsis() string {
	ret := []string{c.Name}
	for _, arg := range c.Arguments {
		var argStr []string
		if !arg.NoKeyword || cmd_schema.ArgType(arg.Type) == cmd_schema.ArgTypeBool {
			argStr = append(argStr, exec.KwToString(arg.Name))
			for _, name := range arg.AlternateNames {
				argStr = append(argStr, "|"+exec.KwToString(name))
			}
		}
		if cmd_schema.ArgType(arg.Type) != cmd_schema.ArgTypeBool {
			argStr = append(argStr, "<"+arg.Name+">")
		}
		if !arg.Required {
			ret = append(ret, "["+strings.Join(argStr, " ")+"]")
		} else {
			ret = append(ret, strings.Join(argStr, " "))
		}
	}
	return strings.Join(ret, " ")
}

func (c *Command) GetExamples() string {
	var ret string
	if len(c.Examples) == 0 {
		ret = c.Name
		for _, arg := range c.Arguments {
			if arg.Default != "" {
				continue
			}
			ex := argExample(arg)
			if len(ex) > 0 {
				ret += " " + ex[0]
			}
		}
		return ret
	}
	for _, ex := range c.Examples {
		ret = "Input: \n```\n" + strings.TrimRight(ex.Input, "\r\n") + "\n```\n"
		ret += "Output: \n```\n" + strings.TrimRight(ex.Output, "\r\n") + "\n```\n"
	}
	return ret
}

func (c *Command) GetSubCommandHelp() []string {
	return cmd_schema.CommandList(c.Subcommands)
}

func argExample(arg *cmd_schema.Argument) []string {
	var ret []string
	examples := arg.Examples
	if arg.Example != "" {
		examples = append([]string{arg.Example}, examples...)
	}
	if len(examples) == 0 {
		var example string
		switch cmd_schema.ArgType(arg.Type) {
		default:
			example = arg.Name + "-1"
		case cmd_schema.ArgTypeTime:
			example = "2019-10-12T07:20:50.52Z"
		case cmd_schema.ArgTypeDuration:
			example = "5 seconds"
		case cmd_schema.ArgTypeBool:
			// nothing to add for bool
		case cmd_schema.ArgTypeInt:
			example = "1"
		case cmd_schema.ArgTypeFloat:
			example = "1.0"
		}
		examples = []string{example}
	}
	prefix := ""
	if !arg.NoKeyword || cmd_schema.ArgType(arg.Type) == cmd_schema.ArgTypeBool {
		prefix = exec.KwToString(arg.Name) + " "
	}
	for _, example := range examples {
		switch cmd_schema.ArgType(arg.Type) {
		default:
			ret = append(ret, prefix+"\""+example+"\"")
		case cmd_schema.ArgTypeBool:
			// nothing to add for bool
			ret = append(ret, prefix)
		case cmd_schema.ArgTypeInt, cmd_schema.ArgTypeFloat:
			ret = append(ret, prefix+example)
		}
	}
	return ret
}
func argDefault(arg *cmd_schema.Argument) string {
	if arg.Default != "" {
		return arg.Default
	}
	return "_None_"
}
func argAttributes(arg *cmd_schema.Argument) string {
	var ret []string
	if arg.Required {
		ret = append(ret, "required")
	} else {
		ret = append(ret, "optional")
	}
	if arg.AllowMultiple {
		ret = append(ret, "multiple allowed")
	}

	return strings.Join(ret, ", ")
}

var funcMap = template.FuncMap{
	"getExample":    argExample,
	"getDefault":    argDefault,
	"getAttributes": argAttributes,
}

func Execute(w io.Writer, cmdName string, cmd *cmd_schema.Command, tmplt string) error {
	t, err := template.New("docs").Funcs(funcMap).Parse(tmplt)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	c := &Command{
		Command: cmd,
		Name:    cmdName,
	}
	return t.Execute(w, c)
}
