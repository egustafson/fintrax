// atlas.hcl -- project configuration

variable "dbuser" {
    type = string
    // default = getenv("ATLAS_DBUSER")
    default = "postgres"
}

variable "dbpass" {
    type = string
    // default = getenv("ATLAS_DBPASS")
    default = "dev"
}

env "local" {
    src = "file://schema.pg.hcl"
    url = "postgres://${var.dbuser}:${var.dbpass}@localhost:5432/postgres?sslmode=disable"
    dev = "docker://postgres/17/dev"
}

variable "dbhost" {
    type = string
    // default = getenv("ATLAS_DBHOST")
    default = "dbdev"
}

env "integration" {
    src = "file://schema.pg.hcl"
    url = "postgres://${var.dbuser}:${var.dbpass}@${var.dbhost}:5432/postgres?sslmode=disable"
    dev = "docker:/postgres/17/dev"
}

env "prod" {
    src = "file://schema.pg.hcl"
    url = "postgres://${var.dbuser}:${var.dbpass}@${var.dbhost}:5432/postgres?sslmode=disable"
    dev = "docker:/postgres/17/dev"
}

