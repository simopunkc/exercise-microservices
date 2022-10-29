package domain

import "time"

type Exercise struct {
	ID          int64      `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Questions   []Question `json:"questions,omitempty"`
}

type Score struct {
	Score int `json:"score"`
}

type Question struct {
	ID            int64     `json:"id,omitempty"`
	ExerciseID    int64     `json:"exercise_id,omitempty"`
	Body          string    `json:"body,omitempty"`
	OptionA       string    `json:"option_a,omitempty"`
	OptionB       string    `json:"option_b,omitempty"`
	OptionC       string    `json:"option_c,omitempty"`
	OptionD       string    `json:"option_d,omitempty"`
	CorrectAnswer string    `json:"correct_answer,omitempty"`
	Score         int       `json:"score,omitempty"`
	CreatorID     int64     `json:"creator_id,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type PublicQuestion struct {
	ID        int64     `json:"id"`
	Body      string    `json:"body"`
	OptionA   string    `json:"option_a"`
	OptionB   string    `json:"option_b"`
	OptionC   string    `json:"option_c"`
	OptionD   string    `json:"option_d"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Answer struct {
	ID         int64     `json:"id,omitempty"`
	ExerciseID int64     `json:"exercise_id,omitempty"`
	QuestionID int64     `json:"question_id,omitempty"`
	UserID     int64     `json:"user_id,omitempty"`
	Answer     string    `json:"answer,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type Repository struct {
	Hash        string   `json:"hash,omitempty"`
	StatusCode  int      `json:"status_code,omitempty"`
	Exercise    Exercise `json:"exercise,omitempty"`
	ListAnswer  []Answer `json:"list_answer,omitempty"`
	RawResponse string   `json:"raw_response,omitempty"`
	Error       error    `json:"error,omitempty"`
}

type Service struct {
	Hash        string   `json:"hash"`
	Exercise    Exercise `json:"exercise,omitempty"`
	Score       Score    `json:"score,omitempty"`
	RawResponse string   `json:"raw_response,omitempty"`
	Error       error    `json:"error,omitempty"`
}

type Handler struct {
	Hash        string    `json:"hash"`
	StatusCode  int       `json:"status_code,omitempty"`
	Exercise    *Exercise `json:"exercise,omitempty"`
	Score       *Score    `json:"score,omitempty"`
	RawResponse string    `json:"raw_response,omitempty"`
	Error       error     `json:"error,omitempty"`
}
