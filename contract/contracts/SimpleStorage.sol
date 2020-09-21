pragma solidity ^0.5.0;

import "./JsmnSolLib.sol";

contract ParseJson {
    uint success;
    JsmnSolLib.Token[] tokens;
    uint tokenCount;

    uint maxNumberElement;

    constructor() public {
        maxNumberElement = 100;
    }
    
    function parse(string memory json, uint numberElements) public {
        require(numberElements <= maxNumberElement,
                "numberElements can't biger than max number of element.");
        JsmnSolLib.Token[] memory tokens_;
        (success, tokens_, tokenCount) = JsmnSolLib.parse(json, numberElements);
        for (uint i = 0; i < tokens_.length; i++) {
            tokens.push(tokens_[i]);
        }
    }

    function setMaxNumElements(uint maxNumber) public {
        maxNumberElement = maxNumber;
    }
}
                
        
            
