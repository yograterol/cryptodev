package pkg

func BitcoinDownload() {
	c := NewCryptocurrency("btc")
	source := downloadTarball("https://bitcoin.org/bin/bitcoin-core-0.13.2/bitcoin-0.13.2-x86_64-linux-gnu.tar.gz", downloadDest)
	ungzip(source, c.FolderBin)
}