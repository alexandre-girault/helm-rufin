# helm-rufin
deploy helm charts with secrets from aws secrets manager

## Install

```sh
helm plugin install https://github.com/alexandre-girault/helm-rufin.git
```

## Usage

Put arn instead of secrets values in helm value files :

```yaml
test:
  token1secret: @secretsmanager/arn:aws:secretsmanager:us-east-1:0123456789:secret:test01-Ftxaat/token1
test2:
  secret: @secretsmanager/arn:aws:secretsmanager:us-east-1:0123456789:secret:test02-Ftxaat
```

Then execute any helm command with the plugin to retrieve secrets and change values

```sh
helm rufin list -f values1.yaml -f values2.yaml 
replacing secret :  @secretsmanager/arn:aws:secretsmanager:us-east-1:0123456789:secret:test01-Ftxaat/token1
replacing secret :  @secretsmanager/arn:aws:secretsmanager:us-east-1:0123456789:secret:test02-Ftxaat
NAME    NAMESPACE       REVISION        UPDATED STATUS  CHART   APP VERSION
```

For each value file containing secrets references, a new value file is created with a "with-secrets-" prefix. It contains the secret values. 

```sh
cat with-secrets-secrets.yaml 
test:
  secret: 'testvalue1'
test2:
  secret2: 'testvalue2'
```

The helm command automatically use the value file created with the plugin.

## Development

### Run locally with test data

```sh
cd src
go run *.go ../testdata/secrets.yaml

̀̀```

### Build

#### Delete binaries

```sh
make clean
```

#### Build binaries for each OS / architecture

```sh
make build
```





