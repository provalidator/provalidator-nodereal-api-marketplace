package model

import "time"

type RpcError struct {
	Error struct {
		Code    int    `json:"code"`
		Data    string `json:"data"`
		Message string `json:"message"`
	} `json:"error"`
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

type RpcValidators struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockHeight string `json:"block_height"`
		Validators  []struct {
			Address string `json:"address"`
			PubKey  struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"pub_key"`
			VotingPower      string `json:"voting_power"`
			ProposerPriority string `json:"proposer_priority"`
		} `json:"validators"`
		Count string `json:"count"`
		Total string `json:"total"`
	} `json:"result"`
}

type RpcTx struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Hash     string `json:"hash"`
		Height   string `json:"height"`
		Index    int    `json:"index"`
		TxResult struct {
			Code      int    `json:"code"`
			Data      string `json:"data"`
			Log       string `json:"log"`
			Info      string `json:"info"`
			GasWanted string `json:"gas_wanted"`
			GasUsed   string `json:"gas_used"`
			Events    []struct {
				Type       string `json:"type"`
				Attributes []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
					Index bool   `json:"index"`
				} `json:"attributes"`
			} `json:"events"`
			Codespace string `json:"codespace"`
		} `json:"tx_result"`
		Tx string `json:"tx"`
	} `json:"result"`
}

type RpcNetInfo struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Listening bool     `json:"listening"`
		Listeners []string `json:"listeners"`
		NPeers    string   `json:"n_peers"`
		Peers     []struct {
			NodeInfo struct {
				ProtocolVersion struct {
					P2P   string `json:"p2p"`
					Block string `json:"block"`
					App   string `json:"app"`
				} `json:"protocol_version"`
				ID         string `json:"id"`
				ListenAddr string `json:"listen_addr"`
				Network    string `json:"network"`
				Version    string `json:"version"`
				Channels   string `json:"channels"`
				Moniker    string `json:"moniker"`
				Other      struct {
					TxIndex    string `json:"tx_index"`
					RPCAddress string `json:"rpc_address"`
				} `json:"other"`
			} `json:"node_info"`
			IsOutbound       bool `json:"is_outbound"`
			ConnectionStatus struct {
				Duration    string `json:"Duration"`
				SendMonitor struct {
					Start    time.Time `json:"Start"`
					Bytes    string    `json:"Bytes"`
					Samples  string    `json:"Samples"`
					InstRate string    `json:"InstRate"`
					CurRate  string    `json:"CurRate"`
					AvgRate  string    `json:"AvgRate"`
					PeakRate string    `json:"PeakRate"`
					BytesRem string    `json:"BytesRem"`
					Duration string    `json:"Duration"`
					Idle     string    `json:"Idle"`
					TimeRem  string    `json:"TimeRem"`
					Progress int       `json:"Progress"`
					Active   bool      `json:"Active"`
				} `json:"SendMonitor"`
				RecvMonitor struct {
					Start    time.Time `json:"Start"`
					Bytes    string    `json:"Bytes"`
					Samples  string    `json:"Samples"`
					InstRate string    `json:"InstRate"`
					CurRate  string    `json:"CurRate"`
					AvgRate  string    `json:"AvgRate"`
					PeakRate string    `json:"PeakRate"`
					BytesRem string    `json:"BytesRem"`
					Duration string    `json:"Duration"`
					Idle     string    `json:"Idle"`
					TimeRem  string    `json:"TimeRem"`
					Progress int       `json:"Progress"`
					Active   bool      `json:"Active"`
				} `json:"RecvMonitor"`
				Channels []struct {
					ID                int    `json:"ID"`
					SendQueueCapacity string `json:"SendQueueCapacity"`
					SendQueueSize     string `json:"SendQueueSize"`
					Priority          string `json:"Priority"`
					RecentlySent      string `json:"RecentlySent"`
				} `json:"Channels"`
			} `json:"connection_status"`
			RemoteIP string `json:"remote_ip"`
		} `json:"peers"`
	} `json:"result"`
}

type RpcAbciInfo struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Response struct {
			Data             string `json:"data"`
			Version          string `json:"version"`
			LastBlockHeight  string `json:"last_block_height"`
			LastBlockAppHash string `json:"last_block_app_hash"`
		} `json:"response"`
	} `json:"result"`
}

type RpcHealth struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
	} `json:"result"`
}

type RpcBlock struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Block struct {
			Data struct {
				Txs []string `json:"txs"`
			} `json:"data"`
			Evidence struct {
				Evidence []interface{} `json:"evidence"`
			} `json:"evidence"`
			Header struct {
				AppHash       string `json:"app_hash"`
				ChainID       string `json:"chain_id"`
				ConsensusHash string `json:"consensus_hash"`
				DataHash      string `json:"data_hash"`
				EvidenceHash  string `json:"evidence_hash"`
				Height        string `json:"height"`
				LastBlockID   struct {
					Hash  string `json:"hash"`
					Parts struct {
						Hash  string `json:"hash"`
						Total int    `json:"total"`
					} `json:"parts"`
				} `json:"last_block_id"`
				LastCommitHash     string    `json:"last_commit_hash"`
				LastResultsHash    string    `json:"last_results_hash"`
				NextValidatorsHash string    `json:"next_validators_hash"`
				ProposerAddress    string    `json:"proposer_address"`
				Time               time.Time `json:"time"`
				ValidatorsHash     string    `json:"validators_hash"`
				Version            struct {
					Block string `json:"block"`
				} `json:"version"`
			} `json:"header"`
			LastCommit struct {
				BlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Hash  string `json:"hash"`
						Total int    `json:"total"`
					} `json:"parts"`
				} `json:"block_id"`
				Height     string `json:"height"`
				Round      int    `json:"round"`
				Signatures []struct {
					BlockIDFlag      int       `json:"block_id_flag"`
					Signature        string    `json:"signature"`
					Timestamp        time.Time `json:"timestamp"`
					ValidatorAddress string    `json:"validator_address"`
				} `json:"signatures"`
			} `json:"last_commit"`
		} `json:"block"`
		BlockID struct {
			Hash  string `json:"hash"`
			Parts struct {
				Hash  string `json:"hash"`
				Total int    `json:"total"`
			} `json:"parts"`
		} `json:"block_id"`
	} `json:"result"`
}

type RpcStatus struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		NodeInfo struct {
			Channels   string `json:"channels"`
			ID         string `json:"id"`
			ListenAddr string `json:"listen_addr"`
			Moniker    string `json:"moniker"`
			Network    string `json:"network"`
			Other      struct {
				RPCAddress string `json:"rpc_address"`
				TxIndex    string `json:"tx_index"`
			} `json:"other"`
			ProtocolVersion struct {
				App   string `json:"app"`
				Block string `json:"block"`
				P2P   string `json:"p2p"`
			} `json:"protocol_version"`
			Version string `json:"version"`
		} `json:"node_info"`
		SyncInfo struct {
			CatchingUp          bool      `json:"catching_up"`
			EarliestAppHash     string    `json:"earliest_app_hash"`
			EarliestBlockHash   string    `json:"earliest_block_hash"`
			EarliestBlockHeight string    `json:"earliest_block_height"`
			EarliestBlockTime   time.Time `json:"earliest_block_time"`
			LatestAppHash       string    `json:"latest_app_hash"`
			LatestBlockHash     string    `json:"latest_block_hash"`
			LatestBlockHeight   string    `json:"latest_block_height"`
			LatestBlockTime     time.Time `json:"latest_block_time"`
		} `json:"sync_info"`
		ValidatorInfo struct {
			Address string `json:"address"`
			PubKey  struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"pub_key"`
			VotingPower string `json:"voting_power"`
		} `json:"validator_info"`
	} `json:"result"`
}

type RpcConsensusState struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		RoundState struct {
			HeightRoundStep string `json:"height/round/step"`
			HeightVoteSet   []struct {
				Precommits         []string `json:"precommits"`
				PrecommitsBitArray string   `json:"precommits_bit_array"`
				Prevotes           []string `json:"prevotes"`
				PrevotesBitArray   string   `json:"prevotes_bit_array"`
				Round              int      `json:"round"`
			} `json:"height_vote_set"`
			LockedBlockHash   string `json:"locked_block_hash"`
			ProposalBlockHash string `json:"proposal_block_hash"`
			Proposer          struct {
				Address string `json:"address"`
				Index   int    `json:"index"`
			} `json:"proposer"`
			StartTime      time.Time `json:"start_time"`
			ValidBlockHash string    `json:"valid_block_hash"`
		} `json:"round_state"`
	} `json:"result"`
}

type RpcDumpConsensusState struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		RoundState struct {
			HeightRoundStep string `json:"height/round/step"`
			HeightVoteSet   []struct {
				Precommits         []string `json:"precommits"`
				PrecommitsBitArray string   `json:"precommits_bit_array"`
				Prevotes           []string `json:"prevotes"`
				PrevotesBitArray   string   `json:"prevotes_bit_array"`
				Round              int      `json:"round"`
			} `json:"height_vote_set"`
			LockedBlockHash   string `json:"locked_block_hash"`
			ProposalBlockHash string `json:"proposal_block_hash"`
			Proposer          struct {
				Address string `json:"address"`
				Index   int    `json:"index"`
			} `json:"proposer"`
			StartTime      time.Time `json:"start_time"`
			ValidBlockHash string    `json:"valid_block_hash"`
		} `json:"round_state"`
	} `json:"result"`
}

type RpcConsensusParams struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Block struct {
			Data struct {
				Txs []string `json:"txs"`
			} `json:"data"`
			Evidence struct {
				Evidence []interface{} `json:"evidence"`
			} `json:"evidence"`
			Header struct {
				AppHash       string `json:"app_hash"`
				ChainID       string `json:"chain_id"`
				ConsensusHash string `json:"consensus_hash"`
				DataHash      string `json:"data_hash"`
				EvidenceHash  string `json:"evidence_hash"`
				Height        string `json:"height"`
				LastBlockID   struct {
					Hash  string `json:"hash"`
					Parts struct {
						Hash  string `json:"hash"`
						Total int    `json:"total"`
					} `json:"parts"`
				} `json:"last_block_id"`
				LastCommitHash     string    `json:"last_commit_hash"`
				LastResultsHash    string    `json:"last_results_hash"`
				NextValidatorsHash string    `json:"next_validators_hash"`
				ProposerAddress    string    `json:"proposer_address"`
				Time               time.Time `json:"time"`
				ValidatorsHash     string    `json:"validators_hash"`
				Version            struct {
					Block string `json:"block"`
				} `json:"version"`
			} `json:"header"`
			LastCommit struct {
				BlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Hash  string `json:"hash"`
						Total int    `json:"total"`
					} `json:"parts"`
				} `json:"block_id"`
				Height     string `json:"height"`
				Round      int    `json:"round"`
				Signatures []struct {
					BlockIDFlag      int       `json:"block_id_flag"`
					Signature        string    `json:"signature"`
					Timestamp        time.Time `json:"timestamp"`
					ValidatorAddress string    `json:"validator_address"`
				} `json:"signatures"`
			} `json:"last_commit"`
		} `json:"block"`
		BlockID struct {
			Hash  string `json:"hash"`
			Parts struct {
				Hash  string `json:"hash"`
				Total int    `json:"total"`
			} `json:"parts"`
		} `json:"block_id"`
	} `json:"result"`
}

type RpcCheckTx struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Code         int           `json:"code"`
		Data         interface{}   `json:"data"`
		Log          string        `json:"log"`
		Info         string        `json:"info"`
		GasWanted    string        `json:"gas_wanted"`
		GasUsed      string        `json:"gas_used"`
		Events       []interface{} `json:"events"`
		Codespace    string        `json:"codespace"`
		Sender       string        `json:"sender"`
		Priority     string        `json:"priority"`
		MempoolError string        `json:"mempoolError"`
	} `json:"result"`
}

type UnconfirmedTxs struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		NTxs       string   `json:"n_txs"`
		Total      string   `json:"total"`
		TotalBytes string   `json:"total_bytes"`
		Txs        []string `json:"txs"`
	} `json:"result"`
}

type NumUnconfirmedTxs struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		NTxs       string      `json:"n_txs"`
		Total      string      `json:"total"`
		TotalBytes string      `json:"total_bytes"`
		Txs        interface{} `json:"txs"`
	} `json:"result"`
}

type AbciQuery struct {
	Error  string `json:"error"`
	Result struct {
		Response struct {
			Log    string `json:"log"`
			Height string `json:"height"`
			Proof  string `json:"proof"`
			Value  string `json:"value"`
			Key    string `json:"key"`
			Index  string `json:"index"`
			Code   string `json:"code"`
		} `json:"response"`
	} `json:"result"`
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}
