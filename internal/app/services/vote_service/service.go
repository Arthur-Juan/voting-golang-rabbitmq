package voteservice

import (
	"github.com/arthur-juan/voting-golang-rabbitmq/config"
	"gorm.io/gorm"
)

type VoteService struct {
	db *gorm.DB
}

func NewVoteService() *VoteService {
	return &VoteService{
		db: config.GetDb(),
	}
}
