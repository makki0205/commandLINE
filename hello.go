// goroutine : 並行処理
// channel

package main

import (
    "fmt"
    // "time"
    "os"
    "os/exec"
    "net"
    // "github.com/morikuni/aec"
)

func main() {
    go listen();
    sned();
}
func sned(){
    var msg string
    for{
        fmt.Scanf("%s",&msg)
        msg = msg+"\n"
        conn, err := net.Dial("tcp", "localhost:"+os.Args[2])
        if err != nil {
            fmt.Printf("Dial error: %s\n", err)
            return
        }
        defer conn.Close()
        conn.Write([]byte(msg))
        conn.Close()
    }
}
func listen(){
    for{
        out,err := exec.Command("nc","-l",os.Args[1]).Output()
        if err !=nil {
            fmt.Println(err)
        }else{
            fmt.Printf("%s",out)
        }
    }
}