<h1 style="display: flex; align-items: center; gap: 10px;">
  <img src="https://raw.githubusercontent.com/Sn0wo2/QuickNote/refs/heads/main/Frontend/public/logo.png" alt="Logo" width="30">
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

Preview(release version): https://note.me0wo.top  **_Full features_**

Vercel(build version): https://demo.qn.me0wo.top **_Only frontend preview(cant save)_**

---

## 📦 **Docker** (RECOMMEND)

`docker pull me0wo/quicknote`

* [DockerHub](https://hub.docker.com/r/me0wo/quicknote)

* [docker-compose.yml](docker-compose.yml)

---

## 📥 **Download**

[![GitHub release](https://img.shields.io/github/v/release/Sn0wo2/QuickNote?logo=github)](https://github.com/Sn0wo2/QuickNote/releases)

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

> _Build docker image please use goreleaser!_

```bash
# Frontend not embedded in the binary
go build -mod=readonly -trimpath \
  -tags="mysql postgres sqlite sqlserver" \
  -o="QuickNote" \
  -ldflags="-s -w -buildid= -extldflags=-static" \
  -gcflags="all=-d=ssa/check_bce/debug=0" \
  -asmflags="-trimpath" main.go

cd Frontend

bun install
bun run build

mv static/* ../

cd ../ && ./QuickNote(.exe)
```

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