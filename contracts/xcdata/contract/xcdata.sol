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

    function getDataLength(string did) external view returns(uint256) {
        return _xcData[did].list.length;
    }

    function getData(string did, uint256 index) external view returns (uint256, string) {
        require(_xcData[did].list.length > 0);
        require(index < _xcData[did].list.length);

        return (_xcData[did].list[index].timestamp, _xcData[did].list[index].datahash);
    }

    function transferDidOwner(string did, address to) external  {
        address owner = _xcOwner[did];

        require(owner != address(0));
        require(msg.sender == owner);

        _xcOwner[did] = to;
    }

    function commitNewOwnerData(string did, uint256 timestamp, string datahash) external {
        address owner = _xcOwner[did];

        require(owner != address(0));
        require(msg.sender == owner);

        _xcData[did].list.push(Data(timestamp, datahash));
    }

    function deletePreOwnerData(string did) external {
        require(_xcOwner[did] == msg.sender);

        delete _xcData[did];
    }
}