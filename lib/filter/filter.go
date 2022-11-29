package filter

import (
	"fmt"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/sirupsen/logrus"
)

type Filter interface {
	Match(string) bool
}

type Engine struct {
	rules []string
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) AddRule(patterns ...string) error {
	for _, rule := range patterns {
		if !doublestar.ValidatePathPattern(rule) {
			return fmt.Errorf("invalid pattern: %s", rule)
		}
	}
	e.rules = append(e.rules, patterns...)
	return nil
}

func (e *Engine) Match(r string) bool {
	for _, rule := range e.rules {
		if rule == r {
			return true
		}
		ok, err := doublestar.PathMatch(rule, r)
		if err != nil {
			logrus.WithField("rule", rule).WithField("target", r).Error(err)
		}
		if ok {
			return true
		}
	}
	return false
}
