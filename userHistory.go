package main

type userHistory struct {
	currentIndex int
	historyLen   int
	history      []string
}

func newUserHistory() *userHistory {
	return &userHistory{currentIndex: -1, history: []string{}, historyLen: 0}
}

func (uH *userHistory) add(element string) {
	uH.history = append(uH.history, element)
	uH.currentIndex++
	uH.historyLen++
}

func (uH *userHistory) decrementIndex() (element string, success bool) {
	if uH.currentIndex > 0 && uH.historyLen > 0 {
		return uH.history[uH.currentIndex], true
	}
	return "", false
}

func (uH *userHistory) incrementIndex() (element string, success bool) {
	if uH.currentIndex+1 < uH.historyLen {
		return uH.history[uH.currentIndex], true
	}
	return "", false
}
