# Creator

## Contents

<!-- START doctoc -->
<!-- END doctoc -->

## Globals

> Note this list contains both internal and external attributes

| Var         | Type    |
| ----------- | ------- |
| admin       | address |
| marketPlace | address |

## Modifiers

### authorized

Restricts `msg.sender` as the only viable caller of a method

#### Declaration

```solidity
modifier authorized
```

## Functions

### create

Allows the owner to create new markets

#### Declaration

```solidity
function create(
uint8 p,
address u,
uint256 m,
address c,
address sw,
string n,
string s,
uint8 d
) external authorized returns
(address, address)
```

#### Modifiers:

| Modifier   |
| ---------- |
| authorized |

#### Args:

| Arg  | Type    | Description                                              |
| ---- | ------- | -------------------------------------------------------- |
| `p`  | uint8   | Protocol associated with the new market                  |
| `u`  | address | Underlying token associated with the new market          |
| `m`  | uint256 | Maturity timestamp of the new market                     |
| `c`  | address | Compounding Token address associated with the new market |
| `sw` | address | Address of the deployed swivel contract                  |
| `n`  | string  | Name of the new market zcToken                           |
| `s`  | string  | Symbol of the new market zcToken                         |
| `d`  | uint8   | Decimals of the new market zcToken                       |

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

### setMarketPlace

We only allow this to be set once

> there is no emit here as it's only done once post deploy by the deploying admin

#### Declaration

```solidity
function setMarketPlace(
address m
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
| `m` | address | Address of the deployed marketPlace contract |

## Events

### SetAdmin

Emitted on a change of the admin
