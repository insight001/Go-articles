package main

//Articles ...
var Articles []Article

func main() {

	Articles = []Article{
		Article{Title: "Hello", Desc: "", Content: ""},
	}

	handleRequests()
}
