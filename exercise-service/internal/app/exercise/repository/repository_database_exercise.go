package repository

import (
	"exercise-service/internal/app/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RepositoryDatabaseExercise struct {
	db *gorm.DB
}

func NewRepositoryDatabaseExercise(db *gorm.DB) *RepositoryDatabaseExercise {
	return &RepositoryDatabaseExercise{
		db: db,
	}
}

func (rde RepositoryDatabaseExercise) GetExerciseByID(id int64) domain.Repository {
	var exercise domain.Exercise
	err := rde.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		return domain.Repository{
			Error: err,
		}
	}

	return domain.Repository{
		Exercise: exercise,
	}
}

func (rde RepositoryDatabaseExercise) GetPublicExerciseByID(id int64) domain.Repository {
	var exercise domain.Exercise
	err := rde.db.Debug().Table("exercises").Where("id = ?", id).First(&exercise).Error
	if err != nil {
		return domain.Repository{
			Error: err,
		}
	}

	var publicQuestions []domain.PublicQuestion
	err = rde.db.Debug().Table("questions").Where("exercise_id = ?", id).Select("id, body, option_a, option_b, option_c, option_d, score, created_at, updated_at").Find(&publicQuestions).Error
	if err != nil {
		return domain.Repository{
			Error: err,
		}
	}

	questions := make([]domain.Question, 0)
	for _, publicQuestion := range publicQuestions {
		questions = append(questions, domain.Question{
			ID:        publicQuestion.ID,
			Body:      publicQuestion.Body,
			OptionA:   publicQuestion.OptionA,
			OptionB:   publicQuestion.OptionB,
			OptionC:   publicQuestion.OptionC,
			OptionD:   publicQuestion.OptionD,
			Score:     publicQuestion.Score,
			CreatedAt: publicQuestion.CreatedAt,
			UpdatedAt: publicQuestion.UpdatedAt,
		})
	}
	exercise.Questions = questions

	return domain.Repository{
		Exercise: exercise,
	}
}

func (rde RepositoryDatabaseExercise) GetAnswerByUserID(userID int64) domain.Repository {
	var answers []domain.Answer
	err := rde.db.Where("user_id = ?", userID).Find(&answers).Error
	if err != nil {
		return domain.Repository{
			Error: err,
		}
	}
	return domain.Repository{
		ListAnswer: answers,
	}
}

func (rde RepositoryDatabaseExercise) CreateExercise(exercise domain.Exercise) domain.Repository {
	err := rde.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&exercise).Error
	return domain.Repository{
		Error: err,
	}
}

func (rde RepositoryDatabaseExercise) CreateQuestion(question domain.Question) domain.Repository {
	err := rde.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&question).Error
	return domain.Repository{
		Error: err,
	}
}

func (rde RepositoryDatabaseExercise) CreateAnswer(answer domain.Answer) domain.Repository {
	err := rde.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&answer).Error
	return domain.Repository{
		Error: err,
	}
}
