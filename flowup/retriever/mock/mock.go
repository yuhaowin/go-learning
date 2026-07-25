package mock

import "bufio"

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	return r.Contents
}

func (r Retriever) Read(p []byte) (n int, err error) {

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {

	}

	return 0, err
}
