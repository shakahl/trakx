# Trakx

Bittorrent tracker written in go.

## How

It uses the go default webserver and MySQL to hold the client list.
It currently uses my own bencode package but I will most likely move to something else eventually.

## Resources

* [Basic spec](https://wiki.theory.org/index.php/BitTorrentSpecification) - super helpful.

## Todo

* Try using https://github.com/go-torrent/bencode
* Docker for easy testing
* Support Ipv6
  * http://www.bittorrent.org/beps/bep_0007.html

## Done

* Support and test peers that join the tracker when they're already complete.
  * Wireshark it with debian torrent
* Comply with compact peer list
* LastSeen timestamp to remove peers with network issues
  * `go tracker.Clean()` should run every minuit and remove peers who haven't been seen in 1 hour
* Auto delete empty tables
* Logging
  * Using zap

## Database layout

It uses the database `bittorrent` and creates a table with the info hash converted to capital hex.

A torrent table looks like this:

```en
+----------+----------------------+------+-----+---------+-------+
| Field    | Type                 | Null | Key | Default | Extra |
+----------+----------------------+------+-----+---------+-------+
| id       | varchar(40)          | YES  |     | NULL    |       |
| peerKey  | varchar(20)          | YES  |     | NULL    |       |
| ip       | varchar(255)         | YES  |     | NULL    |       |
| port     | smallint(5) unsigned | YES  |     | NULL    |       |
| complete | tinyint(1)           | YES  |     | NULL    |       |
| lastSeen | bigint(20) unsigned  | YES  |     | NULL    |       |
+----------+----------------------+------+-----+---------+-------+
```