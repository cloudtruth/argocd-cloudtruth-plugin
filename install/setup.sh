#!/usr/bin/env sh

# fail fast
set -e
remote=0

if test "$0" = "sh"; then
  remote=1
  baseurl=${baseurl:-https://raw.githubusercontent.com/cloudtruth/argocd-cloudtruth-plugin/main/install}
  baseurl_params=${baseurl_params:-}
  basedir="."
else
  basedir=$(cd $(dirname $0) && pwd)
fi

echoFile() {
  if test "$remote" = "1"; then
    curl -s ${baseurl}/${1}${baseurl_params}
  else
    cat ${basedir}/${1}
  fi
}

echo "The requested environment variables control the plugin"
echo "ApiKey is required, and they can all be overriden for each argocd project"
echo

printf "Input Cloudtruth Api Key (required): "
read CLOUDTRUTH_API_KEY < /dev/tty
if test -z ${CLOUDTRUTH_API_KEY}; then echo "Api Key is required"; exit 1; fi

printf "Input Cloudtruth Environment [default]: "
read CLOUDTRUTH_ENVIRONMENT < /dev/tty
if test -z ${CLOUDTRUTH_ENVIRONMENT}; then echo "Using 'default' for environment"; CLOUDTRUTH_ENVIRONMENT=default; fi

printf "Input Cloudtruth Project [MyFirstProject]: "
read CLOUDTRUTH_PROJECT < /dev/tty
if test -z ${CLOUDTRUTH_PROJECT}; then echo "Using 'MyFirstProject' for project"; CLOUDTRUTH_PROJECT=MyFirstProject; fi

printf "Input Cloudtruth Tag []: "
read CLOUDTRUTH_TAG < /dev/tty

printf "Input the namespace that ArgoCD is installed to [argocd]: "
read ARGO_NAMESPACE < /dev/tty
if test -z ${ARGO_NAMESPACE}; then echo "Using 'argocd' for argocd namespace"; ARGO_NAMESPACE=argocd; fi


secret_yaml=$(cat <<EOF
$(echoFile argocd-cloudtruth-plugin-secret.yaml)
  CLOUDTRUTH_API_KEY: $(printf ${CLOUDTRUTH_API_KEY} | base64)
  CLOUDTRUTH_ENVIRONMENT: $(printf ${CLOUDTRUTH_ENVIRONMENT} | base64)
  CLOUDTRUTH_PROJECT: $(printf ${CLOUDTRUTH_PROJECT} | base64)
EOF
)

if test -n "${CLOUDTRUTH_TAG}"; then
secret_yaml=$(cat <<EOF
${secret_yaml}
  CLOUDTRUTH_TAG: $(printf "${CLOUDTRUTH_TAG}" | base64)
EOF
)
fi

echo "${secret_yaml}" | kubectl apply -n ${ARGO_NAMESPACE} -f -

kubectl get -n ${ARGO_NAMESPACE} deployment/argocd-repo-server -o yaml > ${basedir}/argocd-repo-server.original.$(date +%s).yaml
kubectl patch -n ${ARGO_NAMESPACE} deployment/argocd-repo-server --patch "$(echoFile argocd-repo-server.patch.yaml)"

kubectl rollout restart -n ${ARGO_NAMESPACE} deployment/argocd-repo-server
