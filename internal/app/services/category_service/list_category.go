package categoryservice

import (
	"fmt"
	"time"

	"github.com/arthur-juan/voting-golang-rabbitmq/internal/app/types"
)

func (s *CategoryService) ListCategory() ([]*types.CategoryOutput, error) {

	var categories []types.Category

	if err := s.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	fmt.Println(categories)

	var output []*types.CategoryOutput

	for _, item := range categories {
		fmt.Println(item)

		var status string

		switch {
		case item.End.Before(time.Now()):
			status = "Ended"
		case item.End.After(time.Now()):
			status = "Open"
		}

		output = append(output, &types.CategoryOutput{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			End:         item.End,
			Winners:     int(item.Winners),
			Status:      status,
		})
	}

	fmt.Println(output)

	return output, nil
}
