// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MyCounter {
    event Increment(uint by);
    uint256 public count;
    function counterAdd() public{
        count++;
        // 实现链上链下的实时交互，链下如go代码可以订阅，感知合约行为
        emit Increment(1); 
    }
}