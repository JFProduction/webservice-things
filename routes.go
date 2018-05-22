package main

import "net/http"

// Route - route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - list of routes
type Routes []Route

var routes = Routes{
	Route{
		"test",
		"GET",
		"/test",
		test,
	},
	Route{
		"getInfo",
		"GET",
		"/getInfo",
		getInfo,
	},
	Route{
		"killProcesses",
		"GET",
		"/killProcesses",
		killProcesses,
	},
	Route{
		"resetTable",
		"GET",
		"/resetTable",
		resetTable,
	},
	Route{
		"resetFiles",
		"GET",
		"/resetFiles",
		resetFiles,
	},
}
