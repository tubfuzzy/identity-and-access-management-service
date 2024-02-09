package server

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	fiberCache "github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/utils"

	"authentication-and-authorization-service/config"
	apiv1 "authentication-and-authorization-service/internal/app"
	"authentication-and-authorization-service/pkg/cache"
	"authentication-and-authorization-service/pkg/common/exception"
	loggerPkg "authentication-and-authorization-service/pkg/logger"
)

type Server struct {
	app    *fiber.App
	conf   *config.Configuration
	logger loggerPkg.Logger
}

func New() (*Server, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	logger := loggerPkg.NewLogger(conf)
	redisCache, err := cache.NewRedisCache(conf)
	if err != nil {
		return nil, err
	}
	app := NewFiberApp(conf, logger, redisCache)

	return &Server{app, conf, logger}, nil
}

func NewFiberApp(
	conf *config.Configuration,
	logger loggerPkg.Logger,
	cacheEngine cache.Engine,
) *fiber.App {

	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
		ReadTimeout:  time.Second * conf.Server.ReadTimeout,
		WriteTimeout: time.Second * conf.Server.WriteTimeout,
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
	})

	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(recover.New())

	// fiber log
	app.Use(fiberLog.New(fiberLog.Config{
		Next:         nil,
		Done:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stdout,
	}))

	// fiber cache
	app.Use(fiberCache.New(fiberCache.Config{
		Next: func(c *fiber.Ctx) bool {
			if c.Query("refresh") == "true" {
				go cacheEngine.Delete(utils.CopyString(c.Path()) + "_" + c.Method())
				go cacheEngine.Delete(utils.CopyString(c.Path()) + "_" + c.Method() + "_body")

				return true
			}

			return false
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.Path())
		},
		Expiration:   1 * time.Minute,
		CacheControl: true,
		Methods:      []string{fiber.MethodGet},
		Storage:      cacheEngine,
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	apiv1.NewApplication(v1, logger)

	app.Use(func(c *fiber.Ctx) error {
		panic(exception.NotFoundError{Message: "path " + c.Path() + " does not exist."})
	})

	return app
}

func (serv Server) App() *fiber.App {
	return serv.app
}

func (serv Server) Config() *config.Configuration {
	return serv.conf
}

func (serv Server) Logger() loggerPkg.Logger {
	return serv.logger
}
