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
	uH.currentIndex = uH.historyLen
	uH.historyLen++
}

func (uH *userHistory) decrementIndex() (element string, success bool) {
	if uH.currentIndex > -1 && uH.historyLen > 0 {
		e := uH.history[uH.currentIndex]
		uH.currentIndex--
		return e, true
	}
	return "", false
}

func (uH *userHistory) incrementIndex() (element string, success bool) {
	if uH.currentIndex+1 < uH.historyLen {
		e := uH.history[uH.currentIndex]
		uH.currentIndex++
		return e, true
	}
	return "", false
}
