docker:
	docker build -t catboost-serving:latest -f build/Dockerfile .

build:
	cd catboost/catboost/libs/model_interface && ../../../ya make -r --target-platform DEFAULT-LINUX-X86_64 . && \
	cp c_api.h model_calcer_wrapper.h libcatboostmodel.so libcatboostmodel.so.1 wrapped_calcer.h /usr/local/include && \
	cp c_api.h model_calcer_wrapper.h libcatboostmodel.so libcatboostmodel.so.1 wrapped_calcer.h /usr/local/lib && \
	ldconfig && cd ../../../.. && go build . -o build/catboost-serving && go build client/. -o build/benchmark-tool

build_m1:
	cd catboost/catboost/libs/model_interface && ../../../ya make -r --target-platform DEFAULT-DARWIN-ARM64 . && \
    cp c_api.h model_calcer_wrapper.h libcatboostmodel.dylib libcatboostmodel.dylib.1 wrapped_calcer.h /usr/local/include && \
    cp c_api.h model_calcer_wrapper.h libcatboostmodel.dylib libcatboostmodel.dylib.1 wrapped_calcer.h /usr/local/lib && \
    update_dyld_shared_cache && cd ../../../.. && go build . -o build/catboost-serving && go build client/. -o build/benchmark-tool