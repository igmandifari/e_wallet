new_migration:
	migrate create -ext sql -dir migration -seq $(name)

migrateup:
	migrate -path migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migration -database "$(DB_URL)" -verbose down 1