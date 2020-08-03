# ts-wizard

Este é um projeto pessoal em Golang. Para utilizar isso, faça:

1. Instale [golang](https://golang.org) no seu computador;
2. Execute no terminal, prompt, tanto faz:

```bash
go get github.com/ddspog/ts-wizard
```

3. Ainda no terminal, use cd para entrar na sua $GOPATH (pasta onde os projetos Golang são salvos)
4. Execute:

```bash
cd src/github.com/ddspog/ts-wizard
go build .
```

5. Após isso, o executável do projeto estará pronto.
  a) Em Windows:
  ```bash
  ts-wizard.exe filter --null-packets 10 tsdeorigem.ts novotsdedestino.ts
  ```
  b) Em UNIX:
  ```bash
  ./ts-wizard filter --null-packets 10 tsdeorigem.ts novotsdedestino.ts
  ```
