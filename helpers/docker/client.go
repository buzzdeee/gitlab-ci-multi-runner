package docker_helpers


type Client interface {
	InspectImage(name string) (*docker.Image, error)
	PullImage(opts docker.PullImageOptions, auth docker.AuthConfiguration) error
	ImportImage(opts docker.ImportImageOptions) error

	CreateContainer(opts docker.CreateContainerOptions) (*docker.Container, error)
	StartContainer(id string, hostConfig *docker.HostConfig) error
	KillContainer(opts docker.KillContainerOptions) error
	InspectContainer(id string) (*docker.Container, error)
	AttachToContainerNonBlocking(opts docker.AttachToContainerOptions) (docker.CloseWaiter, error)
	RemoveContainer(opts docker.RemoveContainerOptions) error
	DisconnectNetwork(id string, opts docker.NetworkConnectionOptions) error
	ListNetworks() ([]docker.Network, error)
	Logs(opts docker.LogsOptions) error

	Info() (*docker.Env, error)
}
