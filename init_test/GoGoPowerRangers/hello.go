package main

import (
	"encoding/json"
	"html/template"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
  	Temperature int
  	Humidity int
  	Max int
  	Min int
  	Welcome string
}

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

const(
    api = "13a75dbfde99048a6b499bcd9aca260b"
)
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

    p := &Page{
        Welcome: "Welcome to GoLang!",
    }
    t, _ := template.ParseFiles("index.html")
    t.Execute(w,p)

}
 
func resultsHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method)
	fmt.Println(r.URL)

  	z := r.FormValue("zip")
  	fmt.Println(z)
	

    cast,err := getWeather(z)
    if err != nil{
    	fmt.Println(err) 
    }

    p := &Page{
        Temperature: int(cast.Temp.Temp * 9/5 - 459.67),
        Humidity: int(cast.Temp.Humidity),
        Max: int(cast.Temp.TempMax * 9/5 - 459.67),
        Min: int(cast.Temp.TempMin * 9/5 - 459.67),
    }
    t, _ := template.ParseFiles("results.html")
    t.Execute(w, p)
}

func main() {

	getWeather("10128")

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/results", resultsHandler)

	http.ListenAndServe(":8080", nil)
}

// https://github.com/humbhenri/openweather/blob/master/openweather.go
