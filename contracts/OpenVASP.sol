pragma solidity ^0.5.0;

contract OpenVASP {
    string public name;
    bytes8 public code;
    uint8[] public channels;

    string public handshakeKey;
    string public signingKey;

    address public owner;
    address public administrator;

    string public street;
    bytes32 public buildingNo;
    string public addressLine;
    bytes32 public postCode;
    string public town;
    bytes2 public country;

    string public email;
    string public website;

    address[] public trustedPeers;
    address[] public identityClaims;

    function addTrustedPeer(address) public;

    function removeTrustedPeer(address) public;

    function addClaim(address) public;

    function removeClaim(address) public;
}

contract IdentityClaim{
    address public issuer;
    address public vasp;
    bytes32 public claim;
}
