# Terraform Provider Random Joke

Run the following command to build the provider

```shell
go build -o terraform-provider-random-joke
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the configuration.

```shell
terraform init && terraform apply
```