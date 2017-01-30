package pkg

import (
	"path/filepath"
)

type LastPort struct {
	ID		int
	PortNumber	uint16
}

type PortConfig struct {
	Name		string
	PortNumber	[]Port
}

// The minimum numbers of ports for each daemon are 4
// 2 Daemon and 2 RPC
type Port struct{
	Type 		string
	PortNumber	uint16
}

type Cryptocurrency struct {
	ID		int `storm:"id,increment"`
	Name		string `storm:"index"`
	Symbol		string `storm:"unique"`
	FolderBin 	string
	FolderData 	string
	DaemonConfig 	[]PortConfig
}

func NewCryptocurrency (symbol string) *Cryptocurrency{
	var c Cryptocurrency
	err := db.One("Symbol", symbol, &c)
	if err != nil {
		c = Cryptocurrency{}

		c.Name = "Bitcoin"
		c.Symbol = symbol
		c.SetFolders()
		c.SetConfigs()

		db.Save(&c)
	}
	return &c
}

func (c *Cryptocurrency) SetFolders() {
	c.FolderBin = filepath.Join(folderCryptodevBin, c.Symbol)
	c.FolderData = filepath.Join(folderCryptodevData, c.Symbol)

	createDir(c.FolderBin)
	createDir(c.FolderData)
}

func (c *Cryptocurrency) SetConfigs () {
	var lastPort LastPort
	err := db.One("ID", "1", &lastPort)
	if err != nil {
		lastPort = LastPort{ID: 1, PortNumber: 39999}
		db.Save(&lastPort)
	}

	c.DaemonConfig = []PortConfig{
		{Name: "daemon1", PortNumber: []Port{
			{Type: "daemon", PortNumber: lastPort.PortNumber + 1},
			{Type: "rpc", PortNumber: lastPort.PortNumber + 2},
		}},
		{Name: "daemon2", PortNumber: []Port{
			{Type: "daemon", PortNumber: lastPort.PortNumber + 3},
			{Type: "rpc", PortNumber: lastPort.PortNumber + 4},
		}},
	}
	lastPort.PortNumber += 4

	db.Update(&lastPort)
}

func (c *Cryptocurrency) CreateConfig() {

}