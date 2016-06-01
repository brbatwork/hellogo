package main

import (
  "os"
  "log"
  "io/ioutil"
  txt_template "text/template"
  "html/template"
)
type User struct {
  FirstName   string
  LastName    string
  Email       string
  Age         int
}

func (u User) IsOld() bool {
  return u.Age > 30
}

func main() {
  u := User{"Bill", "Jones", "bill.jones@example.com", 37}
  files := []string {
    "/Users/Thunderbird/bbarbour/Dropbox/projects/HelloGo/templates/user.tpl",
    "/Users/Thunderbird/bbarbour/Dropbox/projects/HelloGo/templates/user.html",
    "/Users/Thunderbird/bbarbour/Dropbox/projects/HelloGo/templates/partialTemplate.html",
  }

  body, err := ioutil.ReadFile(files[0])
  if err != nil {
    log.Fatal(err)
  }

  tmpl, err := txt_template.New("MyTemplate").Parse(string(body))
  if err != nil {
    log.Fatal(err)
  }

  tmpl.Execute(os.Stdout, u)

  // HTML Example
  body, err = ioutil.ReadFile(files[1])
  if err != nil {
    log.Fatal(err)
  }

  htmlTmpl, err := template.New("HtmlTemplate").Parse(string(body))
  if err != nil {
    log.Fatal(err)
  }

  // htmlTmpl.Execute(os.Stdout, u)

  // Using partial example
  htmlTmpl, _ = template.ParseFiles(files[1], files[2])
  err = htmlTmpl.ExecuteTemplate(os.Stdout, "partialTemplate.html", u)
  if err != nil {
    log.Fatal(err)
  }
}
