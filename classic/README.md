## Ethereum Classic Plugin

This plugin is designed to implement the [Ethereum Classic](https://github.com/etclabscore/core-geth) protocol when run against foundation PluGeth nodes. 

To run ETC with PluGeth, download a [current release of PluGeth](https://github.com/openrelayxyz/plugeth/releases), and put the corresponding [Classic plugin](https://github.com/openrelayxyz/plugeth-network-plugins/releases) in your plugins folder (by default, this will be `~/.ethereum/plugins`, but can change with the `--datadir` or `--pluginsdir` flags).

_Note: this project is in active development. At this time the plugin has been tested to function once pointed at database with close to current blockchain data. It is not possible to use the plugin for mining at the moment. Also, at this time, the plugin can only be used for the classic chain only, none of its test nets have been implemented._

