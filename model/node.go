package model

type Node struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Latitude  int    `json:"latitude"`
	Longitude int    `json:"longitude"`
}
