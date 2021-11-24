#!/usr/bin/env sh

# fail fast
set -e
typeset -i remote=0

if [[ "$0" == "sh" ]]; then
  remote=1
  baseurl=${baseurl:-https://raw.githubusercontent.com/cloudtruth/argocd-cloudtruth-plugin/main/install}
  baseurl_params=${baseurl_params:-}
  basedir="."
else
  basedir=$(cd $(dirname $0) && pwd)
fi

function echoFile {
  if ((remote)); then
    curl -s ${baseurl}/${1}${baseurl_params}
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
read CLOUDTRUTH_API_KEY < /dev/tty
if [[ -z ${CLOUDTRUTH_API_KEY} ]]; then echo "Api Key is required"; exit 1; fi

printf "Input Cloudtruth Environment [default]: "
read CLOUDTRUTH_ENVIRONMENT < /dev/tty
if [[ -z ${CLOUDTRUTH_ENVIRONMENT} ]]; then echo "Using 'default' for environment"; CLOUDTRUTH_ENVIRONMENT=default; fi

printf "Input Cloudtruth Project [MyFirstProject]: "
read CLOUDTRUTH_PROJECT < /dev/tty
if [[ -z ${CLOUDTRUTH_PROJECT} ]]; then echo "Using 'MyFirstProject' for project"; CLOUDTRUTH_PROJECT=MyFirstProject; fi

printf "Input Cloudtruth Tag []: "
read CLOUDTRUTH_TAG < /dev/tty

secret_yaml=$(cat <<EOF
$(echoFile argocd-cloudtruth-plugin-secret.yaml)
  CLOUDTRUTH_API_KEY: $(printf ${CLOUDTRUTH_API_KEY} | base64)
  CLOUDTRUTH_ENVIRONMENT: $(printf ${CLOUDTRUTH_ENVIRONMENT} | base64)
  CLOUDTRUTH_PROJECT: $(printf ${CLOUDTRUTH_PROJECT} | base64)
EOF
)

if [[ -n ${CLOUDTRUTH_TAG} ]]; then
secret_yaml=$(cat <<EOF
${secret_yaml}
  CLOUDTRUTH_TAG: $(printf "${CLOUDTRUTH_TAG}" | base64)
EOF
)
fi

echo "${secret_yaml}" | kubectl apply -n ${ARGO_NAMESPACE} -f -

kubectl get -n ${ARGO_NAMESPACE} configmap/argocd-cm -o yaml > ${basedir}/argocd-cm..original.$(date +%s).yaml
kubectl patch -n ${ARGO_NAMESPACE} configmap/argocd-cm --patch "$(echoFile argocd-cm.patch.yaml)"

kubectl get -n ${ARGO_NAMESPACE} deployment/argocd-repo-server -o yaml > ${basedir}/argocd-repo-server.original.$(date +%s).yaml
kubectl patch -n ${ARGO_NAMESPACE} deployment/argocd-repo-server --patch "$(echoFile argocd-repo-server.patch.yaml)"
