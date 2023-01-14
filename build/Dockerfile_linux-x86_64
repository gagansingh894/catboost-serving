FROM golang:1.19.3 as build

RUN git clone https://github.com/catboost/catboost.git --branch v1.1.1 && \
    cd catboost/catboost/libs/model_interface && ../../../ya make -r --target-platform DEFAULT-LINUX-X86_64 . && \
    cp c_api.h model_calcer_wrapper.h libcatboostmodel.so libcatboostmodel.so.1 wrapped_calcer.h /usr/local/include && \
    cp c_api.h model_calcer_wrapper.h libcatboostmodel.so libcatboostmodel.so.1 wrapped_calcer.h /usr/local/lib && \
    ldconfig

WORKDIR /build

ADD . .

RUN CGO_ENABLED=1 GOOS=linux go build -o catboost-serving

FROM ubuntu:20.04

COPY --from=build usr/local/lib/c_api.h usr/local/lib/
COPY --from=build usr/local/lib/model_calcer_wrapper.h usr/local/lib/
COPY --from=build usr/local/lib/wrapped_calcer.h usr/local/lib/
COPY --from=build usr/local/lib/libcatboostmodel.so usr/local/lib/
COPY --from=build usr/local/lib/libcatboostmodel.so.1 usr/local/lib/

COPY --from=build usr/local/include/c_api.h usr/local/include/
COPY --from=build usr/local/include/model_calcer_wrapper.h usr/local/include/
COPY --from=build usr/local/include/wrapped_calcer.h usr/local/include/
COPY --from=build usr/local/include/libcatboostmodel.so usr/local/include/
COPY --from=build usr/local/include/libcatboostmodel.so.1 usr/local/include/

COPY --from=build /build/catboost-serving catboost-serving
COPY --from=build /build/artefacts/* artefacts/

RUN ldconfig

CMD ["./catboost-serving"]