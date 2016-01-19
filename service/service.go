package service

import (
	"github.com/antonmaltsev/cabstore/service/resources"
	"github.com/gin-gonic/gin"
	config "github.com/antonmaltsev/cabstore/cfg"
	log "gopkg.in/inconshreveable/log15.v2"
)

type CabifyService struct {
}

func (s *CabifyService) Run(cfg config.Config) error {
	
	cabifyResource := resources.CabifyResource{}

	log.Info("Starting Cabify Store service");

	cabifyResource.Cfg = cfg

	r := gin.Default()

	// Admin config route
	routes := r.Group("/store")
	{
		routes.POST("/totalprice", (&cabifyResource).OrderSum)
	}

	//GIN release mode, either debug, test or release (DebugMode, TestMode, ReleaseMode)
	//by default, runs in releasemode
	gin.SetMode(gin.DebugMode)
	log.Info("Cabify Store service run in", "mode", gin.Mode())

	r.Run(cfg.SvcHost)

	return nil
}
