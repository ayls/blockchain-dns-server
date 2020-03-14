pragma solidity 0.6.3;

contract InetDnsRecord {
    mapping (string => mapping(uint16 => string)) internal _dnsMapping;

    function addRecord(string memory key, uint16 recType, string memory recValue) public {
        _dnsMapping[key][recType] = recValue;
    }

    function getRecord(string memory key, uint16 recType) public view returns (string memory){        
        return _dnsMapping[key][recType];
    }
}