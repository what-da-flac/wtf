# magnet-parser

Listens to messages which contains a `magnet_link` using a torrent payload.

This service converts the magnet link to a valid torrent file, uploads to S3 the torrent file,
and sends another message with updated `torrent` object.

