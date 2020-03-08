pragma solidity 0.6.3;

contract InetDnsRecord {
    mapping (string => mapping(int => string)) internal _dnsMapping;

    function addRecord(string memory key, int8 recType, string memory recValue) public {
        _dnsMapping[key][recType] = recValue;
    }

    function getRecord(string memory key, int8 recType) public view returns (string memory){        
        return _dnsMapping[key][recType];
    }
}