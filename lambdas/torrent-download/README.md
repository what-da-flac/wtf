# torrent-download

This lambda listens to SQS queue. Message is a torrent object as JSON.

Lambda should download all torrent files into `/tmp/downloads` directory.

Once download is completed, sends all files into S3 bucket. Within S3 bucket, it will use a naming convention, so we keep downloaded files identified and grouped together (as they belong to the same torrent anyway).

For instance
