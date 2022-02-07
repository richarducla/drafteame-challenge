package robot

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

var (
	cardenalPoints = []string{"N", "E", "S", "W"}
)

func Process(instructions []string) ([]Robot, error) {
	result := []Robot{}

	coordinates := strings.Split(instructions[0], " ")

	grid, err := getGrid(coordinates)
	if err != nil {
		return result, err
	}

	err = validateCoordinates(grid)
	if err != nil {
		return result, err
	}

	i := 1
	for i < len(instructions) {
		robot, err := getRobot(strings.Split(instructions[i], " "))
		if err != nil {
			return result, err
		}

		executeInstructions(&robot, strings.Split(instructions[i+1], ""), &grid)

		result = append(result, robot)
		i = i + 2
	}
	return result, nil
}

func validateCoordinates(grid Grid) error {
	if grid.X > 50 || grid.Y > 50 {
		return errors.New("the grid is invalid")
	}
	return nil
}

func getGrid(coordinates []string) (Grid, error) {
	x, err := strconv.Atoi(coordinates[0])
	if err != nil {
		log.Print("Error in coordinate X")
		return Grid{}, errors.New("coordinate grid x invalid")
	}

	y, err := strconv.Atoi(coordinates[1])
	if err != nil {
		log.Print("Error in coordinate Y")
		return Grid{}, errors.New("coordinate grid y invalid")
	}

	return Grid{
		X: x,
		Y: y,
	}, nil
}

func getRobot(robotProperties []string) (Robot, error) {
	x, err := strconv.Atoi(robotProperties[0])
	if err != nil {
		log.Print("Error in coordinate X")
		return Robot{}, errors.New("coordinate x invalid")
	}

	y, err := strconv.Atoi(robotProperties[1])
	if err != nil {
		log.Print("Error in coordinate Y")
		return Robot{}, errors.New("coordinate y invalid")
	}

	return Robot{
		X:           x,
		Y:           y,
		Orientation: robotProperties[2],
	}, nil
}

func executeInstructions(robot *Robot, instructions []string, grid *Grid) {
	for i := 0; i < len(instructions); i++ {
		if robot.IsLost == "LOST" {
			break
		}

		if instructions[i] == "F" {
			move(robot, grid)
		} else {
			turn(robot, instructions[i])
		}
	}
}

func move(robot *Robot, grid *Grid) {
	switch robot.Orientation {
	case "N":
		if robot.Y == grid.Y {
			if !checkLost(robot, grid) {
				robot.IsLost = "LOST"
				grid.LostRobot = append(grid.LostRobot, *robot)
			}
		} else {
			robot.Y++
		}
	case "S":
		if robot.Y == 0 {
			robot.IsLost = "LOST"
		} else {
			robot.Y--
		}
	case "E":
		if robot.X == grid.X {
			if !checkLost(robot, grid) {
				robot.IsLost = "LOST"
				grid.LostRobot = append(grid.LostRobot, *robot)
			}
		} else {
			robot.X++
		}
	case "W":
		if robot.X == 0 {
			robot.IsLost = "LOST"
		} else {
			robot.X--
		}

	}
}

func turn(robot *Robot, instruction string) {
	newOrientation := findIndex(robot.Orientation)
	if instruction == "L" {
		newOrientation = (newOrientation + 4 - 1) % 4
	} else {
		newOrientation = (newOrientation + 1) % 4
	}
	robot.Orientation = cardenalPoints[newOrientation]
}

func checkLost(robot *Robot, grid *Grid) bool {
	isLost := false
	for i := 0; i < len(grid.LostRobot); i++ {
		if grid.LostRobot[i].X == robot.X && grid.LostRobot[i].Y == robot.Y {
			isLost = true
			break
		}
	}
	return isLost
}

func findIndex(orientation string) int {
	for i, n := range cardenalPoints {
		if orientation == n {
			return i
		}
	}
	return -1
}
