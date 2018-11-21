package tenniskata

type tennisGame1 struct {
	player1 player
	player2 player
}

type player struct{
	score int
	name string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	player1 := player{
		name: player1Name}
	player2 := player{
		name: player2Name}
	game := &tennisGame1{
		player1: player1,
		player2: player2}

	return game
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.player1.score += 1
	} else {
		game.player2.score += 1
	}
}

func (game *tennisGame1) GetScore() string {
	score := ""
	tempScore := 0
	if game.player1.score == game.player2.score {
		switch game.player1.score {
		case 0:
			score = "Love-All"
		case 1:
			score = "Fifteen-All"
		case 2:
			score = "Thirty-All"
		default:
			score = "Deuce"
		}
	} else if game.player1.score >= 4 || game.player2.score >= 4 {
		minusResult := game.player1.score - game.player2.score
		if minusResult == 1 {
			score = "Advantage player1"
		} else if minusResult == -1 {
			score = "Advantage player2"
		} else if minusResult >= 2 {
			score = "Win for player1"
		} else {
			score = "Win for player2"
		}
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = game.player1.score
			} else {
				score += "-"
				tempScore = game.player2.score
			}
			switch tempScore {
			case 0:
				score += "Love"
			case 1:
				score += "Fifteen"
			case 2:
				score += "Thirty"
			case 3:
				score += "Forty"
			}
		}
	}
	return score
}
