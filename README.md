# Setup
```
cp example.config.yaml config.yaml
```
Update config.yaml with your values

```sh
# ssh into container
docker-compose run app /bin/bash

# Run
go run invgen.go -month $INVOICE_MONTH
```