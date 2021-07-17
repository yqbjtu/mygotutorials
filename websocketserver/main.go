package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"k8s.io/klog"
)

var addr = flag.String("addr", "localhost:19000", "http service address")

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
	//klog.InitFlages(nil)
	flag.Parse()

	klog.Info("start to listen")
	http.HandleFunc("/api/v1/echo", EchoMessage)
	klog.Fatal(http.ListenAndServe(*addr, nil))
}


func EchoMessage(w http.ResponseWriter, r *http.Request) {
    conn, _ := upgrader.Upgrade(w, r, nil) 

    for {
        // 读取客户端的消息
        msgType, msg, err := conn.ReadMessage()
        if err != nil {
            return
        }

        // 把消息打印到标准输出
        fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

        // 把消息写回客户端，完成响应，实际中会在处理对应操作返回请求的结果
        if err = conn.WriteMessage(msgType, msg); err != nil {
            return
        }
    }
}