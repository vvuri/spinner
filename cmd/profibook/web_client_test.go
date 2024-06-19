package profibook

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestWebClient(t *testing.T) {
	URL := "http://ya.ru/"
	data, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer data.Body.Close()
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
