package voteservice

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *VoteService) Vote(user_id uint, input *types.VoteInput) error {

	//check if target is on category
	var candidate *types.CandidateCategory
	if err := s.db.Where(&types.CandidateCategory{UserId: input.TargetId, CategoryId: input.CategoryId, Status: types.Approved}).First(&candidate).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("candidado não pertence a categoria votada, ou não teve candidatura aprovada")
		}
		return err
	}

	//check if user already vote on this category
	var voted *types.Vote

	if err := s.db.Where(&types.Vote{VoterId: user_id, CategoryId: input.CategoryId}).Find(&voted).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
		}
	}
	if voted != nil {
		return errors.New("você já votou nessa categoria")
	}

	message := struct {
		ID          string
		CandidateId uint
		CategoryId  uint
	}{
		ID:          uuid.NewString(),
		CandidateId: input.TargetId,
		CategoryId:  input.CategoryId,
	}

	json, err := json.Marshal(message)
	if err != nil {
		return err
	}

	ch, err := s.queue.Connect(s.queue.ConnStr)
	if err != nil {
		return errors.New(fmt.Sprintf("Erro conexão na fila: %s\n", err.Error()))

	}
	defer s.queue.Disconnect(ch)

	err = s.queue.Publish(ch, string(json))
	if err != nil {
		return errors.New(fmt.Sprintf("Erro ao enviar mensagem: %s\n", err.Error()))
	}

	//insere voto no banco
	vote := *&types.Vote{
		CategoryId: input.CategoryId,
		VoterId:    user_id,
	}

	result := s.db.Create(&vote)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
