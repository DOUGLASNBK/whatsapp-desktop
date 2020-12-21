# WhatsApp Desktop
Web view para execução do WhatsApp Web

## Dependência do linux
```sudo apt-get install libgtk-3-dev```
```sudo apt-get install webkit2gtk-4.0```

### Execução embiente desenvolvimento
Para executar deve-se configurar as variáveis de constance, gerando setando url e hash do ID da aplicação
```go run whatsapp.go```

### Gerando build
```env GOOS=linux GOARCH=arm GOARM=5 go build -ldflags "-s -w" whatsapp.go```
```apt-get install upx```
```upx whatsapp```
