package main

import (
	"github.com/gladiusio/gladius-application-server/pkg/controller"
	"github.com/gladiusio/gladius-application-server/pkg/db"
)

func main() {
	db.Initialize()
	controller.TempDBCalls()
}
