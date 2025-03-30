# services

This directory contains all backend services. File structure supports services on any programming language. A Dockerfile image definition should be present for each service.

Since all services are written in Golang, no specific considerations are for the time being (like unit tests, mocks, etc). If additional programming languages are used in the future, we should adapt Makefile/Codebuild accordingly.

## Execution

Consider UI application is already running. First need to buid services:

```bash
make build-all
```

Then start services

```bash
make local-start
```

To stop services, run this command

```bash
make local-stop
```

## Diagram

TODO

## magnet-parser

Receives a magnet link within a torrent structure. Service converts magnet to torrent file, and file is uploaded to AWS S3.

Sends a message to `torrent-parser` queue.

## torrent-download

Receives a torrent object, then checks its information exists. If it does, checks if files have not already been uploaded to s3. Then it downloads the files from the torrent, sends them to s3 and sends an updated message to `template-info` to update torrent status.

## torrent-info

Receives a torrent object and tries to create/update its information in db. No further messages are sent.

## torrent-parser

Receives a torrent object with a reference to s3 object. Downloads the actual torrent file, extracts its information, and sends it to `torrent-info` queue, so it can be updated in db.


## Volumes

Data is stored in volumes, which end up being docker volumes mapped to a local path.

`MEDIA_PATH` is the name of the environment variable. Default value can be overridden using `.env.credentials` file.
