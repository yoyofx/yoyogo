package YoyoGo

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	// DefaultAddress is used if no other is specified.
	DefaultAddress = ":8080"
)


type YoyoGo struct {

}

func Classic() *YoyoGo{
	return &YoyoGo{}
}

type M = map[string]string


var reqFuncMap = make(map[string]func(ctx *HttpContext))

func (rh *YoyoGo) Map(relativePath string, handler func(ctx *HttpContext)){
	reqFuncMap[relativePath] = handler
}

func (n *YoyoGo) Run(addr ...string) {
	l := log.New(os.Stdout, "[yoyofx] ", 0)
	finalAddr := detectAddress(addr...)
	l.Printf("listening on %s", finalAddr)
	l.Fatal(http.ListenAndServe(finalAddr, n))
}

func detectAddress(addr ...string) string {
	if len(addr) > 0 {
		return addr[0]
	}
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return DefaultAddress
}





func (p *YoyoGo) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.Path)
	ctx := NewContext(w,r)
	fun,ok := reqFuncMap[r.URL.Path]
	if ok{
		fun(ctx)
		return
	}
}


