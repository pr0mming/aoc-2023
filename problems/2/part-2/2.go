package problems_2_2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var answer = 0

	for scanner.Scan() {
		line := scanner.Text()

		answer += getCubesProd(line)
	}

	return strconv.Itoa(answer)
}

func getCubesProd(input string) int {
	var (
		bagPattern          = regexp.MustCompile(`(\d+)\s+(\w+)`) // This regex will extract the groups: [ammount] [color] given each input
		groups     []string = strings.Split(input, ";")           // Get the sets per game
		cubesDict           = make(map[string]int)
	)

	// Iterate over each group and extract numbers per color
	for _, group := range groups {
		// Find all matches in the group
		matches := bagPattern.FindAllStringSubmatch(group, -1)

		// Iterate over matches and update the colorCount map
		for _, match := range matches {
			// Convert the matched number from string to int
			numOfCubes, err := strconv.Atoi(match[1])
			if err != nil {
				panic("Error converting string to int:")
			}

			cubeColorTmp := match[2]
			cubeAmmount, ok := cubesDict[cubeColorTmp]

			// If the item doesn't exist we put it in the map, otherwise we compare it if is the new highest
			if !ok || (ok && numOfCubes > cubeAmmount) {
				cubesDict[cubeColorTmp] = numOfCubes
			}
		}
	}

	var cubeProd int = 1
	for _, value := range cubesDict {
		cubeProd *= value
	}

	return cubeProd
}
