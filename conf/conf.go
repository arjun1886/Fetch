package conf

import "fetch/types"

func GetConf() types.Conf {
	conf := types.Conf{
		TransactionsFileName: "transactions.csv",
	}
	return conf
}
