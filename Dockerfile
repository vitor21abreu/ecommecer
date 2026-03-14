# Imagem base Go
FROM golang:1.26-alpine

# Diretório de trabalho dentro do container
WORKDIR /app

# Copia todos os arquivos do projeto
COPY . .

# Baixa as dependências
RUN go mod download

# Compila o binário do projeto modular
RUN go build -o main ./cmd

# Expõe a porta do Gin
EXPOSE 8080

# Comando para iniciar o app
CMD ["./main"]