package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	bin := binary(name)
    fmt.Fprintf(w, "Hello, %s or %s!", name, bin)
}

func binary(s string) string {
    res := ""
    for _, c := range s {
        res = fmt.Sprintf("%s%.8b", res, c)
    }
    return res
}