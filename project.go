package main

import (
	"fmt"
	"strings"
)

type role struct {
	skill string
	level int
}

func (r role) firstFit(contribs map[string]*contributor) *contributor {
	var fit []contributor
	for _, c := range contribs {
		c := c
		if c.isFit(r) {
			fit = append(fit, *c)
		}
	}
	if len(fit) == 0 {
		return nil
	}
	level := fit[0].skills[r.skill]
	res := fit[0]
	for _, c := range fit {
		cskill := c.skills[r.skill]
		if cskill < level {
			level = cskill
			res = c
		}
	}
	return &res
}

func (r role) String() string {
	return r.skill
}

type project struct {
	name           string
	daysToComplete int
	scoreAward     int
	bestBeforeDay  int
	roles          []role
}

func (p project) String() string {
	roles := strings.Trim(fmt.Sprint(p.roles), "[")
	roles = strings.Trim(fmt.Sprint(roles), "]")
	return fmt.Sprintf("%s\n%s", p.name, roles)
}
