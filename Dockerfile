FROM centos:7.3.1611

RUN yum update -y && \
    yum install -y openssh-server && \
    yum install -y wget && \
    yum install -y rsyslog && \
    yum install -y make && \
    yum install -y gcc && \
    yum install -y pam-devel

WORKDIR /tmp

RUN wget https://storage.googleapis.com/golang/go1.8.1.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.8.1.linux-amd64.tar.gz && \
    ln -s /usr/local/go/bin/go /usr/sbin/go

RUN /usr/bin/ssh-keygen -A

COPY src/ /tmp/pam-src/

WORKDIR /tmp/pam-src

RUN make -f /tmp/pam-src/Makefile && \
    cp pam_custom.so /lib64/security/pam_custom.so && \
    chmod 755 /lib64/security/pam_custom.so && \
    cp /etc/pam.d/sshd /etc/pam.d/sshd.BAK

COPY pamd/sshd_centos /etc/pam.d/sshd

EXPOSE 22

CMD ["/usr/sbin/sshd", "-D"]
