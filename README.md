# Shawarma

## NOTE: Early Experimental Phase

Currently experimenting with different design and implementation to ensure
reliability and customizations

## BYO UI/Secret Engine/Runners

Fully customizable and open source CI/CD with [pickle](https://github.com/apple/pkl)! :D

## Planned Integrations

### Hubs (via Static Plugin)

- [Github](https://github.com/)
- [Forgejo](https://forgejo.org/)
- [Gitlab](https://gitlab.com)
- [Bitbucket](https://bitbucket.org/)

### Auths (via Static Plugin)

- [Vault](https://www.hashicorp.com/products/vault)
- Some LDAP (not sure which one)
- Customizable enough for in-house auths

### Runners (via Static Plugin)

- [SSH](https://man.openbsd.org/ssh)
- Containerizations
    - [Docker](https://www.docker.com/)
    - [Podman](https://podman.io/)
- [Nomad](https://www.nomadproject.io/)
- [Kubernetes](https://kubernetes.io/)

### DBs (via Static Plugin) 

- [Sqlite](https://www.sqlite.org/index.html)
- [Postgres](https://www.postgresql.org/)
