package main

//Article ...
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

//Response ...
type Response struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Article `json:"data"`
}
