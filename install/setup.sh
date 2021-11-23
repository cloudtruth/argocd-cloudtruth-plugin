#!/usr/bin/env bash

# fail fast
set -e


printf "Input Cloudtruth Api Key: "
read CLOUDTRUTH_API_KEY
printf "Input Cloudtruth Environment (Default and can be overridden for each Argo project): "
read CLOUDTRUTH_ENVIRONMENT
printf "Input Cloudtruth Project (Default and can be overridden for each Argo project): "
read CLOUDTRUTH_PROJECT

cat <<EOF | kubectl apply -n argocd -f -
$(cat argocd-cloudtruth-plugin-secret.yaml)
  CLOUDTRUTH_API_KEY: $(echo ${CLOUDTRUTH_API_KEY} | base64)
  CLOUDTRUTH_ENVIRONMENT: $(echo ${CLOUDTRUTH_ENVIRONMENT} | base64)
  CLOUDTRUTH_PROJECT: $(echo ${CLOUDTRUTH_PROJECT} | base64)
EOF

kubectl get -n argocd deployment/argocd-repo-server -o yaml > argocd-repo-server.original.$(date +%s).yaml
kubectl patch -n argocd deployment/argocd-repo-server --patch "$(cat argocd-repo-server.patch.yaml)"

kubectl get -n argocd configmap/argocd-cm -o yaml > argocd-cm..original.$(date +%s).yaml
kubectl patch -n argocd configmap/argocd-cm --patch "$(cat argocd-cm.patch.yaml)"
