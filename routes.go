// routes.go

package main

func initializeRoutes() {
	//handle the index route
	router.GET("/", showIndexPage)
}