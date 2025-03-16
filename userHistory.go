package main

type userHistory struct {
	currentIndex int
	historyLen   int
	history      []string
}

func newUserHistory() *userHistory {
	return &userHistory{currentIndex: -1, history: []string{}, historyLen: 0}
}

func (uH *userHistory) get() (element string, success bool) {
	if uH.historyLen > 0 {
		return uH.history[uH.currentIndex], true
	}
	return "", false
}

func (uH *userHistory) add(element string) {
	uH.history = append(uH.history, element)
	uH.currentIndex++
	uH.historyLen++
}
