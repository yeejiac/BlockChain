package models

type Transaction struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Amounts  string `json:"amounts"`
	Fee      string `json:"fee"`
	Message  string `json:"message"`
}

type Block struct {
	Previous_hash   string `json:"previous_hash"`
	Hash            string `json:"hash"`
	Nonce           string `json:"nonce"`
	Timestamp       string `json:"timestamp"`
	Miner           string `json:"miner"`
	Miner_rewards   string `json:"miner_rewards"`
	Transaction_ary []Transaction
}

type BlockChain struct {
	Block_ary            []Block
	Pending_transactions []Transaction
	Difficulty           string `json:"difficulty"`
	Block_time           string `json:"block_time"`
	Mining_rewards       string `json:"mining_rewards"`
	Block_limitation     string `json:"block_limitation"`
}
