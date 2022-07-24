# crew-api

Technical test for crew

## Run Fetch all talents from existing API
cd <PROJECT_DIRECTORY>
echo "DB_URI=<DB_URI>" > src/fetch/.env
cd src/fetch/ && go run fetchTalents.go

## Run API
cd <PROJECT_DIRECTORY>
echo "DB_URI=<DB_URI>" > src/fetch/.env
docker build . -t crew-api && ./runLocal.sh

## TODO
- CI
- CD
- Automated test
- Fill DB automatically by running fetchTalents
- Better manage errors
- Use docker-composer
- Integrate a mongodb instead of using an external DB URI