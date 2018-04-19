package main
import (
    "fmt"; "net/http"
    "io/ioutil"; "os"
    "bufio"; "strings"
    //"golang.org/x/net/html"
    )

func err_chk(err error, msg string){
    if err != nil{
        fmt.Println("Error Code: "+msg)
        os.Exit(-1)
    }
}

func scrape(url string) []byte{
    resp, err := http.Get("http://"+url)
    err_chk(err, "g10")

    bytes, err := ioutil.ReadAll(resp.Body)
    err_chk(err, "i04")
    resp.Body.Close()

    return bytes
}

func main(){
    fmt.Println("enter a url")
    reader := bufio.NewReader(os.Stdin)
    text, err := reader.ReadString('\n')
    //error check the string
    err_chk(err, "t10")
    text = strings.Trim(text, "\n")
    
    fmt.Println(text)
    out := scrape(text)
    fmt.Println("HTML:\n\n", string(out))
}
