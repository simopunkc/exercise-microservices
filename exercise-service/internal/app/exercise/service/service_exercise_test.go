package service

import (
	"errors"
	"exercise-service/internal/app/domain"
	"reflect"
	"testing"
)

func TestServiceExercise_GetExerciseByID(t *testing.T) {
	type fields struct {
		repositoryExercise RepositoryExercise
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "case exercise is not exists",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					GetPublicExerciseByIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							Error: errors.New("exercise is not exists"),
						}
					},
				},
			},
			args: args{
				id: 0,
			},
			want: "GMvw4Dd7CQ6x",
		},
		{
			name: "case no error",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					GetPublicExerciseByIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							Exercise: domain.Exercise{
								ID: 1,
							},
						}
					},
				},
			},
			args: args{
				id: 1,
			},
			want: "GMlJLX5vZauq",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := NewServiceExercise(tt.fields.repositoryExercise)
			if got := se.GetExerciseByID(tt.args.id); !reflect.DeepEqual(got.Hash, tt.want) {
				t.Errorf("ServiceExercise.GetExerciseByID() = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}

func TestServiceExercise_CalculateUserScore(t *testing.T) {
	type fields struct {
		repositoryExercise RepositoryExercise
	}
	type args struct {
		id     int64
		userID int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantHash  string
		wantScore int
	}{
		{
			name: "case exercise is not exists",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					GetExerciseByIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							Error: errors.New("exercise is not exists"),
						}
					},
				},
			},
			args: args{
				id:     0,
				userID: 0,
			},
			wantHash:  "GMcWGOvA4vWh",
			wantScore: 0,
		},
		{
			name: "case failed get answer user from database",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					GetExerciseByIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							Exercise: domain.Exercise{
								ID: 1,
							},
						}
					},
					GetAnswerByUserIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							Error: errors.New("failed get answer user from database"),
						}
					},
				},
			},
			args: args{
				id:     1,
				userID: 0,
			},
			wantHash:  "GMoIVFp6nc8C",
			wantScore: 0,
		},
		{
			name: "case list answer user from database is empty",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					GetExerciseByIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							Exercise: domain.Exercise{
								ID: 1,
							},
						}
					},
					GetAnswerByUserIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							ListAnswer: []domain.Answer{},
						}
					},
				},
			},
			args: args{
				id:     1,
				userID: 4,
			},
			wantHash:  "GMYQmpAUyq7P",
			wantScore: 0,
		},
		{
			name: "case list question exercise from database is empty",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					GetExerciseByIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							Exercise: domain.Exercise{
								ID:        1,
								Questions: []domain.Question{},
							},
						}
					},
					GetAnswerByUserIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							ListAnswer: []domain.Answer{
								{
									ID:     11,
									Answer: "a",
								},
								{
									ID:     12,
									Answer: "b",
								},
							},
						}
					},
				},
			},
			args: args{
				id:     1,
				userID: 4,
			},
			wantHash:  "GMnozAj8EQ97",
			wantScore: 0,
		},
		{
			name: "case 0 correct answer from 2 question",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					GetExerciseByIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							Exercise: domain.Exercise{
								ID: 1,
								Questions: []domain.Question{
									{
										ID:            11,
										CorrectAnswer: "a",
										Score:         10,
									},
									{
										ID:            12,
										CorrectAnswer: "c",
										Score:         10,
									},
								},
							},
						}
					},
					GetAnswerByUserIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							ListAnswer: []domain.Answer{
								{
									ID:         11,
									ExerciseID: 1,
									QuestionID: 11,
									Answer:     "b",
								},
								{
									ID:         12,
									ExerciseID: 1,
									QuestionID: 12,
									Answer:     "b",
								},
							},
						}
					},
				},
			},
			args: args{
				id:     1,
				userID: 4,
			},
			wantHash:  "GMnozAj8EQ97",
			wantScore: 0,
		},
		{
			name: "case 1 correct answer from 2 question",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					GetExerciseByIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							Exercise: domain.Exercise{
								ID: 1,
								Questions: []domain.Question{
									{
										ID:            11,
										CorrectAnswer: "a",
										Score:         10,
									},
									{
										ID:            12,
										CorrectAnswer: "c",
										Score:         10,
									},
								},
							},
						}
					},
					GetAnswerByUserIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							ListAnswer: []domain.Answer{
								{
									ID:         11,
									ExerciseID: 1,
									QuestionID: 11,
									Answer:     "a",
								},
								{
									ID:         12,
									ExerciseID: 1,
									QuestionID: 12,
									Answer:     "b",
								},
							},
						}
					},
				},
			},
			args: args{
				id:     1,
				userID: 4,
			},
			wantHash:  "GMnozAj8EQ97",
			wantScore: 10,
		},
		{
			name: "case 2 correct answer from 2 question",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					GetExerciseByIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							Exercise: domain.Exercise{
								ID: 1,
								Questions: []domain.Question{
									{
										ID:            11,
										CorrectAnswer: "a",
										Score:         10,
									},
									{
										ID:            12,
										CorrectAnswer: "c",
										Score:         10,
									},
								},
							},
						}
					},
					GetAnswerByUserIDFunc: func(id int64) domain.Repository {
						return domain.Repository{
							ListAnswer: []domain.Answer{
								{
									ID:         11,
									ExerciseID: 1,
									QuestionID: 11,
									Answer:     "a",
								},
								{
									ID:         12,
									ExerciseID: 1,
									QuestionID: 12,
									Answer:     "c",
								},
							},
						}
					},
				},
			},
			args: args{
				id:     1,
				userID: 4,
			},
			wantHash:  "GMnozAj8EQ97",
			wantScore: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := NewServiceExercise(tt.fields.repositoryExercise)
			got := se.CalculateUserScore(tt.args.id, tt.args.userID)
			if !reflect.DeepEqual(got.Hash, tt.wantHash) {
				t.Errorf("ServiceExercise.CalculateUserScore() = %v, want %v", got.Hash, tt.wantHash)
			}
			if !reflect.DeepEqual(got.Score.Score, tt.wantScore) {
				t.Errorf("ServiceExercise.CalculateUserScore() = %v, want %v", got.Score.Score, tt.wantScore)
			}
		})
	}
}

func TestServiceExercise_CreateExercise(t *testing.T) {
	type fields struct {
		repositoryExercise RepositoryExercise
	}
	type args struct {
		exercise domain.Exercise
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "case failed create exercise",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					CreateExerciseFunc: func(exercise domain.Exercise) domain.Repository {
						return domain.Repository{
							Error: errors.New("failed create exercise"),
						}
					},
				},
			},
			args: args{
				exercise: domain.Exercise{
					Title:       "Dev",
					Description: "annual olympiad at school",
				},
			},
			want: "GM9jA7ylQP48",
		},
		{
			name: "case no error",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					CreateExerciseFunc: func(exercise domain.Exercise) domain.Repository {
						return domain.Repository{}
					},
				},
			},
			args: args{
				exercise: domain.Exercise{
					Title:       "Dev",
					Description: "annual olympiad at school",
				},
			},
			want: "GMDVvaZ47LnR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := NewServiceExercise(tt.fields.repositoryExercise)
			if got := se.CreateExercise(tt.args.exercise); !reflect.DeepEqual(got.Hash, tt.want) {
				t.Errorf("ServiceExercise.CreateExercise() = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}

func TestServiceExercise_CheckIsInvalidAnswer(t *testing.T) {
	type fields struct {
		repositoryExercise RepositoryExercise
	}
	type args struct {
		answer string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "case answer from user is a",
			args: args{
				answer: "a",
			},
			want: false,
		},
		{
			name: "case answer from user is b",
			args: args{
				answer: "b",
			},
			want: false,
		},
		{
			name: "case answer from user is c",
			args: args{
				answer: "c",
			},
			want: false,
		},
		{
			name: "case answer from user is d",
			args: args{
				answer: "d",
			},
			want: false,
		},
		{
			name: "case answer from user is A",
			args: args{
				answer: "A",
			},
			want: false,
		},
		{
			name: "case answer from user is B",
			args: args{
				answer: "B",
			},
			want: false,
		},
		{
			name: "case answer from user is C",
			args: args{
				answer: "C",
			},
			want: false,
		},
		{
			name: "case answer from user is D",
			args: args{
				answer: "D",
			},
			want: false,
		},
		{
			name: "case invalid answer from user",
			args: args{
				answer: "e",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := NewServiceExercise(tt.fields.repositoryExercise)
			if got := se.CheckIsInvalidAnswer(tt.args.answer); got != tt.want {
				t.Errorf("ServiceExercise.CheckIsInvalidAnswer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceExercise_CreateQuestion(t *testing.T) {
	type fields struct {
		repositoryExercise RepositoryExercise
	}
	type args struct {
		question domain.Question
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "case failed create question",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					CreateQuestionFunc: func(question domain.Question) domain.Repository {
						return domain.Repository{
							Error: errors.New("failed create question"),
						}
					},
				},
			},
			args: args{
				question: domain.Question{
					Body:          "what is the capital of indonesia?",
					OptionA:       "jakarta",
					OptionB:       "bandung",
					OptionC:       "surabaya",
					OptionD:       "medan",
					CorrectAnswer: "e",
				},
			},
			want: "GMBaOXo2oGjO",
		},
		{
			name: "case no error",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					CreateQuestionFunc: func(question domain.Question) domain.Repository {
						return domain.Repository{}
					},
				},
			},
			args: args{
				question: domain.Question{
					Body:          "what is the capital of indonesia?",
					OptionA:       "jakarta",
					OptionB:       "bandung",
					OptionC:       "surabaya",
					OptionD:       "medan",
					CorrectAnswer: "a",
				},
			},
			want: "GMCbzxwxlNYo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := NewServiceExercise(tt.fields.repositoryExercise)
			if got := se.CreateQuestion(tt.args.question); !reflect.DeepEqual(got.Hash, tt.want) {
				t.Errorf("ServiceExercise.CreateQuestion() = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}

func TestServiceExercise_CreateAnswer(t *testing.T) {
	type fields struct {
		repositoryExercise RepositoryExercise
	}
	type args struct {
		answer domain.Answer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "case failed create answer",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					CreateAnswerFunc: func(answer domain.Answer) domain.Repository {
						return domain.Repository{
							Error: errors.New("failed create answer"),
						}
					},
				},
			},
			args: args{
				answer: domain.Answer{
					ExerciseID: 1,
					QuestionID: 11,
					UserID:     4,
					Answer:     "e",
				},
			},
			want: "GMOAzxBj08Ui",
		},
		{
			name: "case no error",
			fields: fields{
				repositoryExercise: &RepositoryExerciseMock{
					CreateAnswerFunc: func(answer domain.Answer) domain.Repository {
						return domain.Repository{}
					},
				},
			},
			args: args{
				answer: domain.Answer{
					ExerciseID: 1,
					QuestionID: 11,
					UserID:     4,
					Answer:     "a",
				},
			},
			want: "GM0qIcIzIBA9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := NewServiceExercise(tt.fields.repositoryExercise)
			if got := se.CreateAnswer(tt.args.answer); !reflect.DeepEqual(got.Hash, tt.want) {
				t.Errorf("ServiceExercise.CreateAnswer() = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}
