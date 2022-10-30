package handler

import (
	"exercise-service/internal/app/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ServiceExercise interface {
	GetExerciseByID(id int64) domain.Service
	CalculateUserScore(id int64, userID int64) domain.Service
	CreateExercise(exercise domain.Exercise) domain.Service
	CheckIsInvalidAnswer(answer string) bool
	CreateQuestion(question domain.Question) domain.Service
	CreateAnswer(answer domain.Answer) domain.Service
}

type HandlerExercise struct {
	serviceExercise ServiceExercise
}

func NewHandlerExercise(serviceExercise ServiceExercise) *HandlerExercise {
	return &HandlerExercise{serviceExercise}
}

func (he HandlerExercise) GetExerciseByID(c *fiber.Ctx) error {
	paramID := c.Params("id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMYJD6mHdxlA",
			Error: "invalid exercise id",
		})
	}

	exercise := he.serviceExercise.GetExerciseByID(id)
	if exercise.Error != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMuc25aB5HpD",
			Error: "exercise not found",
		})
	}

	return c.Status(200).JSON(domain.Handler{
		Hash:     "GMDkZLh5QJyp",
		Exercise: &exercise.Exercise,
	})
}

func (he HandlerExercise) CalculateUserScore(c *fiber.Ctx) error {
	paramID := c.Params("id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMk6LNXGf9An",
			Error: "invalid exercise id",
		})
	}

	exercise := he.serviceExercise.GetExerciseByID(id)
	if exercise.Error != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMybk5dxHUjp",
			Error: "exercise not found",
		})
	}

	var userID int64 = c.Locals("user_id").(int64)
	answers := he.serviceExercise.CalculateUserScore(id, userID)
	if answers.Error != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMQsb16uYOtU",
			Error: "error when find answers",
		})
	}

	return c.Status(200).JSON(domain.Handler{
		Hash:  "GMyTdADK3jo8",
		Score: &answers.Score,
	})
}

func (he HandlerExercise) CreateExercise(c *fiber.Ctx) error {
	var exercise domain.Exercise
	err := c.BodyParser(&exercise)
	if err != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMZWgqrhx3Ik",
			Error: "invalid input",
		})
	}

	if exercise.Title == "" {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMwLp1hhi581",
			Error: "field title must required",
		})
	}

	if exercise.Description == "" {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GM2x7de7byvF",
			Error: "field description must required",
		})
	}

	service := he.serviceExercise.CreateExercise(exercise)
	if service.Error != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMDBaAcnswBs",
			Error: "failed when create exercise",
		})
	}

	return c.Status(200).JSON(domain.Handler{
		Hash:        "GM47oH52hgKL",
		RawResponse: service.RawResponse,
	})
}

func (he HandlerExercise) CreateQuestion(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMRYox3Mnnbv",
			Error: "invalid exercise id",
		})
	}

	var userID int64 = c.Locals("user_id").(int64)
	var question domain.Question
	question.ExerciseID = id
	question.CreatorID = userID

	err = c.BodyParser(&question)
	if err != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMAxt9cGkA6W",
			Error: "invalid input",
		})
	}

	if question.Body == "" {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMpUACBAKchn",
			Error: "field body must required",
		})
	}

	if question.OptionA == "" {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMsLC5XBPknt",
			Error: "field option_a must required",
		})
	}

	if question.OptionB == "" {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMHr6O8tx0wt",
			Error: "field option_b must required",
		})
	}

	if question.OptionC == "" {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GM67SljLz2Ih",
			Error: "field option_c must required",
		})
	}

	if question.OptionD == "" {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMhLQuldlNvG",
			Error: "field option_d must required",
		})
	}

	if he.serviceExercise.CheckIsInvalidAnswer(question.CorrectAnswer) {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMQgy1xSIFPO",
			Error: "field correct_answer must required",
		})
	}

	if question.Score == 0 {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMprC9t5wiOh",
			Error: "field score must required",
		})
	}

	service := he.serviceExercise.CreateQuestion(question)
	if service.Error != nil {
		return c.Status(500).JSON(domain.Handler{
			Hash:  "GMw3jNyB4w8F",
			Error: "failed when create question",
		})
	}

	return c.Status(200).JSON(domain.Handler{
		Hash:        "GMKL4NFQnyZT",
		RawResponse: service.RawResponse,
	})
}

func (he HandlerExercise) CreateAnswer(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMHedxuc4hUF",
			Error: "invalid exercise id",
		})
	}

	var answer domain.Answer
	answer.ExerciseID = id
	paramIdQuestion := c.Params("qid")
	qid, err := strconv.ParseInt(paramIdQuestion, 10, 64)
	if err != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMe5qvcCFb9r",
			Error: "invalid question id",
		})
	}

	var userID int64 = c.Locals("user_id").(int64)
	answer.QuestionID = qid
	answer.UserID = userID
	err = c.BodyParser(&answer)
	if err != nil {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMV5AmOUQezB",
			Error: "invalid input",
		})
	}

	if he.serviceExercise.CheckIsInvalidAnswer(answer.Answer) {
		return c.Status(400).JSON(domain.Handler{
			Hash:  "GMktQCmK1lZj",
			Error: "field answer must required",
		})
	}

	service := he.serviceExercise.CreateAnswer(answer)
	if service.Error != nil {
		return c.Status(500).JSON(domain.Handler{
			Hash:  "GMiQUgFpkQCN",
			Error: "failed when create answer",
		})
	}

	return c.Status(200).JSON(domain.Handler{
		Hash:        "GMOVL8vFWr0t",
		RawResponse: service.RawResponse,
	})
}
