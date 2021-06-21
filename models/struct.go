package models

type Transaction struct {
	sender   string `json:"sender"`
	receiver string `json:"receiver"`
	amounts  string `json:"amounts"`
	fee      string `json:"fee"`
	message  string `json:"message"`
}

type Block struct {
	previous_hash   string `json:"previous_hash"`
	hash            string `json:"hash"`
	nonce           string `json:"nonce"`
	timestamp       string `json:"timestamp"`
	miner           string `json:"miner"`
	miner_rewards   string `json:"miner_rewards"`
	transaction_ary []Transaction
}

type BlockChain struct {
	block_ary            []Block
	pending_transactions []Transaction
	difficulty           string `json:"difficulty"`
	block_time           string `json:"block_time"`
	mining_rewards       string `json:"mining_rewards"`
	block_limitation     string `json:"block_limitation"`
}
