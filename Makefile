help: ## Show this help
	@ echo 'Usage: make <target>'
	@ echo
	@ echo 'Available targets:'
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

get-aoc-cookie:  ## ensures $AOC_SESSION_COOKIE env var is set
	@ if [ -z "$${AOC_SESSION_COOKIE}" ]; then \
		echo "AOC_SESSION_COOKIE is not set. Attempting to load from .env..."; \
		go run ./scripts/cmd/load/main.go; \
	else \
		echo "AOC_SESSION_COOKIE is loaded."; \
	fi

skeleton: ## Make skeleton main(_test).go files, optional: $DAY and $YEAR
	@ if [[ -n $$DAY ]]; then \
		go run ./scripts/cmd/skeleton/main.go -day $(DAY); \
	else \
		go run ./scripts/cmd/skeleton/main.go; \
	fi

input: ## Get input, requires $AOC_SESSION_COOKIE, optional: $DAY and $YEAR
	@ if [[ -n $$DAY ]]; then \
		go run ./scripts/cmd/input/main.go -day $(DAY) -cookie $(AOC_SESSION_COOKIE); \
	else \
		go run ./scripts/cmd/input/main.go -cookie $(AOC_SESSION_COOKIE); \
	fi
