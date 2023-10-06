export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

# 所有平台
# os-archs=darwin:amd64 darwin:arm64 freebsd:386 freebsd:amd64 linux:386 linux:amd64 linux:arm linux:arm64 windows:386 windows:amd64 windows:arm64 linux:mips64 linux:mips64le linux:mips:softfloat linux:mipsle:softfloat linux:riscv64

# 定义需要编译的目标平台和架构
PLATFORMS := linux_amd64 linux_386 windows_386 windows_amd64

# 定义输出目录
OUTPUT_DIR := release

# 执行编译并保存到指定目录
build:
	@mkdir -p $(OUTPUT_DIR)
	$(foreach platform,$(PLATFORMS),\
		export GOOS=$(word 1,$(subst _, ,$(platform))); \
		export GOARCH=$(word 2,$(subst _, ,$(platform))); \
		CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(OUTPUT_DIR)/$(notdir $(CURDIR))-$(subst _,-,$(platform)); \
	)

.PHONY: build

