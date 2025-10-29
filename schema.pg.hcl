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

// Notes
//
//   Tables:
//     authusers { id, username, passwdhash, realname, prefs(json), enabled }
//     authroles { id, rolename, uid }
//
//     accounts { id, account_num, name, institution, ofx_creds }
//     account_ballances { account_id, date, ballance }
//     account_users { account_id, user_id }
//     transactions { account_id, ??? }                <-- *** FOCUS ***
//     categories { id, name }
//     autocategories { category_id, account_id, desc_pattern, is_re }
