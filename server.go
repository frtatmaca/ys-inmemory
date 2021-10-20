package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.zfu.fintek.local/zfu/ys-inmemory/handler"
	"gitlab.zfu.fintek.local/zfu/ys-inmemory/service"
)

var (
	fileService     service.FileService     = service.NewFileService()
	inMemoryService service.InMemoryService = service.NewInMemoryService(fileService)
	cronService     service.CronService     = service.NewCronService(inMemoryService)
	inMemoryHandler handler.InMemoryHandler = handler.NewInMemoryHandler(inMemoryService)
	fileHandler     handler.FileHandler     = handler.NewFileHandler()
)

func init() {
	go cronService.StartChrone()
	inMemoryService.SetStoreToMemory()
}

func main() {
	router := gin.Default()

	router.GET("/set-key/:key/:value", inMemoryHandler.SetKey)
	router.GET("/get-key/:key", inMemoryHandler.GetKey)
	router.GET("/get-all-key", inMemoryHandler.GetAllKey)

	router.Run(":3001")
}
