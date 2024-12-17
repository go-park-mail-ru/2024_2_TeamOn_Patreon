# Makefile

# Цель для генерации кода easyJSON
generate:
	find . -name '*.go' -exec grep -q '//easyjson:json' {} \; -exec easyjson -all {} \;