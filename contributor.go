package main

import "fmt"

type contributor struct {
	name   string
	skills map[string]int
	delay  int
}

func (c contributor) String() string {
	return fmt.Sprintf("name: %s, skills: %v", c.name, c.skills)
}

func (c contributor) isFit(r role) bool {
	return c.skills[r.skill] >= r.level
}

func skillToContrib(contribs []contributor) map[string][]contributor {
	var skillSet map[string]bool
	for _, c := range contribs {
		for s, _ := range c.skills {
			skillSet[s] = true
		}
	}
	var res map[string][]contributor
	for s, exists := range skillSet {
		if !exists {
			continue
		}
		res[s] = []contributor{}
		for _, c := range contribs {
			if contains(c.skills, s) {
				res[s] = append(res[s], c)
			}
		}
	}
	return res
}
