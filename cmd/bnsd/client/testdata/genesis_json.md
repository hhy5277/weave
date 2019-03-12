### Creating the governance multisigs

In a multisig contract you can define an **activation threshold** which is the number of valid signatures of members required to approve a transaction.
It is the minimum number. If more valid signatures are provided they are ignored.
The **admin threshold** though defines the number of member signatures required to make a change to the contract, like adding/removing members.

To give you an example for a contract with 3 members: Alice, Bert and Charlie. Two of them should be required to approve any transaction but
only when all three together sign a modification to the contract it should be updated:

```json
 "multisig": [
    {
      "activation_threshold": 2,
      "admin_threshold": 3,
      "sigs": [ <-- contains the members addressIDs
        "AAAA..." <-- Alice
        "BBBB..." <-- Bert
        "CCCC..." <-- Charlie
      ]
    }
  ]
```

### Creating the validator distribution

Alice and Bert 1/4 each, Charlie 1/2

```json
  "distribution": [
    {
      "admin": "ZZZZ...",
      "recipients": [
        {
          "address": "AAAA...",
          "weight": 1
        },
        {
          "address": "BBBB...",
          "weight": 1
        },
        {
          "address": "CCCC...",
          "weight": 2
        }
      ]
    }
  ],

```

### Creating the block reward contract

```json
  "escrow": [
    {
      "amount": [
        {
          "ticker": "IOV",
          "whole": 1000000  <-- total amount to distribute
        }
      ],
      "arbiter": "multisig/usage/0000000000000001", <-- multisig contract to relase or burn tokens
      "recipient": "cond:TODO!!!!", <-- a distribution contract
      "sender": "0000000000000000000000000000000000000000", <-- non existing burn address
      "timeout": 9223372036854775807 <-- very very high block height, to never expire
    }
  ],

```