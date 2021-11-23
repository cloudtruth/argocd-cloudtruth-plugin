FROM alpine
ENTRYPOINT ["/argocd-cloudtruth-plugin"]
COPY argocd-cloudtruth-plugin /
