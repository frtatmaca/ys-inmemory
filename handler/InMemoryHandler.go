package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/frtatmaca/ys-inmemory/service"
	"net/http"
)

type InMemoryHandler struct {
	services InMemoryService
}

func NewInMemoryHandler(inMemoryService InMemoryService) InMemoryHandler {
	return InMemoryHandler{inMemoryService}
}

func (p *InMemoryHandler) SetKey(ctx *gin.Context) {
	key := ctx.Param("key")
	value := ctx.Param("value")

	message := p.services.SetKey(key, value)
	ctx.String(http.StatusOK, message)
}

func (p *InMemoryHandler) GetKey(ctx *gin.Context) {
	key := ctx.Param("key")
	val := p.services.GetKey(key)
	ctx.String(http.StatusOK, val)
}

func (p *InMemoryHandler) GetAllKey(ctx *gin.Context) {
	ctx.String(http.StatusOK, fmt.Sprintf("%#v", p.services.GetAllKey()))
}
