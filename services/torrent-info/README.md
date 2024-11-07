# torrent-parser

RabbitMQ listener that receives a `torrent` object which contains a `filename` attribute.

This service downloads torrent file from S3, and reads its metadata. The whole `torrent` object is sent to 
another queue, so someone picks it up for next action.
