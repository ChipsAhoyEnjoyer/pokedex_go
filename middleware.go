package main

func registryMiddleware(f func(*user, string, map[string]cliCommand) error) func(*user, string) error {
	return func(u *user, s string) error {
		commands := generateCommands()
		if err := f(u, s, commands); err != nil {
			return err
		}
		return nil
	}
}
