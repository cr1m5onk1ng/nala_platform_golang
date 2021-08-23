# Nala Platform

This repository contains a prototype version of a platform aimed to make finding and discussing good web content for language learning easier and more fun.

The platform is still in its early stages. For now, only part of the backend is functional and (partially) tested.

All the code for the backend makes use of the awesome [Go](https://golang.org/) programming language.

<a href="https://golang.org/"><img src="https://user-images.githubusercontent.com/3613230/41752586-476b0b24-7596-11e8-95fe-8fd3faa21e8a.png" width="200"/></a>
<br>

Since I'm already somewhat familiar with Express.js, I opted to use the lightning fast [fiber](https://github.com/gofiber/fiber) framework for the REST endpoints, since its workflow is very similar to that of Express.js.

The prototype makes use of Postgres as main database solution, although the actual platform will use a mix of SQL for relational data and noSQL for metadata

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
