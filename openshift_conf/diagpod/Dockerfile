FROM centos
RUN yum update -y
RUN yum install -y telnet
RUN yum install -y nmap-ncat
RUN yum install -y nmap
RUN yum install -y bind-utils
RUN yum install -y curl
RUN yum install -y wget
RUN yum install -y lynx
RUN yum install -y openssh-clients
RUN yum install -y rsync
RUN yum install -y net-tools

RUN useradd gouser

USER gouser
