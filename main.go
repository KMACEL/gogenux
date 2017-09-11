package main

import (
	"fmt"

	"github.com/KMACEL/gogenux/gomq"
	//import the Paho Go MQTT library
)

func main() {

	var loginData gomq.Login
	loginData.SetClientID = "gopher@go"
	loginData.SetUsername = "gophergo"
	loginData.SetPassword = "12345678"

	gomq.Connect(loginData)

	go gomq.SendMessage()

	/*
		gomq.Connect()
		go gomq.SendMessage()*/

	var pressKey string
	fmt.Scan(&pressKey)
}
