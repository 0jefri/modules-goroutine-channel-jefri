package main

import (
	"fmt"

	"github.com/lumoshive-academy/lumo-dev"
	"github.com/modules-goroutin-channel/model"
	"github.com/modules-goroutin-channel/utils"
)

func main() {
	id := utils.GenerateUUID()
	fmt.Println(id)

	text := lumo.NumberToIndonesianText(100)
	fmt.Println(text)

	// newDepedency := depedency.LuasLingkaran()

	userCh := make(chan model.User, 100)
	doneCh := make(chan bool, 100)
	var users []model.User

	go storeData(userCh, &users, doneCh)

	for i := 1; i <= 100; i++ {
		go sendData(userCh, fmt.Sprintf("Nama %d", i), i+20, doneCh)
	}

	for i := 1; i <= 100; i++ {
		<-doneCh
	}

	close(userCh)

	<-doneCh

	fmt.Printf("Total user yang disimpan: %d\n", len(users))
}

func storeData(userCh <-chan model.User, users *[]model.User, doneCh chan<- bool) {
	for user := range userCh {
		*users = append(*users, user)
		fmt.Printf("User %s berhasil disimpan dengan umur %d\n", user.Nama, user.Umur)
	}
	doneCh <- true
}

func sendData(userCh chan<- model.User, nama string, umur int, doneCh chan<- bool) {
	user := model.User{Nama: nama, Umur: umur}
	userCh <- user
	doneCh <- true
}
