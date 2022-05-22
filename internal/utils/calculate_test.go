package utils

import (
	assert "github.com/stretchr/testify/assert"
	"split-costs-bot/internal/domain"
	"testing"
	"time"
)

func TestCalculate(t *testing.T) {
	testTable := []struct {
		input    []domain.Cost
		expected []domain.Payment
	}{
		{
			input: []domain.Cost{
				{
					ID:              1,
					UserID:          1,
					Name:            "Aleksey",
					ParticipantsIds: []int64{2, 3},
					Price:           90000,
					CreatedAt:       time.Now(),
					UpdatedAt:       time.Now(),
				},
			},
			expected: []domain.Payment{
				{
					UserId:      2,
					RecipientID: 1,
					Amount:      30000,
				},
				{
					UserId:      3,
					RecipientID: 1,
					Amount:      30000,
				},
			},
		},
		{
			input: []domain.Cost{
				{
					ID:              1,
					UserID:          1,
					Name:            "Aleksey",
					ParticipantsIds: []int64{2, 3},
					Price:           90000,
					CreatedAt:       time.Now(),
					UpdatedAt:       time.Now(),
				},
				{
					ID:              2,
					UserID:          2,
					Name:            "Aleksey",
					ParticipantsIds: []int64{1, 3},
					Price:           30000,
					CreatedAt:       time.Now(),
					UpdatedAt:       time.Now(),
				},
			},
			expected: []domain.Payment{
				{
					UserId:      2,
					RecipientID: 1,
					Amount:      20000,
				},
				{
					UserId:      3,
					RecipientID: 1,
					Amount:      30000,
				},
				{
					UserId:      3,
					RecipientID: 2,
					Amount:      10000,
				},
			},
		},
		{
			input: []domain.Cost{
				{
					ID:              1,
					UserID:          1,
					Name:            "Aleksey",
					ParticipantsIds: []int64{2},
					Price:           10000,
					CreatedAt:       time.Now(),
					UpdatedAt:       time.Now(),
				},
				{
					ID:              2,
					UserID:          2,
					Name:            "Aleksey",
					ParticipantsIds: []int64{1},
					Price:           30000,
					CreatedAt:       time.Now(),
					UpdatedAt:       time.Now(),
				},
			},
			expected: []domain.Payment{
				{
					UserId:      1,
					RecipientID: 2,
					Amount:      10000,
				},
			},
		},
		{
			input: []domain.Cost{
				{
					ID:              1,
					UserID:          1,
					Name:            "Aleksey",
					ParticipantsIds: []int64{2},
					Price:           10000,
					CreatedAt:       time.Now(),
					UpdatedAt:       time.Now(),
				},
				{
					ID:              2,
					UserID:          2,
					Name:            "Aleksey",
					ParticipantsIds: []int64{1},
					Price:           10000,
					CreatedAt:       time.Now(),
					UpdatedAt:       time.Now(),
				},
			},
			expected: []domain.Payment{},
		},
		{
			input:    []domain.Cost{},
			expected: []domain.Payment{},
		},
	}

	for _, test := range testTable {
		result := Calculate(test.input)
		assert.ElementsMatch(t, result, test.expected, "Wrong costs calculation")
	}
}
