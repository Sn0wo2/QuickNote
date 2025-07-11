from typing import Optional, MutableMapping

import requests


def fetch_latest_release(owner: str, repo: str, proxies: Optional[MutableMapping[str, str]] = None) -> dict:
    resp = requests.get(
        f"https://api.github.com/repos/{owner}/{repo}/releases/latest",
        headers={"Accept": "application/vnd.github.v3+json"},
        proxies=proxies
    )
    resp.raise_for_status()
    return resp.json()
