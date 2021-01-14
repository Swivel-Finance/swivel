pragma solidity ^0.7.3;
pragma experimental ABIEncoderV2;

import "./safeMath.sol";

abstract contract Erc20 {
	function approve(address, uint) virtual external returns (bool);
	function transfer(address, uint) virtual external returns (bool);
	function balanceOf(address _owner) virtual external returns (uint256 balance);
	function transferFrom(address sender, address recipient, uint256 amount) virtual public returns (bool);
}

abstract contract cErc20 is Erc20 {
	function mint(uint) virtual external returns (uint);
	function redeem(uint redeemTokens) virtual external returns (uint);
	function redeemUnderlying(uint redeemAmount) virtual external returns (uint);
	function exchangeRateCurrent() virtual external returns (uint);
}

contract Swivel {

/// Signature struct 
struct sig{
	uint8 v;
	bytes32 r;
	bytes32 s;
}
    
/// Domain struct for sig verification
struct EIP712Domain {
	string  name;
	string  version;
	uint256 chainId;
	address verifyingContract;
}
    
/// Order structure used for maker signature verification (signed off-chain)	
struct order {
	address maker;
	uint256 side;
	address tokenAddress;
	uint256 duration;
	uint256 rate;
	uint256 interest;
	uint256 principal;
	uint256 nonce;
	uint256 expiry;
	bytes orderKey;
}

    
/// Order structure for taker mapping    
struct agreement {
	address maker;
	address taker;
	uint256 side;
	address tokenAddress;
	uint256 duration;
	uint256 rate;
	uint256 interest;
	uint256 principal;
	uint256 state;
	uint256 releaseTime;
	uint256 initialRate;
	bytes orderKey;
}    
    

// Mapping containing a list of agreements for a given order
// @param orderKey: Hash generated by maker to identify order
mapping(bytes => bytes[]) public orderTakers;

// Nested mapping that is connected to a maker
// @param orderKey: Hash generated by maker to identify order
// @param agreementKey: Hash generated by taker to identify agreement
mapping(bytes => mapping(bytes => agreement)) public agreements;

// The amount of a maker's order that has been  filled.
// @param orderKey: Hash generated by maker to identify order
mapping (bytes => uint256) public filled;

// Whether a maker's order has been cancelled.
// @param orderKey: Hash generated by maker to identify order
mapping (bytes => bool) public cancelled;


/// Event on order/agreement activation
event initiated(
	bytes orderKey,
	bytes agreementKey
);

/// Event on maker order cancellation
event cancelledOrder(
	bytes orderKey
);

/// Event on taker order release
event released(
	bytes orderKey,
	bytes agreementKey
);
    

using SafeMath for uint;

constructor () {
	DOMAIN_SEPARATOR = hashDomain(EIP712Domain({
		name: "Swivel",
		version: '1',
		chainId: 5,
		verifyingContract: 0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC
}));
}
    
bytes32 DOMAIN_SEPARATOR;
    
    
// EIP Domain Hash Schema
bytes32 constant EIP712DOMAIN_TYPEHASH = keccak256(
	"EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
);


// Order Hash Schema
bytes32 constant ORDER_TYPEHASH = keccak256(
	"order(address maker,uint256 side,address tokenAddress,uint256 duration,uint256 rate,uint256 interest,uint256 principal,uint256 nonce,uint256 expiry,bytes orderKey)"
);


/// EIP Domain Hash Function
/// @param eip712Domain: EIP-712 Domain information provided in constructor
function hashDomain(EIP712Domain memory eip712Domain) internal pure returns (bytes32) {
	return keccak256(abi.encode(
		EIP712DOMAIN_TYPEHASH,
		keccak256(bytes(eip712Domain.name)),
		keccak256(bytes(eip712Domain.version)),
		eip712Domain.chainId,
		eip712Domain.verifyingContract
        ));
}


/// Order Hash Function
/// @param _order: order struct
function hashOrder(order memory _order)private pure returns(bytes32){
	return keccak256(abi.encode(
		ORDER_TYPEHASH,
		_order.maker,
		_order.side,
		_order.tokenAddress,
		_order.duration,
		_order.rate,
		_order.interest,
		_order.principal,
		_order.nonce,
		_order.expiry,
		keccak256(bytes(_order.orderKey))
	));
}

	
/// Fill the entirety of a maker order's volume
/// @param _order: maker's order 
/// @param agreementKey: off-chain key generated as keccak hash of User + time + nonce
/// @param signature: signature associated with maker order
function fill(order memory _order, bytes memory agreementKey, sig memory signature) public {
	    
	// Check if order already partially filled
	require(filled[_order.orderKey] == 0, "Order Already Partial/Fully Filled");
    
	// Check if order has been cancelled
	require(cancelled[_order.orderKey]==false, "Order Has Been Cancelled");

	// Check if order has already expired
	require(_order.expiry >= block.timestamp, "Order Has Expired");

	// Validate order signature & ensure it was created by maker
	require(_order.maker == ecrecover(
	keccak256(abi.encodePacked(
			"\x19\x01",
			DOMAIN_SEPARATOR,
			hashOrder(_order)
			)),
			signature.v,
			signature.r,
			signature.s), 
	"Invalid Signature");

	// Settle Response
	settle(_order,agreementKey);

}


/// Settle an entire maker order's volume
/// @param _order: maker's order
/// @param agreementKey: off-chain key generated as keccak hash of User + time + nonce
function settle(order memory _order, bytes memory agreementKey) private returns (bool){
	    
	agreement memory _agreement;

	cErc20 cToken = cErc20(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb); //DAI cToken Address

	// Check trades side
	// Transfers funds to Swivel contract
	Erc20 underlying = Erc20(_order.tokenAddress);
	if (_order.side == 0) {
		require(underlying.transferFrom(_order.maker, address(this), _order.principal), "Transfer Failed!");
		require(underlying.transferFrom(msg.sender, address(this), _order.interest), "Transfer Failed!");
		filled[_order.orderKey] = _order.interest;
	}
	if (_order.side == 1) {
		require(underlying.transferFrom(_order.maker, address(this), _order.interest), "Transfer Failed!");
		require(underlying.transferFrom(msg.sender, address(this), _order.principal), "Transfer Failed!");
		filled[_order.orderKey] = _order.principal;
	}
    	    
 	// Mint CToken from Swivel contract
 	mintCToken(_order.tokenAddress,_order.interest.add(_order.principal));
            
	// Instantiate agreement
	_agreement.maker = _order.maker;
	_agreement.taker = msg.sender;
	_agreement.side = _order.side;
	_agreement.tokenAddress = _order.tokenAddress;
	_agreement.duration = _order.duration;
	_agreement.rate = _order.rate;
	_agreement.interest = _order.interest;
	_agreement.principal = _order.principal;
	_agreement.state = 1;
	_agreement.releaseTime = block.timestamp.add(_order.duration);
	_agreement.initialRate = cToken.exchangeRateCurrent();
	_agreement.orderKey = _order.orderKey;
    	    

    // Store taker order	    
	agreements[_order.orderKey][agreementKey]= _agreement;
    	    
	// Push taker order to a list mapped to a given maker's order
	orderTakers[_order.orderKey].push(agreementKey);
	
	    
	emit initiated(_agreement.orderKey,agreementKey);
    	    
	return true;
}


/// Fill partial maker order 
/// @param _order: maker's order 
/// @param takerVolume: amount of currency being taken
/// @param agreementKey: off-chain key generated as keccak hash of User + time + nonce
/// @param signature: signature associated with maker order
function partialFill(order memory _order,uint256 takerVolume, bytes memory agreementKey, sig memory signature) public {

	// Check if order has been cancelled
	require(cancelled[_order.orderKey]==false, "Order Has Been Cancelled");

	// Check if order has already expired
	require(_order.expiry >= block.timestamp, "Order Has Expired");

	// Validate order signature & ensure it was created by maker
	require(_order.maker == ecrecover(
		keccak256(abi.encodePacked(
			"\x19\x01",
			DOMAIN_SEPARATOR,
			hashOrder(_order)
			)),
			signature.v,
			signature.r,
			signature.s), 
	"Invalid Signature");


	// Settle Response
	partialSettle(_order, takerVolume, agreementKey);

}
    
    
/// Settle part of a maker order's volume
/// @param _order: makers order to fill
/// @param takerVolume: amount of currency being taken
/// @param agreementKey: off-chain key generated as keccak hash of User + time + nonce
function partialSettle(order memory _order,uint256 takerVolume, bytes memory agreementKey) private returns (bool){

	agreement memory _agreement;
	
	cErc20 cToken = cErc20(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb); //DAI cToken Address
    Erc20 underlying = Erc20(_agreement.tokenAddress);

	// If order is fixed-side, ensure volume is less than expected interest
	if (_order.side == 0) {

		require (takerVolume <= (_order.interest), "Taker Volume > Maker");

		// If order has already been partially filled
		if (filled[_order.orderKey] != 0) {
		
			require (takerVolume <= (_order.interest - filled[_order.orderKey]), "Taker Volume > Available");
			
		}
		// Calculate taker % of total maker order and set opposing param 
		uint256 orderRatio = (((takerVolume).mul(1e18)).div(_order.interest));
		_agreement.principal = _order.principal.mul(orderRatio).div(1e18);
		_agreement.interest = takerVolume;
		
		// Transfers funds to Swivel contract
		require(underlying.transferFrom(_order.maker, address(this), _agreement.principal), "Transfer Failed!");
		require(underlying.transferFrom(msg.sender, address(this), _agreement.interest), "Transfer Failed!");
	}

	// If order is floating-side, ensure volume is less than expected principal
	if (_order.side == 1) {

		require (takerVolume <= (_order.principal), "Taker Volume > Maker");

		// If order has already been partially filled
		if (filled[_order.orderKey] != 0) {

		require (takerVolume <= (_order.principal - filled[_order.orderKey]), "Taker Volume > Available");
		}
		
		// Calculate taker % of total maker order and set opposing param 
		uint256 orderRatio = (((takerVolume).mul(1e18)).div(_order.principal));
		_agreement.interest = _order.interest.mul(orderRatio).div(1e18);
		_agreement.principal = takerVolume;
		
		// Transfers funds to Swivel contract
		require(underlying.transferFrom(_order.maker, address(this), _agreement.interest), "Transfer Failed!");
		require(underlying.transferFrom(msg.sender, address(this), _agreement.principal), "Transfer Failed!");
	}


	// Mint cToken from Swivel contract
	mintCToken(_agreement.tokenAddress,_agreement.interest.add(_agreement.principal));


	// After mint success fill order params & store agreement in agreements mapping
	_agreement.maker = _order.maker;
	_agreement.taker = msg.sender;
	_agreement.side = _order.side;
	_agreement.tokenAddress = _order.tokenAddress;
	_agreement.duration = _order.duration;
	_agreement.rate= _order.rate;
	_agreement.state = 1;  // Set state to active
	_agreement.releaseTime = block.timestamp.add(_order.duration); // Set locktime
	_agreement.initialRate = cToken.exchangeRateCurrent();  // Get initial exchange rate
	
	agreements[_order.orderKey][agreementKey] = _agreement;

    filled[_order.orderKey] = takerVolume;
    
	// Push agreementKey to orderTakers
	orderTakers[_order.orderKey].push(agreementKey);



    // Emit new agreement
    emit initiated(_agreement.orderKey,agreementKey);

	return true;	    	    
}
    
    
/// Fill partial maker order 
/// @param orders: array of orders
/// @param signatures: array of associated order signatures
/// @param takerVolume: amount of currency being taken
/// @param agreementKey: off-chain key generated as keccak hash of User + time + nonce
function batchFillFixed(order[] memory orders, sig[] memory signatures, uint256 takerVolume, bytes memory agreementKey) public {
    
    uint256 orderCount = orders.length;
    
    uint256 amountFilled;
    
    
    // Loop through orders
    for (uint i=0; i<orderCount; i++) {
        
        // Instantiate order
        order memory _order = orders[i];
        
        // Calculate current available volume
        uint256 availableTaker = takerVolume - amountFilled;
        uint256 availableMaker = _order.interest - filled[_order.orderKey];
        
        // Check if full fill is possible + fill
        if (filled[_order.orderKey] != 0 && _order.interest <= availableTaker) {
            fill(_order, agreementKey, signatures[i]);
            amountFilled = amountFilled + _order.interest;
        }
        
        // If full fill is not possible
        else {
            // Check which side has the limiting available volume + partialFill
            if (availableTaker > availableMaker) {
                partialFill(_order, availableMaker, agreementKey, signatures[i]);
                amountFilled = amountFilled + availableMaker;
            }
            else {
                partialFill(_order, availableTaker, agreementKey, signatures[i]);
                amountFilled = amountFilled + availableTaker;
            }
        }
    }
}


/// Fill partial maker order 
/// @param orders: array of orders
/// @param signatures: array of associated order signatures
/// @param takerVolume: amount of currency being taken
/// @param agreementKey: off-chain key generated as keccak hash of User + time + nonce
function batchFillFloating(order[] memory orders, sig[] memory signatures, uint256 takerVolume, bytes memory agreementKey) public {
    
    
    uint256 orderCount = orders.length;
    
    uint256 amountFilled;
    
    
    // Loop through orders
    for (uint i=0; i<orderCount; i++) {
        
        // Instantiate order
        order memory _order = orders[i];
        
        // Calculate current available volume
        uint256 availableTaker = takerVolume - amountFilled;
        uint256 availableMaker = _order.principal - filled[_order.orderKey];
        
        // Check if full fill is possible
        if (filled[_order.orderKey] != 0 && _order.principal <= availableTaker) {
            fill(_order, agreementKey, signatures[i]);
            amountFilled = amountFilled + _order.principal;
        }
        
        // If full fill is not possible
        else {
            // Check which side has the limiting available volume + partialFill
            if (availableTaker > availableMaker) {
                partialFill(_order, availableMaker, agreementKey, signatures[i]);
                amountFilled = amountFilled + availableMaker;
            }
            else {
                partialFill(_order, availableTaker, agreementKey, signatures[i]);
                amountFilled = amountFilled + availableTaker;
            }
        }
    }
}
    
    
/// Cancel an order
/// @param _order: maker's order
/// @param signature: signature associated with maker's order
function cancel(order memory _order, sig memory signature) public returns(bool){

	// Validate order signature & ensure it was created by maker
	require(msg.sender == ecrecover(
		keccak256(abi.encodePacked(
		"\x19\x01",
		DOMAIN_SEPARATOR,
		hashOrder(_order)
		)),
		signature.v,
		signature.r,
		signature.s), 
	"Invalid Signature");
    
    
	cancelled[_order.orderKey] = true;

    emit cancelledOrder(_order.orderKey);

	return true;
}
    
    
/// Release an agreement if term completed
/// @param orderKey: off-chain key generated as keccak hash of User + time + nonce
/// @param agreementKey: off-chain key generated as keccak hash of User + time + nonce
function release(bytes memory orderKey, bytes memory agreementKey) public returns(uint256){

	// Require swap state to be active
	// Require swap duration to have expired
	require(agreements[orderKey][agreementKey].state == 1, "Already released");
	require(block.timestamp >= agreements[orderKey][agreementKey].releaseTime, "Agreement term not yet complete");

	cErc20 cToken = cErc20(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb);
	Erc20 underlying = Erc20(agreements[orderKey][agreementKey].tokenAddress);

	// Logic for a floating-side maker
	if (agreements[orderKey][agreementKey].side == 1 ) {

		// Calculate annualized interest-rate generated by the swap agreement
		uint total = agreements[orderKey][agreementKey].principal.add(agreements[orderKey][agreementKey].interest);
		uint yield = ((cToken.exchangeRateCurrent().mul(100000000000000000000000000)).div(agreements[orderKey][agreementKey].initialRate)).sub(100000000000000000000000000);
		uint annualizedRate = ((yield.mul(31536000)).div(agreements[orderKey][agreementKey].duration));

		// In order to avoid subtraction underflow, ensures subtraction of smaller annualized rate
		if (agreements[orderKey][agreementKey].rate > annualizedRate) {

			// Calculates difference between annualized expected rate / real rate 
			uint rateDifference = (agreements[orderKey][agreementKey].rate).sub(annualizedRate);

			// Calculates differential in expected currency from previous rate differential
			uint annualFloatingDifference = (rateDifference.mul(total)).div(100000000000000000000000000);

			// De-annualizes the differential for the given time period
			uint floatingDifference = (annualFloatingDifference.div(31536000)).mul(agreements[orderKey][agreementKey].duration);

			// Calculates difference between value and expected interest
			uint floatingReturned = (agreements[orderKey][agreementKey].interest).sub(floatingDifference);

			// Redeems appropriate CTokens
			redeemCToken(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb,(total.add(floatingReturned)));

			// Returns funds to appropriate parties
			underlying.transfer(agreements[orderKey][agreementKey].maker, floatingReturned);
			underlying.transfer(agreements[orderKey][agreementKey].taker, total);

		}

		if (annualizedRate > agreements[orderKey][agreementKey].rate) {
			uint rateDifference = annualizedRate.sub(agreements[orderKey][agreementKey].rate);
			uint annualFloatingDifference = (rateDifference.mul(total)).div(100000000000000000000000000);
			uint floatingDifference = (annualFloatingDifference.div(31536000)).mul(agreements[orderKey][agreementKey].duration);
			uint floatingReturned = (agreements[orderKey][agreementKey].interest).add(floatingDifference);

			redeemCToken(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb,(total.add(floatingReturned)));

			underlying.transfer(agreements[orderKey][agreementKey].maker, floatingReturned);
			underlying.transfer(agreements[orderKey][agreementKey].taker, total);
			}

		if (annualizedRate == agreements[orderKey][agreementKey].rate) {

			redeemCToken(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb,(total.add(agreements[orderKey][agreementKey].interest)));
			
			underlying.transfer(agreements[orderKey][agreementKey].maker, agreements[orderKey][agreementKey].interest);
			underlying.transfer(agreements[orderKey][agreementKey].taker, total);
		}
	}

	// Logic for a fixed-side maker
	if (agreements[orderKey][agreementKey].side == 0 ) {

		// Calculate annualized interest-rate generated by the swap agreement
		uint total = agreements[orderKey][agreementKey].principal.add(agreements[orderKey][agreementKey].interest);
		uint yield = ((cToken.exchangeRateCurrent().mul(100000000000000000000000000)).div(agreements[orderKey][agreementKey].initialRate)).sub(100000000000000000000000000);
		uint annualizedRate = ((yield.mul(31536000)).div(agreements[orderKey][agreementKey].duration));

		// In order to avoid subtraction underflow, ensures subtraction of smaller annualized rate
		if (agreements[orderKey][agreementKey].rate > annualizedRate) {

			// Calculates difference between annualized expected rate / real rate 
			uint rateDifference = (agreements[orderKey][agreementKey].rate).sub(annualizedRate);

			// Calculates differential in expected currency from previous rate differential
			uint annualFloatingDifference = (rateDifference.mul(total)).div(100000000000000000000000000);

			// De-annualizes the differential for the given time period
			uint floatingDifference = (annualFloatingDifference.div(31536000)).mul(agreements[orderKey][agreementKey].duration);

			// Calculates difference between value and expected interest
			uint floatingReturned = (agreements[orderKey][agreementKey].interest).sub(floatingDifference);

			// Redeems appropriate CTokens
			redeemCToken(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb,(total.add(floatingReturned)));

			// Returns funds to appropriate parties
			underlying.transfer(agreements[orderKey][agreementKey].maker, total);
			underlying.transfer(agreements[orderKey][agreementKey].taker, floatingReturned);
		}

		if (annualizedRate > agreements[orderKey][agreementKey].rate) {
			uint rateDifference = annualizedRate.sub(agreements[orderKey][agreementKey].rate);
			uint annualFloatingDifference = (rateDifference.mul(total)).div(100000000000000000000000000);
			uint floatingDifference = (annualFloatingDifference.div(31536000)).mul(agreements[orderKey][agreementKey].duration);
			uint floatingReturned = (agreements[orderKey][agreementKey].interest).add(floatingDifference);

			redeemCToken(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb,(total.add(floatingReturned)));

			underlying.transfer(agreements[orderKey][agreementKey].maker, total);
			underlying.transfer(agreements[orderKey][agreementKey].taker, floatingReturned);
		}

		if (annualizedRate == agreements[orderKey][agreementKey].rate) {

			redeemCToken(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb,(total.add(agreements[orderKey][agreementKey].interest)));

			underlying.transfer(agreements[orderKey][agreementKey].maker, total);
			underlying.transfer(agreements[orderKey][agreementKey].taker, agreements[orderKey][agreementKey].interest);
		}
	}

	// Change state to Expired
	agreements[orderKey][agreementKey].state = 2;

	emit released(orderKey,agreementKey);

	return(agreements[orderKey][agreementKey].state);				
}

/// Release an agreement if term completed
/// @param orderKey: off-chain key generated as keccak hash of User + time + nonce
/// @param agreementKey: off-chain key generated as keccak hash of User + time + nonce
function releaseFixed(bytes memory orderKey, bytes memory agreementKey) public returns(uint256){

	// Require swap state to be active
	// Require swap duration to have expired
	require(agreements[orderKey][agreementKey].state == 1, "Already released");
	require(block.timestamp >= agreements[orderKey][agreementKey].releaseTime, "Agreement term not yet complete");

	cErc20 cToken = cErc20(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb);
	Erc20 underlying = Erc20(agreements[orderKey][agreementKey].tokenAddress);

	// Calculate annualized interest-rate generated by the swap agreement
	uint total = agreements[orderKey][agreementKey].principal.add(agreements[orderKey][agreementKey].interest);
	uint yield = ((cToken.exchangeRateCurrent().mul(100000000000000000000000000)).div(agreements[orderKey][agreementKey].initialRate)).sub(100000000000000000000000000);
	uint annualizedRate = ((yield.mul(31536000)).div(agreements[orderKey][agreementKey].duration));

	// In order to avoid subtraction underflow, ensures subtraction of smaller annualized rate
	if (agreements[orderKey][agreementKey].rate > annualizedRate) {

		// Calculates difference between annualized expected rate / real rate 
		uint rateDifference = (agreements[orderKey][agreementKey].rate).sub(annualizedRate);

		// Calculates differential in expected currency from previous rate differential
		uint annualFloatingDifference = (rateDifference.mul(total)).div(100000000000000000000000000);

		// De-annualizes the differential for the given time period
		uint floatingDifference = (annualFloatingDifference.div(31536000)).mul(agreements[orderKey][agreementKey].duration);

		// Calculates difference between value and expected interest
		uint floatingReturned = (agreements[orderKey][agreementKey].interest).sub(floatingDifference);

		// Redeems appropriate CTokens
		redeemCToken(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb,(total.add(floatingReturned)));

		// Returns funds to appropriate parties
		underlying.transfer(agreements[orderKey][agreementKey].maker, total);
		underlying.transfer(agreements[orderKey][agreementKey].taker, floatingReturned);
	}

	if (annualizedRate > agreements[orderKey][agreementKey].rate) {
		uint rateDifference = annualizedRate.sub(agreements[orderKey][agreementKey].rate);
		uint annualFloatingDifference = (rateDifference.mul(total)).div(100000000000000000000000000);
		uint floatingDifference = (annualFloatingDifference.div(31536000)).mul(agreements[orderKey][agreementKey].duration);
		uint floatingReturned = (agreements[orderKey][agreementKey].interest).add(floatingDifference);

		redeemCToken(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb,(total.add(floatingReturned)));

		underlying.transfer(agreements[orderKey][agreementKey].maker, total);
		underlying.transfer(agreements[orderKey][agreementKey].taker, floatingReturned);
	}

	if (annualizedRate == agreements[orderKey][agreementKey].rate) {

		redeemCToken(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb,(total.add(agreements[orderKey][agreementKey].interest)));

		underlying.transfer(agreements[orderKey][agreementKey].maker, total);
		underlying.transfer(agreements[orderKey][agreementKey].taker, agreements[orderKey][agreementKey].interest);
	}
	

	// Change state to Expired
	agreements[orderKey][agreementKey].state = 2;

	emit released(orderKey,agreementKey);

	return(agreements[orderKey][agreementKey].state);				
}

/// Release an agreement if term completed
/// @param orderKey: off-chain key generated as keccak hash of User + time + nonce
/// @param agreementKey: off-chain key generated as keccak hash of User + time + nonce
function releaseFloating(bytes memory orderKey, bytes memory agreementKey) public returns(uint256){

	// Require swap state to be active
	// Require swap duration to have expired
	require(agreements[orderKey][agreementKey].state == 1, "Already released");
	require(block.timestamp >= agreements[orderKey][agreementKey].releaseTime, "Agreement term not yet complete");

	cErc20 cToken = cErc20(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb);
	Erc20 underlying = Erc20(agreements[orderKey][agreementKey].tokenAddress);

	// Calculate annualized interest-rate generated by the swap agreement
	uint total = agreements[orderKey][agreementKey].principal.add(agreements[orderKey][agreementKey].interest);
	uint yield = ((cToken.exchangeRateCurrent().mul(100000000000000000000000000)).div(agreements[orderKey][agreementKey].initialRate)).sub(100000000000000000000000000);
	uint annualizedRate = ((yield.mul(31536000)).div(agreements[orderKey][agreementKey].duration));

	// In order to avoid subtraction underflow, ensures subtraction of smaller annualized rate
	if (agreements[orderKey][agreementKey].rate > annualizedRate) {

		// Calculates difference between annualized expected rate / real rate 
		uint rateDifference = (agreements[orderKey][agreementKey].rate).sub(annualizedRate);

		// Calculates differential in expected currency from previous rate differential
		uint annualFloatingDifference = (rateDifference.mul(total)).div(100000000000000000000000000);

		// De-annualizes the differential for the given time period
		uint floatingDifference = (annualFloatingDifference.div(31536000)).mul(agreements[orderKey][agreementKey].duration);

		// Calculates difference between value and expected interest
		uint floatingReturned = (agreements[orderKey][agreementKey].interest).sub(floatingDifference);

		// Redeems appropriate CTokens
		redeemCToken(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb,(total.add(floatingReturned)));

		// Returns funds to appropriate parties
		underlying.transfer(agreements[orderKey][agreementKey].maker, floatingReturned);
		underlying.transfer(agreements[orderKey][agreementKey].taker, total);

	}

	if (annualizedRate > agreements[orderKey][agreementKey].rate) {
		uint rateDifference = annualizedRate.sub(agreements[orderKey][agreementKey].rate);
		uint annualFloatingDifference = (rateDifference.mul(total)).div(100000000000000000000000000);
		uint floatingDifference = (annualFloatingDifference.div(31536000)).mul(agreements[orderKey][agreementKey].duration);
		uint floatingReturned = (agreements[orderKey][agreementKey].interest).add(floatingDifference);

		redeemCToken(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb,(total.add(floatingReturned)));

		underlying.transfer(agreements[orderKey][agreementKey].maker, floatingReturned);
		underlying.transfer(agreements[orderKey][agreementKey].taker, total);
		}

	if (annualizedRate == agreements[orderKey][agreementKey].rate) {

		redeemCToken(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb,(total.add(agreements[orderKey][agreementKey].interest)));
		
		underlying.transfer(agreements[orderKey][agreementKey].maker, agreements[orderKey][agreementKey].interest);
		underlying.transfer(agreements[orderKey][agreementKey].taker, total);
	}


	// Change state to Expired
	agreements[orderKey][agreementKey].state = 2;

	emit released(orderKey,agreementKey);

	return(agreements[orderKey][agreementKey].state);				
}	
	
/// Mint cToken
/// @param _erc20Contract: Address of ERC token used in minting
/// @param _numTokensToSupply: Number of tokens to mint with
function mintCToken(address _erc20Contract,uint _numTokensToSupply) internal returns (uint) {
        
	Erc20 underlying = Erc20(_erc20Contract);

	cErc20 cToken = cErc20(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb);

	// Approve transfer on the ERC20 contract
	underlying.approve(0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb, _numTokensToSupply);


	uint mintResult = cToken.mint(_numTokensToSupply);

	return mintResult;
}


/// Redeem cToken
/// @param _cErc20Contract: Address of cERC token to redeem
/// @param _numTokensToRedeem: Number of underlying tokens to redeem
function redeemCToken(address _cErc20Contract, uint _numTokensToRedeem) internal {
    
    cErc20(_cErc20Contract).redeemUnderlying(_numTokensToRedeem);
    
}
    
}