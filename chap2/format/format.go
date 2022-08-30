package main

import (
	"net/http"
	"os"
)

func main() {
	// 引数のio.WriterがWrite()メソッドを持っているので加工して書き出すことができる
	// fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())

	// encoder := json.NewEncoder(os.Stdout)
	// encoder.SetIndent("", "    ")
	// encoder.Encode(map[string]string{
	// 	"example": "encoding/json",
	// 	"hello":   "World",
	// })

	request, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write(os.Stdout)
}
