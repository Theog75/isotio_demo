FROM centos
RUN useradd gouser

FROM scratch
COPY --from=0 /etc/passwd /etc/passwd
USER gouser
ADD  main /
CMD ["/main"]
