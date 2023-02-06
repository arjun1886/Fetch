package conf

import "fetch/types"

// GetConf - Returns the endpoints relevant to the app, in this case the input csv file
func GetConf() types.Conf {
	conf := types.Conf{
		TransactionsFileName: "../data/transactions.csv",
	}
	return conf
}
