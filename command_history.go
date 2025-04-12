package main

type commandHistory struct {
	currentIndex int
	historyLen   int
	history      []string
}

func newUserHistory() *commandHistory {
	return &commandHistory{currentIndex: 0, history: []string{""}, historyLen: 1}
}

func (uH *commandHistory) add(element string) {
	uH.history[uH.historyLen-1] = element
	uH.history = append(uH.history, "")
	uH.historyLen++
	uH.currentIndex = uH.historyLen - 1
}

func (uH *commandHistory) decrementIndex() (element string, success bool) {
	if uH.currentIndex > 0 {
		uH.currentIndex--
		return uH.history[uH.currentIndex], true
	}
	return "", false
}

func (uH *commandHistory) incrementIndex() (element string, success bool) {
	if uH.currentIndex < uH.historyLen-1 {
		uH.currentIndex++
		return uH.history[uH.currentIndex], true
	}
	return "", false
}
