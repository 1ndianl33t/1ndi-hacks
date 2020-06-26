package main

import (
        "bufio"
        "flag"
        "fmt"
    //  "github.com/fatih/color"
        "net"
        "os"
)

func main() {
        var domains []string

        flag.Parse()
        if flag.NArg() > 0 {

                domains = []string{flag.Arg(0)}
                 /// fmt.Println("Got single")
        } else {

               
                sc := bufio.NewScanner(os.Stdin)
                for sc.Scan() {
                        domains = append(domains, sc.Text())
                }

                if err := sc.Err(); err != nil {
                        fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
                }

        }
        for _, fetch := range domains {
                ips, _ := net.LookupIP(fetch)
                // color.HiRed("%s   ", fetch)
                fmt.Printf("%s  ", fetch)
                for _, ip := range ips {
                        //color.HiGreen("%s,", ip.String())
                        fmt.Printf("%s,", ip.String())
                }
                fmt.Printf("\n")
        }

}
