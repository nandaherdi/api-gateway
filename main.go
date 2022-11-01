package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//func main() {
//	router := httprouter.New()
//	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
//		fmt.Fprint(writer, "Hello HttpRouter")
//	})
//
//	server := http.Server{
//		Handler: router,
//		Addr:    "localhost:3000",
//	}
//
//	server.ListenAndServe()
//}

func main() {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		text := "Image : " + image
		fmt.Fprint(writer, text)
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()

	//request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	//recorder := httptest.NewRecorder()

	//router.ServeHTTP(recorder, request)

	//response := recorder.Result()
	//body, _ := io.ReadAll(response.Body)
}
