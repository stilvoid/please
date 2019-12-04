// Package http provides some utility functions for dealing with HTTP requests and responses
package http

type Options struct {
	HeadersIncluded bool
	IncludeHeaders  bool
	IncludeMethod   bool
	IncludePath     bool
	IncludeStatus   bool
}
