migrate-up:
	migrate -database ${DB_URL} -path migrations up
.PHONY:migrate-up

migrate-down:
	migrate -database ${DB_URL} -path migrations down
.PHONY:migrate-down

migrate-force:
	migrate -path migrations -database ${DB_URL} force ${VERSION}
.PHONY:migrate-force

migrate-down-stepback:
	migrate -database ${DB_URL} -path migrations down ${STEPBACK}
.PHONY:migrate-down-stepback

migrate-down-all:
	migrate -database ${DB_URL} -path migrations down -all
.PHONY:migrate-down-all