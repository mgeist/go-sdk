package codefresh

import (
	"encoding/json"
	"time"
)

type (
	ITokenAPI interface {
		GenerateToken(name string, subject string) *Token
		GetTokens() []*Token
	}

	Token struct {
		ID          string    `json:"_id"`
		Name        string    `json:"name"`
		TokenPrefix string    `json:"tokenPrefix"`
		Created     time.Time `json:"created"`
		Subject     struct {
			Type string `json:"type"`
			Ref  string `json:"ref"`
		} `json:"subject"`
		Value string
	}
)

type (
	tokenSubjectType int

	getTokensReponse struct {
		Tokens []*Token
	}
)

const (
	runtimeEnvironment tokenSubjectType = 0
)

func (s tokenSubjectType) String() string {
	return [...]string{"runtime-environment"}[s]
}

func (c *codefresh) GenerateToken(name string, subject string) *Token {
	resp := c.requestAPI(&requestOptions{
		path:   "/api/auth/key",
		method: "POST",
		body: map[string]interface{}{
			"name": name,
		},
		qs: map[string]string{
			"subjectReference": subject,
			"subjectType":      runtimeEnvironment.String(),
		},
	})
	return &Token{
		Name:  name,
		Value: resp.String(),
	}
}

func (c *codefresh) GetTokens() []*Token {
	emptySlice := make([]*Token, 0)
	resp := c.requestAPI(&requestOptions{
		path:   "/api/auth/keys",
		method: "GET",
	})
	tokensAsBytes := []byte(resp.String())
	json.Unmarshal(tokensAsBytes, &emptySlice)

	return emptySlice
}
