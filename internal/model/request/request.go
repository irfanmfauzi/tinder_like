package request

type PremiumRequestRegister struct {
	IsVerified      bool `json:"is_verified"`
	IsInfiniteQuota bool `json:"is_infinite_quota"`
}

type RequestRegister struct {
	Email    string                 `json:"email"`
	Password string                 `json:"password"`
	Premium  PremiumRequestRegister `json:"premium,omitempty"`
	Name     string                 `json:"name"`
	Gender   string                 `json:"gender"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestSwipedMember struct {
	SwipeMemberId int64 `json:"swipe_member_id"`
	IsLiked       bool  `json:"is_liked"`
}
