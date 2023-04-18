# ardentgrowth-internal-api

API for ardentgrowth-internal

## Setup
- Use `.env.sample` to create `.env`
- Run the following command to create postgres database inside a docker container
    ```
    docker-compose up -d
    ```

## Run

- To build before running, execute the following
    ```
    ./scripts/build.sh
    ```

- Use the following command to run the API on the local machine
    ```
    ./scripts/run.sh
    ```
     This command will automatically run any necessary database migrations before starting the API.
