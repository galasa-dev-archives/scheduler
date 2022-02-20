#
# Copyright contributors to the Galasa project 
#
ARG dockerRepository
FROM ${dockerRepository}/library/open-liberty:22.0.0.1-full-java11-openj9

COPY target/dev.galasa.scheduler.api.war /config/apps
COPY apiServer.xml /config/configDropins/overrides/apiServer.xml
COPY git.commit /config/
