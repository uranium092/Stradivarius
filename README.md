## Stradivarius
System that interacts with stocks with solid and secure filters, sorting and recommendations to invest.

 ### How to run
 * Download Go 1.24.2 and Node 20 - only for dev mode
 * Make sure you have an ENV var defined like ```GO_ENV=DEV``` - only for dev mode
 * You will need .ENV file for security (Send on demand) - only for dev mode
 * Run the backend:  on root ```/backend``` run  ```go mod tidy``` and after, ```go run main.go``` - only for dev mode
 * Run the frontend: on root ```/frontend/stradivarius``` run ```npm run dev``` - only for dev mode
 * If you want to deploy this in AWS and automatically launch a server (EC2 instance), you will need:
   * Download AWS CLI and authenticate
   * Download Terraform and config its PATH
   * You will need private credentials to run this infrastructure. (Send on demand)
   * Finally, just run ```terraform apply -var="CREDENTIAL_1" -var="CREDENTIAL_2"```. Remember that it is needed pass the correct CREDENTIALS in this step, so you will need the private credentials.
   * Now, your infrastructure is being configured, wait a few minutes until it finishes.
   * Recommended: monitor the status of the instance in AWS Console, there you can know when it is running. When the instance is up and running this App (Stradivarius), you can access it from Browser with ```http://your-ip-instance:port_http```; You can know the IP of the     instance from AWS Console, and ```:port_http``` is the value that you set in ```:port_http``` Terraform Var (default :8080). Example URL: ```http://98.81.182.209:8080```

