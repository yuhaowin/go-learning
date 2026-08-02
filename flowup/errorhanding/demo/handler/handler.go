package handler

import (
	"fmt"
	"net/http"

	"github.com/yuhaowin/go-learning/flowup/errorhanding/demo/apptype"
)

// 编译期断言：确保 HandleHello 的签名始终满足 apptype.AppHandler，
// 一旦签名改动导致不匹配，这一行会直接编译报错，比运行时才发现更早。
var _ apptype.AppHandler = HandleHello

func HandleHello(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")
	if name == "" {
		return fmt.Errorf("missing query param: name")
	}
	fmt.Fprintf(w, "hello, %s\n", name)
	return nil
}
