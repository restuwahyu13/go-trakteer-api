MIG := @migrate
NODEMON := @nodemon
SIGNAL := SIGTERM
DSN := postgres://restuwahyu13:restuwahyu13@localhost:5433/postgres?sslmode=disable
MIGDIR := file:databases/migrations/

start:
	${NODEMON} -v -x go run main.go --signal ${SIGNAL}

migmake:
ifdef name
		${MIG} --verbose create -ext sql -dir databases/migrations ${name}
endif

migup:
		${MIG} -path databases/migrations --verbose up --all

migdown:
		${MIG} -database ${DSN} -source ${MIGDIR} --verbose down --all

migupspec:
ifdef target
		${MIG} -database ${DSN} -source ${MIGDIR} --verbose down ${target}
endif

migdownspec:
ifdef target
		${MIG} -database ${DSN} -source ${MIGDIR} --verbose down ${target}
endif

migdrop:
		${MIG} -database ${DSN} -source ${MIGDIR} --verbose drop -f