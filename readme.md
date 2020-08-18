# Boosty Test Task

You need to implement a backend application that would allow users to transfer funds between wallets. On every transaction, the system will need to take a commission of 1.5%.

## Requirements

- Create a relational database schema
- Upon application start, the database should be populated with sample data (e.g several wallets that can be used to transfer funds between)
- The system should support two currency: BTC and ETH
- Create a REST endpoint that can be used to transfer funds
- Provide a docker-compose.yml that can be used to run your solution by just doing docker-compose up

## Get started

So, to start project you need to clone it and just run:

```bash
docker-compose up
```

But you should also do one more thing - migrate database. To do that, open one more tab in terminal and run:

```bash
docker-compose exec api bash
```

After that you need to run (this will be inside container) `./migrate`. That's it!

Now, you can go to `http://localhost:8080` and see a json array of generated wallets.

To make a transfer from one wallet to another use next endpoint:

```bash
Method: POST
Endpoint: http://localhost:8080/transaction
```

Request body:

```json
{
    "from": "7b8053f1-f7d9-49ad-9a3f-32c71ffb4aa2",
    "to": "3405b3e4-bc7c-4fe4-a60e-fa43bb7699f7",
    "amount": 0.08
}
```

If transaction was successfully finished you will get transaction info in response:

```json
{
    "success": true,
    "transaction": {
        "id": "5981e606-6557-4573-b168-851cb848eab4",
        "sender": "7b8053f1-f7d9-49ad-9a3f-32c71ffb4aa2",
        "receiver": "3405b3e4-bc7c-4fe4-a60e-fa43bb7699f7",
        "currency": "eth",
        "amount": 0.08,
        "commission": 0.0012,
        "created_at": "2020-08-18T12:10:35.743837Z"
    }
}
```

