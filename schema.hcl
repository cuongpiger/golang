table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "name" {
    null = true
    type = text
  }
  column "email" {
    null = true
    type = text
  }
  column "phone_numbers" {
    null = true
    type = jsonb
  }
  primary_key {
    columns = [column.id]
  }
}
schema "public" {
  comment = "standard public schema"
}
