package main
 
import (
    "html/template"
    "fmt"
    "net/http"
)
 

type Page struct {
  Title string
  Body string
  Image string
  X string
}



func indexHandler(w http.ResponseWriter, r *http.Request) {
    // x := "^^^BITCH^^^^"
    p := &Page{
        Body: "^^^BITCH^^^",
        X: "https://media.licdn.com/mpr/mpr/shrinknp_200_200/AAEAAQAAAAAAAAyjAAAAJDY4NjRhMGQ3LTJlYjQtNDZjMS1hMTU4LWRlMDRhZjA5Njc3Yg.jpg",
    }
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, p)
}
 
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Earth\n")
}




 
func main() {
 
    // Calls for function handlers output to match the directory /
    http.HandleFunc("/", indexHandler)
    
    // Calls for function handler2 output to match directory /earth
    http.HandleFunc("/Welcome", welcomeHandler)
    
    // Listen to port 8080 and handle requests
    http.ListenAndServe(":8080", nil)

}