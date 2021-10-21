package main

import (
	"github.com/frtatmaca/ys-inmemory/core"
	"github.com/frtatmaca/ys-inmemory/handler"
	"github.com/frtatmaca/ys-inmemory/service"
	"github.com/gin-gonic/gin"
)

var (
	repositoryAdapter core.RepositoryAdapter  = core.NewRepositoryAdapter()
	fileService       service.FileService     = service.NewFileService(repositoryAdapter)
	inMemoryService   service.InMemoryService = service.NewInMemoryService(fileService)
	cronService       service.CronService     = service.NewCronService(inMemoryService)
	inMemoryHandler   handler.InMemoryHandler = handler.NewInMemoryHandler(inMemoryService)
	fileHandler       handler.FileHandler     = handler.NewFileHandler()
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
