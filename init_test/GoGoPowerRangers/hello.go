package main

import (
	"encoding/json"
	"html/template"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
  X string
  Temperature float64
  Gif string
}

const(
    // zip = "10128"
    api = "13a75dbfde99048a6b499bcd9aca260b"
)



// https://github.com/humbhenri/openweather/blob/master/openweather.go




type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Temperature struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

type Forecast struct {
	Weathers []Weather   `json:"weather"`
	Temp     Temperature `json:"main"`
}



func getWeather(zip string)(*Forecast, error){
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%s&APPID=%s", zip, api)
	res, err := http.Get(url)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	
	var f Forecast
	err = json.Unmarshal(body, &f)
	if err != nil {
		return nil, err
	}

	t := f.Temp.Temp * 9/5 - 459.67
	fmt.Println(`The temperature is`, t)
	fmt.Println(f)

	return &f, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    // x := "^^^BITCH^^^^"

	fmt.Println(r.Method)
	fmt.Println(r.URL)

  	z := r.FormValue("zip")
  	fmt.Println(z)
	

    cast,err := getWeather(z)
    if err != nil{
    	fmt.Println(err) 
    }

    p := &Page{
        Temperature: cast.Temp.Temp * 9/5 - 459.67,
    }
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, p)
}
 
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Earth\n")
}

func main() {

	getWeather("10128")

	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}