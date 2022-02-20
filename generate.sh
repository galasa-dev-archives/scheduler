openapi generate \
        -i openapi.yaml \
        -o maven/dev.galasa.scheduler.api   \
        -g jaxrs-spec \
        --additional-properties interfaceOnly=true,implFolder=src/main/java,apiPackage=dev.galasa.scheduler.api.openapi,modelPackage=dev.galasa.scheduler.api.openapi.model,generatePom=false,dateLibrary=java8,library=openliberty