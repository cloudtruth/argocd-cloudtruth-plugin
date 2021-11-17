#!/usr/bin/env bash

# fail fast
set -e


print "Input Cloudtruth Api Key: "
read CLOUDTRUTH_API_KEY
print "Input Cloudtruth Environment (Default for all Argo projects): "
read CLOUDTRUTH_ENVIRONMENT
print "Input Cloudtruth Project (Default for all Argo projects): "
read CLOUDTRUTH_PROJECT

secret_yaml <<-EOF
$(cat argocd-cloudtruth-plugin-secret.yaml)
  CLOUDTRUTH_API_KEY: $CLOUDTRUTH_API_KEY
  CLOUDTRUTH_ENVIRONMENT: $CLOUDTRUTH_ENVIRONMENT
  CLOUDTRUTH_PROJECT: $CLOUDTRUTH_PROJECT
EOF

# --type merge ?
#  --dry-run=server -o yaml
echo ${secret_yaml} | kubectl apply -n argocd -f -
kubectl patch -n argocd deployment/argocd-repo-server --patch "$(cat argocd-repo-server.patch.yaml)"
kubectl patch -n argocd configmap/argocd-cm --patch "$(cat argocd-cm.patch.yaml)"
