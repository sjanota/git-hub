package config

type RemotesLister interface {
	List(cfg Config) ([]string, error)
}

type AllRemotesLister struct{}

func (AllRemotesLister) List(cfg Config) ([]string, error) {
	return cfg.ListRemoteNames()
}

type OneRemoteLister struct {
	Remote string
}

func (l OneRemoteLister) List(cfg Config) ([]string, error) {
	return []string{l.Remote}, nil
}
