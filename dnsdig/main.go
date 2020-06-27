package main

import (
        "bufio"
        "flag"
        "fmt"
        "net"
        "os"
)

var (
        Info = Green
        Name = Red
)

var (
        Red   = Color("\033[1;31m%s\033[0m")
        Green = Color("\033[1;32m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
        sprint := func(args ...interface{}) string {
                return fmt.Sprintf(colorString,
                        fmt.Sprint(args...))
        }
        return sprint
}

func main() {
        var domains []string

        flag.Parse()
        if flag.NArg() > 0 {

                domains = []string{flag.Arg(0)}

        } else {

                sc := bufio.NewScanner(os.Stdin)
                for sc.Scan() {
                        domains = append(domains, sc.Text())
                }

                if err := sc.Err(); err != nil {
                        fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
                }

        }
        Ip(domains)

}

func Ip(domains []string) {
        for _, fetch := range domains {
                ips, _ := net.LookupIP(fetch)
                fmt.Printf(Name(fetch) + "        ")
                for _, ip := range ips {

                        fmt.Printf(" " + Info(ip.String()))
                }
                fmt.Printf("\n")
        }
}
