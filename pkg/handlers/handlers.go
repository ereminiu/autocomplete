package handlers

import (
	"github.com/ereminiu/autocomplete/models"
	"github.com/ereminiu/autocomplete/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
)

type Handler struct {
	service *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{service: srv}
}

func (h *Handler) GetTopFive(ctx *gin.Context) {
	prefix := ctx.Query("q")
	addFlag := ctx.Query("add")
	words := h.service.GetTopFive(prefix)
	if addFlag == "true" || addFlag == "t" {
		h.service.AddWord(prefix)
		err := h.service.AddRecord(models.Record{Query: prefix, Freq: 1})
		if err != nil {
			log.Fatal(err)
		}
	}
	ctx.JSON(200, gin.H{
		"requests": words,
	})
}

func (h *Handler) Rebuild(ctx *gin.Context) {
	records, err := h.service.Repository.FetchRecords()
	if err != nil {
		log.Fatal(err)
	}
	h.service.Autocomplete.Rebuild(records)
	ctx.JSON(200, gin.H{
		"message": "Trie has been reinited from database",
	})
}
