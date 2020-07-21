package Mvc

import (
	"fmt"
	"regexp"
	"strings"
)

// ActionFilter Context
type ActionFilterContext struct {
	ActionExecutorContext
}

// ActionFilter Chain
type ActionFilterChain struct {
	pattern      string        // regex
	actionFilter IActionFilter // action filter
}

// New ActionFilter Chain
func NewActionFilterChain(pattern string, filter IActionFilter) ActionFilterChain {
	regex := fmt.Sprintf(`^%s$`, strings.ReplaceAll(strings.Trim(pattern, ""), "*", "[/a-zA-Z0-9]+"))
	return ActionFilterChain{
		pattern:      regex,
		actionFilter: filter,
	}
}

// Set ActionFilter to Chain
func (chain ActionFilterChain) SetFilter(filter IActionFilter) {
	chain.actionFilter = filter
}

// Get ActionFilter of Chain
func (chain ActionFilterChain) GetFilter() IActionFilter {
	return chain.actionFilter
}

// Match path for pattern
func (chain ActionFilterChain) MatchPath(want string) bool {
	b, _ := regexp.MatchString(chain.pattern, want)
	return b
}

//// Match path for pattern to ActionFilter
func (chain ActionFilterChain) MatchFilter(want string) IActionFilter {
	if chain.MatchPath(want) {
		return chain.GetFilter()
	}
	return nil
}
