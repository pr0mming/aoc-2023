package functions

import (
	"errors"
	"fmt"
	"strings"
	"time"

	problems_1_1 "aoc.2023/problems/1/part-1"
	problems_1_2 "aoc.2023/problems/1/part-2"
	problems_2_1 "aoc.2023/problems/2/part-1"
	problems_2_2 "aoc.2023/problems/2/part-2"
	problems_3_1 "aoc.2023/problems/3/part-1"
	problems_3_2 "aoc.2023/problems/3/part-2"
	problems_4_1 "aoc.2023/problems/4/part-1"
	problems_4_2 "aoc.2023/problems/4/part-2"
	problems_5_1 "aoc.2023/problems/5/part-1"
	problems_5_2 "aoc.2023/problems/5/part-2"
	problems_6_1 "aoc.2023/problems/6/part-1"
	problems_6_2 "aoc.2023/problems/6/part-2"
)

// This "Factory Method" triggers the problem solution
// I decided not use interfaces and extra structs of the pattern to avoid complexity
func SolveProblemByKey(args []string) (string, error) {
	var answer string

	key := strings.Join(args, "")
	problemId := args[0]

	// Measure the solution
	start := time.Now()

	switch key {
	case "11":
		answer = problems_1_1.SolveChallenge(problemId)
	case "12":
		answer = problems_1_2.SolveChallenge(problemId)
	case "21":
		answer = problems_2_1.SolveChallenge(problemId)
	case "22":
		answer = problems_2_2.SolveChallenge(problemId)
	case "31":
		answer = problems_3_1.SolveChallenge(problemId)
	case "32":
		answer = problems_3_2.SolveChallenge(problemId)
	case "41":
		answer = problems_4_1.SolveChallenge(problemId)
	case "42":
		answer = problems_4_2.SolveChallenge(problemId)
	case "51":
		answer = problems_5_1.SolveChallenge(problemId)
	case "52":
		answer = problems_5_2.SolveChallenge(problemId)
	case "61":
		answer = problems_6_1.SolveChallenge(problemId)
	case "62":
		answer = problems_6_2.SolveChallenge(problemId)
	default:
		return "", errors.New("The given args aren not in a valid range, try something like: [1 1]")
	}

	duration := time.Since(start)
	fmt.Println(fmt.Sprintf("This solution took: %s", duration))

	return answer, nil
}
