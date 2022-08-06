package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.telegram.org/bot5407477830:AAHRzUQ1_Yt-TQmhgYfqrTfEe3HWgXMNSJM/sendMessage?chat_id=856546075\&parse_mode=HTML&text=hello") //https://api.telegram.org/bot5407477830:AAHRzUQ1_Yt-TQmhgYfqrTfEe3HWgXMNSJM/sendMessage?chat_id=856546075\\&parse_mode=HTML&text=hello
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}
