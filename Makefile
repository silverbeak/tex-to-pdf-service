generate-proto:
	rm -r com_sootsafe_latex
	mkdir com_sootsafe_latex
	protoc -I=./proto/ --go_out=plugins=grpc:com_sootsafe_latex proto/*.proto

docker-build:
	docker build -t eu.gcr.io/sootsafe-app-test/tex-to-pdf-service:latest .

docker-run:
	docker rm -f tex-to-pdf-service
	docker run -d -p 50051:50051 --name tex-to-pdf-service eu.gcr.io/sootsafe-app-test/tex-to-pdf-service:latest
