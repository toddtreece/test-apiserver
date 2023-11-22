local("hack/kind.sh || true")

os = str(local('uname -s')).strip().lower()

local_resource(
  'test-apiserver',
  'go build -gcflags "all=-N -l" -o build/test-apiserver .',
  deps=['./main.go', './pkg'],
  serve_cmd='./build/test-apiserver --secure-port 8443 --kubeconfig ~/.kube/config --authentication-kubeconfig ~/.kube/config --authorization-kubeconfig ~/.kube/config --etcd-servers 127.0.0.1:2379 -v 8',
)

k8s_yaml(kustomize('deploy/local-%s' % os))
