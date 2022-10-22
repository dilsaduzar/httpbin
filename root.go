package main

import (
	"fmt"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintln(w, "Hello World!\n", t.Format("02-01-2006 15:04:05"))
}
