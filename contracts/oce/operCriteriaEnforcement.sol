pragma solidity ^0.4.6;

contract mortal {
    /* Define variable owner of the type address */
    address owner;

    /* This function is executed at initialization and sets the owner of the contract */
    function mortal() public { owner = msg.sender;  }

    /* Function to recover the funds on the contract */
    function kill() public { if (msg.sender == owner) selfdestruct(owner);  }

}

contract operCriteriaEnforcement is mortal {
	
	mapping (string => BinaryMeasure) binaryMeasures;
	mapping (address => User) authorizedUsers;

	enum UserRole {Admin, Operator}

	struct User {
		bool initialized;
		string name;
		UserRole role;
	}
	
	struct BinaryMeasure {
		bool initialized;
		string name;
		uint measurement;
	}

	function operCriteriaEnforcement(string username) public {
		authorizedUsers[msg.sender] = User(true,username,UserRole.Admin);
	}

	function addAuthorizedUser(address usr, string username, UserRole role) public {
		if (!authorizedUsers[msg.sender].initialized) {
			revert();
		}
		if (authorizedUsers[msg.sender].role != UserRole.Admin) {
			revert();
		}
		if (authorizedUsers[usr].initialized) {
			revert();
		}
		authorizedUsers[usr] = User(true,username,role);
	}

	function addBinary(string binName, uint measurement) public {
		
		if (!authorizedUsers[msg.sender].initialized) {
			revert();
		}

		binaryMeasures[binName] = BinaryMeasure(true,binName,measurement);
	}

	function verifyBinary(string binName, uint measurement) view public returns (bool res){
		if(!binaryMeasures[binName].initialized) {
			return false;
		}
		if(binaryMeasures[binName].measurement == measurement) {
			return true;
		}
		return false;
	}
}
