- Create two migrations file in `/postgres/migrations/` folder (which look slike this `number_name.up.sql` and `number_name.down.sql`)

```
$> migrate create -ext sql -dir postgres/migrations initial
```


migrate -source file://postgres/migrations -database postgres://localhost:5432/database up 2
