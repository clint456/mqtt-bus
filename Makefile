# Makefile for mqtt-bus project

# 项目名称和输出二进制文件路径
BINARY_NAME = mqtt-bus
OUTPUT_DIR = mqtt-bus
BINARY_PATH194 = $(OUTPUT_DIR)/$(BINARY_NAME)

# Go 相关命令
GO = go
GO_BUILD = $(GO) build
GO_RUN = $(GO) run
GO_TEST = $(GO) test
GO_CLEAN = $(GO) clean
GO_MOD_TIDY = $(GO) mod tidy

# 编译标志
GO_BUILD_FLAGS = -o $(BINARY_PATH)
GO_BUILD_ENV = CGO_ENABLED=0

# 默认目标
.PHONY: all
all: build

# 编译项目
.PHONY: build
build:
	$(GO_BUILD_ENV) $(GO_BUILD) $(GO_BUILD_FLAGS) $(BINARY_PATH194)
	
# 运行项目
.PHONY: run
run: build
	./$(BINARY_PATH194)

# 测试项目
.PHONY: test
test:
	$(GO_TEST) -v ./...

# 清理构建产物
.PHONY: clean
clean:
	$(GO_CLEAN)
	rm -rf $(BINARY_PATH194)

# 更新 Go 模块依赖
.PHONY: tidy
tidy:
	$(GO_MOD_TIDY)

# 安装依赖
.PHONY: deps
deps:
	$(GO) mod download

# 格式化代码
.PHONY: fmt
fmt:
	$(GO) fmt ./...

# 检查代码规范
.PHONY: lint
lint:
	golangci-lint run ./...

# 帮助信息
.PHONY: help
help:
	@echo "Makefile for mqtt-bus project"
	@echo ""
	@echo "Usage:"
	@echo "  make all        - 编译项目 (默认)"
	@echo "  make build      - 编译项目"
	@echo "  make run        - 编译并运行项目"
	@echo "  make test       - 运行测试"
	@echo "  make clean      - 清理构建产物"
	@echo "  make tidy       - 更新 Go 模块依赖"
	@echo "  make deps       - 下载依赖"
	@echo "  make fmt        - 格式化代码"
	@echo "  make lint       - 检查代码规范"
	@echo "  make help       - 显示此帮助信息"
