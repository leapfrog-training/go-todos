package main

import (
	r "todos/api/routes"
)

/**
 * Main entry point to to-do application
 * @function main
 */
func main() {
	// Goes to routes package and brings all the required API
	r.Route()
}
