package types

type Transactions struct {
	Payer     string `csv:"payer"`
	Points    int    `csv:"points"`
	Timestamp string `csv:"timestamp"`
}

type Conf struct {
	TransactionsFileName string
}
