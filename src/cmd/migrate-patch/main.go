package main

import (

)

var dbURL string

const pgSQLAlterStmt string = `ALTER TABLE schema_migrations ADD COLUMN "dirty" boolean NOT NULL DEFAULT false`
const pqSQLCheckColStmt string `SELECT T1.C1, T2.C2 FROM
`