# DEPOSINATOR

## Webapp/Server Authorization

Bearer token loaded from env at startup

## Endpoints

- POST /signup

> {  
> "username": "string",  
> "email": "string",  
> "password": "string"  
> }

- POST /login

> {  
> "username": "string",  
> "password": "string"  
> }

- POST /deposit

> {  
> "initiator",: "string",  
> "members": []string,  
> "amount": int,  
> "description": "string"  
> }

- POST /withdraw

> {  
> "deposit": int,  
> "amount": int,  
> "description": "string"  
> }

### Implementation ideas

Deposit:

- members is a dropdown selection of usernames in accounts db

Withdraw:

- deposit correlates to a deposit id
- deposit is a dropdown with deposit info (date, amount) to help correlate the withdraw
