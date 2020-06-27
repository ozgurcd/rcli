
TARGET_ARCH=amd64
TARGET_OS=darwin
#TARGET_OS=linux


build:

	GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) go build -o rcli 
