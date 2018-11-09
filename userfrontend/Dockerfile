FROM centos
RUN useradd gouser
RUN mkdir /templates
RUN mkdir /img
ADD  logo.png /
ADD  bg.png /
ADD  fe1.tmpl /
RUN chmod 777 logo.png
RUN chmod 777 bg.png
RUN chmod 777 -R /templates
RUN chmod 777 -R /img
RUN chmod 777 fe1.tmpl

FROM scratch
COPY --from=0 /etc/passwd /etc/passwd
COPY --from=0 --chown=1001 /templates /templates
COPY --from=0 --chown=1001 /img /img
COPY --from=0 --chown=1001 /logo.png /img/logo.png
COPY --from=0 --chown=1001 /bg.png /img/bg.png
COPY --from=0 --chown=1001 /fe1.tmpl /templates/fe1.tmpl
USER 1001
ADD  main /



CMD ["/main"]
