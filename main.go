package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/KMACEL/gogenux/gomq"
	//import the Paho Go MQTT library
)

func main() {
	var loginData gomq.Login
	loginData.SetClientID = "gopher@go"
	loginData.SetUsername = "gophergo"
	loginData.SetPassword = "12345678"

	gomq.Connect(loginData)
	//go gomq.SendMessage()

	//	errs := syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART2)
	cmd := exec.Command("sh", "-c", "mpg321 media/ay_isiginda.mp3")
	go cmd.Run()

	time.Sleep(10 * time.Second)

	fmt.Println(cmd.Process.Kill())
	/*
		gomq.Connect()
		go gomq.SendMessage()*/

	var pressKey string
	fmt.Scan(&pressKey)
}
