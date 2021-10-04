GoBudget
========
Overview
--------
A envelope-based budgeting web app made using Go and (probably) React. My first app in Go and
first SPA, mainly as a learning exercise.

Features should mimic .. _GoodBudget: https://goodbudget.com/ 's main features and
be usable on desktop and mobile on the web.

Goals
-----
- Envelope budgeting system, with at least a set amount of monthly, yearly and
  savings envelopes
- Multiple accounts
- Basic spending/income statistics, per month, per account and per envelope
- Transaction history
- Transaction undoing
- Scheduled recurring transactions
- User system with email/password credentials

- Single-repo fullstack project
- Stateless auth (JWT)
- REST API
- Usable web interface


Nice to haves
-------------
- More envelope intervals
- Multiple currencies and conversions
- Payee/Payer statistics
- Password reset & email verification
- Clean mobile interface


Non-goals
---------
- Debt handling
- Any advanced financing
- Any overly fancy graphs
- Ultratight security / addressing enterprise-level threat model (e.g. non-trivial DDoSes)
