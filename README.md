# cryptodev

Cryptocurrencies bootstrap for developers, this tool MUST be used only in development.

Help the cause sending a donation in Bitcoin: `1Jv1tZZAs6kxhAZxWXWJYkdEThqLFnLuYb`

## Commands

### init

Initialize an environment with the currencies that you choose.

Example:

`$ cryptodev init btc ltc doge zec xmr`

This download the tarball with the currency binaries and save the binaries on
`$HOME/.cryptodev/bin/<CURRENCY_SYMBOL>/`

The data directory will be:

`$HOME/.cryptodev/data/<CURRENCY_SYMBOL>/`

#### Currencies available

- btc: Bitcoin
- ltc: Litecoin
- doge: Dogecoin
- zec: Zcash
- xmr: Monero
- ether: Ethereum

### list

List the clusters with all information about each currency daemon.
Info showed:
- Daemon version, id and port.
- RPC user, password, ip and port.

### start, stop, restart

Start, stop and restart a cryptocurrency of the cluster. By example:

`$ cryptodev start|stop|restart btc`

### startall, stopall, restartall

Start, stop and restart all clusters.

### clean

Delete all cluster data and restart the daemon.

### generate

Generate blocks in the test environment. (Only apply for some currencies)

`$ cryptodev generate <AMOUNT_BLOCKS>`

### mine

Like generate but in this case the mining action happen until you choose.

`$ cryptodev mine`

CTRL + C to exit.

### cli

Wrap the cli of the cryptocurrency and fill the RPC data.

`$ cryptodev cli btc <command> <args>`

Example with `getinfo` of Bitcoin.

`$ cryptodev cli btc getinfo`

### tail

Show and follow the debug.log of the picked daemon on the cluster.

`$ cryptodev tail <DAEMON_ID>`

### update

Get the binaries last version of a cryptocurrency. NOTE: This `clean` the daemon
data. 

`$ cryptodev update btc`
