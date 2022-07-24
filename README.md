# crew-api

Technical test for crew

Some comments:
- The API is running throught Docker
- There is no embedded mongodb instance (no enough time). You must replace DB_URI by a remote mongodb URI (mongodb atlas for example) in run script below
- There is no CI/CD (no enough time)
- I learnt Go for this test, it was very interesting but also challenging

## Run Fetch all talents from existing API
```shell
cd <PROJECT_DIRECTORY>
echo "DB_URI=<DB_URI>" > src/fetch/.env
cd src/fetch/ && go run fetchTalents.go
```

## Run API
```shell
cd <PROJECT_DIRECTORY>
echo "DB_URI=<DB_URI>" > src/fetch/.env
docker build . -t crew-api && ./runLocal.sh
```

## TODO
- CI
- CD
- Embed a mongodb instead of using an external DB URI
- Automated test
- Fill DB automatically by running fetchTalents
- Use docker-composer
- Better manage errors