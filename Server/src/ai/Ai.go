package ai

import (
	"Server/ticTacToe"
	"errors"
	"fmt"
	"log"
	"math"
)

type StateValue struct {
	ChanceOfVictory float32
	Confidence      int32
}

//NumOfStates holds the total possible number of state hashes in a game of tic-tac-toe. this has a few impossible states in it
const NumOfStates = 19_683

//LearningRate holds the learning rate of the ai
const LearningRate = .2

//Alpha holds the step depreciation of value judgements during learning.
const Alpha = .9

type Ai struct {
	// holds a representation of every possible tic-tac-toe board
	StateValues [NumOfStates]StateValue
	Tile        ticTacToe.PlayerTile
}

func HashGameBoard(board ticTacToe.GameBoard) int {
	val := 0
	counter := 0
	for _, row := range board {
		for j := range row {
			val += int(row[j]) * int(math.Pow(3, float64(counter)))
		}
		counter++
	}

	return val
}

func (ai Ai) GetOptimalTurn(board ticTacToe.GameBoard) (int8, int8, error) {
	// try to add a tile to each position, try to calculate the value of that, then take the max value
	var maxValueOfState float32 = -1.0
	xPositionWithMaxValue, yPositionWithMaxValue := -1, -1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[j][i] != ticTacToe.None {
				continue
			}
			boardCopy := board
			boardCopy[j][i] = ai.Tile
			hash := HashGameBoard(board)
			valueOfState := ai.StateValues[hash].ChanceOfVictory
			if valueOfState > maxValueOfState {
				maxValueOfState = valueOfState
				xPositionWithMaxValue, yPositionWithMaxValue = i, j
			}
		}
	}
	if maxValueOfState == -1 {
		return int8(xPositionWithMaxValue), int8(yPositionWithMaxValue), errors.New("couldn't get best move")
	}
	return int8(xPositionWithMaxValue), int8(yPositionWithMaxValue), nil
}

func (ai *Ai) LearnFromGame(game ticTacToe.Game) {
	var valueOfGame float32 = 0.0
	if ai.Tile == game.Winner {
		valueOfGame = 1
	}
	if game.Winner == ticTacToe.None {
		valueOfGame = .5
	}
	boardCopy := game.Board

	gameHash := HashGameBoard(game.Board)
	ai.StateValues[gameHash].Confidence++
	ai.StateValues[gameHash].ChanceOfVictory = valueOfGame

	//todo: refactor loop body into shared method
	for i := len(game.History.History) - 1; i >= 0; i-- {
		// get the value of each step based on alpha depreciation, then update the ai
		historyItem := game.History.History[i]
		boardCopy[historyItem.Position[1]][historyItem.Position[0]] = ticTacToe.None

		gameHash := HashGameBoard(game.Board)
		ai.StateValues[gameHash].Confidence++
		state := ai.StateValues[gameHash]

		updatedAvg := float64(state.ChanceOfVictory - (state.ChanceOfVictory+valueOfGame)/float32(state.Confidence))
		//TODO: discount values based on learning rate
		//avgWithDiscounts := updatedAvg * math.Pow(Alpha, float64(len(game.History.History)-i)) * LearningRate
		ai.StateValues[gameHash].ChanceOfVictory = float32(updatedAvg)
	}
}

func RunSimulation(xAi *Ai, oAi *Ai) {
	game := ticTacToe.NewGame()
	for !game.Ended {
		if game.CurrentPlayerTurn == ticTacToe.X {
			x, y, err := xAi.GetOptimalTurn(game.Board)
			if err != nil {
				panic("there was an error running simulations that could not be recovered from")
			}
			err = game.TakeTurn(ticTacToe.X, x, y)
			if err != nil {
				panic("there was an error running simulations that could not be recovered from")
			}
		} else if game.CurrentPlayerTurn == ticTacToe.O {
			x, y, err := oAi.GetOptimalTurn(game.Board)
			if err != nil {
				panic("there was an error running simulations that could not be recovered from")
			}
			err = game.TakeTurn(ticTacToe.O, x, y)
			if err != nil {
				panic("there was an error running simulations that could not be recovered from")
			}
		}
	}
	xAi.LearnFromGame(game)
	oAi.LearnFromGame(game)
}

func RunSimulations(xAi *Ai, oAi *Ai, numOfSimulations int) {
	for i := 0; i < numOfSimulations; i++ {
		if i%10000 == 0 {
			fmt.Printf("iteration %d of %d\n", i, numOfSimulations)
		}
		RunSimulation(xAi, oAi)
	}
}

// NewAi creates a new Ai. if forceSimulation is false, then it tries to pull the saved simulation from the file system
func NewAi(numOfSimulations int) (xAi Ai, oAi Ai) {
	stateValues := [NumOfStates]StateValue{}
	for i := range stateValues {
		stateValues[i] = StateValue{
			Confidence:      0,
			ChanceOfVictory: .5,
		}
	}
	xAi = Ai{
		StateValues: stateValues,
		Tile:        ticTacToe.X,
	}
	oAi = Ai{
		StateValues: stateValues,
		Tile:        ticTacToe.O,
	}
	log.Println("ai simulations starting")
	RunSimulations(&xAi, &oAi, numOfSimulations)
	log.Println("ai done")

	return xAi, oAi
}
