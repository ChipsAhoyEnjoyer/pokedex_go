package main

func main() {
	commandRegistry := generateCommands()
	user := newUser()
	commandHistoryBuffer := newUserHistory()
	startRepl(
		user,
		commandRegistry,
		commandHistoryBuffer,
	)
}
