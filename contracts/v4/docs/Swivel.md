# Swivel

## Contents

<!-- START doctoc -->
<!-- END doctoc -->

## Globals

> Note this list contains both internal and external attributes

| Var              | Type                                            |
| ---------------- | ----------------------------------------------- |
| cancelled        | mapping(bytes32 => bool)                        |
| filled           | mapping(bytes32 => uint256)                     |
| withdrawals      | mapping(address => uint256)                     |
| approvals        | mapping(address => mapping(address => uint256)) |
| NAME             | string                                          |
| VERSION          | string                                          |
| HOLD             | uint256                                         |
| feeChange        | uint256                                         |
| domain           | bytes32                                         |
| marketPlace      | address                                         |
| admin            | address                                         |
| aaveAddr         | address                                         |
| MIN_FEENOMINATOR | uint16                                          |
| feenominators    | uint16[4]                                       |

## Modifiers

### authorized

Restricts `msg.sender` as the only viable caller of a method

#### Declaration

```solidity
modifier authorized
```

## Functions

### constructor

#### Declaration

```solidity
function constructor(
address m,
address a
) public
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description                                                    |
| --- | ------- | -------------------------------------------------------------- |
| `m` | address | Deployed MarketPlace contract address                          |
| `a` | address | Address of a deployed Aave contract implementing our interface |

### initiate

Allows a user to initiate a position

#### Declaration

```solidity
function initiate(
struct Hash.Order[] o,
uint256[] a,
struct Sig.Components[] c
) external returns
(bool)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type                    | Description                                                         |
| --- | ----------------------- | ------------------------------------------------------------------- |
| `o` | struct Hash.Order[]     | Array of offline Swivel.Orders                                      |
| `a` | uint256[]               | Array of order volume (principal) amounts relative to passed orders |
| `c` | struct Sig.Components[] | Array of Components from valid ECDSA signatures                     |

### initiateVaultFillingZcTokenInitiate

Allows a user to initiate a Vault by filling an offline zcToken initiate order

> This method should pass (underlying, maturity, maker, sender, principalFilled) to MarketPlace.custodialInitiate

#### Declaration

```solidity
function initiateVaultFillingZcTokenInitiate(
struct Hash.Order o,
uint256 a,
struct Sig.Components c
) internal
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type                  | Description                                                     |
| --- | --------------------- | --------------------------------------------------------------- |
| `o` | struct Hash.Order     | Order being filled                                              |
| `a` | uint256               | Amount of volume (premium) being filled by the taker's initiate |
| `c` | struct Sig.Components | Components of a valid ECDSA signature                           |

### initiateZcTokenFillingVaultInitiate

Allows a user to initiate a zcToken by filling an offline vault initiate order

> This method should pass (underlying, maturity, sender, maker, a) to MarketPlace.custodialInitiate

#### Declaration

```solidity
function initiateZcTokenFillingVaultInitiate(
struct Hash.Order o,
uint256 a,
struct Sig.Components c
) internal
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type                  | Description                                                       |
| --- | --------------------- | ----------------------------------------------------------------- |
| `o` | struct Hash.Order     | Order being filled                                                |
| `a` | uint256               | Amount of volume (principal) being filled by the taker's initiate |
| `c` | struct Sig.Components | Components of a valid ECDSA signature                             |

### initiateZcTokenFillingZcTokenExit

Allows a user to initiate zcToken? by filling an offline zcToken exit order

> This method should pass (underlying, maturity, maker, sender, a) to MarketPlace.p2pZcTokenExchange

#### Declaration

```solidity
function initiateZcTokenFillingZcTokenExit(
struct Hash.Order o,
uint256 a,
struct Sig.Components c
) internal
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type                  | Description                                                       |
| --- | --------------------- | ----------------------------------------------------------------- |
| `o` | struct Hash.Order     | Order being filled                                                |
| `a` | uint256               | Amount of volume (principal) being filled by the taker's initiate |
| `c` | struct Sig.Components | Components of a valid ECDSA signature                             |

### initiateVaultFillingVaultExit

Allows a user to initiate a Vault by filling an offline vault exit order

> This method should pass (underlying, maturity, maker, sender, principalFilled) to MarketPlace.p2pVaultExchange

#### Declaration

```solidity
function initiateVaultFillingVaultExit(
struct Hash.Order o,
uint256 a,
struct Sig.Components c
) internal
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type                  | Description                                                  |
| --- | --------------------- | ------------------------------------------------------------ |
| `o` | struct Hash.Order     | Order being filled                                           |
| `a` | uint256               | Amount of volume (interest) being filled by the taker's exit |
| `c` | struct Sig.Components | Components of a valid ECDSA signature                        |

### exit

Allows a user to exit (sell) a currently held position to the marketplace.

#### Declaration

```solidity
function exit(
struct Hash.Order[] o,
uint256[] a,
struct Sig.Components[] c
) external returns
(bool)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type                    | Description                                                         |
| --- | ----------------------- | ------------------------------------------------------------------- |
| `o` | struct Hash.Order[]     | Array of offline Swivel.Orders                                      |
| `a` | uint256[]               | Array of order volume (principal) amounts relative to passed orders |
| `c` | struct Sig.Components[] | Components of a valid ECDSA signature                               |

### exitZcTokenFillingZcTokenInitiate

Allows a user to exit their zcTokens by filling an offline zcToken initiate order

> This method should pass (underlying, maturity, sender, maker, principalFilled) to MarketPlace.p2pZcTokenExchange

#### Declaration

```solidity
function exitZcTokenFillingZcTokenInitiate(
struct Hash.Order o,
uint256 a,
struct Sig.Components c
) internal
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type                  | Description                                                  |
| --- | --------------------- | ------------------------------------------------------------ |
| `o` | struct Hash.Order     | Order being filled                                           |
| `a` | uint256               | Amount of volume (interest) being filled by the taker's exit |
| `c` | struct Sig.Components | Components of a valid ECDSA signature                        |

### exitVaultFillingVaultInitiate

Allows a user to exit their Vault by filling an offline vault initiate order

> This method should pass (underlying, maturity, sender, maker, a) to MarketPlace.p2pVaultExchange

#### Declaration

```solidity
function exitVaultFillingVaultInitiate(
struct Hash.Order o,
uint256 a,
struct Sig.Components c
) internal
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type                  | Description                                                   |
| --- | --------------------- | ------------------------------------------------------------- |
| `o` | struct Hash.Order     | Order being filled                                            |
| `a` | uint256               | Amount of volume (principal) being filled by the taker's exit |
| `c` | struct Sig.Components | Components of a valid ECDSA signature                         |

### exitVaultFillingZcTokenExit

Allows a user to exit their Vault filling an offline zcToken exit order

> This method should pass (underlying, maturity, maker, sender, a) to MarketPlace.exitFillingExit

#### Declaration

```solidity
function exitVaultFillingZcTokenExit(
struct Hash.Order o,
uint256 a,
struct Sig.Components c
) internal
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type                  | Description                                                   |
| --- | --------------------- | ------------------------------------------------------------- |
| `o` | struct Hash.Order     | Order being filled                                            |
| `a` | uint256               | Amount of volume (principal) being filled by the taker's exit |
| `c` | struct Sig.Components | Components of a valid ECDSA signature                         |

### exitZcTokenFillingVaultExit

Allows a user to exit their zcTokens by filling an offline vault exit order

> This method should pass (underlying, maturity, sender, maker, principalFilled) to MarketPlace.exitFillingExit

#### Declaration

```solidity
function exitZcTokenFillingVaultExit(
struct Hash.Order o,
uint256 a,
struct Sig.Components c
) internal
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type                  | Description                                                  |
| --- | --------------------- | ------------------------------------------------------------ |
| `o` | struct Hash.Order     | Order being filled                                           |
| `a` | uint256               | Amount of volume (interest) being filled by the taker's exit |
| `c` | struct Sig.Components | Components of a valid ECDSA signature                        |

### cancel

Allows a user to cancel an order, preventing it from being filled in the future

#### Declaration

```solidity
function cancel(
struct Hash.Order[] o
) external returns
(bool)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type                | Description                             |
| --- | ------------------- | --------------------------------------- |
| `o` | struct Hash.Order[] | Array of offline orders being cancelled |

### setAdmin

Sets the admin

#### Declaration

```solidity
function setAdmin(
address a
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type    | Description            |
| --- | ------- | ---------------------- |
| `a` | address | Address of a new admin |

### scheduleWithdrawal

Allows the admin to schedule the withdrawal of tokens

#### Declaration

```solidity
function scheduleWithdrawal(
address e
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type    | Description                          |
| --- | ------- | ------------------------------------ |
| `e` | address | Address of (ERC20) token to withdraw |

### scheduleApproval

Allows the admin to schedule the approval of tokens

#### Declaration

```solidity
function scheduleApproval(
address e,
address a
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type    | Description                         |
| --- | ------- | ----------------------------------- |
| `e` | address | Address of (ERC20) token to approve |
| `a` | address | Address of the contract to approve  |

### scheduleFeeChange

allows the admin to schedule a change to the fee denominators

#### Declaration

```solidity
function scheduleFeeChange(
uint16[4] f
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type      | Description                                                                                                  |
| --- | --------- | ------------------------------------------------------------------------------------------------------------ |
| `f` | uint16[4] | array of length 4 holding values which suggest replacing any at the same index for the current feenominators |

### blockWithdrawal

Emergency function to block unplanned withdrawals

#### Declaration

```solidity
function blockWithdrawal(
address e
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type    | Description                          |
| --- | ------- | ------------------------------------ |
| `e` | address | Address of token withdrawal to block |

### blockApproval

Emergency function to block unplanned approvals

#### Declaration

```solidity
function blockApproval(
address e,
address a
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type    | Description                                  |
| --- | ------- | -------------------------------------------- |
| `e` | address | Address of token approval to block           |
| `a` | address | Address of the contract to block approval of |

### blockFeeChange

Emergency function to block unplanned changes to fee structure

#### Declaration

```solidity
function blockFeeChange(
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

### withdraw

Allows the admin to withdraw the given token, provided the holding period has been observed

#### Declaration

```solidity
function withdraw(
address e
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type    | Description                  |
| --- | ------- | ---------------------------- |
| `e` | address | Address of token to withdraw |

### changeFee

allows the admin to set new fee denominators

> note that, since 0 values are allowable the way to leave a feenominator value unchanged is to pass the existing value

#### Declaration

```solidity
function changeFee(
uint16[4] f
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type      | Description                                                                                            |
| --- | --------- | ------------------------------------------------------------------------------------------------------ |
| `f` | uint16[4] | array of length 4 holding values which will replace any at the same index in the current feenominators |

### approveUnderlying

Allows the admin to bulk approve given compounding addresses at the underlying token, saving marginal approvals,
providing the holding period has been observed

#### Declaration

```solidity
function approveUnderlying(
address[] u,
address[] c
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type      | Description                         |
| --- | --------- | ----------------------------------- |
| `u` | address[] | array of underlying token addresses |
| `c` | address[] | array of compound token addresses   |

### splitUnderlying

Allows users to deposit underlying and in the process split it into/mint
zcTokens and vault notional. Calls mPlace.mintZcTokenAddingNotional

#### Declaration

```solidity
function splitUnderlying(
uint8 p,
address u,
uint256 m,
uint256 a
) external returns
(bool)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description                                               |
| --- | ------- | --------------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market pair      |
| `u` | address | Underlying token address associated with this market pair |
| `m` | uint256 | Maturity timestamp of this associated market              |
| `a` | uint256 | Amount of underlying being deposited                      |

### combineTokens

Allows users deposit/burn 1-1 amounts of both zcTokens and vault notional,
in the process "combining" the two, and redeeming underlying. Calls mPlace.burnZcTokenRemovingNotional.

#### Declaration

```solidity
function combineTokens(
uint8 p,
address u,
uint256 m,
uint256 a
) external returns
(bool)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description                                          |
| --- | ------- | ---------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market pair |
| `u` | address | Underlying token address associated with the market  |
| `m` | uint256 | Maturity timestamp of the market                     |
| `a` | uint256 | Amount of zcTokens being redeemed                    |

### authRedeem

Allows MarketPlace to complete its contractual obligation as IRedeemer, redeeming zcTokens and withdrawing underlying
p Protocol Enum value associated with this market pair

> Note that this bubbles up from the zcToken instead of starting on Swivel (as per the ERC5095)

#### Declaration

```solidity
function authRedeem(
uint8 u,
address c,
address t,
address a
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type    | Description                                               |
| --- | ------- | --------------------------------------------------------- |
| `u` | uint8   | Underlying token address associated with this market pair |
| `c` | address | Compound token address associated with this market pair   |
| `t` | address | Address of the user receiving the underlying tokens       |
| `a` | address | Amount of underlying being redeemed                       |

### redeemZcToken

Allows zcToken holders to redeem their tokens for underlying tokens after maturity has been reached (via MarketPlace).

#### Declaration

```solidity
function redeemZcToken(
uint8 p,
address u,
uint256 m,
uint256 a
) external returns
(bool)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description                                          |
| --- | ------- | ---------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market pair |
| `u` | address | Underlying token address associated with the market  |
| `m` | uint256 | Maturity timestamp of the market                     |
| `a` | uint256 | Amount of zcTokens being redeemed                    |

### redeemVaultInterest

Allows Vault owners to redeem any currently accrued interest (via MarketPlace)

#### Declaration

```solidity
function redeemVaultInterest(
uint8 p,
address u,
uint256 m
) external returns
(bool)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description                                          |
| --- | ------- | ---------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market pair |
| `u` | address | Underlying token address associated with the market  |
| `m` | uint256 | Maturity timestamp of the market                     |

### redeemSwivelVaultInterest

Allows Swivel to redeem any currently accrued interest (via MarketPlace)

#### Declaration

```solidity
function redeemSwivelVaultInterest(
uint8 p,
address u,
uint256 m
) external returns
(bool)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description                                          |
| --- | ------- | ---------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market pair |
| `u` | address | Underlying token address associated with the market  |
| `m` | uint256 | Maturity timestamp of the market                     |

### validOrderHash

Verifies the validity of an order and it's signature.

#### Declaration

```solidity
function validOrderHash(
struct Hash.Order o,
struct Sig.Components c
) internal returns
(bytes32)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type                  | Description                           |
| --- | --------------------- | ------------------------------------- |
| `o` | struct Hash.Order     | An offline Swivel.Order               |
| `c` | struct Sig.Components | Components of a valid ECDSA signature |

#### Returns:

| Type  | Description   |
| ----- | ------------- |
| `the` | hashed order. |

### deposit

Use the Protocol Enum to direct deposit type transactions to their specific library abstraction

> This functionality is an abstraction used by `IVFZI`, `IZFVI` and `splitUnderlying`

#### Declaration

```solidity
function deposit(
uint8 p,
address u,
address c,
uint256 a
) internal returns
(bool)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description                                       |
| --- | ------- | ------------------------------------------------- |
| `p` | uint8   | Protocol Enum Value                               |
| `u` | address | Address of an underlying token (used by Aave)     |
| `c` | address | Compounding token address                         |
| `a` | uint256 | Amount to deposit todo compounding or underlying? |

### withdraw

Use the Protocol Enum to direct withdraw type transactions to their specific library abstraction

> This functionality is an abstraction used by `EVFZE`, `EZFVE`, `combineTokens`, `redeemZcToken` and `redeemVaultInterest`.
> Note that while there is an external method `withdraw` also on this contract the unique method signatures (and visibility)
> exclude any possible clashing

#### Declaration

```solidity
function withdraw(
uint8 p,
address u,
address c,
uint256 a
) internal returns
(bool)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description                                   |
| --- | ------- | --------------------------------------------- |
| `p` | uint8   | Protocol Enum Value                           |
| `u` | address | Address of an underlying token (used by Aave) |
| `c` | address | Compounding token address                     |
| `a` | uint256 | Amount to withdraw                            |

## Events

### Cancel

Emitted on order cancellation

### Initiate

Emitted on any initiate\*

> filled is 'principalFilled' when (vault:false, exit:false) && (vault:true, exit:true)
> filled is 'premiumFilled' when (vault:true, exit:false) && (vault:false, exit:true)

### Exit

Emitted on any exit\*

> filled is 'principalFilled' when (vault:false, exit:false) && (vault:true, exit:true)
> filled is 'premiumFilled' when (vault:true, exit:false) && (vault:false, exit:true)

### ScheduleWithdrawal

Emitted on token withdrawal scheduling

### ScheduleApproval

Emitted on token approval scheduling

### ScheduleFeeChange

Emitted on fee change scheduling

### BlockWithdrawal

Emitted on token withdrawal blocking

### BlockApproval

Emitted on token approval blocking

### BlockFeeChange

Emitted on fee change blocking

### ChangeFee

Emitted on a change to the fee structure

### SetAdmin

Emitted on a change of the admin
