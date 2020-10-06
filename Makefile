# Color
RED		:= $(shell printf "\033[31m")
GREEN	:= $(shell printf "\033[32m")
YELLOW	:= $(shell printf "\033[33m")
BLUE	:= $(shell printf "\033[34m")
BOLD	:= $(shell printf "\033[1m")
RESET	:= $(shell printf "\033[m")

DOCKER_COMPOSE ?= docker-compose

.PHONY: help
help: ## Provides help information on available commands
	@printf "Usage: make <command>\n\n"
	@printf "Commands:\n"
	@awk -F ':(.*)## ' '/^[a-zA-Z0-9%\\\/_.-]+:(.*)##/ { \
	  printf "  \033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)

.PHONY: compose/build
compose/build: ## Build all Docker images of the project
	@$(DOCKER_COMPOSE) build

.PHONY: compose/up
compose/up: ## Start all containers (in the background)
	@$(DOCKER_COMPOSE) up -d 
	@make about 
	@make urls

.PHONY: compose/down
compose/down: ## Stops and deletes containers and networks created by "up".
	@$(DOCKER_COMPOSE) down


.PHONY: compose/purge
compose/purge: ## Stops and deletes containers, volumes, images (all) and networks created by "up".
	@$(DOCKER_COMPOSE) down -v --rmi all

.PHONY: compose/rebuild
compose/rebuild: compose/down compose/build compose/up ## Rebuild the project

.PHONY: about
about:
	@echo "\n"  
	@echo "${BOLD} _____           _                 _       _			"       
	@echo "${BOLD}|  __ \         | |               | |     | |			"      
	@echo "${BOLD}| |  \/ ___  ___| |__   ___  _ __ | | __ _| |_ ___	" 
	@echo "${BOLD}| | __ / _ \/ __|  _ \ / _ \|  _ \| |/ _  | __/ _ \	"
	@echo "${BOLD}| |_\ \  __/ (__| | | | (_) | |_) | | (_| | ||  __/	"
	@echo "${BOLD} \____/\___|\___|_| |_|\___/|  __/|_|\__ _|\__\___|	"
	@echo "${BOLD}                            | |						"                    
	@echo "${BOLD}                            |_|				${RESET}" 
	@echo "\n"  

DOCKER_INPECT_FORMAT__AWK ?= "'\''{{.Name}} : {{range $$p, $$conf := .NetworkSettings.Ports}}{{$$p}} -> {{(index $$conf 0).HostIp}}:{{(index $$conf 0).HostPort}}\t{{end}}'\''"

.PHONY: urls
urls: ## Get project's URL
	@echo "------------------------------------------------------------"
	@echo "${GREEN}You can access your project at the following URLS:${RESET}"
	@echo "------------------------------------------------------------"
	@echo "\n"  
	@$(DOCKER_COMPOSE) ps -q | awk '{ \
		cmd_docker_inspect = sprintf("docker inspect --format=%s %s", ${DOCKER_INPECT_FORMAT__AWK}, $$0) ; \
		cmd_docker_inspect | getline docker_inspect ; close(cmd_docker_inspect) ; \
		gsub(/0.0.0.0/, "http://localhost", docker_inspect) ; \
		split(docker_inspect, urls, "\t") ; \
		printf "%s\n", urls[1] ; \
		i = 2 ; while (i <= length(urls)) { \
		index_tab = index(docker_inspect,":") ; \
		printf "%*s %*s\n", index_tab, "", index_tab, urls[i]; i++ \
		} ; \
	}'

ENV_FILE ?= "APP_URL=http:localhost:1323 \n%s\n
			APP_NAME=Gechoplate \n%s\n
			DB_HOST=db \n%s\n
			DB_PORT=3306 \n%s\n
			DB_NAME=database \n%s\n
			DB_USER=user \n%s\n
			DB_PASSWORD=password"

.PHONY: env
env: ## Generate env file
	@echo "Environment file generation...."
	@touch .env
	