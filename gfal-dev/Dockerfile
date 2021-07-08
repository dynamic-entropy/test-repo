FROM gitlab-registry.cern.ch/linuxsupport/cc7-base:latest

RUN yum -y group install "Development Tools"
RUN yum -y install man-pages
RUN yum -y install git-all
RUN yum -y install cmake doxygen glib2-devel libattr-devel openldap-devel zlib-devel json-c-devel
RUN yum -y install lfc-devel dpm-devel srm-ifce-devel dcap-devel globus-gass-copy-devel davix-devel xrootd-client-devel libssh2-devel gtest-devel 
RUN yum -y install e2fsprogs-devel libuuid-devel
RUN yum -y install boost-devel
RUN yum -y install python-devel
RUN yum -y install python-pip
RUN yum -y install pugixml-devel

RUN mkdir /home/dev
WORKDIR /home/dev

#clone the gfal2 repositories
RUN git clone https://github.com/cern-fts/gfal2.git
RUN git clone https://github.com/cern-fts/gfal2-python.git
RUN git clone https://github.com/cern-fts/gfal2-util.git

#build and install gfal2
RUN mkdir gfal2/build && cd gfal2/build && \
    cmake -Wno-dev -DCMAKE_INSTALL_PREFIX=/usr -DSYSCONF_INSTALL_DIR=/etc .. && \
    make && \
    make install

#install gfal-python
RUN cd gfal2-python && python setup.py install

#install gfal-util
RUN cd gfal2-util && python setup.py install

CMD ["/bin/bash"]
