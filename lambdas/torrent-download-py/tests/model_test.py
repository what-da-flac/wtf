import json

from models.model import Torrent


def test_torrent_from_json():
    payload = """
{
    "id": "abc-123",
    "hash": "123",
    "total_size": "2 Gb",
    "name": "my torrent"
}
    """
    data = json.loads(payload)
    torrent = Torrent(**data)
    print()
    print(torrent)
    print(json.dumps(torrent.model_dump(exclude_unset=True)))
    pass
