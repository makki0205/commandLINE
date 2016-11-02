// goroutine : 並行処理
// channel

package main

import (
    "fmt"
    "os"
    "os/exec"
    "net"
    "./myfunc"
)

func main() {
    myfunc.Cls();
    go listen();
    sned();
}
func sned(){
    var msg string
    for{
        myfunc.Locate(30,0)
        fmt.Printf("                                                               ")
        myfunc.Locate(30,0)
        fmt.Printf(">")
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
    msgList := make([]string,1)
    var msg string
    for{
        out,err := exec.Command("nc","-l",os.Args[1]).Output()
        if err !=nil {
            fmt.Println(err)
        }else{
            msg = string(out)
            msgList = append(msgList,msg)
            myfunc.Locate(0,0)
            myfunc.ArrayPrint(msgList)
            myfunc.Locate(30,2)

        }
    }
}