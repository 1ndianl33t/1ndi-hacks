package main

import (
        "bufio"
        "fmt"
        "net/url"
        "os"
)

func main() {
        sc := bufio.NewScanner(os.Stdin)
        for sc.Scan() {
                u, err := url.Parse(sc.Text())
                if err != nil {
                        //fmt.Fprintf(os.Stderr, "failed to parse url %s [%s]\n", sc.Text(), err)
                }
                m, _ := url.ParseQuery(u.RawQuery)

                for key, _ := range m {
                        fmt.Println(key)
                }
        }
}
