# go health 
### This contains 2 services
1. CRUD for Hospitals
2. CRUD for Medical Procedures

## To run
- Change .envsample to .env
- Add DB, DIALECT in .env
- Run `go build`
- Run `./go-health`

## DB Tables/Schema
Tested on sqlite3
- `CREATE TABLE medprocs ( id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, description TEXT, cost REAL, created_at DATETIME DEFAULT CURRENT_DATE );`
- `CREATE TABLE hospitals ( id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, address TEXT, city TEXT, country TEXT, pin TEXT, created_at DATETIME DEFAULT CURRENT_DATE );`
