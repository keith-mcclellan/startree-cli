# startree-cli

This is a naive CLI implementation for StarTree using the swagger.json file from the Apache Pinot server

## Installation Instructions
To use, copy releases/startree-arm64 to /usr/local/bin and then `chmod +x /usr/local/bin/startree-arm64`
You will likely want to run `ln -s /usr/local/bin/startree-arm64 /usr/local/bin/startree` to make it easier to run

To avoid having to manually set the hostname and auth token each time, create the following directory `mkdir -p /Users/$USER/.config/startree` then copy the config from `releases/config.yaml` to that new directory

Edit the config.yaml and replace `$HOSTNAME` with the hostname of your cluster (something like pinot.startree.cloud, do not include https://) and replace ${TOKEN} with your API token. Note that this field is encapsulated and requires the auth type prepended, so it should look something like:
        `Authorization: 'Basic OIFDGsjfe24234sdfdfsdv...='`

## Usage instructions

This "should" work identically to the swagger UI that's hosted in the cluster. All functions include any help text that was included in the original swagger.json. 

If something isn't working as expected, the `--debug` flag is useful to see what is actually going on.

### Enable Shell Completion on Mac

To enable shell completion on mac, first you will need to enable Shell Completion in zsh. Run the following command once:

`echo "autoload -U compinit; compinit" >> ~/.zshrc`

After that, you can run the following command to enable shell completion for the current session.

`source <(startree completion zsh)`

If you would like to add this permanently to the system, you can do the following:
```
mkdir -p ${fpath[1]}
chown $USER ${fpath[1]}
sudo startree completion zsh > "${fpath[1]}/_startree"
```

## Building / Development

This CLI was built using the swagger build tools from https://goswagger.io/generate/cli.html

To install swagger, you can download the appropriate binary from https://github.com/go-swagger/go-swagger/releases/ and copy it to `/usr/local/bin`

Make sure your clone of this repo is in your $GOPATH, `cd src`, and then run the following commands:
```
go mod tidy
swagger generate cli -f swagger.json --cli-app-name startree
go build -o cmd/startree/main.go 
```

