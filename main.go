package main

import (
	"flag"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func capture(writer http.ResponseWriter, request *http.Request) {
	var builder strings.Builder
	builder.WriteString("Method: " + request.Method + "\n")
	builder.WriteString("Host: " + request.Host + "\n")
	builder.WriteString("URL: " + request.URL.RequestURI() + "\n")

	builder.WriteString("\nHeaders: \n")
	headers := make([]string, 0, len(request.Header))
	for k := range request.Header {
		headers = append(headers, k)
	}
	sort.Strings(headers)
	for _, k := range headers {
		builder.WriteString(k)
		builder.WriteString(":")
		builder.WriteString(strings.Join(request.Header[k], "|"))
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
	cookies := make([]string, 0, len(request.Cookies()))
	for _, c := range request.Cookies() {
		cookies = append(cookies, c.String())
	}
	for _, cookie := range cookies {
		builder.WriteString(cookie)
		builder.WriteString("\n")
	}
	writer.Header().Add("Content-Type", "text/plain")
	str := builder.String()
	fmt.Fprintf(writer, str)
	fmt.Print(str)
}

func main() {
	port := flag.Int("port", 8080, "listening on")
	flag.Parse()
	fmt.Printf("Server listen on: " + strconv.Itoa(*port))
	http.HandleFunc("/", capture)
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
