# Dicebot

Dicebot is a simple chatbot. It has a kind of framework for adding new
verbs, and supports stateful verbs (via GCP Cloud Datastore, or local
files.)

There is a simple CLI, for easy tweaking and development:
```
go run ./cmd/db/*go
```

There is also a Google Appengine app that implements a LINE webhook:
```
gcloud app deploy ./app/dicebot --project=${YOUR_GCP_PROJECT}
```

To get the bot running inside your LINE groups, you'll need to:
* setup an official "LINE Business ID" account (instructions
[here](https://respond.io/blog/the-ultimate-guide-to-line-for-business/#8crh6)
* get the channel secret for that account
* get a long-lived channel token for that account
* `mv ./pkg/config/myconfig.sample ./pkg/config/myconfig.go`
* edit the new file, putting in your secret & token
* deploy the app
* update your LINE biz acount config, point it to your new webhook (`https://host/line` - the `host` is your new appengine service URL)

The main verb is `roll`, for rolling dice. There are some other verbs
that help out with running an RPG in a chat group:
```
db help

# Setup things
db bot claim lleldron      # First: claim a character name
db char                    # review your character 'sheet'
db inv                     # review your inventory
db inv stash some rope     # populate your inventory

# Things while playing
db roll 5d6                # roll some dice !
db save vs str             # test your strength !
db inv use 2               # assert you will use item #2
db vow to eat less         # make a vow to do something
db spell memorise 1:1 cure # memorize a spell
db spell cast 1:1          # cast a spell !
db history                 # review a log of important events
```

## Misc notes

On LINE, it needs to maintain a mapping between the LINE user IDs and
the names that identify characters. The easiest way is to get users to
claim their name once they're logged into LINE: `db bot claim NAME`.
Note that until the user has agreed to the bot's Terms of Use, the
user IDs won't be visible to the bot; the user needs to add the bot to
their friend list and agree to the ToU.

If you are an admin user (see `./config/myconfig.go`), you will need
to claim your name too :) Once you have it, you can then masquerade as
other users, by prefixing all your bot commands with `as USER`, e.g.
`db as foobar save vs INT`.
