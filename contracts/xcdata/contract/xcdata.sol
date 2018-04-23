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

    event NewCommitData(address indexed from, string did, string datahash, uint256 to);
    event Authorize(address indexed from, address indexed to, string datahash);
    event TransferDid(address indexed from, string did, address indexed to);
    event DeleteDid(address indexed from, string did);

    address _owner;
    mapping (string => DataList) _xcData;
    mapping (string => address) _xcOwner;

    mapping (address => mapping (string => bytes)) _HashToAuthorizedAESKey;
    mapping (address => uint256) _AuthorizedDataLength;

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

        uint256 index = _xcData[did].list.length;

        NewCommitData(msg.sender,did,datahash,index);
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

        TransferDid(msg.sender,did,to);
    }

    function commitNewOwnerData(string did, string datahash, bytes encryptedAESKey) external {
        address owner = _xcOwner[did];

        require(owner != address(0));
        require(msg.sender == owner);

        _xcData[did].list.push(Data(now, datahash, encryptedAESKey));

        uint256 index = _xcData[did].list.length;

        NewCommitData(msg.sender,did,datahash,index);
    }

    //TODO, how to ensure the AES key has been re-encrypted by counterparty public key
    //TODO, In theroy, patient can hack xci and use wrong public key deliberately
    function authorizeData(address to, string did, uint256 index, bytes encryptedAESKey) external {
        address owner = _xcOwner[did];

        require(owner != address(0));
        require(owner != to);
        require(msg.sender == owner);

        uint256 count = _AuthorizedDataLength[to];

        _HashToAuthorizedAESKey[to][_xcData[did].list[index].datahash]=encryptedAESKey;
        _AuthorizedDataLength[to]=count+1;

        Authorize(msg.sender,to,_xcData[did].list[index].datahash);
    }

    function getAuthorizedDataLength(address addr) external view returns (uint256) {
        return _AuthorizedDataLength[addr];
    }

    function getAuthorizedAESKeyByHash(address addr, string datahash) external view returns (bytes) {
        return _HashToAuthorizedAESKey[addr][datahash];
    }

    function deletePreOwnerData(string did) external {
        require(_xcOwner[did] == msg.sender);

        delete _xcData[did];

        DeleteDid(msg.sender,did);
    }
}