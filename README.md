# go-simple-bank
under development

add mockgen to path

sudo nano ~/.bashrc
```
export PATH=$PATH:~/go/bin
```

Create new db migration:
```bash
migrate create -ext sql -dir db/migration -seq <migration_name>
```