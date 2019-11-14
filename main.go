package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func capture(writer http.ResponseWriter, request *http.Request) {
	var builder strings.Builder
	builder.WriteString("Method: " + request.Method + "\n")
	builder.WriteString("Host: " + request.Host + "\n")
	builder.WriteString("URL: " + request.URL.RequestURI() + "\n")

	builder.WriteString("\nHeaders: \n")
	for k, v := range request.Header {
		builder.WriteString(k)
		builder.WriteString(":")
		builder.WriteString(strings.Join(v, "|"))
		builder.WriteString("\n")
	}
	builder.WriteString("\nParameters: \n")
	request.ParseForm()
	for parameter, value := range request.Form {
		builder.WriteString(parameter)
		builder.WriteString(":")
		builder.WriteString(strings.Join(value, "|"))
		builder.WriteString("\n")
	}
	builder.WriteString("\nCookies: \n")
	for _, cookie := range request.Cookies() {
		builder.WriteString(cookie.String())
		builder.WriteString("\n")
	}
	writer.Header().Add("Content-Type", "text/plain")
	str := builder.String()
	fmt.Fprintf(writer, str)
	fmt.Print(str)
}

func main() {
	port := flag.Int("port", 8080, "listening on")
	http.HandleFunc("/", capture)
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
