package main

import (
	"bytes"
	"container/ring"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func part1(dat []byte) {
	var numList []string
	for _, num := range dat {
		numList = append(numList, string(num))
	}

	r := ring.New(len(numList))
	for i := 0; i < r.Len(); i++ {
		r.Value = numList[i]
		r = r.Next()
	}

	var matchSum int64
	i := 1
	for i <= r.Len() {
		if r.Value == r.Next().Value {
			//fmt.Printf("[match] cur: %s next: %s\n", r.Value, r.Next().Value)
			val, err := strconv.ParseInt(r.Value.(string), 10, 64)
			if err != nil {
				log.Fatalf("NaN wtf: %s", err)
			}
			//fmt.Printf("adding: %d\n", val)
			matchSum = matchSum + val
			//fmt.Printf("matchsum = %d\n", matchSum)
		}
		i++
		r = r.Next()
	}

	fmt.Println("[part1]", matchSum)
}

func part2(dat []byte) {
	var numList []string
	for _, num := range dat {
		numList = append(numList, string(num))
	}

	r := ring.New(len(numList))
	for i := 0; i < r.Len(); i++ {
		r.Value = numList[i]
		r = r.Next()
	}

	var matchSum int64
	i := 1
	for i <= r.Len() {
		var nextVal *ring.Ring
		j := 1
		nextVal = r
		for j <= r.Len()/2 {
			nextVal = nextVal.Next()
			j++
		}
		if r.Value == nextVal.Value {
			//fmt.Printf("[match] cur: %s next: %s\n", r.Value, r.Next().Value)
			val, err := strconv.ParseInt(r.Value.(string), 10, 64)
			if err != nil {
				log.Fatalf("NaN wtf: %s", err)
			}
			//fmt.Printf("adding: %d\n", val)
			matchSum = matchSum + val
			//fmt.Printf("matchsum = %d\n", matchSum)
		}
		i++
		r = r.Next()
	}

	fmt.Println("[part2]", matchSum)
}

func main() {
	dat, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	dat = bytes.TrimSpace(dat)

	part1(dat)
	part2(dat)
}
