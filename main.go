package main

func main() {
	commandRegistry := generateCommandRegistry()
	user := newUser()
	commandHistoryBuffer := newUserHistory()
	startRepl(
		user,
		commandRegistry,
		commandHistoryBuffer,
	)
}
