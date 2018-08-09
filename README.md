# Kn
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FKxrr%2Fkn.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FKxrr%2Fkn?ref=badge_shield)


Send a markdown message via udp then get notified on Telegram.

`Kn` is a udp service that can forward markdown message to telegram.


## Usage

1. Download the kn binary file from https://github.com/Kxrr/kn/releases

2. Launch kn

```
$ ./kn -host=0.0.0.0 -port=9996 -token=6666666:AAG-GGGGGGGGGGGGG -userid=3000000
```

2. Now, let's send a udp message to the server

```
$ echo "hi" | nc -u -w0 127.0.0.1 9996
```


## FAQ

Q: A `i/o timeout` error?
A: Make sure you can connect to `https://api.telegram.org/`, the api server is blocked in some area.

Q: How to create telegram bot?
A: https://core.telegram.org/bots#3-how-do-i-create-a-bot

Q: How can I get my telegram user id?
A: A easy way is to talk to a telegram bot called `@get_id_bot`.


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FKxrr%2Fkn.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FKxrr%2Fkn?ref=badge_large)