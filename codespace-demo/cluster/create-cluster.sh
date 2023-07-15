cat <<EOF | kind create cluster --name=skyramp-local-kind -v7 --wait 1m --retain --config=-
apiVersion: kind.x-k8s.io/v1alpha4
kind: Cluster
nodes:
  - role: control-plane
    extraMounts:
      - hostPath: ../../src/checkoutservice
        containerPath: /checkoutservice
      - hostPath: ../../src/cartservice
        containerPath: /cartservice
      - hostPath: ../../src/productcatalogservice
        containerPath: /productcatalogservice
EOF