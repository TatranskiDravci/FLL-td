package main

import (
	"os"
	"io/ioutil"
	"io/fs"
	"strconv"
	"fmt"
)

func RequestValue(key string) (string, bool) {
	data, err := ioutil.ReadFile("../data/env")
	keyRead := ""
	valueRead := ""
	keyNow := true
	foundKey := false
	if err == nil {
		datas := string(data)
		for i := 0; i < len(datas); i++ {
			if string(datas[i]) == "=" {
				if keyRead == key {
					foundKey = true
				}
				keyNow = false
				continue
			} else if string(datas[i]) == "\n" {
				keyNow = true
				if foundKey {
					return valueRead, true
				}
				keyRead = ""
				valueRead = ""
				continue
			}
			if keyNow {
				keyRead += string(datas[i])
			} else {
				valueRead += string(datas[i])
			}
		}
	}
	return "", false
}

func ClearKeys() {
	ioutil.WriteFile("../data/env", []byte(""), fs.ModePerm)
}

func SetColor(color [3]int, name string) {
	R := strconv.Itoa(color[0])
	G := strconv.Itoa(color[1])
	B := strconv.Itoa(color[2])
	nameR := name + "R"
	nameG := name + "G"
	nameB := name + "B"

	vars := nameR + "=" + R + "\n" + nameG + "=" + G + "\n" + nameB + "=" + B + "\n"

	data, err := ioutil.ReadFile("../data/env")
	datas := string(data)
	if err != nil {
		os.Create("../data/env")
	}

	ioutil.WriteFile("../data/env", []byte(vars + datas), fs.ModePerm)
}

func GetColor(name string) ([3]int, bool) {
	nameR := name + "R"
	nameG := name + "G"
	nameB := name + "B"
	var R, G, B int

	Rs, okR := RequestValue(nameR)
	if !okR {
		fmt.Println("COLOR NOT SET")
		return [3]int{-1, -1, -1}, okR
	} else {
		R, _ = strconv.Atoi(Rs)
	}

	Gs, okG := RequestValue(nameG)
	if !okG {
		fmt.Println("COLOR NOT SET")
		return [3]int{-1, -1, -1}, okG
	} else {
		G, _ = strconv.Atoi(Gs)
	}

	Bs, okB := RequestValue(nameB)
	if !okB {
		fmt.Println("COLOR NOT SET")
		return [3]int{-1, -1, -1}, okB
	} else {
		B, _ = strconv.Atoi(Bs)
	}

	return [3]int{R, G, B}, true
}