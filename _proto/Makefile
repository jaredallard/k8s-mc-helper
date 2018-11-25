space := $(empty) $(empty)
comma := ,
empty :=

import_path = github.com/jaredallard/k8s-mc-helper/pkg/proto/mchelper

dst 	= ../pkg/proto
target  = go
plugins = grpc

release_ias  = $(subst $(space),$(comma),$(addsuffix =$(import_path)/$(release_pkg),$(addprefix M,$(release_pbs))))
release_pbs  = $(sort $(wildcard mchelper/*.proto))
release_pkg  = mchelper

google_deps	 = Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any

.PHONY: all
all: mchelper

.PHONY: mchelper
mchelper:
	PATH=../bin:$(PATH) protoc --$(target)_out=plugins=$(plugins),$(google_deps),$(chart_ias),$(version_ias):$(dst) $(release_pbs)

.PHONY: clean
clean:
	@rm -rf $(dst)/hapi 2>/dev/null