<!DOCTYPE html>
<html>
  
  <body>
    <h1>Bot para Discord em Golang</h1>
    <p>Este é um bot para Discord desenvolvido em Golang usando a biblioteca <a href="https://github.com/bwmarrin/discordgo">discordgo</a>.</p>
    <p>O bot possui as seguintes funcionalidades:</p>
    <ul>
      <li>Responder a mensagens do usuário com uma mensagem pré-definido</li>
      <li><em>Em breve mais funcionalidades...</em></li>
    </ul>
<h2>Pré-requisitos</h2>
<ul>
  <li><a href="https://golang.org/dl/">Go</a> instalado na sua máquina</li>
  <li>Uma conta no <a href="https://discord.com/">Discord</a></li>
</ul>

<h2>Configurando o bot</h2>
<ol>
  <li>Crie um aplicativo no Discord Developer Portal e adicione um bot ao aplicativo</li>
  <li>Copie o token do bot e salve em uma variável de ambiente chamada <code>DISCORD_TOKEN</code></li>
  <li>Conceda as permissões necessárias ao bot para que ele possa se conectar ao servidor do Discord e ler e enviar mensagens</li>
  <li>Adicione o bot a um servidor do Discord</li>
</ol>

<h2>Executando o bot</h2>
<ol>
  <li>Clone este repositório para a sua máquina</li>
  <li>No terminal, vá para o diretório do projeto e execute o comando <code>go build</code></li>
  <li>Execute o binário gerado com o comando <code>go run main.go</code></li>
</ol>

<h2>Uso</h2>
<p>O bot responderá automaticamente as palavras pré-definidas. Atualmente, os seguintes comandos estão disponíveis:</p>
    <ul>
      <li>java</li>
      <li>node</li>
      <li>python</li>
      <li>c++</li>
      <li>c#</li>
      <li>c</li>
      <li>javascript</li>
      <li>matrix</li>
    </ul>
    <p>Se a mensagem contém uma dessas palavras, o bot responderá com uma mensagem.</p>
<h2>Contribuição</h2>
<p>Se você quiser contribuir com o bot, sinta-se à vontade para abrir um pull request.</p>
  </body>
</html>
