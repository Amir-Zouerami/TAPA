# TAPA

![tapa-cover-github-min](https://github.com/user-attachments/assets/8f7bf40a-bcd4-492e-8526-257252d1118b)

## 🚀 Introduction
TAPA (The Actual Postman Alternative) was born out of frustration with existing HTTP client tools. Postman feels bloated and heavy, while other alternatives that tried to replace it either lack essential features, lack a modern UI/UX, or Feel bloated with unnecessary functionalities. I wanted something that was:

- **Fast and lightweight** ⚡
- **Minimal but beautiful** 🎨
- **Highly efficient and keyboard-friendly** ⌨️
- **Fully offline-first, with no forced cloud integration** 🌍
- **Focused solely on core functionality, without unnecessary clutter** 🧹

## 🎯 Philosophy

TAPA was made specifically for developers and tries to stay out of your way when working. Working as a full-stack developer, I have transformed all my frustrations with other HTTP clients into a few core principles for TAPA:

### ✨ **Minimalist & Modern UI**

Unlike other HTTP clients that overload the screen with excessive options (API mocks, etc.), TAPA keeps it simple and distraction-free. It focuses on the essentials:

- **A sleek, clean interface** with a modern design
- **Keyboard shortcut-friendly** (inspired by VSCode) for quick navigation
- **Clean and intuitive design**—nothing should feel overwhelming or cluttered
- **Dark mode support** for a comfortable midnight development experience!

### ⚡ **Fast & Efficient**

TAPA has a focus on performance for a smooth development experience:

- Built with **Go + React 19** (using React Compiler for performance optimizations)
- Uses **Wails instead of Electron**, near native performance with a much smaller footprint
- Optimized for **speed and efficiency**

### 🔒 **Full Data Ownership**

If you use TAPA, it is assumed you're a developer who WANTS to mess with things and modify your data freely:

- **No forced accounts** — TAPA works entirely offline.
- **No telemetry, data collection, or tracking**—safe for enterprise use.
- **Your data, your way** — all requests and collections are stored in **a single SQLite database**, which is fully accessible.
- **Easy import/export** — you can always export data as JSON or SQLite for backup or sharing.


### 🔌 **Offline-First, No Cloud Nonsense**

Many modern tools push **cloud lock-in** and **AI-based features** at every step. **I hate this approach with a burning passion.**

- **TAPA does neither** — everything is local-first.
- **No AI-powered nonsense** — Just a simple, efficient tool for developers
- If you need synchronization, TAPA takes a **BYOB (Bring Your Own Backend)** approach, meaning:
  - You **host your own** backend.
  - TAPA follows a simple **API contract** to interact with it.

_Encryption_: Currently, TAPA stores data unencrypted for full accessibility, but encryption **may be added in a future stable release**.

## 🚧 Current Status: Work in Progress
🔹 **There is no stable or working version of TAPA yet.**  
🔹 This README serves as a **vision** for its first release.  
🔹 If you’re curious about the progress:
- Check out the **TAPA GitHub Project Page**.
- **Build the app yourself** using Wails.

## ❌ No PRs (For Now)

TAPA is a solo project with a **highly opinionated** design, meant to solve personal frustrations. **For now, PRs won’t be accepted** for two reasons:

1. **TAPA is still under heavy development.** Accepting PRs would slow down progress.
2. **It’s a personal tool first.** TAPA is built to solve my own frustrations. If you resonate with its philosophy, great! Otherwise, it may not be for you.

TAPA is not being built as a community-driven project (at least for now). Once things stabilize, I may consider contributions, but for now, TAPA remains a personal project shared with the open-source community.


## 🛠️ Building TAPA
To build TAPA yourself, you need **Go, Wails, and Node.js** installed. Once dependencies are ready:
```bash
# Clone the repo
git clone https://github.com/Amir-Zouerami/TAPA.git && cd TAPA

# Install frontend dependencies
cd frontend && npm install && cd ..

# Build & run the app
wails dev
```

⚠️ **WARNING**: Setting `TAPA_ENV` environment variable to development (i.e. `TAPA_ENV=development wails dev`) will **use the development database** which has a `_DEV` postfix to its name. In development mode TAPA will seed this database with placeholder data! Backup up everything before you attempt to do this (unless you are a developer, you should NOT).

## 📜 License
TAPA is open-source and available under the **Apache License**.

---

Stay tuned for updates! 🚀
