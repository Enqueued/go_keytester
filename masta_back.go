package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	//"golang.org/x/net/html"
	)

func main(){
	fmt.Println("Hello der")
	resp, err := http.Get("https://www.google.com")
	
	if err != nil{
		fmt.Println("error code: g10")
		os.Exit(-1)
	}
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		fmt.Println("error code: ioR")
		os.Exit(-1)
	}

	fmt.Println("HTML:\n\n", string(bytes))
	resp.Body.Close()
}
