package main

import (
	"os"
	"strconv"
	"fmt"
)

func SetColor(color [3]int, name string) {
	R := string(color[0])
	G := string(color[1])
	B := string(color[2])
	Rname := name + "R"
	Gname := name + "G"
	Bname := name + "B"

	os.Setenv(Rname, R)
	os.Setenv(Gname, G)
	os.Setenv(Bname, B)
}

func GetColor(name string) ([3]int, bool) {
	Rname := name + "R"
	Gname := name + "G"
	Bname := name + "B"
	R := 0
	G := 0
	B := 0

	Rs, okR := os.LookupEnv(Rname)
	if !okR {
		fmt.Println("COLOR NOT SET")
		return [3]int{-1, -1, -1}, okR
	} else {
		R, _ = strconv.Atoi(Rs)
	}

	Gs, okG := os.LookupEnv(Gname)
	if !okG {
		fmt.Println("COLOR NOT SET")
		return [3]int{-1, -1, -1}, okG
	} else {
		G, _ = strconv.Atoi(Gs)
	}

	Bs, okB := os.LookupEnv(Bname)
	if !okB {
		fmt.Println("COLOR NOT SET")
		return [3]int{-1, -1, -1}, okB
	} else {
		B, _ = strconv.Atoi(Bs)
	}

	return [3]int{R, G, B}, true
}