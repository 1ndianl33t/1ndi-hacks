package main

import (
        "bufio"
        "fmt"
        "github.com/tomnomnom/gahttp"
        "net/http"
        "os"
        "time"
)

func printStatus(req *http.Request, resp *http.Response, err error) {
        if err != nil {
                return
        }
        fmt.Printf("[%d] : %s\n", resp.StatusCode, req.URL)
}

func main() {
        p := gahttp.NewPipeline()
        p.SetConcurrency(20)
        p.SetRateLimit(time.Second * 10)
        urls := gahttp.Wrap(printStatus, gahttp.CloseBody)
        sc := bufio.NewScanner(os.Stdin)
        for sc.Scan() {
                p.Get(sc.Text(), urls)
        }
        
        p.Done()

        p.Wait()
}
