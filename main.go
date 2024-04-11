package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {
	var students []student
	file, err := os.Open(filePath)
	check(err)

	defer file.Close()

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	check(err)

	for i, row := range records {
		if i > 0 {
			info, err := createStudentFromRow(row)
			check(err)
			students = append(students, info)
		}
	}

	return students
}

func calculateGrade(students []student) []studentStat {
	return nil
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	return studentStat{}
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	parseCSV("grades.csv")
}

func createStudentFromRow(row []string) (student, error) {
	var st student
	if len(row) > 7 {
		fmt.Println("Invalid row")
		return student{}, errors.New("invalid csv")
	}
	st.firstName = row[0]
	st.lastName = row[1]
	st.university = row[2]
	t1Score, errT1 := strconv.Atoi(row[3])
	t2Score, errT2 := strconv.Atoi(row[4])
	t3Score, errT3 := strconv.Atoi(row[5])
	t4Score, errT4 := strconv.Atoi(row[6])
	if errT1 != nil || errT2 != nil || errT3 != nil || errT4 != nil {
		return student{}, errors.New("invalid csv row")
	}

	st.test1Score = t1Score
	st.test2Score = t2Score
	st.test3Score = t3Score
	st.test4Score = t4Score
	return st, nil
}
