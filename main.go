package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>NO</title>
		<style>
			body {
				display: flex;
				justify-content: center;
				align-items: center;
				flex-direction: column;
				height: 100vh;
				margin: 0;
				background-color: white;
			}
			h1 {
				font-size: 10em;
				color: black;
				margin: 0;
			}
			div {
				margin-top: 20px;
			}
			a {
				font-size: 1.2em;
				color: blue;
				text-decoration: none;
			}
			a:hover {
				text-decoration: underline;
			}
		</style>
	</head>
	<body>
		<h1>NO</h1>
		<div>
			<a href="https://github.com/joswayski/areweauthyet" target="_blank" rel="noopener noreferrer">
				https://github.com/joswayski/areweauthyet
			</a>
		</div>
	</body>
	</html>`
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
