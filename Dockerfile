FROM alpine
ARG CMP_ROOT=/home/argocd/cmp-server

COPY argocd-cloudtruth-plugin /usr/bin/
COPY plugin.yaml ${CMP_ROOT}/config/
