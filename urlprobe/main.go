

package main

import (
        "bufio"
        "fmt"
        "net/http"
        "os"
        "time"
        "github.com/fatih/color"
        "github.com/tomnomnom/gahttp"
)

func Banner() {
        color.HiCyan(`
 ____ _____________.____   __________              ___.
|    |   \______   \    |  \______   \_______  ____\_ |__   ____
|    |   /|       _/    |   |     ___/\_  __ \/  _ \| __ \_/ __ \
|    |  / |    |   \    |___|    |     |  | \(  <_> ) \_\ \  ___/
|______/  |____|_  /_______ \____|     |__|   \____/|___  /\___  >
                 \/        \/                           \/     \/   V 1.0`)

        color.HiYellow("URLProbe:- Urls Status Code Checker")

        color.HiRed("https://github.com/1ndianl33t")
}
func printStatus(req *http.Request, resp *http.Response, err error) {
        if err != nil {
                return
        }
        fmt.Printf("[%d] L %d : %s\n", resp.StatusCode, resp.ContentLength, req.URL)
}

func main() {
        Banner ()
        p := gahttp.NewPipeline()
        p.SetConcurrency(100)
        p.SetRateLimit(time.Second * 05)
        urls := gahttp.Wrap(printStatus, gahttp.CloseBody)
        sc := bufio.NewScanner(os.Stdin)
        for sc.Scan() {
                p.Get(sc.Text(), urls)
        }
        
        p.Done()

        p.Wait()
}


