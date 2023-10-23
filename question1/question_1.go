package main

import "fmt"

type Order struct {
	OrderId        int
	ShippingMethod int // 1, 2, 3
	Printed        bool
}

func Print1() string {
	return "PrintData_1"
}

func Print2() string {
	return "PrintData_2"
}

func Print3() string {
	return "PrintData_3"
}

func Process1(orderId int) string {
	return fmt.Sprintf("ProcessData_1_%d", orderId)
}

func Process2(orderId int) string {
	return fmt.Sprintf("ProcessData_2_%d", orderId)
}

func Process3(orderId int) string {
	return fmt.Sprintf("ProcessData_3_%d", orderId)
}

func GetOrders() []Order {
	return []Order{
		{1, 1, false},
		{2, 2, false},
		{3, 3, true},
		{4, 1, false},
		{5, 2, true},
		{6, 3, false},
		{7, 1, true},
		{8, 2, false},
		{9, 3, false},
		{10, 1, true},
		{11, 2, true},
		{12, 3, false},
		{13, 1, false},
		{14, 2, true},
		{15, 3, false},
		{16, 1, true},
		{17, 2, true},
		{18, 3, false},
		{19, 1, false},
		{20, 2, false},
	}
}

type Result struct {
	Method1Data string
	Method2Data string
	Method3Data string
}

func (r *Result) Data() string {
	return fmt.Sprintf("Method 1: %s\nMethod 2: %s\nMethod 3: %s\n", r.Method1Data, r.Method2Data, r.Method3Data)
}

func (r *Result) Print() {
	fmt.Println(r.Data())
}

func (r *Result) Check() {
	ans := "Method 1: PrintData_1:ProcessData_1_1,PrintData_1:ProcessData_1_4,ProcessData_1_7,ProcessData_1_10,PrintData_1:ProcessData_1_13,ProcessData_1_16,PrintData_1:ProcessData_1_19\nMethod 2: PrintData_2:ProcessData_2_2,ProcessData_2_5,PrintData_2:ProcessData_2_8,ProcessData_2_11,ProcessData_2_14,ProcessData_2_17,PrintData_2:ProcessData_2_20\nMethod 3: ProcessData_3_3,PrintData_3:ProcessData_3_6,PrintData_3:ProcessData_3_9,PrintData_3:ProcessData_3_12,PrintData_3:ProcessData_3_15,PrintData_3:ProcessData_3_18\n"
	fmt.Println("Pass: ", r.Data() == ans)
}

func Process(orders []Order) Result {
	// TODO: implement this function
	// 1. Print the data if the order is not printed
	// 2. Process the data
	// 3. add the data to the result (the result need to be sorted)
	// 4. return the result
	// result data format
	// - if data not printed: [PrintData:ProcessData]
	// - if data printed: [ProcessData]
	// each data combine with comma (,)
	// example: [PrintData:ProcessData,PrintData:ProcessData,ProcessData]

	return Result{}
}

func main() {
	orders := GetOrders()
	result := Process(orders)
	result.Print()
	result.Check()
}
