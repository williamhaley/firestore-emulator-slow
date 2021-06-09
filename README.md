# firestore-emulator-slow

## Run

* clone this repo
* `npm i`
* `GCLOUD_PROJECT=fake-project-id ./node_modules/.bin/firebase emulators:start --import=emulator-data --export-on-exit --debug --project fake-project-id`

### Basic Write (one document)

* `GCLOUD_PROJECT=fake-project-id FIRESTORE_EMULATOR_HOST=localhost:8080 go run demo.go`
* `GCLOUD_PROJECT=fake-project-id FIRESTORE_EMULATOR_HOST=localhost:8080 node demo.js`

_Note: To see a [failure scenariorio](https://github.com/firebase/firebase-tools/issues/3475), mix and match the `GCLOUD_PROJECT` ids. With the consistent ids provided above, writes should work as expected and appear in the UI._

### Batch Write (many documents)

* `NUM_RECORDS=10000000 GCLOUD_PROJECT=fake-project-id FIRESTORE_EMULATOR_HOST=localhost:8080 go run demo.go`

_Note: To see a [failure scenario](), note the write time for batches predictably decrease over time with a large `NUM_RECORDS` value.

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

# GitHub Issue

https://github.com/firebase/firebase-tools/issues/3475
