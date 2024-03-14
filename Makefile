run:
	go run ./app/main.go;

mod-tidy:
	./scripts/go_mod_tidy.sh;

cleanup:
	./scripts/cleanup.sh;

mod-add:
	./scripts/go_mod_add.sh;

mod-lint:
	./scripts/golangci-lint.sh;

go-upgrade:
	./scripts/upgrade_go_mods.sh

mod-run:
	./scripts/go_run_parallel.sh;

mod-sync:
	go work sync;


generate:
	cd .proto;
	buf generate proto;

service-run:
	cd ./beservice && ./service_cmd.sh;

feature-run:
	cd ./features && ./features_cmd.sh;

mirror-folder:
	./scripts/mirror_folder.sh beservice api/protobuf;
	./scripts/mirror_folder.sh features api/protobuf;


.PHONY: run tidy generate cleanup modadd lint go-upgrade mod-sync mod-run service-run feature-run mirror-folder