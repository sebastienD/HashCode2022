package main

import (
	"fmt"
	"strings"
)

type filledRole struct {
	contributor *contributor
	role        role
}

func (fr filledRole) String() string {
	return fr.contributor.name
}

type output struct {
	project      string
	contributors []filledRole
}

func (o output) String() string {
	contrib := strings.Trim(fmt.Sprint(o.contributors), "[")
	contrib = strings.Trim(fmt.Sprint(contrib), "]")
	return fmt.Sprintf("%s\n%s", o.project, contrib)
}

func formatOutput(outs []output) string {
	result := fmt.Sprintf("%d\n", len(outs))
	for i, v := range outs {
		result += fmt.Sprintf("%s", v)
		if i != len(outs)-1 {
			result += "\n"
		}
	}
	return result
}
