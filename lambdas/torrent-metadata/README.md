# torrent-metadata

This lambda reads incoming torrent from payload, and then:

1. Uses `aria2` to convert a magnet link to a torrent file.
2. Uses `transmission-cli` to extract metadata and files info from torrent.
3. Updates status of torrent to `parsed` and sends to SQS queue for next process to pick it up.
