# MarketPlace

## Contents

<!-- START doctoc -->
<!-- END doctoc -->

## Globals

> Note this list contains both internal and external attributes

| Var     | Type                                                                                |
| ------- | ----------------------------------------------------------------------------------- |
| markets | mapping(uint8 => mapping(address => mapping(uint256 => struct MarketPlace.Market))) |
| paused  | mapping(uint8 => bool)                                                              |
| admin   | address                                                                             |
| swivel  | address                                                                             |
| creator | address                                                                             |

## Modifiers

### authorized

Restricts `msg.sender` as the only viable caller of a method

#### Declaration

```solidity
modifier authorized
```

### unpaused

Returns whether a given market is paused or not

#### Declaration

```solidity
modifier unpaused
```

## Functions

### constructor

#### Declaration

```solidity
function constructor(
address c
) public
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description                              |
| --- | ------- | ---------------------------------------- |
| `c` | address | Address of the deployed creator contract |

### setSwivel

We only allow this to be set once

> there is no emit here as it's only done once post deploy by the deploying admin

#### Declaration

```solidity
function setSwivel(
address s
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type    | Description                             |
| --- | ------- | --------------------------------------- |
| `s` | address | Address of the deployed swivel contract |

### setAdmin

No description

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

### createMarket

Allows the owner to create new markets

> the memory allocation of `s` is for alleviating STD err, there's no clearly superior scoping or abstracting alternative.

#### Declaration

```solidity
function createMarket(
uint8 p,
uint256 m,
address c,
string n,
string s
) external authorized unpaused returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |
| unpaused   |

#### Args:

| Arg | Type    | Description                                              |
| --- | ------- | -------------------------------------------------------- |
| `p` | uint8   | Protocol associated with the new market                  |
| `m` | uint256 | Maturity timestamp of the new market                     |
| `c` | address | Compounding Token address associated with the new market |
| `n` | string  | Name of the new market zcToken                           |
| `s` | string  | Symbol of the new market zcToken                         |

### matureMarket

Can be called after maturity, allowing all of the zcTokens to earn floating interest on Compound until they release their funds

#### Declaration

```solidity
function matureMarket(
uint8 p,
address u,
uint256 m
) public unpaused returns
(bool)
```

#### Modifiers:

| Modifier |
| -------- |
| unpaused |

#### Args:

| Arg | Type    | Description                                                  |
| --- | ------- | ------------------------------------------------------------ |
| `p` | uint8   | Protocol Enum value associated with the market being matured |
| `u` | address | Underlying token address associated with the market          |
| `m` | uint256 | Maturity timestamp of the market                             |

### mintZcTokenAddingNotional

Allows Swivel caller to deposit their underlying, in the process splitting it - minting both zcTokens and vault notional.

#### Declaration

```solidity
function mintZcTokenAddingNotional(
uint8 p,
address u,
uint256 m,
address t,
uint256 a
) external authorized unpaused returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |
| unpaused   |

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |
| `t` | address | Address of the depositing user                      |
| `a` | uint256 | Amount of notional being added                      |

### burnZcTokenRemovingNotional

Allows Swivel caller to deposit/burn both zcTokens + vault notional. This process is "combining" the two and redeeming underlying.

#### Declaration

```solidity
function burnZcTokenRemovingNotional(
uint8 p,
address u,
uint256 m,
address t,
uint256 a
) external authorized unpaused returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |
| unpaused   |

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |
| `t` | address | Address of the combining/redeeming user             |
| `a` | uint256 | Amount of zcTokens being burned                     |

### authRedeem

Implementation of authRedeem to fulfill the IRedeemer interface for ERC5095

#### Declaration

```solidity
function authRedeem(
uint8 p,
address u,
uint256 m,
address f,
address t,
uint256 a
) external authorized unpaused returns
(uint256)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |
| unpaused   |

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |
| `f` | address | Address of the user having their zcTokens burned    |
| `t` | address | Address of the user receiving underlying            |
| `a` | uint256 | Amount of zcTokens being redeemed                   |

#### Returns:

| Type     | Description                                            |
| -------- | ------------------------------------------------------ |
| `Amount` | of underlying being withdrawn (needed for 5095 return) |

### redeemZcToken

Allows (via swivel) zcToken holders to redeem their tokens for underlying tokens after maturity has been reached.

#### Declaration

```solidity
function redeemZcToken(
uint8 p,
address u,
uint256 m,
address t,
uint256 a
) external authorized unpaused returns
(uint256)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |
| unpaused   |

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |
| `t` | address | Address of the redeeming user                       |
| `a` | uint256 | Amount of zcTokens being redeemed                   |

### redeemVaultInterest

Allows Vault owners (via Swivel) to redeem any currently accrued interest

#### Declaration

```solidity
function redeemVaultInterest(
uint8 p,
address u,
uint256 m,
address t
) external authorized unpaused returns
(uint256)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |
| unpaused   |

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |
| `t` | address | Address of the redeeming user                       |

### calculateReturn

Calculates the total amount of underlying returned including interest generated since the `matureMarket` function has been called

#### Declaration

```solidity
function calculateReturn(
uint8 p,
address u,
uint256 m,
uint256 a
) internal returns
(uint256)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |
| `a` | uint256 | Amount of zcTokens being redeemed                   |

### cTokenAddress

Return the compounding token address for a given market

#### Declaration

```solidity
function cTokenAddress(
uint8 p,
address u,
uint256 m
) external returns
(address)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |

### exchangeRate

Return the exchange rate for a given protocol's compounding token

#### Declaration

```solidity
function exchangeRate(
uint8 p,
address c
) external returns
(uint256)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description               |
| --- | ------- | ------------------------- |
| `p` | uint8   | Protocol Enum value       |
| `c` | address | Compounding token address |

### rates

Return current rates (maturity, exchange) for a given vault. See VaultTracker.rates for details

> While it's true that Compounding exchange rate is not strictly affiliated with a vault, the 2 data points are usually needed together.

#### Declaration

```solidity
function rates(
uint8 p,
address u,
uint256 m
) external returns
(uint256, uint256)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |

### custodialInitiate

Called by swivel IVFZI && IZFVI

> Call with protocol, underlying, maturity, mint-target, add-notional-target and an amount

#### Declaration

```solidity
function custodialInitiate(
uint8 p,
address u,
uint256 m,
address z,
address n,
uint256 a
) external authorized unpaused returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |
| unpaused   |

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |
| `z` | address | Recipient of the minted zcToken                     |
| `n` | address | Recipient of the added notional                     |
| `a` | uint256 | Amount of zcToken minted and notional added         |

### custodialExit

Called by swivel EVFZE FF EZFVE

> Call with protocol, underlying, maturity, burn-target, remove-notional-target and an amount

#### Declaration

```solidity
function custodialExit(
uint8 p,
address u,
uint256 m,
address z,
address n,
uint256 a
) external authorized unpaused returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |
| unpaused   |

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |
| `z` | address | Owner of the zcToken to be burned                   |
| `n` | address | Target to remove notional from                      |
| `a` | uint256 | Amount of zcToken burned and notional removed       |

### p2pZcTokenExchange

Called by swivel IZFZE, EZFZI

> Call with underlying, maturity, transfer-from, transfer-to, amount

#### Declaration

```solidity
function p2pZcTokenExchange(
uint8 p,
address u,
uint256 m,
address f,
address t,
uint256 a
) external authorized unpaused returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |
| unpaused   |

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |
| `f` | address | Owner of the zcToken to be burned                   |
| `t` | address | Target to be minted to                              |
| `a` | uint256 | Amount of zcToken transfer                          |

### p2pVaultExchange

Called by swivel IVFVE, EVFVI

> Call with protocol, underlying, maturity, remove-from, add-to, amount

#### Declaration

```solidity
function p2pVaultExchange(
uint8 p,
address u,
uint256 m,
address f,
address t,
uint256 a
) external authorized unpaused returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |
| unpaused   |

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |
| `f` | address | Owner of the notional to be transferred             |
| `t` | address | Target to be transferred to                         |
| `a` | uint256 | Amount of notional transfer                         |

### transferVaultNotional

External method giving access to this functionality within a given vault

> Note that this method calculates yield and interest as well

#### Declaration

```solidity
function transferVaultNotional(
uint8 p,
address u,
uint256 m,
address t,
uint256 a
) external unpaused returns
(bool)
```

#### Modifiers:

| Modifier |
| -------- |
| unpaused |

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |
| `t` | address | Target to be transferred to                         |
| `a` | uint256 | Amount of notional to be transferred                |

### transferVaultNotionalFee

Transfers notional fee to the Swivel contract without recalculating marginal interest for from

#### Declaration

```solidity
function transferVaultNotionalFee(
uint8 p,
address u,
uint256 m,
address f,
uint256 a
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type    | Description                                         |
| --- | ------- | --------------------------------------------------- |
| `p` | uint8   | Protocol Enum value associated with this market     |
| `u` | address | Underlying token address associated with the market |
| `m` | uint256 | Maturity timestamp of the market                    |
| `f` | address | Owner of the amount                                 |
| `a` | uint256 | Amount to transfer                                  |

### pause

Called by admin at any point to pause / unpause market transactions in a specified protocol

#### Declaration

```solidity
function pause(
uint8 p,
bool b
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type  | Description                                                  |
| --- | ----- | ------------------------------------------------------------ |
| `p` | uint8 | Protocol Enum value of the protocol to be paused             |
| `b` | bool  | Boolean which indicates the (protocol) markets paused status |

## Events

### Create

Emitted upon the creation of a market

### Mature

Emitted upon the maturation of a market

### RedeemZcToken

Emitted upon the redemption of a ZC token

### RedeemVaultInterest

Emitted upon the redemption of vault interest

### CustodialInitiate

Emitted upon a custodial initiate

### CustodialExit

Emitted upon a custodial exit

### P2pZcTokenExchange

Emitted upon a P2P ZC token swap

### P2pVaultExchange

Emitted upon a vault swap

### TransferVaultNotional

Emitted upon a transfer of notional from a vault

### SetAdmin

Emitted on a change of the admin
