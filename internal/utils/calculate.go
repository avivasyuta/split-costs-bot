package utils

import (
	"split-costs-bot/internal/domain"
	"strconv"
	"strings"
)

func makeIndex(id int64, id2 int64) string {
	return strconv.FormatInt(id, 10) + ":" + strconv.FormatInt(id2, 10)
}

func splitIndex(index string) (int64, int64) {
	ids := strings.Split(index, ":")
	firstId, err := strconv.ParseInt(ids[0], 10, 64)
	if err != nil {
		panic(err)
	}

	secondId, err := strconv.ParseInt(ids[1], 10, 64)
	if err != nil {
		panic(err)
	}

	return firstId, secondId
}

func Calculate(costs []domain.Cost) []domain.Payment {
	var result []domain.Payment
	payments := make(map[string]int64)

	for _, cost := range costs {
		amount := int64(cost.Price / (len(cost.ParticipantsIds) + 1))

		for _, participantId := range cost.ParticipantsIds {
			index := makeIndex(cost.UserID, participantId)
			_, ok := payments[index]
			if ok {
				payments[index] = payments[index] + amount
				continue
			}

			index = makeIndex(participantId, cost.UserID)
			payments[index] = payments[index] - amount
		}
	}

	for index, amount := range payments {
		firstUserId, secondUserId := splitIndex(index)

		if amount > 0 {
			result = append(result, domain.Payment{
				UserId:      secondUserId,
				RecipientID: firstUserId,
				Amount:      amount,
			})
		}

		if amount < 0 {
			result = append(result, domain.Payment{
				UserId:      firstUserId,
				RecipientID: secondUserId,
				Amount:      -amount,
			})
		}
	}

	return result
}
