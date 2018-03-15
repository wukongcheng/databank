pragma solidity ^0.4.18;

contract ipfs {

    address _owner;
    //address: transaction sender address
    //first string: timestamp(yyyy-mm-dd)+"-"+filename(length no more than 10)
    //second string: ipfs url
    mapping (address => mapping (string => string) ) _ipfsUrl;
    //address: transaction sender address
    //first string: date(yyyy-mm-dd)
    //second string: fileName(length no more than 10)
    mapping (address => mapping (string => string) ) _dateFileList;

    function ipfs () public {
        _owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == _owner);
        _;
    }

    //Currently this function is not efficient, later I will optimize it.
    function addNewIpfsUrl(string date, string fileName, string url) external {
        string memory key = strConcat(date,"-",fileName);
        string memory oldIpfsUrl = _ipfsUrl[msg.sender][key];
        bytes memory oldIpfsUrlBytes = bytes(oldIpfsUrl);
        assert(oldIpfsUrlBytes.length==0);

        //save ipfs url
        _ipfsUrl[msg.sender][key]=url;

        //Add file name to file name list, each date(yyyy-mm-dd) may have corresponding fileName list
        string memory fileNameList = _dateFileList[msg.sender][date];
        bytes memory fileNameListBytes = bytes(fileNameList);

        string memory newFileNameList;
        if(fileNameListBytes.length==0){
            newFileNameList=fileName;
        } else {
            newFileNameList = strConcat(fileNameList,",",fileName);
        }

        _dateFileList[msg.sender][date]=newFileNameList;
    }

    function getFileListByDate(string date) public constant returns (string) {
        return _dateFileList[msg.sender][date];
    }

    //key=date+fileName
    function getIpfsUrl(string date, string fileName) public constant returns (string) {
        string memory key = strConcat(date,"-",fileName);
        return _ipfsUrl[msg.sender][key];
    }

    function getOwner() public constant returns (address) {
        return _owner;
    }

    function strConcat(string _a, string _b, string _c) internal returns (string){
        bytes memory _ba = bytes(_a);
        bytes memory _bb = bytes(_b);
        bytes memory _bc = bytes(_c);

        string memory abc = new string(_ba.length + _bb.length + _bc.length);
        bytes memory abcBytes = bytes(abc);
        uint k = 0;
        uint i;
        for (i = 0; i < _ba.length; i++)
            abcBytes[k++] = _ba[i];
        for (i = 0; i < _bb.length; i++)
            abcBytes[k++] = _bb[i];
        for (i = 0; i < _bc.length; i++)
            abcBytes[k++] = _bc[i];

        return string(abcBytes);
    }

}