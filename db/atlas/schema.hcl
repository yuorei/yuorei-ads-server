table "ad_groups" {
  schema = schema.ads
  column "ad_group_id" {
    null = false
    type = varchar(255)
  }
  column "campaign_id" {
    null = false
    type = varchar(255)
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null      = false
    type      = timestamp
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "deleted_at" {
    null = true
    type = timestamp
  }
  column "is_approval" {
    null = true
    type = bool
  }
  primary_key {
    columns = [column.ad_group_id]
  }
  foreign_key "ad_groups_ibfk_1" {
    columns     = [column.campaign_id]
    ref_columns = [table.campaigns.column.campaign_id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "campaign_id" {
    columns = [column.campaign_id]
  }
}
table "ads" {
  schema = schema.ads
  column "ad_id" {
    null = false
    type = varchar(255)
  }
  column "ad_group_id" {
    null = false
    type = varchar(255)
  }
  column "type" {
    null = false
    type = varchar(255)
  }
  column "content" {
    null = false
    type = text
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null      = false
    type      = timestamp
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "deleted_at" {
    null = true
    type = timestamp
  }
  column "is_approval" {
    null = true
    type = bool
  }
  primary_key {
    columns = [column.ad_id]
  }
  foreign_key "ads_ibfk_1" {
    columns     = [column.ad_group_id]
    ref_columns = [table.ad_groups.column.ad_group_id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "ad_group_id" {
    columns = [column.ad_group_id]
  }
}
table "campaigns" {
  schema = schema.ads
  column "campaign_id" {
    null = false
    type = varchar(255)
  }
  column "user_id" {
    null = false
    type = varchar(255)
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  column "budget" {
    null = false
    type = int
  }
  column "start_date" {
    null = false
    type = timestamp
  }
  column "end_date" {
    null = false
    type = timestamp
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null      = false
    type      = timestamp
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "deleted_at" {
    null = true
    type = timestamp
  }
  column "is_approval" {
    null = true
    type = bool
  }
  primary_key {
    columns = [column.campaign_id]
  }
  foreign_key "campaigns_ibfk_1" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.user_id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "user_id" {
    columns = [column.user_id]
  }
}
table "impressions" {
  schema = schema.ads
  column "impression_id" {
    null = false
    type = varchar(255)
  }
  column "ad_id" {
    null = false
    type = varchar(255)
  }
  column "date" {
    null = false
    type = date
  }
  column "impressions" {
    null = false
    type = int
  }
  column "clicks" {
    null = false
    type = int
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null      = false
    type      = timestamp
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "deleted_at" {
    null = true
    type = timestamp
  }
  primary_key {
    columns = [column.impression_id]
  }
  foreign_key "impressions_ibfk_1" {
    columns     = [column.ad_id]
    ref_columns = [table.ads.column.ad_id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "ad_id" {
    columns = [column.ad_id]
  }
}
table "organizations" {
  schema = schema.ads
  column "organization_id" {
    null = false
    type = varchar(255)
  }
  column "organization_name" {
    null = false
    type = varchar(255)
  }
  column "representative_name" {
    null = false
    type = varchar(255)
  }
  column "representative_email" {
    null = false
    type = varchar(255)
  }
  column "purpose" {
    null = false
    type = varchar(255)
  }
  column "category" {
    null = false
    type = varchar(255)
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null      = false
    type      = timestamp
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "deleted_at" {
    null = true
    type = timestamp
  }
  primary_key {
    columns = [column.organization_id]
  }
}
table "organizations_users" {
  schema = schema.ads
  column "organization_id" {
    null = false
    type = varchar(255)
  }
  column "user_id" {
    null = false
    type = varchar(255)
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null      = false
    type      = timestamp
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "deleted_at" {
    null = true
    type = timestamp
  }
  primary_key {
    columns = [column.organization_id, column.user_id]
  }
  foreign_key "fk_organization_id" {
    columns     = [column.organization_id]
    ref_columns = [table.organizations.column.organization_id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "fk_user_id" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.user_id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "fk_user_id" {
    columns = [column.user_id]
  }
}
table "roles" {
  schema = schema.ads
  column "role_id" {
    null = false
    type = varchar(255)
  }
  column "name" {
    null = false
    type = varchar(255)
  }
  column "description" {
    null = true
    type = text
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null      = false
    type      = timestamp
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "deleted_at" {
    null = true
    type = timestamp
  }
  primary_key {
    columns = [column.role_id]
  }
  index "name" {
    unique  = true
    columns = [column.name]
  }
}
table "targeting" {
  schema = schema.ads
  column "targeting_id" {
    null = false
    type = varchar(255)
  }
  column "ad_id" {
    null = false
    type = varchar(255)
  }
  column "type" {
    null = false
    type = varchar(255)
  }
  column "value" {
    null = false
    type = varchar(255)
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null      = false
    type      = timestamp
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "deleted_at" {
    null = true
    type = timestamp
  }
  primary_key {
    columns = [column.targeting_id]
  }
  foreign_key "targeting_ibfk_1" {
    columns     = [column.ad_id]
    ref_columns = [table.ads.column.ad_id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "ad_id" {
    columns = [column.ad_id]
  }
}
table "user_roles" {
  schema = schema.ads
  column "user_id" {
    null = false
    type = varchar(255)
  }
  column "role_id" {
    null = false
    type = varchar(255)
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null      = false
    type      = timestamp
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "deleted_at" {
    null = true
    type = timestamp
  }
  primary_key {
    columns = [column.user_id, column.role_id]
  }
  foreign_key "user_roles_ibfk_1" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.user_id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "user_roles_ibfk_2" {
    columns     = [column.role_id]
    ref_columns = [table.roles.column.role_id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "role_id" {
    columns = [column.role_id]
  }
}
table "users" {
  schema = schema.ads
  column "user_id" {
    null = false
    type = varchar(255)
  }
  column "username" {
    null = false
    type = varchar(255)
  }
  column "email" {
    null = false
    type = varchar(255)
  }
  column "hashed_password" {
    null = false
    type = char(60)
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null      = false
    type      = timestamp
    default   = sql("CURRENT_TIMESTAMP")
    on_update = sql("CURRENT_TIMESTAMP")
  }
  column "deleted_at" {
    null = true
    type = timestamp
  }
  primary_key {
    columns = [column.user_id]
  }
  index "email" {
    unique  = true
    columns = [column.email]
  }
  index "username" {
    unique  = true
    columns = [column.username]
  }
}
schema "ads" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
