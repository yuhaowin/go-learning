package apptype

import "net/http"

// AppHandler 是 main 包和 handler 包都会用到的公共类型，
// 放在单独的包里可以避免 main <-> handler 之间的循环导入。
type AppHandler func(w http.ResponseWriter, r *http.Request) error
