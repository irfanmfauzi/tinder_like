package entity

type User struct {
	Id              int64  `db:"id" json:"id"`
	Email           string `db:"email" json:"email"`
	Password        string `db:"password,omitempty" json:"-"`
	IsPremium       bool   `db:"is_premium" json:"is_premium"`
	IsVerified      bool   `db:"is_verified" json:"is_verified"`
	IsInfiniteQuota bool   `db:"is_infinite_quota" json:"is_infinite_quota"`
}

type Member struct {
	Id     int64  `db:"id" json:"id"`
	UserId int64  `db:"user_id" json:"user_id"`
	Name   string `db:"name" json:"name"`
	Gender string `db:"gender" json:"gender"`
}
