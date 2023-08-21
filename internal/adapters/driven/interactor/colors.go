package interactor

type color string

const (
	colorRed    = color("\033[31m")
	colorGreen  = color("\033[32m")
	colorYellow = color("\033[33m")
	colorBlue   = color("\033[34m")
	colorPurple = color("\033[35m")
	colorCyan   = color("\033[36m")
	colorWhite  = color("\033[37m")
)
