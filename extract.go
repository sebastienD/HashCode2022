package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func extractContributors(filename string) ([]contributor, []project) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	contributors := make([]contributor, 0)
	projects := make([]project, 0)

	line := scanAndSplit(scanner)
	nbContr, _ := strconv.Atoi(line[0])
	nbProj := AtoI(line[1])
	for nbContr > 0 {
		contr := extractContributor(scanner)
		contributors = append(contributors, contr)
		nbContr--
	}

	for nbProj > 0 {
		proj := extractProject(scanner)
		projects = append(projects, proj)
		nbProj--
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}
	return contributors, projects
}

func extractContributor(scanner *bufio.Scanner) contributor {
	line := scanAndSplit(scanner)
	contr := contributor{
		name:   line[0],
		skills: map[string]int{},
	}
	nbSkills := AtoI(line[1])

	for i := 0; i < nbSkills; i++ {
		skill := scanAndSplit(scanner)
		level, _ := strconv.Atoi(skill[1])
		contr.skills[skill[0]] = level
	}
	return contr
}

func extractProject(scanner *bufio.Scanner) project {
	line := scanAndSplit(scanner)
	proj := project{
		name:           line[0],
		daysToComplete: AtoI(line[1]),
		scoreAward:     AtoI(line[2]),
		bestBeforeDay:  AtoI(line[3]),
		roles:          []role{},
	}
	for i := 0; i < AtoI(line[4]); i++ {
		skill := scanAndSplit(scanner)
		proj.roles = append(proj.roles, role{skill[0], AtoI(skill[1])})
	}
	return proj
}
