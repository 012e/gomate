package permmanager

import (
	"errors"

	"github.com/casbin/casbin/v2"
	"github.com/sirupsen/logrus"
)

type Context struct {
	Username      string
	GroupID       int64
	GroupIDStr    string
	Route         string
	UserHaveGroup bool
}

type PermManager struct {
	*casbin.Enforcer
	Context
}

func (p *PermManager) AddPoliciesForUser(route string, newActions ...string) error {
	if len(newActions) == 0 {
		return errors.New("policy manager: no policies specified")
	}
	actions := make([][]string, len(newActions))
	for i, action := range newActions {
		actions[i] = []string{p.Username, route, action}
		logrus.Debugf("adding %s permisson for %s, at %s", action, p.Username, action)
	}
	_, err := p.AddPolicies(actions)
	if err != nil {
		return err
	}
	return nil
}

// add actions for user at the current route.
func (p *PermManager) AddPoliciesForUserHere(newActions ...string) error {
	if len(newActions) == 0 {
		return errors.New("policy manager: no policies specified")
	}
	err := p.AddPoliciesForUser(p.Route, newActions...)
	if err != nil {
		return err
	}
	return nil
}

func (p *PermManager) AddPoliciesForGroup(route string, newActions ...string) error {
	if len(newActions) == 0 {
		return errors.New("policy manager: no policies specified")
	}
	if !p.Context.UserHaveGroup {
		return errors.New("policy manager: user don't have a group")
	}
	actions := make([][]string, len(newActions))
	for i, action := range newActions {
		actions[i] = []string{p.GroupIDStr, route, action}
		logrus.Debugf("adding %s permisson for %s, at %s", action, p.GroupIDStr, action)
	}
	_, err := p.AddPolicies(actions)
	if err != nil {
		return err
	}
	return nil
}

// add actions for user at the current route.
func (p *PermManager) AddPoliciesForGroupHere(newActions ...string) error {
	if len(newActions) == 0 {
		return errors.New("policy manager: no policies specified")
	}
	err := p.AddPoliciesForGroup(p.Route, newActions...)
	if err != nil {
		return err
	}
	return nil
}

