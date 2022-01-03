# SQL Clause

- Category: `web`
- Challenge author: **Pontus Norrström**

## Description

He's making a database\
He's sorting it twice\
SELECT * FROM contacts WHERE behavior = 'nice'\
SQL Clause is coming to town

### Connection info

<http://2021.santahack.xyz:42204>

## Writeup

This is a classic SQL-injection challenge.

To solve it but **admin** in the username field, and int the password field you put:
```sql
' OR 1=1 --
```
