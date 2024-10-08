# Card_validator

 A simple card validation service using Luhn algorithm to validate card number.

 Steps to launch:
1. Rename Environment Files:
    Change .example.env to .env in both the configs and deployments directories.
2. Adjust HOST_PORT:
    Open the newly created .env file in both configs and deployments.
    Set HOST_PORT to your desired port number (e.g., HOST_PORT=8080).
3. Run Deployment:
    Navigate to the project root (where your Makefile is located).
    Run the command: **make deploy**
   
All service logs are stored in **logs/app.log**
