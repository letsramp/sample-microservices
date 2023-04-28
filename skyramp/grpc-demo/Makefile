# This target creates a worker using Skyramp's Helm chart
.PHONY: create_worker 
create_worker:
	helm upgrade --atomic --install skyramp-worker skyramp/worker -n test-grpc-demo --create-namespace

#This target deletes worker using Skyramp's Helm chart
.PHONY: delete_worker
delete_worker: 
	helm uninstall skyramp-worker -n test-grpc-demo

# This target can be used to install helm chart for sample microservices
.PHONY: run_sample_microservices
run_sample_microservices:
	helm upgrade --atomic --install sample-microservices ../demo/sample-microservices -n test-grpc-demo -f ../grpc-demo/helm/values.yaml

#This target uninstalls sample microservices helm chart
.PHONY: uninstall_sample_microservices
uninstall_sample_microservices: 
	helm uninstall sample-microservices -n test-grpc-demo

#This target can be used to run functional test created inside /tests folder
.PHONY: run-functional-test
run-functional-test: 
	skyramp tester start checkout-test -n test-grpc-demo

#This target can be used to run load test created inside /tests folder
.PHONY: run-load-test
run-load-test: 
	skyramp tester start test-load -n test-grpc-demo

#This target creates mock service for Payment service
.PHONY: mock-payment-service
mock-payment-service: 
	skyramp mocker create --input ./pb/demo.proto --kubernetes-service payment-service \
	--port 50051 --proto-service PaymentService  --protocol protobuf  

#This target can be used to deploy mock service to the worker 
.PHONY: deploy-mock-services
deploy-mock-services:
	skyramp mocker deploy -n test-grpc-demo

#This target can be used to run functional test created inside /tests folder
.PHONY: run-functional-test-with-mocked-payment-service
run-functional-test-with-mocked-payment-service:
	skyramp up checkout-system -n test-grpc-demo &&\
	skyramp tester start checkout-test-with-mocks -n test-grpc-demo