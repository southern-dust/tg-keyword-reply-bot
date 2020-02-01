package main

import (
	"bufio"
	"os"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func readText() []string {
	file, err := os.Open("rules.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// This is our buffer now
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}


func addRuleFromText(){
	var gid int64
	var gid_int int
	var rule string
	var err error

	var lines=readText()
	for _, line := range lines {
		flysnowRegexp := regexp.MustCompile(`-?[0-9A-Za-z]+`)
		params := flysnowRegexp.FindStringSubmatch(line)

		if gid_int, err = strconv.Atoi(params[0]); err != nil {
			//err, ignore this line
			continue
		} else{
			gid = int64(gid_int)
			rule = strings.TrimLeft(line, params[0])
			addRule(gid, rule)
		}
	}
}

