package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

// 这是 file listing 包里 userError（这是一个具体类型，和 web.go 里 main 包中同名的 userError 接口不是一回事，只是恰好同名）。
// 1. type userError string：基于 string 定义了一个新类型，它本质就是一个字符串，可以直接用 userError("xxx") 构造。
// 2. func (e userError) Error() string：给这个类型实现 Error() string 方法，内容直接调用 Message()。这满足了内置的 error 接口。
// 3. func (e userError) Message() string：把底层的字符串值转回 string 返回。这满足了 web.go 里 userError 接口要求的 Message() string 方法。
// 因为同时有 Error() 和 Message() 两个方法，这个类型天然就满足了 web.go 中定义的 userError 接口。
type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println()
	if strings.Index(
		request.URL.Path, prefix) != 0 {
		return userError(fmt.Sprintf("path %s must start with %s", request.URL.Path, prefix))
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
