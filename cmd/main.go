package main

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/configs"
)

func main() {
	// Initialize logs
	configs.InitLogs()

	// Initialize env
	configs.InitializeEnv()

	// Initialize database
	db, err := configs.InitializeSQLite()
	if err != nil {
		panic("Failed to initialize database: " + err.Error())
	}

	// Initialize redis
	redis, err := configs.InitializeRedisClient()
	if err != nil {
		panic("Failed to initialize redis: " + err.Error())
	}
	
	// Initialize server
	configs.Initialize(db, redis)
}