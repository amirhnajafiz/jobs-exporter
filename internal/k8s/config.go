package k8s

type Config struct {
	Path      string `koanf:"kubeconfig"`
	Namespace string `koanf:"namespace"`
}
