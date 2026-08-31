package memo_test

import (
	"testing"

	"github.com/yuhaowin/go-learning/ch09/memo/memotest"
	memo "github.com/yuhaowin/go-learning/ch09/memo/version4"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
