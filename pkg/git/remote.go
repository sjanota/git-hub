package git

type RemotesLister interface {
	List(cfg Repo) ([]string, error)
}

type AllRemotesLister struct{}

func (AllRemotesLister) List(cfg Repo) ([]string, error) {
	return cfg.ListRemoteNames()
}

type OneRemoteLister struct {
	Remote string
}

func (l OneRemoteLister) List(cfg Repo) ([]string, error) {
	return []string{l.Remote}, nil
}
