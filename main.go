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
	var studentStats []studentStat
	for _, row := range students {
		var st studentStat
		total := float32(row.test1Score+row.test2Score+row.test3Score+row.test4Score) / 4
		st.student = row
		st.finalScore = total
		st.grade = getGrade(total)
		studentStats = append(studentStats, st)
	}

	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	topper := gradedStudents[0]
	for _, row := range gradedStudents {
		if row.finalScore > topper.finalScore {
			topper = row
		}
	}
	return topper
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	students := parseCSV("grades.csv")
	calculateGrade(students)
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

func getGrade(score float32) Grade {
	if score >= 70 {
		return A
	} else if score < 70 && score >= 50 {
		return B
	} else if score < 50 && score >= 35 {
		return C
	} else {
		return F
	}
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	topperPerUni := make(map[string]studentStat)
	studentsUniMap := createStudentUniMap(gs)
	for uni, students := range studentsUniMap {
		t := findOverallTopper(students)
		topperPerUni[uni] = t
	}
	return topperPerUni
}
func createStudentUniMap(gs []studentStat) map[string][]studentStat {
	studentsUniMap := make(map[string][]studentStat)
	for _, stStat := range gs {
		studentsUniMap[stStat.university] = append(studentsUniMap[stStat.university], stStat)
	}
	return studentsUniMap
}
