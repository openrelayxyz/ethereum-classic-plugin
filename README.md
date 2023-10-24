# PluGeth Network Plugins

These plugins are designed to implement networks other than the ethereum mainnet and its test nets. 

Please visit our flagship [repository](https://github.com/openrelayxyz/plugeth) for general information about PluGeth. 

### Build and Deploy

Along with our [general utility plugins](https://github.com/openrelayxyz/plugeth-plugins), the plugins in this repository are built by first navigating to the root of the specific plugin and running:

`go build -buildmode=plugin`

This will produce a binary file, `nameOfPlugin.so`. 

PluGeth will then look into a directory at `path-to-ethereum-DataDir/plugins`. If no such directory exists it will need to be made. Move the plugin binary into this directoy and start PluGeth as you would a typical go-etereum node.

Please [contact us](https://docs.plugeth.org/en/latest/contact.html) with any questions or ideas.
