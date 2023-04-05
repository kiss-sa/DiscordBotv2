package commands

import (
	"bufio"
	"math/rand"
	"os"

	"discordv2.at/m/v2/config"
)

func getPun() (string, error) {
	punFile, err := os.Open(config.Config.FilePath + "puns.txt")
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
