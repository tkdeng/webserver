package main

//!htmlc: --src="src/templates" --dist="dist"

func main() {
	//todo: handle server
	// also compile with separate htmlc module
	// routes will be compiled by server (not by htmlc)

	compile()

	//todo: create method to run dist files in 'routes' directory
	// also include dist route handlers in routes.go
}
