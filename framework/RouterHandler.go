package YoyoGo
//
//import (
//	"fmt"
//	"net/http"
//)
//
//type RouterHandler struct {
//
//}
//
//var Router * RouterHandler = new (RouterHandler)
//
//var reqFuncMap = make(map[string]func(w http.ResponseWriter, r *http.Request))
//
//func (rh *RouterHandler) Add(relativePath string, handler func(http.ResponseWriter, *http.Request)){
//	reqFuncMap[relativePath] = handler
//}
//
//
//
//func (p *RouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
//	fmt.Println(r.URL.Path)
//	fun,ok := reqFuncMap[r.URL.Path]
//	if ok{
//		fun(w,r)
//		return
//	}
//}