# Setup
```
cp example.config.yaml config.yaml
```
Update config.yaml with your values

```sh
# build and ssh into container
docker-compose up --build && docker-compose run app /bin/bash

# Run
go run invgen.go -month $INVOICE_MONTH
```