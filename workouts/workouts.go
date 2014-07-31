package workouts

import "time"

type ITS_A_SWOLE_WORKOUT struct {
	Kind string
	Amount int
}

func GET_ME_THE_WORKOUT_FOR_RIGHT_NOW(t time.Time) ITS_A_SWOLE_WORKOUT {
	return ITS_A_SWOLE_WORKOUT{
		Kind: "PUSHUPS",
		Amount: -1,
	}
}