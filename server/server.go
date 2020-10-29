package main

//필요 패키지 임포트
import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"os"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()

	r.ParseForm()
	fmt.Fprintln(w, "Hostname: " + hostname)
	//Get 파라미터 및 정보 출력
	fmt.Fprintln(w, "default : ", r.Form)
	fmt.Fprintln(w, "path", r.URL.Path)
	fmt.Fprintln(w, "param : ", r.Form["test_param"])
	//Parameter 전체 출력
	for k, v := range r.Form {
		fmt.Fprintln(w, k, " - ", strings.Join(v, ""))
	}
	//기본 출력
	fmt.Fprintln(w, "Golang WebServer Working!")
}

func main() {
	//기본 Url 핸들러 메소드 지정
	http.HandleFunc("/", defaultHandler)
	//서버 시작
	err := http.ListenAndServe(":5000", nil)
	//예외 처리
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Println("ListenAndServe Started! -> Port(5000)")
	}
}
