package robot

type Robot struct {
	X           int
	Y           int
	Orientation string
	IsLost      string
}

type Grid struct {
	X         int
	Y         int
	LostRobot []Robot
}
