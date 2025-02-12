package app

func (a *App) Wait() error {
	return a.server.WaitAndClose()
}
