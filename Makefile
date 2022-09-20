MIG := @migrate
NODEMON := @nodemon
SIGNAL := SIGTERM
DSN := postgres://restuwahyu13:restuwahyu13@localhost:5432/postgres?sslmode=disable
MIGDIR := databases/migrations/

dev:
	${NODEMON} -V -x go run --race *.go --signal ${SIGNAL}

build:
	go build -v -race main.go

migmake:
ifdef name
		${MIG} -verbose create -ext sql -dir ${MIGDIR} ${name}
endif

migup:
		${MIG} -database ${DSN} -path ${MIGDIR} -verbose up

migupf:
ifdef id
		${MIG} -database ${DSN} -path ${MIGDIR} -verbose force ${id}
endif

migdown:
		${MIG} -database ${DSN} -path ${MIGDIR} -verbose down -all

migupspec:
ifdef target
		${MIG} -database ${DSN} -path ${MIGDIR} -verbose down ${target}
endif

migdownspec:
ifdef target
		${MIG} -database ${DSN} -path ${MIGDIR} -verbose down ${target}
endif

migdrop:
		${MIG} -database ${DSN} -path ${MIGDIR} -verbose drop -f