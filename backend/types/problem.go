package types

type Problem struct {
	Id          int    `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
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
