package main

import "github.com/JoaoBarreto255/GoBorrowMeABook/internal/router"
import "github.com/JoaoBarreto255/GoBorrowMeABook/configs"

func main() {
	logger := configs.GetLogger()
	router.Initializer()
}
