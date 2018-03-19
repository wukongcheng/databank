pragma solidity ^0.4.18;

contract XCData {

    struct Data {
        uint256 timestamp;
        string datahash;
    }

    struct DataList {
        Data[] list;
    }

    address _owner;
    mapping (string => DataList) _xcData;
    mapping (string => address) _xcOwner;

    function XCData () public {
        _owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == _owner);
        _;
    }

    function commitData(string did, string datahash) external  {
        address owner = _xcOwner[did];
        if(owner != address(0)) {
            require(msg.sender == owner);
        }
        else {
            _xcOwner[did] = msg.sender;
        }

        _xcData[did].list.push(Data(now, datahash));
    }

}