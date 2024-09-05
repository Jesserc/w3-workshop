pragma solidity ^0.8.4;

contract Store {
    uint number;

    function store(uint256 num) public {
        number = num;
    }

    function retrieve() public returns(uint){
        return number;
    }
}