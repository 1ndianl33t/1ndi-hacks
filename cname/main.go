package main

import (                                                                                 "fmt"
        "net"
        "bufio"
        "os"
)

func main() {
         sc := bufio.NewScanner(os.Stdin)
        for sc.Scan() {
        domain := sc.Text()
        cname, _ := net.LookupCNAME(domain)
        fmt.Println(cname)
 }
}
