package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"time"
)

func calculateTime(pastTime int64) int64 {
	timestamp := time.Now().Unix()
	timestamp -= pastTime
	timeleft := timestamp % 100 
	var section int64
	if timeleft <= 25 {
		section = 1
	} else if timeleft <= 50 {
		section = 2
	} else if timeleft <= 75 {
		section = 3
	} else {
		section = 4
	}
	computedtime := (timestamp / 100) + section

	return computedtime
}

func sha256String(input string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashed := hasher.Sum(nil)
	hasher.Reset()
	return hashed
}

func getCodeOutput(hashSize int, lenPassword int, hashBytes []byte) string {
	hashRange := hashSize / lenPassword
	output_code := ""

	for i := 0; i < lenPassword; i++ {
		sum := 0
		for _, j := range hashBytes[i : i+hashRange] {
			sum += int(j)
		}
		output_code += strconv.Itoa(sum % 10) // get last number
	}

	return output_code
}

func main() {

	inputPassword := os.Args[1]
	lenPassword, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("invalid number.")
		os.Exit(1)
	}

	customTime := calculateTime(0)
	hashed := sha256String(inputPassword+strconv.FormatInt(customTime,10))

	output_code := getCodeOutput(32,lenPassword,hashed) // 32 for sha256 length
	fmt.Println("Your code:", output_code)
}
