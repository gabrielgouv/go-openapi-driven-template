PACKAGE = openapi
STATIK_PACKAGE = swaggerui
OPENAPI_FILE = api.yaml
SWAGGER_UI_DIR = .swagger-ui

build:
	@echo "-------------------------------------------------------------------------------------"
	@echo "BUILDING SERVICE"
	@echo "-------------------------------------------------------------------------------------"
	@echo "-> Cleaning up generated files..."
	@rm -rf generated
	@echo "-> Installing temporary dependencies..."
	@make install
	@echo "-> Creating generated directories..."
	@mkdir -p generated/$(PACKAGE)
	@mkdir -p generated/$(STATIK_PACKAGE)
	@echo "-> Generating OpenAPI interfaces and types..."
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -generate chi-server -package $(PACKAGE) -o generated/$(PACKAGE)/interfaces.gen.go $(OPENAPI_FILE)
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -generate types -package $(PACKAGE) -o generated/$(PACKAGE)/models.gen.go $(OPENAPI_FILE)
	@echo "-> Copying $(OPENAPI_FILE) file to Swagger UI..."
	@cp $(OPENAPI_FILE) $(SWAGGER_UI_DIR)
	@echo "-> Packaging Swagger UI..."
	@go run github.com/rakyll/statik -src=$(SWAGGER_UI_DIR) -dest=generated -p $(STATIK_PACKAGE)
	@echo "-> Cleaning up Swagger UI..."
	@rm $(SWAGGER_UI_DIR)/$(OPENAPI_FILE)
	@echo "-> Cleaning up temporary dependencies..."
	@go mod tidy
	@echo ""
	@echo "BUILD SUCCEEDED"
	@echo "-------------------------------------------------------------------------------------"
	@echo ""

build-run:
	@make build
	@go build -o generated/service cmd/main.go
	@generated/service

run:
	@go run cmd/main.go

install:
	@go get -d github.com/deepmap/oapi-codegen/pkg/codegen@v1.8.2
	@go get -d github.com/rakyll/statik


