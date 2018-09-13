GO = godep go
ENTRY = main.go
RUNTIME_CONFIG_FILE = ./runtime.config
BINARY_NAME = ./binary

prodEnv = production
devEnv = development
# config
devConfig:
	@NODE_ENV=${devEnv} node -e "var cfg=require('config');console.log(JSON.stringify(cfg, null, 2))" > ${RUNTIME_CONFIG_FILE}.json

prodConfig:
	@NODE_ENV=${prodEnv} node -e "var cfg=require('config');console.log(JSON.stringify(cfg, null, 2))" > ${RUNTIME_CONFIG_FILE}.json

# compile

binary: config
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO} build -o ${BINARY_NAME}


# server
development: devConfig 
	${GO} run ${ENTRY} -conf ${RUNTIME_CONFIG_FILE} 
	
production: prodConfig 
	${GO} run ${ENTRY} -conf ${RUNTIME_CONFIG_FILE} env=production


# alias
dev: development
prod: production

# PHONY的作用是防止目录下同文件的出现
.PHONY: devConfig prodConfig static binary package development server dev
