pragma solidity ^0.4.18;

contract ipfs {

    address _owner;
    //address: transaction sender address
    //first string: fileName
    //second string: ipfs url
    mapping (address => mapping (string => string) ) _ipfsUrl;
    //address: transaction sender address
    //first string: fileIndex
    //second string: fileName
    mapping (address => mapping (uint256 => string) ) _fileName;

    mapping (address => uint256 ) _fileCount;

    function ipfs () public {
        _owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == _owner);
        _;
    }

    //Currently this function is not efficient, later I will optimize it.
    function addNewIpfsUrl(string fileName, string ipfsUrl) external {
        uint256 fileIndex = _fileCount[msg.sender];
        string memory existValue = _ipfsUrl[msg.sender][fileName];
        bytes memory valueBytes = bytes(existValue);
        assert(valueBytes.length==0);
        //save ipfs url
        _ipfsUrl[msg.sender][fileName]=ipfsUrl;
        _fileName[msg.sender][fileIndex]=fileName;

        fileIndex++;
        _fileCount[msg.sender]=fileIndex;
    }

    function getFileQuantity(address account) public constant returns (uint256) {
        return _fileCount[account];
    }

    function getFileNameByIndex(address account, uint256 index) public constant returns (string) {
        return _fileName[account][index];
    }

    function getIpfsUrl(address account, string fileName) public constant returns (string) {
        return _ipfsUrl[account][fileName];
    }

    function getOwner() public constant returns (address) {
        return _owner;
    }

}