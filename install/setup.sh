#!/usr/bin/env bash

# fail fast
set -e
typeset -i remote=0

if [[ "$0" == "sh" ]]; then
  remote=1
  baseurl="https://raw.githubusercontent.com/cloudtruth/argocd-cloudtruth-plugin/main/install"
  basedir="."
else
  basedir=$(cd $(dirname $0) && pwd)
fi

function echoFile {
  if ((remote)); then
    curl -s ${baseurl}/${1}
  else
    cat ${basedir}/${1}
  fi
}
export -f echoFile

ARGO_NAMESPACE=${ARGO_NAMESPACE:-argocd}
echo "Installing to namespace '${ARGO_NAMESPACE}', override with: ARGO_NAMESPACE=xyz setup.sh"
echo
echo "The requested environment variables control the plugin"
echo "ApiKey is required, and they can all be overriden for each argocd project"
echo

printf "Input Cloudtruth Api Key (required): "
read CLOUDTRUTH_API_KEY
if [[ -z ${CLOUDTRUTH_API_KEY} ]]; then echo "Api Key is required"; exit 1; fi

printf "Input Cloudtruth Environment [default]: "
read CLOUDTRUTH_ENVIRONMENT
if [[ -z ${CLOUDTRUTH_ENVIRONMENT} ]]; then echo "Using 'default' for environment"; CLOUDTRUTH_ENVIRONMENT=default; fi

printf "Input Cloudtruth Project [MyFirstProject]: "
read CLOUDTRUTH_PROJECT
if [[ -z ${CLOUDTRUTH_PROJECT} ]]; then echo "Using 'MyFirstProject' for project"; CLOUDTRUTH_PROJECT=MyFirstProject; fi

printf "Input Cloudtruth Tag []: "
read CLOUDTRUTH_TAG

secret_yaml=$(cat <<EOF
$(echoFile argocd-cloudtruth-plugin-secret.yaml)
  CLOUDTRUTH_API_KEY: $(echo -n ${CLOUDTRUTH_API_KEY} | base64)
  CLOUDTRUTH_ENVIRONMENT: $(echo -n ${CLOUDTRUTH_ENVIRONMENT} | base64)
  CLOUDTRUTH_PROJECT: $(echo -n ${CLOUDTRUTH_PROJECT} | base64)
EOF
)

if [[ -n ${CLOUDTRUTH_TAG} ]]; then
secret_yaml=$(cat <<EOF
${secret_yaml}
  CLOUDTRUTH_TAG: $(echo -n "${CLOUDTRUTH_TAG}" | base64)
EOF
)
fi

echo "${secret_yaml}" | kubectl apply -n ${ARGO_NAMESPACE} -f -

kubectl get -n ${ARGO_NAMESPACE} configmap/argocd-cm -o yaml > ${basedir}/argocd-cm..original.$(date +%s).yaml
kubectl patch -n ${ARGO_NAMESPACE} configmap/argocd-cm --patch "$(echoFile argocd-cm.patch.yaml)"

kubectl get -n ${ARGO_NAMESPACE} deployment/argocd-repo-server -o yaml > ${basedir}/argocd-repo-server.original.$(date +%s).yaml
kubectl patch -n ${ARGO_NAMESPACE} deployment/argocd-repo-server --patch "$(echoFile argocd-repo-server.patch.yaml)"
