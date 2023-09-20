// SPDX-License-Identifier: AGPL-3.0

pragma solidity ^0.8.13;

import 'src/lib/Hash.sol';

library TestHash {
    /// @param o A Swivel Order
    function order(Hash.Order memory o) internal pure returns (bytes32) {
        // TODO assembly
        return
            keccak256(
                abi.encode(
                    Hash.ORDER_TYPEHASH,
                    o.key,
                    o.protocol,
                    o.maker,
                    o.underlying,
                    o.vault,
                    o.exit,
                    o.principal,
                    o.premium,
                    o.maturity,
                    o.expiry
                )
            );
    }
}
