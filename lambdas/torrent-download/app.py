import json
import os
import subprocess
import time
from datetime import datetime, timedelta

from aws_lambda_powertools.utilities.typing import LambdaContext

from models.model import Torrent


def lambda_handler(event: dict, context: LambdaContext):
    wait_seconds = 900
    data = event
    torrent = Torrent(**data)
    return torrent.model_dump(exclude_unset=True)
    target_dir = "/Users/mau/Downloads/xxx"
    # download_torrent(target_dir, torrent_file)
    # filenames = [
    #     "Coldplay - Moon Music (2024) FLAC [PMEDIA] ⭐️/01. Moon Music (Feat. Jon Hopkins).flac",
    #     "Coldplay - Moon Music (2024) FLAC [PMEDIA] ⭐️/02. Feelslikeimfallinginlove.flac",
    #     "Coldplay - Moon Music (2024) FLAC [PMEDIA] ⭐️/03. We Pray (Feat. Little Simz, Burna Boy, Elyanna & Tini).flac",
    #     "Coldplay - Moon Music (2024) FLAC [PMEDIA] ⭐️/04. Jupiter.flac",
    #     "Coldplay - Moon Music (2024) FLAC [PMEDIA] ⭐️/05. Good Feelings (Feat. Ayra Starr).flac",
    #     "Coldplay - Moon Music (2024) FLAC [PMEDIA] ⭐️/06. Rainbow.flac",
    #     "Coldplay - Moon Music (2024) FLAC [PMEDIA] ⭐️/07. Iaam.flac",
    #     "Coldplay - Moon Music (2024) FLAC [PMEDIA] ⭐️/08. Aeterna.flac",
    #     "Coldplay - Moon Music (2024) FLAC [PMEDIA] ⭐️/09. All My Love.flac",
    #     "Coldplay - Moon Music (2024) FLAC [PMEDIA] ⭐️/10. One World.flac",
    #     "Coldplay - Moon Music (2024) FLAC [PMEDIA] ⭐️/INFO.nfo",
    #     "Coldplay - Moon Music (2024) FLAC [PMEDIA] ⭐️/cover.jpg",
    # ]
    # filenames_with_path = [os.path.join(target_dir, filename) for filename in filenames]
    # if wait_for_download_completion(wait_seconds=wait_seconds, interval_check_seconds=10):
    #     print("download completed successfully.")
    #     upload_files_to_s3(filenames_with_path)
    #     print("completed uploading files to s3")
    # else:
    #     print(f"download did not complete within the {wait_seconds} seconds limit.")


def upload_files_to_s3(filenames: list[str]):
    for filename in filenames:
        file_info = os.stat(filename)
        print(f"file name: {filename}  size: {file_info.st_size} bytes")


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
