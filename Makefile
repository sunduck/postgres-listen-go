-include .env

gen-migration:
	docker run --rm -it -u `id -u $(USER)` -v $(PWD)/tools:/db -e DBMATE_MIGRATIONS_DIR=/db/migrations amacneil/dbmate new playground-postgres-listen