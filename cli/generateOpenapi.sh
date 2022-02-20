openapi generate -i ../openapi.yaml                              \
                 -g go                                           \
                 -o pkg/openapi                                  \
                 --additional-properties=packageName=openapi

rm pkg/openapi/go.mod
rm pkg/openapi/go.sum
