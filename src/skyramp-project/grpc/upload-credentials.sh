#!/bin/bash

DOCKER_CONFIG=$HOME/.docker
KIND_CLUSTER_NAME=skyramp-local-local-cluster
# setup credentials on each node
echo "Moving credentials to kind cluster name='${KIND_CLUSTER_NAME}' nodes ..."
for node in $(kind get nodes --name "${KIND_CLUSTER_NAME}"); do
  # the -oname format is kind/name (so node/name) we just want name
  node_name=${node#node/}
  # copy the config to where kubelet will look
  docker cp "${DOCKER_CONFIG}/secret" "${node_name}:/var/lib/kubelet/config.json"
  # restart kubelet to pick up the config
  docker exec "${node_name}" systemctl restart kubelet.service
done
