# GO-ECHO-TUTORIAL
Learn Go lang with echo framewrok + dig.In + gomock + mysql

mock generator:
mockgen -destination=mock/mock_employee_repository.go -package=mock GO-ECHO-TUTORIAL/repository EmployeeRepositoryInterface
mockgen -destination=mock/mock_employee_service.go -package=mock GO-ECHO-TUTORIAL/service EmployeeServiceInterface
mockgen -destination=mock/mock_cache_service.go -package=mock GO-ECHO-TUTORIAL/service CacheServiceInterface