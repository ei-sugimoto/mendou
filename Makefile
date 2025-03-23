.PHONY inspector:

inspector:
	@echo "inspector..... 🚀"
	@go build .
	@npx @modelcontextprotocol/inspector -e PATH=${PWD}/tmp main