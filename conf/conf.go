package conf

import "fetch/types"

func GetConf() types.Conf {
	conf := types.Conf{
		TransactionsFileName: "../data/transactions.csv",
	}
	return conf
}
