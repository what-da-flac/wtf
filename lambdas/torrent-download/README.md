# torrent-metadata

Receives a torrent JSON as payload. Then:

1. Download torrent file from S3
2. Starts transmission-daemon in background
2. Download all torrent contents into provided directory
3. Checks torrent status to be Done 100%
4. Sends all downloaded files from torrent to S3
5. Sends SQS message to next listener
