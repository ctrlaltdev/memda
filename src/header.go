package memda

import (
	"fmt"
	"math/rand"
	"time"
)

func printHeader() {
	headers := [][7]string{}

	h1 := [...]string{
		"                           _       ",
		"  /\\/\\   ___ _ __  ___  __| | __ _ ",
		" /    \\ / _ \\ '_ \\` _ \\/ _\\`|/ _\\`|",
		"/ /\\/\\ \\  __/ | | | | | (_| | (_| |",
		"\\/    \\/\\___|_| |_| |_|\\__,_|\\__,_|",
		"",
		"",
	}

	h2 := [...]string{
		".___  ___.  _______ .___  ___.  _______       ___      ",
		"|   \\/   | |   ____||   \\/   | |       \\     /   \\     ",
		"|  \\  /  | |  |__   |  \\  /  | |  .--.  |   /  ^  \\    ",
		"|  |\\/|  | |   __|  |  |\\/|  | |  |  |  |  /  /_\\  \\   ",
		"|  |  |  | |  |____ |  |  |  | |  '--'  | /  _____  \\  ",
		"|__|  |__| |_______||__|  |__| |_______/ /__/     \\__\\ ",
		"                                                       ",
	}

	h3 := [...]string{
		"                                                           ",
		" _|      _|                                  _|            ",
		" _|_|  _|_|    _|_|    _|_|_|  _|_|      _|_|_|    _|_|_|  ",
		" _|  _|  _|  _|_|_|_|  _|    _|    _|  _|    _|  _|    _|  ",
		" _|      _|  _|        _|    _|    _|  _|    _|  _|    _|  ",
		" _|      _|    _|_|_|  _|    _|    _|    _|_|_|    _|_|_|  ",
		"                                                           ",
	}

	headers = append(headers, h1)
	headers = append(headers, h2)
	headers = append(headers, h3)

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(headers))

	for _, l := range headers[i] {
		fmt.Println(l)
	}
}
