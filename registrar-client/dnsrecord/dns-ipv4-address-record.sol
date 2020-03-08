pragma solidity >=0.5.2 <0.6.4;

contract DnsIPv4AddressRecord {
    mapping (string => string) internal _dnsMapping;

    function addRecord(string memory domainname, string memory ip) public {
        _dnsMapping[domainname] = ip;
    }

    function getRecord(string memory domainname) public view returns (string memory){        
        return _dnsMapping[domainname];
    }
}