# **iBeers**

## **Descrição do projeto**

**O projeto tem por objetivo fornecer uma API REST básica para servir de modelo em um portifólio voltado para um SRE.**
  
Ele é baseado no modelo de Clean Architecture e busca utilizar as melhores práticas adotadas pela comunidade.

Conceitos de Engenharia de software também serão empregados.

As entregas ou (deliverys) serão feitas através de HTTP e HTML.

Sendo os endpoints http:

- **/v1/beers**
- **/v1/beers{id}**

## **Bibliotecas**

- **Gorilla Mux**
- **Negroni**
- **client_golang/prometheus**
- **client_golang/prometheus/promauto**
- **client_golang/prometheus/promhttp**

## **Executando localmente**

Para executar localmente é necessário ter os recursos abaixo:
 - Docker

 O Compose irá subir dois containers sendo eles:
  - DB: postgresql
  - APP: iBeers

Execute o comando:
````
docker compose up --build
````

![](https://github.com/oxosix/iBeers/blob/master/utils/ibeers.gif)
