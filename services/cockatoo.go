package services

import (
	"context"
	"fmt"
	"time"

	"github.com/ansharw/rest-api/api"
)

type Cockatoo struct {
	session ISession
}

func (ckt *Cockatoo) SetSession(session ISession) {
	ckt.session = session
}

type sessionReq struct {
	BrowserSession string `json:"browser_session"`
	UserId         string `json:"user_id"`
	UserAgent      string `json:"user_agent"`
	Type           string `json:"type"`
}

type ResponseSession struct {
	Message string `json:"message"`
}

func (ckt *Cockatoo) ProcessSetSession(ctx context.Context, ses *sessionReq) (response *ResponseSession, err error) {
	sesReq := &api.Session{
		BrowserSession: ses.BrowserSession,
		UserId:         ses.UserId,
		UserAgent:      ses.UserAgent,
		Type:           ses.Type,
	}

	errRedisCreateSession := ckt.session.CreateSession(ctx, "user-session", ses.UserId, sesReq, time.Duration(env.Adapter.Redis.Expire))
	if errRedisCreateSession != nil {
		return &ResponseSession{
			Message: "Failed create session",
		}, fmt.Errorf("[Session]ProcessSetSession - failed create session from redis [%s]", errRedisCreateSession)
	} 

	return &ResponseSession{
		Message: "Success created session",
	}, nil
}
