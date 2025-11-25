package domain

type Token struct {
    Token     string `json:"token"`
    ExpiresAt int64  `json:"expires_at"`
    UserID    int32  `json:"user_id"`
}