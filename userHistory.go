package main

type userHistory struct {
	currentIndex int
	historyLen   int
	history      []string
}

func newUserHistory() *userHistory {
	return &userHistory{currentIndex: 0, history: []string{""}, historyLen: 1}
}

func (uH *userHistory) add(element string) {
	uH.history[uH.historyLen-1] = element
	uH.history = append(uH.history, "")
	uH.historyLen++
	uH.currentIndex = uH.historyLen - 1
}

func (uH *userHistory) decrementIndex() (element string, success bool) {
	if uH.currentIndex > 0 {
		uH.currentIndex--
		return uH.history[uH.currentIndex], true
	}
	return "", false
}

func (uH *userHistory) incrementIndex() (element string, success bool) {
	if uH.currentIndex < uH.historyLen-1 {
		uH.currentIndex++
		return uH.history[uH.currentIndex], true
	}
	return "", false
}
