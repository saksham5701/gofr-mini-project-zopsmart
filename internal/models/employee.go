package models


type Employee struct {
	Eid            int    `json:"eid"`
	Empname        string `json:"Empname"`
	Salary         int    `json:"salary"`
	Email          string `json:"email"`
}
