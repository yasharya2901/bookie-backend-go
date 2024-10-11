# Bookie Backend
This is the backend for the Bookie project. It is a RESTful API that provides endpoints for the project [Bookie](https://www.github.com/yasharya2901/bookie).


## Installation
1. Clone the repository
    ```bash
    git clone https://github.com/yasharya2901/bookie-backend-go/
    ```

2. Change the directory
    ```bash
    cd bookie-backend-go
    ```

3. Install the dependencies
    ```bash
    go mod download
    ```

4. Copy the `.env.example.ps1` or `.env.example.sh`

    #### For MacOS/Linux:
    ```bash
    cp .env.example.sh .env.sh
    ```

    #### For Windows:
    ```powershell
    cp .env.example.ps1 .env.ps1
    ```

5. Update the environment variables and run the script
    #### For MacOS/Linux:
    ```bash
    source .env.sh
    ```

    #### For Windows:
    ```powershell
    .\.env.ps1
    ```

6. Run the server
    ```bash
    go run main.go
    ```
    The server will start on `localhost:8080` unless you have changed the `PORT` in the environment variables.
