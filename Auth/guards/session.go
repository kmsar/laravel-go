package guards

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IAuth"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/ISession"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

func SessionGuard(name string, config Support.Fields, ctx Support.Context, provider IAuth.UserProvider) IAuth.Guard {
	if guard, ok := ctx.Get("guard:" + name).(IAuth.Guard); ok {
		return guard
	}
	guard := &Session{
		session:    ctx.Get("session").(ISession.Session),
		ctx:        ctx,
		users:      provider,
		sessionKey: config["session_key"].(string),
	}

	ctx.Set("guard:"+name, guard)

	return guard
}

type Session struct {
	sessionKey string
	isVerified bool
	session    ISession.Session
	ctx        Support.Context
	users      IAuth.UserProvider
	current    IAuth.Authenticatable
}

func (this *Session) Logout() error {
	this.session.Remove(this.sessionKey)
	this.current = nil
	return nil
}

func (this *Session) Error() error {
	return nil
}

func (this *Session) Once(user IAuth.Authenticatable) {
	this.current = user
	this.isVerified = true
}

func (this *Session) Login(user IAuth.Authenticatable) interface{} {
	this.session.Put(this.sessionKey, user.GetId())

	this.Once(user)

	return true
}

func (this *Session) User() IAuth.Authenticatable {
	if !this.isVerified {
		this.isVerified = true
		if userId := this.session.Get(this.sessionKey, ""); userId != "" {
			this.current = this.users.RetrieveById(userId)
		}
	}

	return this.current
}

func (this *Session) GetId() (id string) {
	if user := this.User(); user != nil {
		id = user.GetId()
	}
	return
}

func (this *Session) Check() bool {
	return this.User() != nil
}

func (this *Session) Guest() bool {
	return this.User() == nil
}
