<h1 style="display: flex; align-items: center; gap: 10px;">
  <img src="https://raw.githubusercontent.com/Sn0wo2/QuickNote/refs/heads/main/Frontend/public/quicknote.svg" alt="Logo" width="77">
  <span><strong>QuickNote</strong></span>
</h1>

> Create and share notes quickly and easily.

[![GitHub License](https://img.shields.io/github/license/Sn0wo2/QuickNote)](LICENSE)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/Sn0wo2/QuickNote/total)
![Docker Pulls](https://img.shields.io/docker/pulls/me0wo/quicknote)
![Docker Stars](https://img.shields.io/docker/stars/me0wo/quicknote)
![Docker Image Size](https://img.shields.io/docker/image-size/me0wo/quicknote)


[![CodeQL](https://github.com/Sn0wo2/QuickNote/actions/workflows/codeql.yml/badge.svg)](https://github.com/Sn0wo2/QuickNote/actions/workflows/codeql.yml)
[![Dependabot Updates](https://github.com/Sn0wo2/QuickNote/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/Sn0wo2/QuickNote/actions/workflows/dependabot/dependabot-updates)
[![Go CI](https://github.com/Sn0wo2/QuickNote/actions/workflows/go.yml/badge.svg)](https://github.com/Sn0wo2/QuickNote/actions/workflows/go.yml)
[![React CI](https://github.com/Sn0wo2/QuickNote/actions/workflows/react.yml/badge.svg)](https://github.com/Sn0wo2/QuickNote/actions/workflows/react.yml)
[![Release](https://github.com/Sn0wo2/QuickNote/actions/workflows/release.yml/badge.svg)](https://github.com/Sn0wo2/QuickNote/actions/workflows/release.yml)

---

## 🎉 **Demo**

**🔗 Preview (Release Version):** [https://note.me0wo.top](https://note.me0wo.top)  
如果你在中国大陆，可以使用这个加速CDN访问：[https://qnote.me0wo.top](https://qnote.me0wo.top)  
*✨ Full features available (updated manually, may be slightly delayed)*

> We’re working on enabling [**automatic deployment**](https://github.com/Sn0wo2/QuickNote/tree/main/Updater) for this preview demo.

**🛠 Vercel (Build Preview):** [https://demo.qn.me0wo.top](https://demo.qn.me0wo.top)  
*🚧 Frontend-only preview — saving is not supported*

---

## 📦 **Docker** (RECOMMEND)
```bash
docker pull me0wo/quicknote
docker compose up -d
```

* [DockerHub](https://hub.docker.com/r/me0wo/quicknote)

* [docker-compose.yml](docker-compose.yml)

---

## 📥 **Download**

[![GitHub release](https://img.shields.io/github/v/release/Sn0wo2/QuickNote?logo=github)](https://github.com/Sn0wo2/QuickNote/releases)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FSn0wo2%2FQuickNote.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2FSn0wo2%2FQuickNote?ref=badge_shield)

---

## 🚀 **Project Status**

| Status | `Developing` |
|--------|--------------|

---

## ✅ **Features**

* ✔️ No login required
* ✔️ High performance
* ✔️ Simple UI
* ✔️ Markdown preview support
* ✔️ Dark Mode
* ✔️ Note sharing
* ️️✔️ Compression

**Planned:**

* 🔒 Encryption
* 🕑 Note history

---

## 🗃️ **Supported Databases**

* **Relational:**

    * MySQL, MariaDB, TiDB, Aurora
    * PostgreSQL, CockroachDB, AlloyDB
    * SQLite3
    * Microsoft SQL Server

---

## 📚 **Docs**

| Status | `Developing` |
|--------|--------------|

---

## ⚙️ **Build Instructions**

### ✅ **Using `GitHub Actions` and `goreleaser`** (RECOMMEND)

Check:

* [`release.yml`](.github/workflows/release.yml)
* [`.goreleaser.yml`](LICENSE)
* [`Dockerfile`](Dockerfile)

---

### 🔧 **Manual Build**

Requires:

* [`Go SDK`](https://go.dev/dl) latest version
* [`Node.js`](https://nodejs.org/zh-cn/download) latest version
* [`pnpm`](https://pnpm.io/installation) latest version
* [`GoReleaser`](https://github.com/goreleaser/goreleaser/releases) latest version

Check:

* [`Makefile`](Makefile)
* [`.goreleaser.yml`](LICENSE)
* [`Dockerfile`](Dockerfile)

> Run:
> ```bash
> make test_release
> ```

---

## 👥 **Contributors**

![Contributors](https://contrib.rocks/image?repo=Sn0wo2/QuickNote)

---

## ⭐ **Star History**

<a href="https://www.star-history.com/#Sn0wo2/QuickNote&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=Sn0wo2/QuickNote&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=Sn0wo2/QuickNote&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=Sn0wo2/QuickNote&type=Date" />
 </picture>
</a>

---

## 📄 **License**

Licensed under [GPL 3.0](LICENSE).


[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FSn0wo2%2FQuickNote.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2FSn0wo2%2FQuickNote?ref=badge_large)