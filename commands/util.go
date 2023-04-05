package commands

import (
	"bufio"
	"math/rand"
	"os"
)

func getPun() (string, error) {
	punFile, err := os.Open("./src/puns.txt")
	if err != nil {
		return "", err
	}

	fileScanner := bufio.NewScanner(punFile)
	fileScanner.Split(bufio.ScanLines)
	var puns []string

	for fileScanner.Scan() {
		puns = append(puns, fileScanner.Text())
	}

	ind := rand.Intn(len(puns))

	return puns[ind], nil
}
