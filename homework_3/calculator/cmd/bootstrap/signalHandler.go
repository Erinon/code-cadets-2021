package bootstrap

import "code-cadets-2021/homework_3/calculator/internal/tasks"

func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
