# Healthcare organisations (HCO) command line client
---------------------------------------

  * [Intro](#intro)
  * [Requirements](#requirements)
  * [Usage](#usage)
    * [Options '1-30' : search](#options-1-to-30) 
    * [Option '90' : create database tables](#option-90--create-database-tables) 
    * [Option '91' : drop database tables](#option-91--drop-database-tables) 
    * [Option '92' : import healthcare data](#option-92--import-healthcare-data) 
    * [Option '99' : display database statistics](#option-99--display-database-statistics) 
  * [Data Model](#data-model)
  * [License](#license)

---------------------------------------

## Intro

This is a command line interface (CLI), implemented in Golang, enabling the import and processing of public data of Estonian Health Board. The data is available on their [web page](https://medre.tehik.ee/open-data) and includes the registries of healthcare organisations (HCO-s), general practitioners (GP-s), healthcare professions and healthcare services.

The CLI can create data structures in PostgreSQL and complete the import of the registries. 

## Requirements

* Go 1.18 or higher
* PostgreSQL 14.6 or higher

p.s. If you want to create the tables on Postgre database and import HCO data (menu items from 90 till 99), you need to have DATABASE_URL environment variable assigned, specifying the host, database and DB credentials. If you do not need database import, the Postgre nor the DATABASE_URL variable are not needed.

```
export DATABASE_URL="host=localhost user=hco password=hco dbname=mydb"
```

## Usage

To run the command line client

```
go run hcocli.go  
```

It displays a menu and waits for your choice of action 

```
= = = = = = = = = MAIN MENU = = = = = = = = = = = =

 0 => quit!

 1 => search for HCO (Health Care Organisation)
 2 => search for HCO (activity fields and employees are also displayed)
 3 => search for HCO employee

 10 => search for a General Practitioner (GP)
 20 => search for a General Practitioner (employees are also displayed)
 30 => search for a GP employee

 90 => create database tables
 91 => drop database tables
 92 => import healthcare data into the database

 99 => display database statistics

= = = = = = = = = = = = = = = = = = = = = = = = = =

Your choice => 

```
### Options '1' to '30'

Different dialogs for searching health care organisations, GP-s or employees.

NB! note, that the search is run against XML files, not the Postgre database. The XML files are located on the "/data" directory and can be refreshed from the website of [Estonian Health Board public registries](https://medre.tehik.ee/open-data).

### Option '90' : create database tables

Creates all the needed HCO tables in the database.

### Option '91' : drop database tables

Drops all HCO tables from the database.

### Option '92' : import healthcare data

Fills the tables with HCO data loaded from the registry XML files.

### Option '99' : display database statistics

Displays database statistics (as given below), showing if the HCO tables exist and how many records they hold.

```
* * * * * * * * * * * * * * DATABASE STATISTICS * * * * * * * * * * * * * * 

Table 'hco_healthcare_organisation' exists => true | Row count : 1649
Table 'hco_employee' exists => true | Row count : 968
Table 'hco_employee_profession' exists => true | Row count : 730
Table 'hco_license_residence_service' exists => true | Row count : 4812
Table 'hco_license_residence' exists => true | Row count : 3352
Table 'hco_license' exists => true | Row count : 2638
Table 'hco_service' exists => true | Row count : 247
Table 'hco_profession' exists => true | Row count : 65

* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
```

## Data model

![Data Model](/data/data-model.png "Data model")

## License

[MIT](https://opensource.org/license/mit/)
