package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context, user *User) error {
	/* todo: test existing of user */
	session := sessions.Default(c)

	session.Set(USERKEY, user.ID)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}
