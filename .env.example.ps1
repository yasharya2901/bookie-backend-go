# PowerShell script to set environment variables for Database and Appwrite Webhook Secret

$env:DB_HOST="your_db_host"
$env:DB_PORT="your_db_port"
$env:DB_USER="your_db_user"
$env:DB_PASSWORD="your_db_password"
$env:DB_SSLMODE="your_db_sslmode"
$env:DB_NAME="your_db_name"
$env:DB_TIMEZONE="your_db_timezone"
$env:APPWRITE_WEBHOOK_SECRET="your_secret_value"
$env:PORT="8182" # Change this to your desired port number
$env:GIN_MODE="debug" # Change this to "release" for production