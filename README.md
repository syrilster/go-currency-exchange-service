# Currency-Exchange-Service

This is an API layer built to connect to the openexchangerates.org and get the latest currency rates. The service will then
perform some internal calculations and return the conversion multiple for fromCurrency to toCurrency conversion.

#Debugging
* Once the app is in Running status (check kubectl get pods), exec (execute a command directly inside the Pod) into 
the Pod to check its environment variables using command. This will give the HOST IP details
    ```
    kubectl exec <pod_name> -- env
    kubectl exec currency-exchange-app-7966496848-c7tf6 -- env
    ```
* Let's access this application from another Pod using the environment variables.
    ```
    kubectl run --rm --generator=run-pod/v1 curl --image=radial/busyboxplus:curl -i --tty
    ```
* Use curl to hit the endpoint and verify the results
    ```
    curl -vvv --request GET 'http://10.0.14.113:8000/v1/currency-exchange/from/EUR/to/INR'
    ```
 * Using the DNS:
    ```
   The FQDN format is <service-name>.<namespace>.<cluster-domain-suffix>
   Ex: curl -vvv --request GET 'http://currency-exchange-service.default.svc.cluster.local:8000/v1/currency-exchange/from/EUR/to/INR'
   ```