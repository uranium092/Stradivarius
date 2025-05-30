## Stradivarius
System that interacts with stocks with solid and secure filters, sorting and recommendations to invest.

 ### How to run
 * Download Go 1.24.2 and Node 20 - only for dev mode
 * Make sure you have an ENV var defined like ```GO_ENV=DEV``` - only for dev mode
 * You will need .ENV file for security (Send on demand) - only for dev mode
 * Run the backend:  on root ```/backend``` run  ```go run main.go``` - only for dev mode
 * Run the frontend: on root ```/frontend/stradivarius``` run ```npm run dev``` - only for dev mode
 * If you want to deploy this in AWS and automatically launch a server (EC2 instance), you will need:
   * Download AWS CLI and authenticate
   * Download Terraform and config its PATH
   * You will need private credentials to run this infrastructure. (Send on demand)
   * Finally, just run ```terraform apply -var"CREDENTIAL_1" -var"CREDENTIAL_2"```. Remember that it is needed pass the correct CREDENTIALS in this step, so you will need the private credentials.
