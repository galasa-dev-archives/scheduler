#
# Copyright contributors to the Galasa project 
#
ARG dockerRepository
FROM ${dockerRepository}/library/open-liberty:22.0.0.1-full-java11-openj9

COPY target/dev.galasa.scheduler.api.war /config/apps
COPY apiServer.xml /config/configDropins/overrides/apiServer.xml
COPY jpaServer.xml /config/configDropins/overrides/jpaServer.xml
RUN  mkdir /config/postgresql
COPY target/liberty/wlp/usr/servers/guideServer/jdbc/postgresql.jar /config/postgresql/
COPY git.commit /config/
