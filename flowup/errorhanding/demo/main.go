package main

import (
	"golearning/flowup/errorhanding/demo/apptype"
	"golearning/flowup/errorhanding/demo/handler"
	"log"
	"net/http"
)

func errWrapper(h apptype.AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func main() {
	// 这里显式传入的是 apptype.AppHandler 类型的 handler.HandleHello，
	// 而 handler 包里已经用 var _ apptype.AppHandler = HandleHello 显式关联过了。
	http.HandleFunc("/hello", errWrapper(handler.HandleHello))
	log.Println("listening on :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
