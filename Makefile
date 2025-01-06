# 远程仓库
REMOTE_REPOSITORY = registry.cn-hangzhou.aliyuncs.com/schwarzsail/docker

# 项目 MODULE 名
MODULE = github.com/schwarzsail/tiktok

SERVICES := api user video

DIR = $(shell pwd)
CMD = $(DIR)/cmd
IDL_PATH = $(DIR)/pkg/idl

# 使用 kitex 更新和生成 Kitex (微服务)代码
.PHONY: rpc-gen-%:
rpc-gen-%:
	@echo "Generating rpc service code for $*..."
	kitex -module "$(MODULE)" -gen-path ./pkg/kitex_gen $(IDL_PATH)/$*.thrift
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run --config=./.golangci.yml
