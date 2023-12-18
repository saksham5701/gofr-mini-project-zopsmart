package models


type Employee struct {
	Eid            int    `json:"eid"`
	Empname        string `json:"empname"`
	Salary         int    `json:"salary"`
	Email          string `json:"email"`
}
