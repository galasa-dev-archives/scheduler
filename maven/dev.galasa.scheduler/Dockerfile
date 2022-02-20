#
# Copyright contributors to the Galasa project 
#
ARG dockerRepository
FROM ${dockerRepository}/library/openjdk:8

RUN mkdir -p /galasa/cp

COPY target/dev.galasa.scheduler-0.21.0.jar /galasa/cp/
COPY target/deps/* /galasa/cp/

ENTRYPOINT ["java","-cp","/galasa/cp/*","dev.galasa.scheduler.Scheduler"]

WORKDIR /galasa
