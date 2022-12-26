package main

import (
	"part2/internal/adaptor"
	"part2/internal/consts"
	"part2/internal/model"
	"part2/internal/service"
	"time"
)

func main() {
	date, _ := time.Parse(consts.LAYOUT, "2000-04-28")
	date2 := time.Now()
	date2.Date()
	u := model.NewUser("alex", "alexlogin", "qwe", "88005553535", date)
	//fmt.Print(u)
	storage := adaptor.CreateConcurrentUserStorage()
	storage.Add(u)
	router := service.SetupRouter(*storage)

	router.Run("localhost:8080")
}
