# DevBook

Aplicação de rede social completa desenvolvida em Go.

## Sobre o Projeto

DevBook é uma aplicação de rede social criada para fins didáticos e de demonstração, utilizando principalmente a linguagem Go no backend, com HTML, CSS, JavaScript, jQuery e Bootstrap no frontend. O objetivo é permitir aos usuários compartilhar publicações, seguir outros perfis, comentar e interagir, simulando as principais funcionalidades de uma rede social moderna.

## Funcionalidades

- Cadastro e autenticação de usuários
- Criação, edição e exclusão de publicações
- Seguir e deixar de seguir outros usuários
- Curtir publicações
- Perfil de usuário
- Feed de publicações dos usuários seguidos
- API RESTful desenvolvida em Go
- Interface web responsiva com Bootstrap e interatividade com jQuery

## Tecnologias Utilizadas

- **Go** (backend, API REST)
- **HTML** (estrutura das páginas)
- **CSS** (estilização e responsividade)
- **JavaScript** (interatividade no frontend)
- **jQuery** (manipulação dinâmica do DOM e requisições AJAX)
- **Bootstrap** (estilização e responsividade avançada)

## Pré-requisitos

- [Go](https://golang.org/dl/) instalado na máquina (versão recomendada: 1.18+)
- [Git](https://git-scm.com/)
- (Opcional) [Docker](https://www.docker.com/) para ambiente isolado

## Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/marcoantoniogava/DevBook.git
   cd DevBook
   ```

2. Instale as dependências do Go:
   ```bash
   go mod tidy
   ```

3. Garanta que os arquivos do frontend estejam utilizando as bibliotecas jQuery e Bootstrap. Caso utilize CDN, não há necessidade de instalação adicional:
   ```html
   <!-- Exemplo de inclusão via CDN -->
   <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
   <script src="https://code.jquery.com/jquery-3.5.1.min.js"></script>
   <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.min.js"></script>
   ```

4. Configure o banco de dados (caso necessário):
   - O projeto pode utilizar um banco de dados relacional (ex: MySQL ou PostgreSQL).
   - Crie o banco de dados e ajuste as variáveis de ambiente para conexão (`DB_HOST`, `DB_USER`, `DB_PASSWORD`, etc).

## Como Usar

1. Execute a API em um terminal e a aplicação web em outro: <br>
   Terminal 1:
   ```bash
   cd api
   ```
   
   ```bash
   go run main.go
   ```
   Terminal 2:
   ```bash
   cd webapp
   ```
   
   ```bash
   go run main.go
   ```

3. Acesse o sistema pelo navegador:
   ```
   http://localhost:xxxx (Endereço do webapp)
   ```

4. Realize o cadastro ou login e comece a utilizar as funcionalidades da rede social.

## Estrutura Geral do Projeto

- `/models` - Modelos de dados (usuários, publicações, etc)
- `/controllers` - Lógica das rotas e endpoints
- `/views` - Páginas da aplicação
- `/sql` - Scripts de migração e conexão
- `/router` - Definição das rotas da API


## Autor

Desenvolvido por [Marco Antônio de Freitas Gava](https://github.com/marcoantoniogava).
