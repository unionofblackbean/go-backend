package dao

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/unionofblackbean/go-backend/services/smartbots/api/entities"
)

func CreateBot(bot *entities.Bot) (err error) {
	err = pool.Validate()
	if err != nil {
		return
	}

	err = pool.Exec(
		"INSERT INTO bots VALUES ($1, $2);",
		pgtype.UUID{
			Bytes:  bot.UUID,
			Status: pgtype.Present,
		},
		bot.Name,
	)
	return
}

func GetBot(uuid uuid.UUID) (*entities.Bot, error) {
	bot := new(entities.Bot)

	var rawUUID pgtype.UUID
	err := pool.QueryRow(
		"SELECT uuid, name FROM bots WHERE uuid=$1;",
		pgtype.UUID{
			Bytes:  uuid,
			Status: pgtype.Present,
		},
	).Scan(
		&rawUUID,
		&bot.Name,
	)
	if err != nil {
		return nil, err
	}

	bot.UUID = rawUUID.Bytes

	return bot, nil
}

func GetAllBots() ([]entities.Bot, error) {
	err := pool.Validate()
	if err != nil {
		return nil, err
	}

	rows, err := pool.Query(
		"SELECT uuid, name FROM bots;",
	)
	if err != nil {
		return nil, err
	}

	var bots []entities.Bot
	for rows.Next() {
		var rawBotUUID pgtype.UUID
		var botName string
		err = rows.Scan(
			&rawBotUUID,
			&botName,
		)
		if err != nil {
			return nil, err
		}

		botUUID, err := uuid.FromBytes(rawBotUUID.Bytes[:])
		if err != nil {
			return nil, fmt.Errorf("failed to process UUID obtained from database -> %v", err)
		}

		bots = append(bots, entities.Bot{
			UUID: botUUID,
			Name: botName,
		})
	}

	return bots, nil
}

func UpdateBot(bot *entities.Bot) (err error) {
	err = pool.Validate()
	if err != nil {
		return
	}

	err = pool.Exec(
		"UPDATE bots SET name=$2 WHERE uuid=$3;",
		bot.Name,
		pgtype.UUID{
			Bytes:  bot.UUID,
			Status: pgtype.Present,
		},
	)
	return

}

func DeleteBot(uuid uuid.UUID) (err error) {
	err = pool.Validate()
	if err != nil {
		return
	}

	err = pool.Exec(
		"DELETE FROM bots WHERE uuid=$1;",
		pgtype.UUID{
			Bytes:  uuid,
			Status: pgtype.Present,
		},
	)
	return
}
