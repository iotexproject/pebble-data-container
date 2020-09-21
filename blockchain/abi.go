package blockchain

const (
	ABI = `[
    {
        "constant": false,
        "inputs": [
            {
                "name": "maxNumber",
                "type": "uint256"
            }
        ],
        "name": "setMaxNumElements",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": false,
        "inputs": [
            {
                "name": "json",
                "type": "string"
            },
            {
                "name": "numberElements",
                "type": "uint256"
            }
        ],
        "name": "parse",
        "outputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "constructor"
    }
]`
)
