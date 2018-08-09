# Kn
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FKxrr%2Fkn.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FKxrr%2Fkn?ref=badge_shield)


Send a markdown message via udp then get notified on Telegram.

![Screenshot](https://i.imgur.com/Fuhxhvw.gif)

## What's kn

`Kn` is a udp service that can forward markdown message to telegram.


## Example usage

1. Get your telegram bot token and your telegram user id

2. Download the kn binary file from https://github.com/Kxrr/kn/releases

3. Start kn

```
$ ./kn -host=0.0.0.0 -port=9996 -token=<YOUR_TELEGRAM_BOT_TOKEN> -userid=<YOUR_USER_ID>
```

4. Let's send a udp message to the kn server

```
$ echo "hi" | nc -u -w0 127.0.0.1 9996
```

5. Now you should received the hi message on telegram from you telegram bot

## FAQ

Q: A `i/o timeout` error?

A: Make sure you can connect to https://api.telegram.org, the telegram api server is blocked in some area.

---

Q: How to create telegram bot and get the token?

A: https://core.telegram.org/bots#3-how-do-i-create-a-bot

---

Q: How can I get my telegram user id?

A: A easy way is to talk to a telegram bot called `@get_id_bot`.


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FKxrr%2Fkn.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FKxrr%2Fkn?ref=badge_large)
