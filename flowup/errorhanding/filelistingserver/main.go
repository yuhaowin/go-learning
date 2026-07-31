package main

import (
	"golearning/flowup/errorhanding/filelistingserver/filelisting"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

// userError 标记可直接展示给用户的错误：内嵌 error（Error() string 供日志记录），
// 另加 Message() string 作为对外展示的提示文案，与内部系统错误区分开。
// 相当于 userError 要求同时满足 error 接口（即有 Error() string 方法）和自己新增的 Message() string 方法。
type userError interface {
	error // 接口嵌入（interface embedding）
	Message() string
}

type appHandler func(http.ResponseWriter, *http.Request) error

func wrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			log.Printf("Error occurred handling request: %s", err.Error())
			// user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", wrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
