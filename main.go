package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

type sortCookie []*http.Cookie

func (s sortCookie) Len() int {
	return len(s)
}
func (s sortCookie) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortCookie) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

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
	sort.Sort(sortCookie(request.Cookies()))
	for _, cookie := range request.Cookies() {
		builder.WriteString(cookie.Name)
		builder.WriteString(":")
		s, _ := url.QueryUnescape(cookie.Value)
		builder.WriteString(s)
		// builder.WriteString(fmt.Sprintf("%#v", cookie.Value))
		builder.WriteString("\n")
	}
	writer.Header().Add("Content-Type", "text/plain")
	str := builder.String()
	fmt.Fprint(writer, str)
	fmt.Print(str)
}

func main() {
	port := flag.Int("port", 8080, "listening on")
	flag.Parse()
	fmt.Printf("Server listen on: " + strconv.Itoa(*port))
	http.HandleFunc("/", capture)
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
