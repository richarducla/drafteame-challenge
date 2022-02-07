package main

import (
	"drafteame/internal/robot"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileBytes, err := ioutil.ReadFile("./files/template.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	instructions := strings.Split(string(fileBytes), "\n")

	result, err := robot.Process(instructions)
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(result); i++ {
		robot := result[i]
		fmt.Println(robot.X, robot.Y, robot.Orientation, robot.IsLost)
	}
}
