package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
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

	var correctAns []ans
	var incorrentAns []ans

	for _, value := range records {
		fmt.Print(value[0])
		var input string
		fmt.Scanf("%s", &input)
		if value[1] == input {
			correctAns = append(correctAns, ans{userAns: input, correctAns: value[1]})
		} else {
			incorrentAns = append(incorrentAns, ans{userAns: input, correctAns: value[1]})
		}
		fmt.Println("")
	}

	fmt.Println("List of correct ans")
	print(correctAns)
	fmt.Println("List of incorrent ans")
	print(incorrentAns)

}
