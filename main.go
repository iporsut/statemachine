package main

import (
	"fmt"
)

type Player string

const (
	A Player = "A"
	B Player = "B"
)

type ScoreFN func(nextPlayerScore Player) *ScoreState

type ScoreState struct {
	Score     string
	NextScore ScoreFN
}

var (
	LoveAll *ScoreState = New("Love All")

	FifteenZero *ScoreState = New("Fifteen - Zero")
	ZeroFifteen *ScoreState = New("Zero - Fifteen")

	ThirtyZero *ScoreState = New("Thirty - Zero")
	FifteenAll *ScoreState = New("Fifteen All")
	ZeroThirty *ScoreState = New("Zero - Thirty")

	FortyZero     *ScoreState = New("Forty - Zero")
	ThirtyFifteen *ScoreState = New("Thirty - Fifteen")
	FifteenThirty *ScoreState = New("Fifteen - Thirty")
	ZeroForty     *ScoreState = New("Zero - Forty")

	Awin         *ScoreState = New("A Win")
	FortyFifteen *ScoreState = New("Forty - Fifteen")
	ThirtyAll    *ScoreState = New("Thirty All")
	FifteenForty *ScoreState = New("Fifteen - Forty")
	Bwin         *ScoreState = New("B Win")

	FortyThirty *ScoreState = New("Forty - Thirty")
	ThirtyForty *ScoreState = New("Thirty - Forty")

	Deuce *ScoreState = New("Deuce")

	Aadv *ScoreState = New("A Adv.")
	Badv *ScoreState = New("B Adv.")
)

func makeNextScore(aScoreState, bScoreState *ScoreState) ScoreFN {
	return func(nextPlayerScore Player) (nextState *ScoreState) {
		switch nextPlayerScore {
		case A:
			nextState = aScoreState
		case B:
			nextState = bScoreState
		}
		return
	}
}

func New(score string) *ScoreState {
	return &ScoreState{Score: score}
}

func init() {
	LoveAll.NextScore = makeNextScore(FifteenZero, ZeroFifteen)

	FifteenZero.NextScore = makeNextScore(ThirtyZero, FifteenAll)
	ZeroFifteen.NextScore = makeNextScore(FifteenAll, ZeroThirty)

	ThirtyZero.NextScore = makeNextScore(FortyZero, ThirtyFifteen)
	FifteenAll.NextScore = makeNextScore(ThirtyFifteen, FifteenThirty)
	ZeroThirty.NextScore = makeNextScore(FifteenThirty, ZeroForty)

	FortyZero.NextScore = makeNextScore(Awin, FortyFifteen)
	ThirtyFifteen.NextScore = makeNextScore(FortyFifteen, ThirtyAll)
	FifteenThirty.NextScore = makeNextScore(ThirtyAll, FifteenForty)
	ZeroForty.NextScore = makeNextScore(FifteenForty, Bwin)

	FortyFifteen.NextScore = makeNextScore(Awin, FortyThirty)
	ThirtyAll.NextScore = makeNextScore(FortyThirty, ThirtyForty)
	FifteenForty.NextScore = makeNextScore(ThirtyForty, Bwin)

	FortyThirty.NextScore = makeNextScore(Awin, Deuce)
	ThirtyForty.NextScore = makeNextScore(Deuce, Bwin)

	Deuce.NextScore = makeNextScore(Aadv, Badv)

	Aadv.NextScore = makeNextScore(Awin, Deuce)
	Badv.NextScore = makeNextScore(Deuce, Bwin)
}

func main() {
	fmt.Println(LoveAll.
		NextScore("A").
		NextScore("B").
		NextScore("A").
		NextScore("B").
		Score)
}
