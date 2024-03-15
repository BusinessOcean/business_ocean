# Target: run
# Description: Runs the main.go file using the 'go run' command.
run:
	go run ./app/main.go;

# Target: mod-tidy
# Description: Runs the go_mod_tidy.sh script to tidy up the Go module.
mod-tidy:
	./scripts/go_mod_tidy.sh;

# Target: cleanup
# Description: Runs the cleanup.sh script to perform cleanup tasks.
cleanup:
	./scripts/cleanup.sh;

# Target: mod-add
# Description: Runs the go_mod_add.sh script to add a new dependency to the Go module.
mod-add:
	./scripts/go_mod_add.sh;

# Target: mod-lint
# Description: Runs the golangci-lint.sh script to perform linting on the Go module.
mod-lint:
	./scripts/golangci-lint.sh;

# Target: go-upgrade
# Description: Runs the upgrade_go_mods.sh script to upgrade the Go module dependencies.
go-upgrade:
	./scripts/upgrade_go_mods.sh

# Target: mod-run
# Description: Runs the go_run_parallel.sh script to run the Go module in parallel.
mod-run:
	./scripts/go_run_parallel.sh;

# Target: mod-sync
# Description: Runs the 'go work sync' command to synchronize the Go module.
mod-sync:
	go work sync;

# Target: proto-gen
# Description: Generates API files using the gen-api-folder.sh script and buf generate command.
proto-gen:
	./scripts/gen-api-folder.sh beservice api/protobuf;
	./scripts/gen-api-folder.sh features api/protobuf;
	buf generate;

# Target: service-run
# Description: Changes directory to ./beservice and runs the service_cmd.sh script.
service-run:
	cd ./beservice && ./service_cmd.sh;

# Target: feature-run
# Description: Changes directory to ./features and runs the features_cmd.sh script.
feature-run:
	cd ./features && ./features_cmd.sh;

# Target: mirror-folder
# Description: Runs the mirror_folder.sh script to mirror folders for beservice and features.
# mirror-folder:
# 	./scripts/mirror_folder.sh beservice api/protobuf;
# 	./scripts/mirror_folder.sh features api/protobuf;

# PHONY targets
.PHONY: run tidy generate cleanup modadd lint go-upgrade mod-sync mod-run service-run feature-run mirror-folder