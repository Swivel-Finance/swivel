// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.4;

/**
 * @dev These functions deal with verification of Merkle trees (hash trees),
 */
library MerkleProof {
  /// @dev Returns true if a `leaf` can be proved to be a part of a Merkle tree defined by `root`.
  /// For this, a `proof` must be provided, containing sibling hashes on the branch from the leaf to the root of the 
  /// tree. Each pair of leaves and each pair of pre-images are assumed to be sorted.
  /// @param p Array of merkle proofs
  /// @param r Merkle root
  /// @param l Candidate 'leaf' of the merkle tree
  function verify(bytes32[] memory p, bytes32 r, bytes32 l) internal pure returns (bool) {
    uint256 len = p.length;
    bytes32 hash = l;

    for (uint256 i = 0; i < len; i++) {
      bytes32 element = p[i];

      if (hash <= element) {
        // Hash(current computed hash + current element of the proof)
        hash = keccak256(abi.encodePacked(hash, element));
      } else {
        // Hash(current element of the proof + current computed hash)
        hash = keccak256(abi.encodePacked(element, hash));
      }
    }

    // Check if the computed hash (root) is equal to the provided root
    return hash == r;
  }
}
