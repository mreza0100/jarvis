package runnerport

type Runner interface {
	RunScript(script string) (string, error)
}
