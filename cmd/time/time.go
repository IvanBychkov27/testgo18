package main

import (
	"fmt"
	"time"
)

func main() {
	workDate := getWorkDateAtStartup()
	fmt.Printf("work date: %d\n", workDate)
}

// получить рабочую дату при запуске программы и проверить ее заполнение для всех Networks
func getWorkDateAtStartup() int {
	t := time.Now().UTC()
	nowDate := buildDate(t)
	fmt.Println(nowDate)

	lastWorkingDay := lastWorkingDayMonth(t.Year()+2, 1)
	fmt.Println(lastWorkingDay)

	var workDate int
	if len(lastWorkingDay) > 0 {
		workDate = lastWorkingDay[0]
	}

	for _, d := range lastWorkingDay {
		if d >= nowDate {
			break
		}
		workDate = d
	}
	return workDate
}

// последний рабочий день месяца от текущего времени до заданной даты (year, month)
func lastWorkingDayMonth(year, month int) []int {
	lastWorkingDays := make([]int, 0)
	t := time.Now().UTC().AddDate(0, -1, 0) // + один предыдущий месяца
	for {
		if t.Year() >= year && t.Month() >= time.Month(month) {
			break
		}
		tmp := t
		t = t.Add(time.Hour * 24)
		if tmp.Month() != t.Month() {
			if tmp.Weekday() == time.Saturday {
				lastWorkingDays = append(lastWorkingDays, buildDate(tmp.AddDate(0, 0, -1)))
				continue
			}
			if tmp.Weekday() == time.Sunday {
				lastWorkingDays = append(lastWorkingDays, buildDate(tmp.AddDate(0, 0, -2)))
				continue
			}
			lastWorkingDays = append(lastWorkingDays, buildDate(tmp))
		}
	}
	return lastWorkingDays
}

func buildDate(t time.Time) int {
	return t.Year()*10000 + int(t.Month())*100 + t.Day()
}

// последовательность дней до заданной даты (year, month)
func sequenceDays(year, month int) {
	t := time.Now().UTC()
	fmt.Println(t.Format(time.RFC1123))
	for {
		t = t.Add(time.Hour * 24)
		fmt.Println(t.Format(time.RFC1123))
		if t.Year() == year && t.Month() == time.Month(month) {
			break
		}
	}
}

func checkTime() {
	tMy := time.Date(2022, 10, 16, 11, 45, 0, 0, time.UTC)
	t := time.Now().UTC()

	fmt.Println("String:", t.String())
	fmt.Printf("Time  : %d:%d:%d\n", t.Hour(), t.Minute(), t.Second())
	fmt.Println("Day    :", t.Day())
	fmt.Println("Month :", t.Month())
	fmt.Println("Year  :", t.Year())
	fmt.Println("YearDay:", t.YearDay())
	fmt.Println("Format RFC1123:", t.Format(time.RFC1123))
	fmt.Println("Format RFC1123:", tMy.Format(time.RFC1123))
	fmt.Println("Month:", tMy.Format("1"))
	fmt.Println("Weekday:", tMy.Weekday())
}
