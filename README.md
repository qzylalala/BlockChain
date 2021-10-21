### Transaction

![Transactions](https://jeiwan.net/images/transactions-diagram.png)

1. **There are outputs that are not linked to inputs.**

2. **In one transaction, inputs can reference outputs from multiple transactions.**

3. **An input must reference an output.**

4. **Output comes first**



### Merkle Tree

![Merkle tree diagram](https://jeiwan.net/images/merkle-tree-diagram.png)

1. A Merkle tree is built for each block, and it starts with leaves (the bottom of the tree), where a leaf is a transaction hash (Bitcoins uses double `SHA256` hashing). The number of leaves must be even, but not every block contains an even number of transactions. In case there is an odd number of transactions, the last transaction is duplicated (in the Merkle tree, not in the block!).

2. Moving from the bottom up, leaves are grouped in pairs, their hashes are concatenated, and a new hash is obtained from the concatenated hashes. The new hashes form new tree nodes. This process is repeated until there’s just one node, which is called the root of the tree. The root hash is then used as the unique representation of the transactions, is saved in block headers, and is used in the proof-of-work system.

3. The benefit of Merkle trees is that a node can verify membership of certain transaction without downloading the whole block. Just a transaction hash, a Merkle tree root hash, and a Merkle path are required for this.
