from datetime import datetime
from typing import List, Optional

from pydantic import BaseModel


class User(BaseModel):
    id: Optional[str]
    name: Optional[str]
    email: Optional[str]
    image: Optional[str]
    created: Optional[datetime]
    last_login: Optional[datetime]
    is_deleted: Optional[bool]


class TorrentFile(BaseModel):
    id: str
    file_name: str
    file_size: str
    torrent_id: str


class Torrent(BaseModel):
    id: str
    magnet_link: Optional[str] = None
    hash: Optional[str] = None
    piece_count: Optional[int] = None
    piece_size: Optional[str] = None
    total_size: Optional[str] = None
    privacy: Optional[str] = None
    files: Optional[List[TorrentFile]] = None
    name: str
    created: Optional[datetime] = None
    updated: Optional[datetime] = None
    user: Optional[User] = None
    status: Optional[str] = None
    filename: Optional[str] = None
    last_error: Optional[str] = None
