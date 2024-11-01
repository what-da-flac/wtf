import os
import subprocess
import time
from datetime import datetime, timedelta

from aws_lambda_powertools.utilities.typing import LambdaContext


def lambda_handler(event: dict, context: LambdaContext):
    wait_seconds = 900
    torrent_file = event.get("file")
    # name matches folder where files are downloaded
    name = event.get("name")
    target_dir = "/Users/mau/Downloads/xxx"
    download_torrent(target_dir, torrent_file)
    # TODO: filenames should come in torrent payload, this is just for local development
    filenames = [
        os.path.join(target_dir, name, "David Guetta Ft. SIA - Titanium Anky.mp3"),
    ]
    if wait_for_download_completion(wait_seconds=wait_seconds, interval_check_seconds=10):
        print("download completed successfully.")
        upload_files_to_s3(filenames)
        # print("completed uploading files to s3")
    else:
        print(f"download did not complete within the {wait_seconds} seconds limit.")


def upload_files_to_s3(filenames: list[str]):
    for filename in filenames:
        # TODO: read file info, sends error now
        print(f"filename to be uploaded: {filename}")
    #     file_info = os.stat(filename)
    #     print(f"file name: {filename}  size: {file_info.st_size} bytes")


def download_torrent(target_dir: str, torrent_file: str):
    """
    Downloads a torrent file using Transmission's CLI.

    Parameters:
    - torrent_file (str): Path to torrent file.
    """
    command = [
        "transmission-remote",
        "--download-dir",
        target_dir,
        "--add",
        torrent_file
    ]

    # Run the transmission-remote command
    result = subprocess.run(command, capture_output=True, text=True, check=True)
    print(f"torrent added successfully: {result.stdout}")

    command = [
        "transmission-remote",
        "-t all",
        "--start",
    ]
    result = subprocess.run(command, capture_output=True, text=True, check=True)
    print(f"download started: {result.stdout}")


def check_download_status() -> bool:
    """
    Checks if the torrent download is complete using Transmission's CLI.

    Returns:
    - bool: True if all torrents are downloaded, False otherwise.
    """
    command = ["transmission-remote", "-l"]
    result = subprocess.run(command, capture_output=True, text=True, check=True)

    # Check if any torrents are still downloading
    for line in result.stdout.splitlines():
        if "100%" in line:
            return True
    return False


def wait_for_download_completion(wait_seconds: int, interval_check_seconds: int) -> bool:
    """
    Polls the torrent download status every minute until it completes or 15 minutes have passed.

    Returns:
    - bool: True if the torrent is downloaded within 15 minutes, False otherwise.
    """
    start_time = datetime.now()
    timeout = timedelta(seconds=wait_seconds)

    while datetime.now() - start_time < timeout:
        if check_download_status():
            return True
        print(f"download still in progress, checking again in {interval_check_seconds} seconds...")
        time.sleep(interval_check_seconds)  # Wait for 1 minute before checking again

    return False
