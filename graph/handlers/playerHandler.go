package handler

import "github.com/PParist/go-graphql/graph/model"

var players = []*model.Player{
	{
		ID:    "Player-01",
		Name:  "Pew",
		Level: 1,
		Class: &model.AllPlayerClass[0],
		Item: []*model.Item{
			{
				ID:   "Item-01",
				Name: "Sword",
			},

			{
				ID:   "Item-02",
				Name: "Potion",
			},
		},
	},
	{
		ID:    "Player-02",
		Name:  "Pew2",
		Level: 10,
		Class: &model.AllPlayerClass[1],
		Item: []*model.Item{
			{
				ID:   "Item-01",
				Name: "Rod",
			},

			{
				ID:   "Item-02",
				Name: "Sp-Potion",
			},
		},
	},
}

func GetPlayer() []*model.Player {
	return players
}

func CreatePlayer(req *model.NewPlayer) *model.Player {
	newPlayer := &model.Player{
		ID:    "PLAYER-03",
		Name:  req.Name,
		Level: 1,
		Class: &req.Class,
		Item: []*model.Item{
			{
				ID:   "Item-01",
				Name: "Potion",
			},
		},
	}
	players = append(players, newPlayer)
	return newPlayer
}
