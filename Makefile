run:
	go run ./app/main.go;

tidy:
	./scripts/go_mod_tidy.sh;

cleanup:
	./scripts/cleanup.sh;

modadd:
	./scripts/go_mod_add.sh;

lint:
	./scripts/golangci-lint.sh;

go-upgrade:
	./scripts/go_upgrade.sh;

mod-run:
	./scripts/go_run_parallel.sh;

mod-sync:
	go work sync;


generate:
	cd .proto;
	buf generate proto;


.PHONY: run tidy generate cleanup modadd lint go-upgrade mod-sync