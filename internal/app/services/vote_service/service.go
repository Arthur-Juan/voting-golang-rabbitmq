package voteservice

import (
	"github.com/arthur-juan/voting-golang-rabbitmq/config"
	"github.com/arthur-juan/voting-golang-rabbitmq/pkg/queue"
	"gorm.io/gorm"
)

type VoteService struct {
	db    *gorm.DB
	queue *queue.RabbitMq
}

func NewVoteService() *VoteService {
	return &VoteService{
		db:    config.GetDb(),
		queue: queue.NewQueue(),
	}
}
