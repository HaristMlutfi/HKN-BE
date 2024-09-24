pkgs      = $(shell go list ./... | grep -v /tests | grep -v /vendor/ | grep -v /common/)
datetime	= $(shell date +%s)

build:
	@echo "Building HKN-BE Project"
	@gox -os="linux" -arch="amd64" -output="hkn-be"

deploy-dev:build
	rsync -a hkn-be admin@103.217.144.72:/home/admin/hkn-be/hkn-be-$(datetime) -v --stats --progress
	rsync -a templates admin@103.217.144.72:/home/admin/hkn-be -v --stats --progress
	ssh admin@103.217.144.72 "cd /home/admin/hkn-be && sudo service hkn-be stop && sudo unlink hkn-be && sudo ln -s hkn-be-$(datetime) hkn-be && sudo service hkn-be start"




