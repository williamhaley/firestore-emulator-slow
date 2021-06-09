# firestore-emulator-slow

## Run

* clone this repo
* `npm i`
* `./node_modules/.bin/firebase emulators:start --import=emulator-data --export-on-exit --debug`
* `FIRESTORE_EMULATOR_HOST=localhost:8080 go run demo.go`
* `FIRESTORE_EMULATOR_HOST=localhost:8080 node demo.js`

## Info

```bash
% firebase --version
9.12.1
```

## How this repo was originally created

* `firebase init`
 * Only added `firestore` and the `emulators`
* Explicitly set up the emulator
    ```
     % firebase setup:emulators:firestore
    i  firestore: downloading cloud-firestore-emulator-v1.12.0.jar...
    ```
