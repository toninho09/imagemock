package main

import "fmt"

const htmlMainPage = `
	<!doctype html>
    <html lang="pt-br">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport"
              content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Image Mock</title>
    </head>
	<style>
		body {
			font-family: monospace
		}
		.error {
			font-weight: 600;
			color: red;
			margin: 0 0 25px 0;
		}
		.error:before {
			content: "Erro: ";
		}
		.error:empty {
			display: none;
		}
	</style>
    <body>
		<h1>Image Mock</h1>
		<p>Uma simples aplicação para mock de imagens para seus protótipos web</p>
		<br>
		<p class="error">%s</p>
		<h3>Utilizando</h3>
		<p> <a href="/">/</a>: Esta ajuda. </p>
		<p> <a href="/300">/{tamanho}</a>: Gera uma imagem de {tamanho} x {tamanho} </p>
		<p> <a href="/300x200">/{largura}x{altura}</a>: Gera uma imagem de {largura} x {altura} </p>
		<br>
		<h3>Configurações</h3>
		<p> <a href="/300?t=Exemplo">t={texto}</a>: Altera o texto da imagem para {texto} </p>
		<p> <a href="/300?bg=FF0000">bg={cor}</a>: Altera a cor do fundo para {cor} (Utilize hexadecimal sem #) </p>
		<p> <a href="/300?c=000000">c={cor}</a>: Altera a cor do texto para {cor} (Utilize hexadecimal sem #) </p>
    </body>
    </html>
`

func renderHelp() []byte {
	return []byte(fmt.Sprintf(htmlMainPage, ""))
}

func renderError(message string) []byte {
	return []byte(fmt.Sprintf(htmlMainPage, message))
}
