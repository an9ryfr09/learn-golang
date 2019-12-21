FROM centos

LABEL name="an9ryfr09@gmail.com" \
version=1.0

WORKDIR /go

ADD ./testgo ./

ENTRYPOINT [ "./testgo" ]