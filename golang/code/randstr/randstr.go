package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r *rand.Rand

func main() {
	fmt.Println(RandString(24))
	fmt.Println(RandString(64))
}

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
