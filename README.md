# Kobiton Execute Test Buildkite Plugin

A Buildkite Plugin to (synchronously) execute an automated test script on Kobiton service.

## Example

Add the following to your `pipeline.yml`:

```yml
steps:
  - label: "Kobiton Execute Test"
    plugins:
      - kobiton/kobiton-execute-test#v1.0.0:
           kobi-username: 'your kobiton username'
           kobi-api-key: "your kobiton api key"
           executor-url: 'https://executor-demo.kobiton.com'
           executor-username: 'your kobiton executor server username'
           executor-password: "your kobiton executor server password"
           git-repo-url: 'https://github.com/Nhattd97/azure-devops-sample-java-prod.git'
           git-repo-branch: 'master'
           git-repo-ssh-key: ''
           app-id: 'kobiton-store:91041'
           root-directory: "/"
           command: 'mvn test'
           device-name: 'Galaxy S10'
           device-platform-version: '10'
           use-custom-device: 'false'
           device-platform: 'android'
           wait-for-execution: 'true'
           log-type: 'combined'
```

## Configuration

### `kobiton-username` (Required, string)

Kobiton Username to upload to Kobiton, for example `"kobitonadmin"`.

### `kobi-api-key` (Required, string)

API key to access Kobiton API, for example `"2c8n41e4-b30d-4f19-ba63-6596016c9e58"`.

### `executor-url` (Required, string)

Kobiton Automation Test Executor URL, please contact our Support Team to get this.

### `executor-username` (Required, string)

The Username for Kobiton Automation Test Executor, please contact our Support Team to get this.

### `executor-password` (Required, string)

The Password Kobiton Automation Test Executor, please contact our Support Team to get this.

### `git-repo-url` (Required, string)

Link to your Git repository.

### `git-repo-branch` (Required, string)

The branch of your Git repository you want to execute automation test with.

### `git-repo-ssh-key` (Optional, string)

This is required if your Git Repository is private.

### `kobiton-app-id` (Optional, string)

The App ID or App URL to use in your test script, for example `"kobiton-store:91041"`.

### `root-directory` (Required, string)

Input the root directory of your Git repository, for example `"\"`.

### `command` (Required, string)

Command lines to install dependencies and execute your automation test script. These commands will run from the root directory of your Git repository. For example `"mvn test"`.

### `use-custom-device` (Optional, boolean)

Check if you want to execute one or some test cases with a specific Kobiton Cloud Device. If you already set your device information in your test script, leave this field `false`.

### `device-name` (Optional, string)

This value will be consumed by the `KOBITON_DEVICE_NAME` environment variable in your test script.

### `device-platform` (Optional, string)

This value will be consumed by the `KOBITON_DEVICE_PLATFORM_NAME` environment variable in your test script.

### `device-platform-version` (Optional, string)

This value will be consumed by the `KOBITON_SESSION_PLATFORM_VERSION` environment variable in your test script.

### `wait-for-execution` (Optional, boolean)

Check if your want the release pipeline to wait until your automation testing is completed or failed, then print out the console log and test result.

### `log-type` (Optional, string)

Your desired log type to be showed. Choose `"combined"` to show logs in chronological order, or Separated for single type of log (`"ouput"` or `"error"`).

## Developing

To run the tests:

```shell
docker-compose run --rm tests
```

To validate the `plugin.yml`:
```shell
docker-compose run --rm lint
```

## Contributing

1. Fork the repo
2. Make the changes
3. Run the tests
4. Commit and push your changes
5. Send a pull request
