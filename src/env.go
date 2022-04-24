/*
	devnote:
		This codes is a mess beyond saving. The legacy code cleanup was
		largely unsuccessful 'round here, thus much of the old, unmaintainable
		codebase still poisons this file. There is no reason to comment individual
		functions in this file, as such effort would be beyond useless. Hopefully
		this piece of code dies and disappears after the competition. It is truly
		marvelous how horrendous and unreadable it is. But at least it works, somewhat...

		This code is designed to interface with environment files located in data/. These
		files are primitive key=value pair data files, meant for storage of repeatedly
		used data such as color packets.
*/
package main

import (
	"io/ioutil"
	"io/fs"
	"strconv"
)

func RequestValue(key string, fname string) (string, bool) {
	data, err := ioutil.ReadFile("../data/" + fname)
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

func ClearKeys(fname string) {
	ioutil.WriteFile("../data/" + fname, []byte(""), fs.ModePerm)
}

func ClearAll() {
	ClearKeys("col")
	ClearKeys("err")
}

func GetColor(name, fname string) ([3]int, bool) {
	r, ok := RequestValue(name + "R", fname)
	if !ok { return [3]int{-1, -1, -1}, ok }
	R, _ := strconv.Atoi(r)

	g, ok := RequestValue(name + "G", fname)
	if !ok { return [3]int{-1, -1, -1}, ok }
	G, _ := strconv.Atoi(g)

	b, ok := RequestValue(name + "B", fname)
	if !ok { return [3]int{-1, -1, -1}, ok }
	B, _ := strconv.Atoi(b)

	return [3]int{R, G, B}, true
}

func GetColor2(name string) ([2][3]int, bool) {
	col, ok1 := GetColor(name, "col")
	err, ok2 := GetColor(name, "err")

	return [2][3]int{col, err}, ok1 && ok2
}
