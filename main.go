package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/artem-shestakov/go/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handler("/bye", gh)

	err := http.ListenAndServe(":8080", sm)
	if err != nil {
		fmt.Println(err)
	}

}
