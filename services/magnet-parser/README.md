# magnet-parser

RabbitMQ listener that receives a `torrent` object which contains a `magnet_link`.

This service converts the magnet link to a valid torrent file, uploads to S3 the torrent file,
and sends another message with updated `torrent` object.

