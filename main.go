package main

import (
	"github.com/ereminiu/autocomplete/pkg/handlers"
	"github.com/ereminiu/autocomplete/pkg/repository"
	"github.com/ereminiu/autocomplete/pkg/service"
	"github.com/ereminiu/autocomplete/trie"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	rep, err := repository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}
	tr := trie.NewTrie()
	srv := service.NewService(rep, tr)
	handler := handlers.NewHandler(srv)

	router := gin.Default()
	router.POST("/search", handler.GetTopFive)
	router.POST("/rebuild", handler.Rebuild)
	router.Run(":6000")
}
