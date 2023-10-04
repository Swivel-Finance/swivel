# VaultTracker

## Contents

<!-- START doctoc -->
<!-- END doctoc -->

## Globals

> Note this list contains both internal and external attributes

| Var          | Type                                          |
| ------------ | --------------------------------------------- |
| vaults       | mapping(address => struct VaultTracker.Vault) |
| cTokenAddr   | address                                       |
| marketPlace  | address                                       |
| swivel       | address                                       |
| maturity     | uint256                                       |
| maturityRate | uint256                                       |
| protocol     | uint8                                         |

## Modifiers

### authorized

Restricts msg.sender as the only viable caller of a method

#### Declaration

```solidity
modifier authorized
```

## Functions

### constructor

#### Declaration

```solidity
function constructor(
uint8 m,
uint256 c,
address s,
address mp
) public
```

#### Modifiers:

No modifiers

#### Args:

| Arg  | Type    | Description                                                                                     |
| ---- | ------- | ----------------------------------------------------------------------------------------------- |
| `m`  | uint8   | Maturity timestamp associated with this vault                                                   |
| `c`  | uint256 | Compounding Token address associated with this vault                                            |
| `s`  | address | Address of the deployed swivel contract                                                         |
| `mp` | address | Address of the designated admin, which is the Marketplace addess stored by the Creator contract |

### addNotional

Adds notional to a given address

#### Declaration

```solidity
function addNotional(
address o,
uint256 a
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type    | Description               |
| --- | ------- | ------------------------- |
| `o` | address | Address that owns a vault |
| `a` | uint256 | Amount of notional added  |

### removeNotional

Removes notional from a given address

#### Declaration

```solidity
function removeNotional(
address o,
uint256 a
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
| `o` | address | Address that owns a vault    |
| `a` | uint256 | Amount of notional to remove |

### redeemInterest

Redeem's interest accrued by a given address

#### Declaration

```solidity
function redeemInterest(
address o
) external authorized returns
(uint256)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type    | Description               |
| --- | ------- | ------------------------- |
| `o` | address | Address that owns a vault |

### matureVault

Matures the vault

#### Declaration

```solidity
function matureVault(
uint256 c
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type    | Description                      |
| --- | ------- | -------------------------------- |
| `c` | uint256 | The current cToken exchange rate |

### transferNotionalFrom

Transfers notional from one address to another

#### Declaration

```solidity
function transferNotionalFrom(
address f,
address t,
uint256 a
) external authorized returns
(bool)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg | Type    | Description             |
| --- | ------- | ----------------------- |
| `f` | address | Owner of the amount     |
| `t` | address | Recipient of the amount |
| `a` | uint256 | Amount to transfer      |

### transferNotionalFee

Transfers, in notional, a fee payment to the Swivel contract without recalculating marginal interest for the owner

#### Declaration

```solidity
function transferNotionalFee(
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

| Arg | Type    | Description         |
| --- | ------- | ------------------- |
| `f` | address | Owner of the amount |
| `a` | uint256 | Amount to transfer  |

### rates

Return both the current maturityRate if it's > 0 (or exchangeRate in its place) and the Compounding exchange rate

> While it may seem unnecessarily redundant to return the exchangeRate twice, it prevents many kludges that would otherwise be necessary to guard it

#### Declaration

```solidity
function rates(
) public returns
(uint256, uint256)
```

#### Modifiers:

No modifiers

#### Returns:

| Type           | Description                                             |
| -------------- | ------------------------------------------------------- |
| `exchangeRate` | if maturityRate > 0, exchangeRate, exchangeRate if not. |

### balancesOf

Returns both relevant balances for a given user's vault

#### Declaration

```solidity
function balancesOf(
address o
) external returns
(uint256, uint256)
```

#### Modifiers:

No modifiers

#### Args:

| Arg | Type    | Description               |
| --- | ------- | ------------------------- |
| `o` | address | Address that owns a vault |
