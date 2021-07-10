# Nala Platform 

This repository contains a prototype version of a platform aimed to make finding and discussing good web resources for language learning easier and more fun.

The platform is still in its early stages. For now, only part of the backend is functional and (partially) tested.

All the code for the backend makes use of the awesome Go programming language [Go](https://golang.org/).

<a href="https://golang.org/"><img src="https://mpng.subpng.com/20180430/clw/kisspng-go-programming-language-computer-programming-progr-programming-language-5ae6e800efbb03.589382971525082112982.jpg" width="200"/></a>
<br>

Since I'm already somewhat familiar with Express.js, I opted to use the lightning fast [fiber](https://github.com/gofiber/fiber) framework for the REST endpoints, since its workflow is very similar to that of Express.js.

The prototype makes use of good old reliable Postgres, although the actual platform will use a mix of SQL for relational data and noSQL for metadata

<a href="https://www.postgresql.org/"><img src="https://icon2.cleanpng.com/20180315/ifq/kisspng-postgresql-logo-computer-software-database-open-source-vector-images-5aaa26e1a38cf4.7370214515211005136699.jpg" width="200" height="200"/></a>
<br>

## TODO

- [x] Relational database schema definition
- [x] Routes for users handling
- [x] Endpoints for core use cases
- [ ] Endpoints for edge use cases
- [ ] Endpoints for users interactions
- [ ] Testing all routes
- [ ] Discussions functionalities
- [ ] React.js frontend prototype
- [ ] Search engine service implementation


### Resources that helped me get started
[golang with postgres](https://dev.to/techschoolguru/series/7172)

[fiber REST API](https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j#create-validators-for-a-model-fields) 

## Author

**Mirco Cardinale**

