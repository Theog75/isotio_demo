FROM centos
RUN useradd gouser
RUN mkdir /uploads

FROM scratch
COPY --from=0 /etc/passwd /etc/passwd
COPY --from=0 --chown=1001 /uploads /uploads
USER 1001
ADD  main /
ADD  upload.html /
CMD ["/main"]
