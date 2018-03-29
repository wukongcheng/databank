pragma solidity ^0.4.18;

contract XCData {

    struct Data {
        uint256 timestamp;
        string datahash;
        bytes encryptedAESKey;
    }

    struct DataList {
        Data[] list;
    }

    event Autherize(address to, string datahash);

    address _owner;
    mapping (string => DataList) _xcData;
    mapping (string => address) _xcOwner;

    mapping (address => mapping (string => bytes)) _HashToAutherizedAESKey;
    mapping (address => uint256) _AutherizedDataLength;

    function XCData () public {
        _owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == _owner);
        _;
    }

    function commitData(string did, string datahash, bytes encryptedAESKey) external  {
        address owner = _xcOwner[did];
        if(owner != address(0)) {
            require(msg.sender == owner);
        }
        else {
            _xcOwner[did] = msg.sender;
        }

        _xcData[did].list.push(Data(now, datahash,encryptedAESKey));
    }

    function getDataLength(string did) external view returns(uint256) {
        return _xcData[did].list.length;
    }

    function getData(string did, uint256 index) external view returns (uint256, string, bytes) {
        require(_xcData[did].list.length > 0);
        require(index < _xcData[did].list.length);

        return (_xcData[did].list[index].timestamp, _xcData[did].list[index].datahash,_xcData[did].list[index].encryptedAESKey);
    }

    function transferDidOwner(string did, address to) external  {
        address owner = _xcOwner[did];

        require(owner != address(0));
        require(msg.sender == owner);

        _xcOwner[did] = to;
    }

    function commitNewOwnerData(string did, string datahash, bytes encryptedAESKey) external {
        address owner = _xcOwner[did];

        require(owner != address(0));
        require(msg.sender == owner);

        _xcData[did].list.push(Data(now, datahash, encryptedAESKey));
    }

    //TODO, how to ensure the AES key has been re-encrypted by counterparty public key
    //TODO, In theroy, patient can hack xci and use wrong public key deliberately
    function autherizeData(address to, string did, uint256 index, bytes encryptedAESKey) external {
        address owner = _xcOwner[did];

        require(owner != address(0));
        require(owner != to);
        require(msg.sender == owner);

        uint256 count = _AutherizedDataLength[to];

        _HashToAutherizedAESKey[to][_xcData[did].list[index].datahash]=encryptedAESKey;
        _AutherizedDataLength[to]=count+1;

        Autherize(to,_xcData[did].list[index].datahash);
    }

    function getAutherizedDataLength(address addr) external view returns (uint256) {
        return _AutherizedDataLength[addr];
    }

    function getAutherizedAESKeyByHash(address addr, string datahash) external view returns (bytes) {
        return _HashToAutherizedAESKey[addr][datahash];
    }

    function deletePreOwnerData(string did) external {
        require(_xcOwner[did] == msg.sender);

        delete _xcData[did];
    }
}