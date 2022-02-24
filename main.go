package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filenames := []string{
		"a_an_example.in.txt",
		"b_better_start_small.in.txt",
		"c_collaboration.in.txt",
		"d_dense_schedule.in.txt",
		"e_exceptional_skills.in.txt",
		"f_find_great_mentors.in.txt",
	}
	for _, v := range filenames {
		contributors, projects := extractContributors("./data-input/" + v)
		fmt.Println("solved ", v)
		content := solve(projects, contributors)
		WriteFile(v, content)
	}
}

func solve(projects []project, contributors []contributor) string {
	var res []output

	mapCr := toMap(contributors)
	for _, proj := range projects {
		filledRoles := findContributors(proj, mapCr)
		if filledRoles == nil {
			continue
		}
		res = append(res, output{
			project:      proj.name,
			contributors: filledRoles,
		})
		upgradeLevel(filledRoles)
	}

	return formatOutput(res)
}

func upgradeLevel(frole []filledRole) {
	for _, v := range frole {
		if v.contributor.skills[v.role.skill] == v.role.level {
			v.contributor.skills[v.role.skill]++
		}
	}
}

func findContributors(proj project, cr map[string]*contributor) []filledRole {
	copyCr := copyMap(cr)
	var reserved []filledRole
	for _, r := range proj.roles {
		c := r.firstFit(copyCr)
		if c == nil {
			return nil
		}
		reserved = append(reserved, filledRole{
			contributor: c,
			role:        r,
		})
		delete(copyCr, c.name)
	}
	return reserved
}

func copyMap(myMap map[string]*contributor) map[string]*contributor {
	res := map[string]*contributor{}
	for i, v := range myMap {
		res[i] = v
	}
	return res
}

func toMap(contributors []contributor) map[string]*contributor {
	cr := map[string]*contributor{}
	for i, _ := range contributors {
		cr[contributors[i].name] = &contributors[i]
	}
	return cr
}

func contains(skills map[string]int, skill string) bool {
	for s, _ := range skills {
		if s == skill {
			return true
		}
	}
	return false
}

func WriteFile(filename, content string) {
	s := []byte(content)
	output := fmt.Sprintf("./submissions/%s.out.txt", filename[:len(filename)-7])
	if err := ioutil.WriteFile(output, s, 0644); err != nil {
		panic(err.Error())
	}
}
