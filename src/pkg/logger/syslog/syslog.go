package syslog

import (
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"log"
)

type syslog struct {
}

func New() logger.Logger {
	return &syslog{}
}

func (s *syslog) Error(msg logger.ErrorType) {
	log.Println(msg)
}

func (s *syslog) Fatal(msg logger.ErrorType) {
	log.Println(msg)
}

func (s *syslog) Warn(msg logger.ErrorType) {
	log.Println(msg)
}

func (s *syslog) Info(msg logger.ErrorType) {
	log.Println(msg)
}
