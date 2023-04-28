# This target creates a worker using Skyramp's Helm chart
.PHONY: create_worker 
create_worker:
	helm upgrade --atomic --install skyramp-worker skyramp/worker -n test-thrift-demo --create-namespace

#This target deletes worker using Skyramp's Helm chart
.PHONY: delete_worker
delete_worker: 
	helm uninstall skyramp-worker -n test-thrift-demo

# This target can be used to install helm chart for sample microservices
.PHONY: run_sample_microservices
run_sample_microservices:
	helm upgrade --atomic --install sample-microservices ../demo/sample-microservices -n test-thrift-demo -f ../thrift-demo/helm/values.yaml

#This target uninstalls sample microservices helm chart
.PHONY: uninstall_sample_microservices
uninstall_sample_microservices: 
	helm uninstall sample-microservices -n test-thrift-demo

#This target can be used to run functional test created inside /tests folder
.PHONY: run-functional-test
run-functional-test: 
	skyramp tester start checkout-test -n test-thrift-demo

#This target can be used to run load test created inside /tests folder
.PHONY: run-load-test
run-load-test: 
	skyramp tester start test-load -n test-thrift-demo