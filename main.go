package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name string
	Gender string
	Age int
}
	


func sayHello(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Println("Parse template failed, err%v", err)
	}
	
	name := "mmm"

	user := User{
		Name : name,
		Gender: "男",
		Age: 23,
	}

	m := map[string]interface{}{
		"name": name,
		"gender": "男",
		"age": 24,
	}

	carList := []string{
		"汽車",
		"火車",
		"貨車",
	}

	err = t.Execute(w, map[string]interface{}{
		"m":		m,
		"user":		user,
		"carList": 	carList,
	})
	if err != nil {
		log.Println("render template failed, err%v", err)
		return 
	}
}

func main(){
	http.HandleFunc("/",sayHello)
	err := http.ListenAndServe(":8099",nil)
	if err != nil {
		log.Println("http server start failed,err:%v", err)
	}
}