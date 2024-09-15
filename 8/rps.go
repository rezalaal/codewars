package kata8

/*
#Rock Paper Scissors!

Rock Paper Scissors
Let's play! You have to return which player won! In case of a draw return Draw!.

Examples(Input1, Input2 --> Output):

"scissors", "paper" --> "Player 1 won!"
"scissors", "rock" --> "Player 2 won!"
"paper", "paper" --> "Draw!"

<img src="https://i.imgur.com/aimOQVX.png"/>
*/
func Rps(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}
	if (p1 == "rock" && p2 == "scissors") || (p1 == "paper" && p2 == "rock") || (p1 == "scissors" && p2 == "paper") {
		return "Player 1 won!"
	}

	return "Player 2 won!"
}

var m = map[string]string{"rock": "paper", "paper": "scissors", "scissors": "rock"}

func Rps2(a, b string) string {
	if a == b {
		return "Draw!"
	}
	if m[a] == b {
		return "Player 2 won!"
	}
	return "Player 1 won!"
}
