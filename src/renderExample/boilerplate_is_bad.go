package main

// import (
//   "net/http"
//   "encoding/json"
//   "io/ioutil"
//   "fmt"
//   "text/template"
// )
// type User struct {
//   Name string `json:"name" xml:"name"`
//   Email string `json:"email" xml:"email"`
// }

// func main() {
//   http.HandleFunc("/json", func(res http.ResponseWriter, req *http.Request) {
//     user := userFromReq(req)
//     bytes, _ := json.Marshal(user)
//     res.Header().Add("Content-Type", "application/json")
//     res.WriteHeader(200)
//     res.Write(bytes)
//   })

//   http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
//     user := userFromReq(req)
//     body, _ := ioutil.ReadFile("/templates/index.html")
//     tmpl, _ := template.New("Some Name").Parse(string(body))
//     tmpl.Execute(res, user)
//     res.Header().Add("Content-Type", "text/html")
//     res.WriteHeader(200)
//   })

//   http.ListenAndServe(":4000", nil)
// }

// func userFromReq(req *http.Request) *User {
//   name := req.URL.Query().Get("name")
//   user := &User{
//     Name: name,
//     Email: fmt.Sprintf("%s@example.com", name),
//   }

//   return user
// }
