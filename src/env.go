package main

import (
	"io/ioutil"
	"io/fs"
	"strconv"
	"fmt"
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

func SetProfile(k [3]float64, l [3]int, id string) {
	varsK := "kR=" + fmt.Sprintf("%f", k[0]) + "\nkG=" + fmt.Sprintf("%f", k[1]) + "\nkB=" + fmt.Sprintf("%f", k[2]) + "\n"
	varsL := "lR=" + strconv.Itoa(l[0]) + "\nlG=" + strconv.Itoa(l[1]) + "\nlB=" + strconv.Itoa(l[2]) + "\n"

	vars := varsK + varsL
	ioutil.WriteFile("../data/prof" + id, []byte(vars), fs.ModePerm)
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

func GetColorf(name, fname string) ([3]float64, bool) {
	r, ok := RequestValue(name + "R", fname)
	if !ok { return [3]float64{-1., -1., -1.}, ok }
	R, _ := strconv.ParseFloat(r, 64)

	g, ok := RequestValue(name + "G", fname)
	if !ok { return [3]float64{-1., -1., -1.}, ok }
	G, _ := strconv.ParseFloat(g, 64)

	b, ok := RequestValue(name + "B", fname)
	if !ok { return [3]float64{-1., -1., -1.}, ok }
	B, _ := strconv.ParseFloat(b, 64)

	return [3]float64{R, G, B}, true
}

func GetColor2(name string) ([2][3]int, bool) {
	col, ok1 := GetColor(name, "col")
	err, ok2 := GetColor(name, "err")

	return [2][3]int{col, err}, ok1 && ok2
}

func GetProfile(id string) ([3]float64, [3]int, bool) {
	k, ok1 := GetColorf("k", "prof" + id)
	l, ok2 :=  GetColor("l", "prof" + id)

	return k, l, ok1 && ok2
}
