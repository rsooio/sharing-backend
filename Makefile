api:
	goctl api go -api backend.api -dir . --home .goctl/1.5.0
secret:
	openssl rand -hex 16	

.PHONY: api secret