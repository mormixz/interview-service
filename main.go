package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"mormixz.com/interview-service/interview"
	"mormixz.com/interview-service/models"
	"mormixz.com/interview-service/store"
)

var (
	config     = Config{}
	configPath = flag.String("configPath", ".", "-configPath=.")
	configName = flag.String("configName", "config", "-configName=config")
)

type Config struct {
	HttpPort  string `mapstructure:"http_port"`
	AllowCORS string `mapstructure:"allow_cors"`
	MongoDB   struct {
		ConnectionString  string `mapstructure:"connection_string"`
		ConnectionTimeout int    `mapstructure:"connection_timeout"`
		DBName            string `mapstructure:"db_name"`
	} `mapstructure:"mongo_db"`
}

type handler struct {
	interviewService *interview.Service
}

func initConfig(path, name string) {
	viper.SetConfigName(name)
	viper.SetConfigType("json")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err.Error())
	}
	log.Println("Read Config From -> " + viper.ConfigFileUsed())

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unmarshal config error %s", err.Error())
	}
}

func initHandler() *handler {
	dbStore, err := store.NewStore(config.MongoDB.ConnectionString, config.MongoDB.DBName, config.MongoDB.ConnectionTimeout)
	if err != nil {
		log.Fatalf("Failed to create database store: %s", err.Error())
	}

	return &handler{
		interview.NewService(dbStore),
	}
}

func initRoutes(app *fiber.App, h *handler) {
	interviewGroup := app.Group("/interview")
	interviewGroup.Get("/all", h.GetInterviewAll)
	interviewGroup.Get("/:id", h.GetInterview)
	interviewGroup.Put("/:id/comment", h.CommentInterview)
	interviewGroup.Put("/:id/update", h.UpdateInterview)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON("api path not found")
	})
}

func main() {
	flag.Parse()

	// Read Configutaion Using Viper
	initConfig(*configPath, *configName)

	app := fiber.New(fiber.Config{
		ServerHeader: "Interview Service", // add custom server header
		Immutable:    true,
	})

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: config.AllowCORS,
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "",
	}))

	initRoutes(app, initHandler())

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.HttpPort)))
}

func (h *handler) GetInterview(c *fiber.Ctx) error {
	interviewID := c.Params("id")

	interview, err := h.interviewService.GetInterview(interviewID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return c.Status(http.StatusNotFound).JSON(fmt.Sprintf("%s not found", interviewID))
		default:
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
	}

	return c.Status(http.StatusOK).JSON(interview)
}

func (h *handler) GetInterviewAll(c *fiber.Ctx) error {
	status := c.Query("status")
	limit := c.Query("limit")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	interviews, err := h.interviewService.GetInterviewAll(status, limitInt)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(interviews)
}

func (h *handler) CommentInterview(c *fiber.Ctx) error {
	interviewID := c.Params("id")

	comment := models.Comment{}
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	if err := h.interviewService.UpdateCommentInterview(interviewID, comment); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return c.Status(http.StatusNotFound).JSON(fmt.Sprintf("%s not found", interviewID))
		default:
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
	}

	return c.Status(http.StatusOK).JSON("success")
}

func (h *handler) UpdateInterview(c *fiber.Ctx) error {
	interviewID := c.Params("id")

	updateInterview := &models.Interview{}
	if err := c.BodyParser(&updateInterview); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	if err := h.interviewService.UpdateInterview(interviewID, updateInterview); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return c.Status(http.StatusNotFound).JSON(fmt.Sprintf("%s not found", interviewID))
		default:
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
	}

	return c.Status(http.StatusOK).JSON("success")
}
