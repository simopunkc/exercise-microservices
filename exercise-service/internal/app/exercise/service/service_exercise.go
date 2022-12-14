package service

import (
	"exercise-service/internal/app/domain"
	"log"
	"strings"
)

//go:generate moq -out service_exercise_mock_test.go . RepositoryExercise
type RepositoryExercise interface {
	GetExerciseByID(id int64) domain.Repository
	GetPublicExerciseByID(id int64) domain.Repository
	GetAnswerByUserID(userID int64) domain.Repository
	CreateExercise(exercise domain.Exercise) domain.Repository
	CreateQuestion(question domain.Question) domain.Repository
	CreateAnswer(answer domain.Answer) domain.Repository
}

type ServiceExercise struct {
	repositoryExercise RepositoryExercise
}

func NewServiceExercise(repositoryExercise RepositoryExercise) *ServiceExercise {
	return &ServiceExercise{
		repositoryExercise: repositoryExercise,
	}
}

func (se ServiceExercise) GetExerciseByID(id int64) domain.Service {
	exercise := se.repositoryExercise.GetPublicExerciseByID(id)
	if exercise.Error != nil {
		log.Println(exercise.Error)
		return domain.Service{
			Hash:  "GMvw4Dd7CQ6x",
			Error: exercise.Error,
		}
	}

	return domain.Service{
		Hash:     "GMlJLX5vZauq",
		Exercise: exercise.Exercise,
	}
}

func (se ServiceExercise) CalculateUserScore(id int64, userID int64) domain.Service {
	exercise := se.repositoryExercise.GetExerciseByID(id)
	if exercise.Error != nil {
		log.Println(exercise.Error)
		return domain.Service{
			Hash:  "GMcWGOvA4vWh",
			Error: exercise.Error,
		}
	}

	answers := se.repositoryExercise.GetAnswerByUserID(userID)
	if answers.Error != nil {
		log.Println(exercise.Error)
		return domain.Service{
			Hash:  "GMoIVFp6nc8C",
			Error: exercise.Error,
		}
	}

	if len(answers.ListAnswer) == 0 {
		return domain.Service{
			Hash: "GMYQmpAUyq7P",
			Score: domain.Score{
				Score: 0,
			},
		}
	}

	mapQuestion := make(map[int64]domain.Question)
	for _, question := range exercise.Exercise.Questions {
		mapQuestion[question.ID] = question
	}

	var score int
	for _, answer := range answers.ListAnswer {
		if strings.EqualFold(answer.Answer, mapQuestion[answer.QuestionID].CorrectAnswer) {
			score += mapQuestion[answer.QuestionID].Score
		}
	}

	return domain.Service{
		Hash: "GMnozAj8EQ97",
		Score: domain.Score{
			Score: score,
		},
	}
}

func (se ServiceExercise) CreateExercise(exercise domain.Exercise) domain.Service {
	repo := se.repositoryExercise.CreateExercise(exercise)
	if repo.Error != nil {
		log.Println(repo.Error)
		return domain.Service{
			Hash:  "GM9jA7ylQP48",
			Error: repo.Error,
		}
	}
	return domain.Service{
		Hash:        "GMDVvaZ47LnR",
		RawResponse: "exercise created successfully",
	}
}

func (se ServiceExercise) CheckIsInvalidAnswer(answer string) bool {
	validAnswer := []string{"a", "b", "c", "d"}
	invalidAnswer := true
	lowerCaseAnswer := strings.ToLower(answer)
	for _, v := range validAnswer {
		if v == lowerCaseAnswer {
			invalidAnswer = false
			break
		}
	}
	return invalidAnswer
}

func (se ServiceExercise) CreateQuestion(question domain.Question) domain.Service {
	repo := se.repositoryExercise.CreateQuestion(question)
	if repo.Error != nil {
		log.Println(repo.Error)
		return domain.Service{
			Hash:  "GMBaOXo2oGjO",
			Error: repo.Error,
		}
	}

	return domain.Service{
		Hash:        "GMCbzxwxlNYo",
		RawResponse: "question created successfully",
	}
}

func (se ServiceExercise) CreateAnswer(answer domain.Answer) domain.Service {
	repo := se.repositoryExercise.CreateAnswer(answer)
	if repo.Error != nil {
		log.Println(repo.Error)
		return domain.Service{
			Hash:  "GMOAzxBj08Ui",
			Error: repo.Error,
		}
	}

	return domain.Service{
		Hash:        "GM0qIcIzIBA9",
		RawResponse: "answer created successfully",
	}
}
