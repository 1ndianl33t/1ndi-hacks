package main

import (
        "bufio"
        "flag"
        "fmt"
        "net"
        "os"
)

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
        for _, fetch := range domains {
                cname, _ := net.LookupCNAME(fetch)
                fmt.Printf("%s            ", fetch)
                fmt.Printf("%s\n ", cname)
        }
        fmt.Printf("\n")

}
