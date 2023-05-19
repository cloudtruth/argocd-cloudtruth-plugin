FROM alpine
ARG CMP_ROOT=/home/argocd/cmp-server

COPY dist/argocd-cloudtruth-plugin /usr/bin/
COPY plugin.yaml ${CMP_ROOT}/config/
