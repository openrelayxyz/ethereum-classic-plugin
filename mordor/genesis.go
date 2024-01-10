package main

func GenesisBlock() []byte {
	return []byte(`{
        "config": {
            "networkId": 7,
            "chainId": 63,
            "homesteadBlock": 0,
			"daoForkBlock": null,
			"daoForkSupport": false,
			"eip150Block": 0,
			"eip155Block": 0,
			"eip158Block": 0,
			"byzantiumBlock": 0,
			"constantinopleBlock": 301243,
			"petersburgBlock": 301243,
			"istanbulBlock": 999983,
			"berlinBlock": 3985893,
			"londonBlock": 5520000,
            "shanghaiBlock": 9957000,
            "disposalBlock": 0,
            "ethash": {},
            "requireBlockHashes": {
                "840013": "0x2ceada2b191879b71a5bcf2241dd9bc50d6d953f1640e62f9c2cee941dc61c9d",
                "840014": "0x8ec29dd692c8985b82410817bac232fc82805b746538d17bc924624fe74a0fcf"
            }
        },
        "nonce": "0x0",
        "timestamp": "0x5d9676db",
        "extraData": "0x70686f656e697820636869636b656e206162737572642062616e616e61",
        "gasLimit": "0x2fefd8",
        "difficulty": "0x20000",
        "mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
        "coinbase": "0x0000000000000000000000000000000000000000",
        "alloc": {},
        "number": "0x0",
        "gasUsed": "0x0",
        "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
    }`)

}