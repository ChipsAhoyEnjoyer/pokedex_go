package main

func registryMiddleware(f func(*user, string, map[string]cliCommand) error) func(*user, string) error {
	return func(u *user, s string) error {
		commands := generateCommands() // TODO: Don't want to regenerate the commands list everytime the user calls help
		if err := f(u, s, commands); err != nil {
			return err
		}
		return nil
	}
}
