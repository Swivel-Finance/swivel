pragma solidity ^0.5.9;
pragma experimental ABIEncoderV2;

import "./safeMath.sol";

contract Erc20 {
  function approve(address, uint) external returns (bool);
  function transfer(address, uint) external returns (bool);
  function balanceOf(address _owner) external returns (uint256 balance);
  function transferFrom(address sender, address recipient, uint256 amount) public returns (bool);
}

contract CErc20 is Erc20 {
  function mint(uint) external returns (uint);
  function redeem(uint redeemTokens) external returns (uint);
  function redeemUnderlying(uint redeemAmount) external returns (uint);
  function exchangeRateCurrent() external returns (uint);
}

contract CEth {
  function mint() external payable;
  function redeem(uint redeemTokens) external returns (uint);
  function balanceOf(address _owner) external returns (uint256 balance);
  function approve(address spender, uint tokens) public returns (bool success);
  function redeemUnderlying(uint redeemAmount) external returns (uint);
  function exchangeRateCurrent() external returns (uint);
}

contract DefiHedge {
	
	struct RPCSig{
        uint8 v;
        bytes32 r;
        bytes32 s;
    }
    
    struct EIP712Domain {
        string  name;
        string  version;
        uint256 chainId;
        address verifyingContract;
    }
	
	struct Order {
    	address maker;
        uint256 side;
        address tokenAddress;
        uint256 duration;
        uint256 rate;
        uint256 interest;
        uint256 base;
        uint256 makerNonce;
        uint256 expiryTime;
        bytes32 makerOrderKey;
    }
    
    struct signedOrder {
    	address maker;
        uint256 side;
        address tokenAddress;
        uint256 duration;
        uint256 rate;
        uint256 interest;
        uint256 base;
        uint256 makerNonce;
        bytes32 orderKey;
        uint256 state;
        uint256 initialRate;
        uint256 expiryTime;
    }
    
    
 struct activeOrder {
        address maker;
        address taker;
        uint256 side;
        address tokenAddress;
        uint256 duration;
        uint256 rate;
        uint256 interest;
        uint256 base;
        uint256 makerNonce;
        uint256 state;
        uint256 lockTime;
        uint256 initialRate;
        bytes32 takerOrderKey;
    }    
    
    /// makerOrderMapping 
    ///@param maker order hash key
    mapping(bytes32 => signedOrder) public orderMapping;
    
    bytes32[] public makerList;
    
    mapping(bytes32 => bytes32[]) public takerListMapping;
    
    mapping(bytes32 => mapping(bytes32 => activeOrder)) takerMapping;
    
    /// @param Order hash.
    /// @return The amount of taker asset filled.
    mapping (bytes32 => uint256) public filled;

    /// @param Order hash.
    /// @return Whether the order was cancelled.
    mapping (bytes32 => bool) public cancelled;
    
    event newLockedOrder(
        address maker,
        address taker,
        uint256 side,
        address tokenAddress,
        uint256 duration,
        uint256 rate,
        uint256 interest,
        uint256 base,
        bytes32 makerOrderKey,
        bytes32 takerOrderKey
    );
    
    event Aborted(
        bytes32 orderKey
        );
        
    event orderReleased(
        bytes32 orderKey
        );
    

    using SafeMath for uint;

	constructor () public {
        DOMAIN_SEPARATOR = hashDomain(EIP712Domain({
            name: "DefiHedge",
            version: '1',
            chainId: 3,
            verifyingContract: 0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC
        }));
    }
    
    bytes32 DOMAIN_SEPARATOR;
    
    // Offer + EIP Domain Hash Schema
    bytes32 constant EIP712DOMAIN_TYPEHASH = keccak256(
        "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
    );

    bytes32 constant OFFER_TYPEHASH = keccak256(
        "Offer(address maker,address taker,uint256 side,address tokenAddress,uint256 duration,uint256 rate,uint256 interest,uint256 base)"
    );
    
    function hashDomain(EIP712Domain memory eip712Domain) internal pure returns (bytes32) {
        return keccak256(abi.encode(
            EIP712DOMAIN_TYPEHASH,
            keccak256(bytes(eip712Domain.name)),
            keccak256(bytes(eip712Domain.version)),
            eip712Domain.chainId,
            eip712Domain.verifyingContract
        ));
    }
    
    function hashOrder(Order memory _order)private pure returns(bytes32){
        return keccak256(abi.encode(
            OFFER_TYPEHASH,
            _order.maker,
            _order.side,
            _order.tokenAddress,
            _order.duration,
            _order.rate,
            _order.interest,
            _order.base
        ));
    }
    
    
    function getActiveOffer(bytes32 makerOrderKey,bytes32 takerOrderKey)
    public
    view
    returns (address maker, address taker, uint256 side, address tokenAddress, uint256 duration, uint256 rate, uint256 base, uint256 interest, uint256 state, uint256 lockTime, uint256 initialRate)
    {
        // Returns all offer details
        maker = takerMapping[makerOrderKey][takerOrderKey].maker;
        side = takerMapping[makerOrderKey][takerOrderKey].side;
        tokenAddress = takerMapping[makerOrderKey][takerOrderKey].tokenAddress;
        duration = takerMapping[makerOrderKey][takerOrderKey].duration;
        rate = takerMapping[makerOrderKey][takerOrderKey].rate;
        interest = takerMapping[makerOrderKey][takerOrderKey].interest;
        base = takerMapping[makerOrderKey][takerOrderKey].base;
        state = takerMapping[makerOrderKey][takerOrderKey].state;
        lockTime = takerMapping[makerOrderKey][takerOrderKey].lockTime;
        initialRate = takerMapping[makerOrderKey][takerOrderKey].initialRate;
        
        return (maker, taker, side, tokenAddress, duration, rate, base, interest, state, lockTime, initialRate);
    }
    
	
	
	function orderSettle(Order memory order, bytes32 takerOrderKey) private returns (bool){
	    

	    
	    signedOrder memory makerOrder;
	    
	    activeOrder memory takerOrder;
	    
	    CErc20 cToken = CErc20(0xdb5Ed4605C11822811a39F94314fDb8F0fb59A2C); //DAI cToken Address

	        // Check trades side
	        // Transfers funds to DefiHedge contract
    	    Erc20 underlying = Erc20(order.tokenAddress);
    	    if (order.side == 0) {
	        require(underlying.transferFrom(order.maker, address(this), order.base), "Transfer Failed!");
            require(underlying.transferFrom(msg.sender, address(this), order.interest), "Transfer Failed!");
    	    }
    	    if (order.side == 1) {
	        require(underlying.transferFrom(order.maker, address(this), order.interest), "Transfer Failed!");
            require(underlying.transferFrom(msg.sender, address(this), order.base), "Transfer Failed!");
    	    }
    	    
    	    // Mint CToken from DefiHedge contract
            mintCToken(order.tokenAddress,order.interest.add(order.base));
            
    	    
    	    // Instantiate makerOrder
    	    makerOrder.maker = order.maker;
    	    makerOrder.side = order.side;
    	    makerOrder.tokenAddress = order.tokenAddress;
    	    makerOrder.duration = order.duration;
    	    makerOrder.rate = order.rate;
    	    makerOrder.interest = order.interest;
    	    makerOrder.base = order.base;
    	    makerOrder.state = 1;
    	    makerOrder.initialRate = cToken.exchangeRateCurrent();
    	    
    	    // Instantiate takerOrder
    	    takerOrder.maker = order.maker;
    	    takerOrder.taker = msg.sender;
    	    takerOrder.side = order.side;
    	    takerOrder.tokenAddress = order.tokenAddress;
    	    takerOrder.duration = order.duration;
    	    takerOrder.rate = order.rate;
    	    takerOrder.interest = order.interest;
    	    takerOrder.base = order.base;
    	    takerOrder.makerNonce = order.makerNonce;
    	    takerOrder.state = 1;
    	    takerOrder.lockTime = now.add(order.duration);
    	    takerOrder.initialRate = cToken.exchangeRateCurrent();
    	    
    	    
    	    
    	    orderMapping[order.makerOrderKey] = makerOrder;
    	    
    	    takerMapping[order.makerOrderKey][takerOrderKey]= takerOrder;
    	    
	    
	        makerList.push(order.makerOrderKey);
	    
	        takerListMapping[order.makerOrderKey].push(takerOrderKey);
	    
	   emit newLockedOrder(takerOrder.maker,msg.sender,takerOrder.side,takerOrder.tokenAddress,takerOrder.duration,takerOrder.rate,takerOrder.interest,takerOrder.base,order.makerOrderKey,takerOrderKey);
    	    
        return true;
	    
	    
	}
	
	
	/// @param
	/// makerOrder = maker's order 
	/// takerOrderKey = off-chain key generated as keccak hash of User + time + nonce
	/// makerSignature = signature associated with order param
	///Fill the entirety of a maker order's volume
	function fillOffer(Order memory makerOrder, bytes32 takerOrderKey,bytes memory makerSignature) 	public returns (uint256){
	    
	    //Check if order already partially filled
	    require(orderMapping[makerOrder.makerOrderKey].state != 1, "Order Already Partial/Fully Filled");
	    
	    //Check if order has been cancelled
	    require(cancelled[makerOrder.makerOrderKey]==false, "Order Has Been Cancelled");
	    
	    //Check if order has already expired
	    require(orderMapping[makerOrder.makerOrderKey].expiryTime >= now, "Order Has Expired");
	    
	    /// Instantiate offer
	    Order memory filledOrder = Order(
	                                                     makerOrder.maker,
	                                                     makerOrder.side,
	                                                     makerOrder.tokenAddress,
	                                                     makerOrder.duration,
	                                                     makerOrder.rate,
	                                                     makerOrder.interest,
	                                                     makerOrder.base,
	                                                     makerOrder.makerNonce,
	                                                     makerOrder.expiryTime,
	                                                     makerOrder.makerOrderKey
	                                                     );
	                                                     
	     // Parse signature into R,S,V                        
	    RPCSig memory RPCsig = signatureRPC(makerSignature);
	    
	     // Validate offer signature & ensure it was created by maker
	    require(makerOrder.maker == ecrecover(
	        keccak256(abi.encodePacked(
            "\x19\x01",
            DOMAIN_SEPARATOR,
            hashOrder(filledOrder)
            )),
            RPCsig.v,
            RPCsig.r,
            RPCsig.s), "Invalid Signature");

	    // Settle Response
	    orderSettle(filledOrder,takerOrderKey);
        
	    
	}
	
	
	/// @param
	/// makerOrder = maker's order 
	/// takerVolume = amount of currency being taken
	/// takerOrderKey = off-chain key generated as keccak hash of User + time + nonce
	/// makerSignature = signature associated with order param
	/// Fill partial maker order 
	function partialFillOffer(Order memory makerOrder,uint256 takerVolume, bytes32 takerOrderKey, bytes memory makerSignature ) public returns (uint256){
	    
    	//instantiate taker's order
        activeOrder memory takerorder;
	    
	    //Check if order has been cancelled
	    require(cancelled[makerOrder.makerOrderKey]==false, "Order Has Been Cancelled");
	    
	    //Check if order has already expired
	    require(makerOrder.expiryTime >= now, "Order Has Expired");
	    

	    //If order is fixed-side, ensure <= expected total interest and then instantiate order
        if (orderMapping[makerOrder.makerOrderKey].side == 0) {
        
        require (takerVolume <= (orderMapping[makerOrder.makerOrderKey].interest), "Taker Volume > Maker");
        
            // if order has already been partially filled
    	    if (orderMapping[makerOrder.makerOrderKey].state != 0) {
    	        
            require (takerVolume <= (orderMapping[makerOrder.makerOrderKey].interest - filled[makerOrder.makerOrderKey]), "Taker Volume > Available");
    	    }
    	
    	
            takerorder.maker=makerOrder.maker;
            takerorder.taker=msg.sender;
            takerorder.side=makerOrder.side;
            takerorder.tokenAddress=makerOrder.tokenAddress;
            takerorder.duration=makerOrder.duration;
            takerorder.rate=makerOrder.rate;
            takerorder.interest= takerVolume;
            //calculate taker % of total maker order and set opposing param 
            uint256 orderRatio = (((takerVolume).mul(100000000000000000000000000)).div(makerOrder.interest)).div(100000000000000000000000000);
            takerorder.base=makerOrder.base.mul(orderRatio);
            takerorder.takerOrderKey=takerOrderKey;
    	
        }
        
        //If order is floating-side, ensure < expected principal
        if (orderMapping[makerOrder.makerOrderKey].side == 1) {
        
        require (takerVolume <= (orderMapping[makerOrder.makerOrderKey].base), "Taker Volume > Maker");
        
            // if order has already been partially filled
            if (orderMapping[makerOrder.makerOrderKey].state != 0) {
    	        
            require (takerVolume <= (orderMapping[makerOrder.makerOrderKey].base - filled[makerOrder.makerOrderKey]), "Taker Volume > Available");
    	    }
    	    
    	    	
            takerorder.maker=makerOrder.maker;
            takerorder.taker=msg.sender;
            takerorder.side=makerOrder.side;
            takerorder.tokenAddress=makerOrder.tokenAddress;
            takerorder.duration=makerOrder.duration;
            takerorder.rate=makerOrder.rate;
            //calculate taker % of total maker order and set opposing param 
            uint256 orderRatio = (((takerVolume).mul(100000000000000000000000000)).div(makerOrder.base)).div(100000000000000000000000000);
            takerorder.interest= makerOrder.interest.mul(orderRatio);
            takerorder.base= takerVolume;
            takerorder.takerOrderKey=takerOrderKey;
        }
    
                                             
	     // Parse signature into R,S,V                        
	    RPCSig memory RPCsig = signatureRPC(makerSignature);
	    
	     // Validate offer signature & ensure it was created by maker
	    require(makerOrder.maker == ecrecover(
	        keccak256(abi.encodePacked(
            "\x19\x01",
            DOMAIN_SEPARATOR,
            hashOrder(makerOrder)
            )),
            RPCsig.v,
            RPCsig.r,
            RPCsig.s), "Invalid Signature");
        

	    // Settle Response
	    partialOrderSettle(takerorder, makerOrder.makerOrderKey);
	    
    }
    
    
    ///settlement function for 
    function partialOrderSettle(activeOrder memory takerOrder, bytes32 makerOrderKey) private returns (bool){
	    

	    signedOrder memory makerOrder;
	    
	    CErc20 cToken = CErc20(0xdb5Ed4605C11822811a39F94314fDb8F0fb59A2C); //DAI cToken Address

	        // Check trades side
	        // Transfers funds to DefiHedge contract
    	    Erc20 underlying = Erc20(takerOrder.tokenAddress);
    	    if (takerOrder.side == 0) {
    	        require(underlying.transferFrom(takerOrder.maker, address(this), takerOrder.base), "Transfer Failed!");
                require(underlying.transferFrom(msg.sender, address(this), takerOrder.interest), "Transfer Failed!");
    	    }
    	    if (takerOrder.side == 1) {
    	        require(underlying.transferFrom(takerOrder.maker, address(this), takerOrder.interest), "Transfer Failed!");
                require(underlying.transferFrom(msg.sender, address(this), takerOrder.base), "Transfer Failed!");
    	    }
    	    
    	    // Mint CToken from DefiHedge contract
            mintCToken(takerOrder.tokenAddress,takerOrder.interest.add(takerOrder.base));
            
            
            // After mint success fill order params & store takerOrder in takerMapping
    	    takerOrder.state = 1;  /// Set state to active
    	    takerOrder.lockTime = now.add(takerOrder.duration); /// Set locktime
    	    takerOrder.initialRate = cToken.exchangeRateCurrent();  /// Get initial exchange rate
    	    
    	    takerMapping[makerOrderKey][takerOrder.takerOrderKey]= takerOrder;
            
            
            // Instantiate & store makerOrder if this is first fill
            if (orderMapping[makerOrderKey].maker == address(0x0000000000000000000000000000000000000000)) {
    	    
        	    makerOrder.maker = takerOrder.maker;
        	    makerOrder.side = takerOrder.side;
        	    makerOrder.tokenAddress = takerOrder.tokenAddress;
        	    makerOrder.duration = takerOrder.duration;
        	    makerOrder.rate = takerOrder.rate;
        	    makerOrder.interest = takerOrder.interest;
        	    makerOrder.base = takerOrder.base;
        	    makerOrder.state = 1; /// Set state to active
        	    makerOrder.initialRate = takerOrder.initialRate; /// Get initial exchange rate
        	   
        	    orderMapping[makerOrderKey] = makerOrder;
            }
            
            
    	    //push makerOrderKey to general order list (for testing)
	        makerList.push(makerOrderKey);
	        
	        //push takerOrderKey to takerOrderList nested in mapping
	        takerListMapping[makerOrderKey].push(takerOrder.takerOrderKey);
	    
	   emit newLockedOrder(takerOrder.maker,msg.sender,takerOrder.side,takerOrder.tokenAddress,takerOrder.duration,takerOrder.rate,takerOrder.interest,takerOrder.base,makerOrderKey,takerOrder.takerOrderKey);
    	    
        return true;
	    
	    
	}
    
    
    function cancelOrder(bytes32 makerOrderKey, Order memory order, bytes memory makerSignature) public returns(bool){
	    
	    // Parse signature into R,S,V                        
	    RPCSig memory RPCsig = signatureRPC(makerSignature);
	    
	     // Validate offer signature & ensure it was created by maker
	    require(msg.sender == ecrecover(
	        keccak256(abi.encodePacked(
            "\x19\x01",
            DOMAIN_SEPARATOR,
            hashOrder(order)
            )),
            RPCsig.v,
            RPCsig.r,
            RPCsig.s), "Invalid Signature");
            
        cancelled[makerOrderKey]== true;
	    
	    
	    return true;
	}
    
    
	 /// Release an Erc bond once if term completed
	function releaseErcOrder(bytes32 makerOrderKey, bytes32 takerOrderKey)
	public
	returns(uint256)
	{

        
        // Require swap state to be active
        // Require swap duration to have expired
        require(takerMapping[makerOrderKey][takerOrderKey].state == 1, "Invalid State");
		require(now >= takerMapping[makerOrderKey][takerOrderKey].lockTime, "Invalid Time");
		
	    CErc20 cToken = CErc20(0xdb5Ed4605C11822811a39F94314fDb8F0fb59A2C);
	    Erc20 underlying = Erc20(orderMapping[makerOrderKey].tokenAddress);
	    
	    if (orderMapping[makerOrderKey].side == 0 ) {
	    
    	    // Calculate annualized interest-rate generated by the swap agreement
    		uint total = takerMapping[makerOrderKey][takerOrderKey].base.add(takerMapping[makerOrderKey][takerOrderKey].interest);
    		uint yield = ((cToken.exchangeRateCurrent().mul(100000000000000000000000000)).div(takerMapping[makerOrderKey][takerOrderKey].initialRate)).sub(100000000000000000000000000);
    		uint annualizedRate = ((yield.mul(31536000)).div(orderMapping[makerOrderKey].duration));
    
    		// In order to avoid subtraction underflow, ensures subtraction of smaller annualized rate
    		if (orderMapping[makerOrderKey].rate > annualizedRate) {
    		    
    		    // Calculates difference between annualized expected rate / real rate 
        	    uint rateDifference = (orderMapping[makerOrderKey].rate).sub(annualizedRate);
        	    
        	    // Calculates differential in expected currency from previous rate differential
        	    uint annualFloatingDifference = (rateDifference.mul(total)).div(100000000000000000000000000);
        	    
        	 	// De-annualizes the differential for the given time period
        	    uint floatingDifference = (annualFloatingDifference.div(31536000)).mul(orderMapping[makerOrderKey].duration);
        	    
        	    // Calculates difference between value and expected interest
                uint floatingReturned = (takerMapping[makerOrderKey][takerOrderKey].interest).sub(floatingDifference);
                
                // Redeems appropriate CTokens
    		    redeemCToken(0xdb5Ed4605C11822811a39F94314fDb8F0fb59A2C,(total.add(floatingReturned)));
    		    
    		    // Returns funds to appropriate parties
                if (orderMapping[makerOrderKey].side == 0){
        		    underlying.transfer(orderMapping[makerOrderKey].maker, total);
        		    underlying.transfer(takerMapping[makerOrderKey][takerOrderKey].taker, floatingReturned);
    		    }
    		    if (orderMapping[makerOrderKey].side == 1){
        		    underlying.transfer(orderMapping[makerOrderKey].maker, floatingReturned);
        		    underlying.transfer(takerMapping[makerOrderKey][takerOrderKey].taker, total);
    		    }
    		}
    		
    		if (annualizedRate > orderMapping[makerOrderKey].rate) {
        	    uint rateDifference = annualizedRate.sub(orderMapping[makerOrderKey].rate);
        	    uint annualFloatingDifference = (rateDifference.mul(total)).div(100000000000000000000000000);
        	    uint floatingDifference = (annualFloatingDifference.div(31536000)).mul(orderMapping[makerOrderKey].duration);
                uint floatingReturned = (orderMapping[makerOrderKey].interest).add(floatingDifference);
    
        	    redeemCToken(0xdb5Ed4605C11822811a39F94314fDb8F0fb59A2C,(total.add(floatingReturned)));
        	    
                if (orderMapping[makerOrderKey].side == 0){
        		    underlying.transfer(orderMapping[makerOrderKey].maker, total);
        		    underlying.transfer(takerMapping[makerOrderKey][takerOrderKey].taker, floatingReturned);
    		    }
    		    if (orderMapping[makerOrderKey].side == 1){
        		    underlying.transfer(orderMapping[makerOrderKey].maker, floatingReturned);
        		    underlying.transfer(takerMapping[makerOrderKey][takerOrderKey].taker, total);
    		    }
    		}
    		
    		if (annualizedRate == orderMapping[makerOrderKey].rate) {
    		    
    		    redeemCToken(0xdb5Ed4605C11822811a39F94314fDb8F0fb59A2C,(total.add(orderMapping[makerOrderKey].interest)));
    		    
                if (orderMapping[makerOrderKey].side == 0){
        		    underlying.transfer(orderMapping[makerOrderKey].maker, total);
        		    underlying.transfer(takerMapping[makerOrderKey][takerOrderKey].taker, takerMapping[makerOrderKey][takerOrderKey].interest);
    		    }
    		    if (orderMapping[makerOrderKey].side == 1){
        		    underlying.transfer(orderMapping[makerOrderKey].maker, takerMapping[makerOrderKey][takerOrderKey].interest);
        		    underlying.transfer(takerMapping[makerOrderKey][takerOrderKey].taker, total);
    		    }
    		}
	    }
	    
	    if (orderMapping[makerOrderKey].side == 1 ) {
	    
    	    // Calculate annualized interest-rate generated by the swap agreement
    		uint total = takerMapping[makerOrderKey][takerOrderKey].base.add(takerMapping[makerOrderKey][takerOrderKey].interest);
    		uint yield = ((cToken.exchangeRateCurrent().mul(100000000000000000000000000)).div(takerMapping[makerOrderKey][takerOrderKey].initialRate)).sub(100000000000000000000000000);
    		uint annualizedRate = ((yield.mul(31536000)).div(orderMapping[makerOrderKey].duration));
    
    		// In order to avoid subtraction underflow, ensures subtraction of smaller annualized rate
    		if (orderMapping[makerOrderKey].rate > annualizedRate) {
    		    
    		    // Calculates difference between annualized expected rate / real rate 
        	    uint rateDifference = (orderMapping[makerOrderKey].rate).sub(annualizedRate);
        	    
        	    // Calculates differential in expected currency from previous rate differential
        	    uint annualFloatingDifference = (rateDifference.mul(total)).div(100000000000000000000000000);
        	    
        	 	// De-annualizes the differential for the given time period
        	    uint floatingDifference = (annualFloatingDifference.div(31536000)).mul(orderMapping[makerOrderKey].duration);
        	    
        	    // Calculates difference between value and expected interest
                uint floatingReturned = (takerMapping[makerOrderKey][takerOrderKey].interest).sub(floatingDifference);
                
                // Redeems appropriate CTokens
    		    redeemCToken(0xdb5Ed4605C11822811a39F94314fDb8F0fb59A2C,(total.add(floatingReturned)));
    		    
    		    // Returns funds to appropriate parties
                if (orderMapping[makerOrderKey].side == 0){
        		    underlying.transfer(orderMapping[makerOrderKey].maker, total);
        		    underlying.transfer(takerMapping[makerOrderKey][takerOrderKey].taker, floatingReturned);
    		    }
    		    if (orderMapping[makerOrderKey].side == 1){
        		    underlying.transfer(orderMapping[makerOrderKey].maker, floatingReturned);
        		    underlying.transfer(takerMapping[makerOrderKey][takerOrderKey].taker, total);
    		    }
    		}
    		
    		if (annualizedRate > orderMapping[makerOrderKey].rate) {
        	    uint rateDifference = annualizedRate.sub(orderMapping[makerOrderKey].rate);
        	    uint annualFloatingDifference = (rateDifference.mul(total)).div(100000000000000000000000000);
        	    uint floatingDifference = (annualFloatingDifference.div(31536000)).mul(orderMapping[makerOrderKey].duration);
                uint floatingReturned = (orderMapping[makerOrderKey].interest).add(floatingDifference);
    
        	    redeemCToken(0xdb5Ed4605C11822811a39F94314fDb8F0fb59A2C,(total.add(floatingReturned)));
        	    
                if (orderMapping[makerOrderKey].side == 0){
        		    underlying.transfer(orderMapping[makerOrderKey].maker, total);
        		    underlying.transfer(takerMapping[makerOrderKey][takerOrderKey].taker, floatingReturned);
    		    }
    		    if (orderMapping[makerOrderKey].side == 1){
        		    underlying.transfer(orderMapping[makerOrderKey].maker, floatingReturned);
        		    underlying.transfer(takerMapping[makerOrderKey][takerOrderKey].taker, total);
    		    }
    		}
    		
    		if (annualizedRate == orderMapping[makerOrderKey].rate) {
    		    
    		    redeemCToken(0xdb5Ed4605C11822811a39F94314fDb8F0fb59A2C,(total.add(orderMapping[makerOrderKey].interest)));
    		    
                if (orderMapping[makerOrderKey].side == 0){
        		    underlying.transfer(orderMapping[makerOrderKey].maker, total);
        		    underlying.transfer(takerMapping[makerOrderKey][takerOrderKey].taker, takerMapping[makerOrderKey][takerOrderKey].interest);
    		    }
    		    if (orderMapping[makerOrderKey].side == 1){
        		    underlying.transfer(orderMapping[makerOrderKey].maker, takerMapping[makerOrderKey][takerOrderKey].interest);
        		    underlying.transfer(takerMapping[makerOrderKey][takerOrderKey].taker, total);
    		    }
    		}
	    }
	    
	    
	    
	    
		// Change state to Expired
		takerMapping[makerOrderKey][takerOrderKey].state = 2;
		
    	emit orderReleased(takerOrderKey);

		
		return(orderMapping[makerOrderKey].state);
		
		
	}
	
		/// Mint cToken
	function mintCToken(
      address _erc20Contract,
      uint _numTokensToSupply
    ) internal returns (uint) {
        
      Erc20 underlying = Erc20(_erc20Contract);

      CErc20 cToken = CErc20(0xdb5Ed4605C11822811a39F94314fDb8F0fb59A2C);
      
      // Approve transfer on the ERC20 contract
      underlying.approve(0xdb5Ed4605C11822811a39F94314fDb8F0fb59A2C, _numTokensToSupply);
      
      
      uint mintResult = cToken.mint(_numTokensToSupply);
      
      return mintResult;
    }
    	    
	    /// Redeem cToken
	function redeemCToken(
        address _cErc20Contract, uint _numTokensToRedeem) internal {
        CErc20(_cErc20Contract).redeemUnderlying(_numTokensToRedeem);
	}
	
	// Splits signature into RSV
	function signatureRPC(bytes memory sig)internal pure returns (RPCSig memory RPCsig){
        bytes32 r;
        bytes32 s;
        uint8 v;
        
        if (sig.length != 65) {
          return RPCSig(0,'0','0');
        }
    
        assembly {
          r := mload(add(sig, 32))
          s := mload(add(sig, 64))
          v := and(mload(add(sig, 65)), 255)
        }
        
        if (v < 27) {
          v += 27;
        }
                
        if (v == 39 || v == 40) {
          v = v-12;
        }
        
        if (v != 27 && v != 28) {
          return RPCSig(0,'0','0');
        }
        
        return RPCSig(v,r,s);
    }

}

