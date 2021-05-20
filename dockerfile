FROM golang:1.16

# make work directory
RUN mkdir -p /opt/am \
    && chmod 777 -R /opt/am

# copy service
COPY AmService /opt/am/
RUN chmod 500 /opt/am/*

# change owner
RUN chown root:root -R /opt/am

EXPOSE 8080
WORKDIR /opt/am
CMD ["./AmService"]

