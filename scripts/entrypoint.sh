#!/bin/bash

# Leia o arquivo de configuração e defina as variáveis de ambiente
source config.sh

# Execute o programa Go
go run ../cmd/ibeers/main.go
