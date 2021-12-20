package controllers

import (
	"bytes"
	"encoding/json"
	"strings"

	"fmt"
	"regexp"

	// "fmt"
	"io/ioutil"

	// "fmt"
	// "io/ioutil"
	"log"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type user_info struct{
	fname string
	lname string
	phone string
	email string
	password string
	bdate string
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	flash := beego.NewFlash()
	c.TplName = "index.tpl"



	fname:=c.GetString("firstname")
	lname:=c.GetString("lastname")
	phone:=c.GetString("phonenumber")
	email:=c.GetString("email")
	password:=c.GetString("password")
	dob:=c.GetString("birthdate")


	rep := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	p:= rep.MatchString(phone)
	red := regexp.MustCompile("((?:19|20)\\d\\d)-(0?[1-9]|1[012])-([12][0-9]|3[01]|0?[1-9])")
	bdate := red.MatchString(dob)
	rem := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	em := rem.MatchString(email)
	
	fmt.Println(p,bdate,em)

	if p && bdate && em {
		jsondata,_:=json.Marshal(map[string]string{
			"FirstName":fname,
			"LastName":lname,
			"Phonenumber":phone,
			"Email":email,
			"Password":password,
			"DateOfBirth":dob,
		})
		// json:=string(jsondata)
		// fmt.Println(json)
		responseBody := bytes.NewBuffer(jsondata)
		// fmt.Println(json)
		response,err:= http.Post("http://localhost:8080/v1/object","application/json",responseBody)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		 }
		 defer response.Body.Close()
	  //Read the response body
		 body, err := ioutil.ReadAll(response.Body)
		 if err != nil {
			log.Fatalln(err)
		 }
		 sb := string(body)
		 nb:=strings.Split(sb,":")
		 msg:=strings.Split(nb[1],"}")
		 log.Printf(msg[0])
		 flash.Error(msg[0])
		 flash.Store(&c.Controller)
	  return
	
	} else {

		if !p {
			
			flash.Error("Phone Number is not valid")
      	 	flash.Store(&c.Controller)
        	return
			
		}
		if !bdate {
			flash.Error("Birthdate is not valid")
      	 	flash.Store(&c.Controller)
        	return
		}
		if !em {
			flash.Error("email is not valid")
			flash.Store(&c.Controller)
		 return
		}
		
		
		}
	
  }
	// fmt.Println(firstname,lastname,password,phonenumber,email,birthdate)

