package main

import "fmt"

type SalaryCalc interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

type Contract struct {
	empId    int
	basicPay int
}

type FreeLancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicPay
}

func (f FreeLancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

func totalExpense(s []SalaryCalc) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total salary per month $%d\n", expense)
}

func main() {
	pemp := Permanent{
		empId:    101,
		basicPay: 8000,
		pf:       20,
	}
	cemp := Contract{
		empId:    201,
		basicPay: 10000,
	}
	femp := FreeLancer{
		empId:       301,
		ratePerHour: 90,
		totalHours:  100,
	}
	employees := []SalaryCalc{pemp, cemp, femp}
	totalExpense(employees)
}
