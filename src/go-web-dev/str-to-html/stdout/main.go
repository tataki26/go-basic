package main

import "fmt"

func main() {
	greeting := "Hello, World!"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>My First GO Web Page</title>
		</head>
		<body>
			<h1>` + greeting + `</h1>
		</body>
	</html>
	`

	fmt.Println(tpl)
}
