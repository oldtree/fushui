package auth

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Info("auth package init")
}

type Auth struct {
}

func (auth *Auth) Auth() (bool, error) {
	return true, nil
}

func (auth *Auth) AuthThirdType() (bool, error) {
	return true, nil
}

func (auth *Auth) AuthWithoutRandom() (bool, error) {
	return true, nil
}
