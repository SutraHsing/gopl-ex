package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	for _, port := range []string{"8010", "8020", "8030"} {
		conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
		if err != nil {
			continue
		}
		defer conn.Close()
		mustCopy(os.Stdout, conn)
		io.WriteString(os.Stdout, "\t")
	}
}

func mustCopy(writer io.Writer, reader io.Reader) {
	if _, err := io.Copy(writer, reader); err != nil {
		log.Fatalln(err)
	}
}
