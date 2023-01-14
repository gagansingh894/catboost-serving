FROM golang:1.19.3 as build

WORKDIR /build

ADD . .

RUN cp build/shared/linux-x86_64/* /usr/local/include && \
    cp build/shared/linux-x86_64/* /usr/local/lib
RUN CGO_ENABLED=1 GOOS=linux go build -o catboost-serving

FROM ubuntu:20.04

COPY --from=build /build/build/shared/linux-x86_64/* usr/local/include/
COPY --from=build /build/build/shared/linux-x86_64/* usr/local/lib/
COPY --from=build /build/catboost-serving catboost-serving
COPY --from=build /build/artefacts/* artefacts/

RUN ldconfig

CMD ["./catboost-serving"]