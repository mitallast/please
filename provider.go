package please

type Provider interface {

    search(arg ...string) ([]byte, error)

    install(arg ...string) ([]byte, error)
}

