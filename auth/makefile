documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

create_auth_mocks:
	mockery --dir=domain --name=AuthRepository --filename=auth_repository_mock.go --output=domain/mocks --outpkg=auth
	mockery --dir=domain --name=AuthUseCase --filename=auth_usecase_mock.go --output=domain/mocks --outpkg=auth
