package main

import (
  "net/http"
  "os"
  "fmt"
  "log"
  "github.com/gorilla/pat"
  "github.com/unrolled/render"
  "github.com/markbates/goth"
  "github.com/markbates/goth/providers/twitter"
  "github.com/markbates/goth/providers/facebook"
  "github.com/markbates/goth/gothic"
  "github.com/gorilla/sessions"
)

var AppKey = "dsfgdfhgshrstjrjrawertgerhhmmtjr"
var Store = sessions.NewCookieStore([]byte(AppKey))

func main() {
  r := render.New(render.Options{
    Directory: "templates",
    Extensions: []string{".html"},
    IsDevelopment: true,
  })

  goth.UseProviders(
    twitter.New(os.Getenv("TWITTER_K"), os.Getenv("TWITTER_S"), "http://localhost:4000/auth/twitter/callback"),
    facebook.New(os.Getenv("FB_K"), os.Getenv("FB_S"), "http://localhost:4000/auth/facebook/callback"),
  )
  p := pat.New()
  p.Get("/", func(resp http.ResponseWriter, req *http.Request) {
    r.HTML(resp, http.StatusOK, "index", nil)
  })

  // p.Get("/auth/{provider}", func(resp http.ResponseWriter, req *http.Request) {
  //   providerName := req.URL.Query().Get(":provider")
  //   log.Println("Handling a request to provider: ", providerName)
  //   provider, err := goth.GetProvider(providerName)
  //   if err != nil {
  //     fmt.Fprint(resp, err)
  //     return
  //   }

  //   sess, err := provider.BeginAuth("play_state")
  //   if err != nil {
  //     fmt.Fprint(resp, err)
  //     return
  //   }

  //   url, err := sess.GetAuthURL()
  //   if err != nil {
  //     fmt.Fprint(resp, err)
  //     return
  //   }

  //   session, _ := Store.Get(req, "goth-session")
  //   session.Values["goth-session"] = sess.Marshal()
  //   err = session.Save(req, resp)


  //   http.Redirect(resp, req, url, http.StatusFound)
  // })

  p.Get("/auth/{provider}", gothic.BeginAuthHandler)

  p.Get("/auth/{provider}/callback", func(resp http.ResponseWriter, req *http.Request) {
    user, err := gothic.CompleteUserAuth(resp, req)
    if err != nil {
      fmt.Fprint(resp, err)
      log.Println(err)
      return
    }

    r.HTML(resp, http.StatusOK, "user", user)
  })

  http.ListenAndServe(":4000", p)
}
