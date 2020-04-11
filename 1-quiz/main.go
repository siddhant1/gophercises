package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type ans struct {
	correctAns string
	userAns    string
}

var strFlag = flag.String("problem", "problem.csv", "Provide the csv problem file")

func print(a []ans) {
	for _, value := range a {
		fmt.Println(value)
	}
}

var correctAns []ans
var incorrentAns []ans

func takeInput(c chan string) {
	var input string
	fmt.Scanf("%s", &input)
	c <- input
}

func main() {
	flag.Parse()

	f, err := os.Open(*strFlag)

	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}

	timer := time.NewTimer(2 * time.Second)

problemLoop:
	for _, value := range records {
		fmt.Print(value[0])
		c := make(chan string)
		go takeInput(c)
		select {
		case input := <-c:
			if value[0] == input {
				correctAns = append(correctAns, ans{userAns: input, correctAns: value[1]})
			} else {
				incorrentAns = append(incorrentAns, ans{userAns: input, correctAns: value[1]})
			}
		case <-timer.C:
			fmt.Printf("\nEXITING BRO")
			break problemLoop
		}
	}

	fmt.Println("\nList of correct ans")
	print(correctAns)
	fmt.Println("\nList of incorrent ans")
	print(incorrentAns)

}
