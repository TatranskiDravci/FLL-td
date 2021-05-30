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
	nameR := name + "R"
	nameG := name + "G"
	nameB := name + "B"

	os.Setenv(nameR, R)
	os.Setenv(nameG, G)
	os.Setenv(nameB, B)
}

func GetColor(name string) ([3]int, bool) {
	nameR := name + "R"
	nameG := name + "G"
	nameB := name + "B"
	R := 0
	G := 0
	B := 0

	Rs, okR := os.LookupEnv(nameR)
	if !okR {
		fmt.Println("COLOR NOT SET")
		return [3]int{-1, -1, -1}, okR
	} else {
		R, _ = strconv.Atoi(Rs)
	}

	Gs, okG := os.LookupEnv(nameG)
	if !okG {
		fmt.Println("COLOR NOT SET")
		return [3]int{-1, -1, -1}, okG
	} else {
		G, _ = strconv.Atoi(Gs)
	}

	Bs, okB := os.LookupEnv(nameB)
	if !okB {
		fmt.Println("COLOR NOT SET")
		return [3]int{-1, -1, -1}, okB
	} else {
		B, _ = strconv.Atoi(Bs)
	}

	return [3]int{R, G, B}, true
}