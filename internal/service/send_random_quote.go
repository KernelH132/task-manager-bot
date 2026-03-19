package service

import (
	"math/rand"
)

// SendRandomQuote sends a random quote from a predefined list
func SendRandomQuote(chatID int64) {
	quotes := []string{
		"Keep going, you're closer than you think.",
		"Small steps still move you forward.",
		"Progress beats perfection.",
		"You've handled worse, you'll handle this.",
		"Start now, figure it out later.",
		"Discipline creates freedom.",
		"Your future is built today.",
		"Stay consistent, results will come.",
		"Hard things build strong people.",
		"Do it tired, do it unmotivated, just do it.",
		"Every day is a new chance to improve.",
		"Focus on what you can control.",
		"You don’t need permission to grow.",
		"Action removes doubt.",
		"Keep showing up, that’s the secret.",
		"Greatness starts with small habits.",
		"Your effort is never wasted.",
		"Make today count.",
		"Comfort zones don’t build success.",
		"Be better than yesterday.",
		"You are capable of more.",
		"Stay patient, stay working.",
		"Momentum comes from movement.",
		"Nothing changes if nothing changes.",
		"Push through, you're almost there.",
		"Believe it, then build it.",
		"Growth feels uncomfortable for a reason.",
		"Win the day, one task at a time.",
		"Your mindset shapes your reality.",
		"Done is better than perfect.",
		"Chase progress, not approval.",
		"Turn effort into results.",
		"Don’t quit before it works.",
		"Make yourself proud.",
		"Stay focused, ignore distractions.",
		"You’re building something bigger.",
		"Energy flows where focus goes.",
		"Level up quietly.",
		"Consistency beats talent.",
		"Work now, shine later.",
		"Keep it simple, keep it moving.",
		"Results come to those who persist.",
		"Trust the process.",
		"Pressure creates diamonds.",
		"Do the work, reap the reward.",
		"Your limits are not fixed.",
		"Stay hungry, stay driven.",
		"Every effort counts.",
		"Build, learn, repeat.",
		"Success starts with trying.",
	}

	randomIndex := rand.Intn(len(quotes))
	randomQuote := quotes[randomIndex]

	SendMessage(chatID, randomQuote)
}
