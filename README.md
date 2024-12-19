Distributed File Storage

Descrição

O projeto Distributed File Storage é uma solução para armazenamento, recuperação e gerenciamento de arquivos distribuídos entre vários servidores ou nós. Foi projetado para oferecer:

Escalabilidade: Suporte ao crescimento e aumento de nós sem perda de desempenho.

Resiliência: Alta disponibilidade e tolerância a falhas.

Eficiência: Otimização de recursos para maior desempenho.

O projeto também inclui funcionalidades de:

P2P (peer-to-peer): Para conexões distribuídas diretas entre nós.

Criptografia: Garantindo segurança dos dados durante o armazenamento e transferência.

Funcionalidades

Armazenamento distribuído: Os arquivos são divididos e armazenados entre diferentes nós.

Recuperação de arquivos: Arquivos podem ser acessados de forma eficiente.

Conexões P2P: Para transferência direta de arquivos entre os nós.

Criptografia de dados: Proporcionando segurança e confidencialidade.

Alta disponibilidade: Tolerância a falhas nos nós.

Estrutura do Projeto

main.go: Ponto de entrada da aplicação.

crypto.go: Implementação das funcionalidades de criptografia.

server.go: Configuração e gerenciamento dos servidores.

store.go: Gerenciamento do armazenamento distribuído.

store_test.go e crypto_test.go: Testes unitários para funcionalidades chave.

Makefile: Automatização de tarefas do projeto.

Arquivos de dependências (go.mod, go.sum): Gerenciamento de pacotes no Go.

Tecnologias Utilizadas

Linguagem de programação: Go (Golang).

Arquitetura P2P para comunicação entre nós.

Criptografia para segurança de dados.

Como Executar

Clone o repositório:

git clone https://github.com/seu_usuario/distributed-file-storage.git
cd distributed-file-storage

Instale as dependências:

go mod tidy

Execute a aplicação:

go run main.go

Testes podem ser executados com:

go test ./...

Licença

Este projeto é licenciado sob a Licença MIT. Veja o arquivo LICENSE para mais detalhes.

Contribuições

Contribuições são bem-vindas! Siga os passos abaixo:

Realize um fork do projeto.

Crie uma branch para a sua funcionalidade ou correção:

git checkout -b minha-nova-feature

Envie um pull request com sua contribuição.
