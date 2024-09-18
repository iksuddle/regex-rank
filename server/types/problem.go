package types

import (
	"time"
)

type Problem struct {
	Id        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Statement struct {
	Id        int    `json:"id" db:"id"`
	ProblemId int    `json:"problem_id" db:"problem_id"`
	Match     string `json:"match" db:"match"`
	Literal   string `json:"literal" db:"literal"`
}

type Solution struct {
	Id        int    `json:"id" db:"id"`
	UserId    int    `json:"user_id" db:"user_id"`
	ProblemId int    `json:"problem_id" db:"problem_id"`
	Literal   string `json:"literal" db:"literal"`
}

func NewProblem() Problem {
	return Problem{
		CreatedAt: time.Now().UTC(),
	}
}

func NewStatement(problemId int, match string, literal string) Statement {
	return Statement{
		ProblemId: problemId,
		Match:     match,
		Literal:   literal,
	}
}
