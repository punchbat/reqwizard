install_dependencies_client:
	cd ./client && npm install

install_dependencies_server:
	cd ./server && go get ./...

start_client:
	cd ./client && export NODE_ENV=development && npm run dev

start_server:
	cd ./server && make start_dev

install_dependencies: install_dependencies_client install_dependencies_server

start: install_dependencies start_client start_server