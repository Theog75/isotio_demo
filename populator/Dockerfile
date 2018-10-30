FROM centos
RUN useradd gouser
RUN mkdir /uploads
RUN chown -R gouser /uploads

FROM scratch
COPY --from=0 /etc/passwd /etc/passwd
COPY --from=0 /uploads /uploads
USER gouser
ADD  main /
CMD ["/main"]
