schema "public" {
  comment = "public schema"
}

table "placeholder" {
    schema = schema.public
    column "id" {
      null = false
      type = int 
      identity {
        generated = ALWAYS
        start = 1
      }
    }
    column "value" {
      null = false
      type = varchar(255)
    }
}
