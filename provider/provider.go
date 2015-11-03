package provider

type Provider interface {

	Search(arg ...string) ([]string, error)

	Install(arg ...string) ([]byte, error)
}
