package main
 
import (
	"fmt"
        "net"
        "bufio"
        "os"
)

func main() {
         sc := bufio.NewScanner(os.Stdin)
        for sc.Scan() {
        domain := sc.Text()
	txtrecords, _ := net.LookupTXT(domain)
 
	for _, txt := range txtrecords {
		fmt.Println(txt)
	}
 }
}	
