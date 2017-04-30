FROM scratch

ADD mock-esi /

ENTRYPOINT ["/mock-esi"]

EXPOSE 8080