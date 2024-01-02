# <img src="https://images.projectceleste.com/Art/UserInterface/Icons/Units/AvatarVillagers_ua.png" width="40" height="40" alt=":villager:" class="emoji" title=":villager:"/> villager-db

> Villager is a civilian unit in Age of Empires that can be trained at the Town Center. Villagers are the backbone of every civilization. Their purpose is to construct buildings and gather resources.  
> 
> We use villager to gather our database migration and construct them to our main database.


This is an improved Kitabisa migration service as a replacement for our old flyway service [https://github.com/kitabisa/flyway-db-migration](https://github.com/kitabisa/flyway-db-migration).

*__Note__ : our old flyway service is now deprecated and any new migration should be added here.*

## Table of contents
* [Why we need to migrate from flyway-db to villager-db](#why-we-need-to-migrate-from-flyway-db-to-villager-db)
* [How to contribute](#how-to-contribute)
* [Migration rules](#migration-rules)
* [Requirement](#requirement)
* [How to configure migration on your local machine](#how-to-configure-migration-on-your-local-machine)
* [How to run migration](#how-to-run-migration)
* [How to add dummy data to database](#how-to-add-dummy-data-to-database)
* [Troubleshooting](#troubleshooting)
* [Need more help?](#need-more-help)

## Why we need to migrate from flyway-db to villager-db

1. We are using k8s now. So we want to unify all deployment process. Including db migration. So, we don’t need Jenkins to do migration anymore.
2. Our flyway-db cannot do migration down (rollback). Pay money for flyway enterprise to use this.
3. Flyway use checksum restriction for migration history. This is good. But we noticed problem when we wanted to deploy DB to Google Cloud SQL. Because we wanted to change current table engine to InnoDB.
4. We can use unix timestamp for migration versioning instead of integer number. So we can prevent migration version conflict from other developers.

## How to contribute

Read [CONTRIBUTING.md](https://github.com/kitabisa/villager-db/blob/master/CONTRIBUTING.md) to know more about how to contribute to this repo and how to deploy migration to our database. If you are new to this repo, it is mandatory to read this file first.

## Migration rules

- Use InnoDB as a storage engine. It's a must because Google Cloud SQL currently `(on Nov 2019)` only support InnoDB.
- Reserve your migration version here before creating migration file. This must be done to prevent conflict version with other developers.
  ```
  https://docs.google.com/spreadsheets/d/1m_vyJB-OjfQUlHHvKxtVw99JIqv5SZHHaBpmgiM3I6k/edit#gid=0
  ```

## Requirement

1. Go 1.13 or above (download [here](https://golang.org/dl/)).
2. Git (download [here](https://git-scm.com/downloads)).
3. MySQL version 5.7.23

## How to configure migration on your local machine

1. Clone this repostory to your local.
   ```bash
   $ git clone git@github.com:kitabisa/villager-db.git
   ```

2. Change working directory to `villager-db` folder.
   ```bash
   $ cd villager-db
   ```

3. Create configuration files.
   ```bash
   $ cp params/.env.example params/.env
   ```

4. Edit configuration values in `params/.env` according to your database setting.

## How to run migration

This migration can do these actions:

1. Migration up  

   This command will migrate the database to the most recent version available. Migration files can be seen in this folder `migrations/sql/`.
   ```bash
   $ go run main.go up
   ```

2. Migration down  
   
   This command will undo/rollback database migration up to one step backwards.
   ```bash
   $ go run main.go down
   ```

    If you wish to rollback migration more than one step (up to N step backward), add flag option `-n` to specify how many step will do undo. Not specifying this flag means only rollback one step migration. The example below will rollback migration 7 step backwards.
    ```bash
    $ go run main.go down -n 7
    ```

3. Migration new  
   
   This command will generate new migration file based on unix timestamp format
   ```bash
   $ go run main.go new create_table_A
   2019/11/28 13:08:01 New migration file has been created: migrations/sql/1574921281_create_table_A.sql
   ```

4. Migration status  
   
   Use this command to print out migration status. Which migration had been migrated and which was not will be printed by this command.
   ```bash
   $ go run main.go status
   ```

To get any help about these command, add `--help` at the end of the command.
```bash
$ go run main.go --help
$ go run main.go up --help
$ go run main.go down --help
$ go run main.go new --help
$ go run main.go status --help
```

## How to add dummy data to database

Just import any `*.sql` seeds data from this folder `migration/seeds/`.

## Troubleshooting

- [Move current database schema history created by flyway-db to villager-db](https://github.com/kitabisa/villager-db/wiki/move-current-database-schema-history-created-by-flyway-db-to-villager-db)  

## Need more help

If you need more help or anything else, please ask Kitabisa backend engineer team. We would be happy to help you.
