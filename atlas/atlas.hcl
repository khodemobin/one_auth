table "users" {
  schema = schema.auth
  column "id" {
    null     = false
    type     = int
    unsigned = true
  }
  column "phone" {
    null = false
    type = varchar(20)
  }
  column "password" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null     = false
    type     = tinyint
    default  = 2
    unsigned = true
  }
  column "created_at" {
    null = false
    type = datetime
  }
  column "updated_at" {
    null = false
    type = datetime
  }
  primary_key {
    columns = [table.users.column.id]
  }
  index "users_password_index" {
    unique  = false
    columns = [table.users.column.password]
  }
  index "users_phone_unique" {
    unique  = true
    columns = [table.users.column.phone]
  }
}
schema "auth" {
  charset   = "utf8mb4"
  collation = "utf8mb4_general_ci"
}
