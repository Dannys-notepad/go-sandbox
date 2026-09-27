package main

import (
  "fmt"
  "math/rand/v2"
)

type Grade struct {
  Course string
  Unit int
  Score int
}

type Student struct {
  ID int
  Name string
  Age int
  CGPA float64
  Grades []Grade
}

// Student Methods

func (student *Student) calculateCGPA() {
  totalPoints, totalUnits := 0, 0

  for _, g := range student.Grades {
    if g.Score >= 70 {

      totalPoints += g.Unit * 5
      totalUnits += g.Unit
      
    } else if g.Score >= 60 {

      totalPoints += g.Unit * 4
      totalUnits += g.Unit 

    } else if g.Score >= 50 {

      totalPoints += g.Unit * 3
      totalUnits += g.Unit

    } else if g.Score >= 45 {

      totalPoints += g.Unit * 2
      totalUnits += g.Unit

    } else if g.Score >= 40 {

      totalPoints += g.Unit * 1
      totalUnits += g.Unit

    } else if g.Score >= 0 {

      totalPoints += g.Unit * 0
      totalUnits += g.Unit

    }
  }

  CGPA := 0.0

  if totalPoints > 0 && totalUnits > 0 {
    CGPA = float64(totalPoints) / float64(totalUnits)
  }

  student.CGPA = CGPA
}

func (student Student) displayINFO() {

  fmt.Println("----------------------")
  fmt.Printf("ID: %d\n", student.ID)
  fmt.Printf("Name: %s\n", student.Name)
  fmt.Printf("Age: %d\n", student.Age)
  fmt.Printf("CGPA: %.2f\n", student.CGPA)
  fmt.Println("Grades....")
  fmt.Println("Course. Unit. Score.")
  displayGrades(student.Grades)
  fmt.Println("----------------------")

}

// Other Functions

func displayGrades (grades []Grade) {
  for _, g := range grades {
    fmt.Println(g.Course, g.Unit, g.Score)
  }
}

func displayAllStudents (record []Student) {
  for _, student := range record {
    student.displayINFO()
  }
}

func findStudentByID (record []Student, id int) (any, bool) {

  low := 0
  high := len(record) - 1
  mid := 0
  
  for low <= high {
    mid = low + (high - low) / 2
    guess := record[mid]

    if guess.ID == id {
      return guess, true
    }

    if guess.ID > id {
      high = mid - 1
    } else {
      low = mid + 1
    }
  }
  
  return -1, false
}

func generateStudentsRecord (number int, record *[]Student) {

  recordSize := len(*record)
  id := 1000 + recordSize
  name := "Test Student"

  for i := 0; i < number; i++ {
    studentTemplate := Student {
      ID: id,
      Name: name,
      Age: 19 + rand.IntN(25 - 19 + 1),
      CGPA: 0.0,
      Grades: []Grade{
        {
          Course: "CHM 208.",
          Unit: 1,
          Score: rand.IntN(100),
        },
        {
          Course: "CHM 202.",
          Unit: 2,
          Score: rand.IntN(100),
        },
        {
          Course: "ICH 252.",
          Unit: 3,
          Score: rand.IntN(100),
        },
      },
    }
    *record = append(*record, studentTemplate)
    id += 1
  }
}

func calculateStudentsCGPA (record *[]Student) {
  for _, student := range *record {
    student.calculateCGPA()
  }
}

func findStudentWithHighestCGPA (record []Student) Student {
  n := len(record)
  highest := record[0]

  for i := 0; i < n; i++ {
    if record[i].CGPA > highest.CGPA {
      highest = record[i]
    }
  }

  return highest
}

func main() {
  studentsRecord := []Student{}
  generateStudentsRecord(3, &studentsRecord)
  calculateStudentsCGPA(&studentsRecord)
  fmt.Println(studentsRecord)
  //displayAllStudents(studentsRecord)
}