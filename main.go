package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
) 

type Problem struct{
	q string
	a string
}

func main(){
	csvFileName := flag.String("csv", "questions.csv", "csv file in the format of questions and answers")
	timeFrame := flag.Int64("time", 30, "Problem expiring time")
	flag.Parse()
	//Open file
	file, err := os.Open(*csvFileName)
	if err != nil{
		exit(fmt.Sprintf("Cannot open the file: %s\n", *csvFileName))
	}
	//Read file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil{
		exit("Something went wrong in reading the file")
	}

	timer := time.NewTimer(time.Duration(*timeFrame) * time.Second)
	
	
	
	problems := parseLines(lines)
	
	correct := 0

	for i, problem := range problems{
		fmt.Printf("%d. %s = ",i+1, problem.q)
		
		answerCh := make(chan string)
		
		go func ()  {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select{
			case <-timer.C:
				fmt.Printf("\nYou scored %d out of %d \n", correct, len(problems))
				return
			case answer := <- answerCh:
				if answer == problem.a{
					correct++	
				}
		}
	}
	fmt.Printf("\nYou scored %d out of %d \n", correct, len(problems))
}

func parseLines(lines [][]string)[]Problem{
	res := make([]Problem, len(lines))
	for i, line := range lines{
		res[i] = Problem{
			q: line[0],
			a: line[1],
		}
	}
	return res
}

func exit(str string){
	fmt.Println(str)
	os.Exit(1)
}

