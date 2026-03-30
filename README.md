# 🛠️ WinShaper

<p align="center">
  <img src="icon.ico" width="128" height="128" alt="WinShaper Logo">
</p>

<p align="center">
  <strong>Shape Your Windows Experience.</strong><br>
  A lightweight, modern Windows 11 tweaker built with Go. Fast, open-source, and clutter-free.
</p>

<p align="center">
  <a href="https://github.com/sh1neez/WinShaper/releases">
    <img src="https://img.shields.io/github/v/release/sh1neez/WinShaper?style=for-the-badge&color=3b82f6" alt="Release">
  </a>
  <a href="https://github.com/sh1neez/WinShaper/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/sh1neez/WinShaper?style=for-the-badge&color=d946ef" alt="License">
  </a>
</p>

---

## 📖 About
**WinShaper** was born out of a desire for a cleaner Windows 11 experience without the overhead of heavy "optimizer" suites. It's a straightforward tool designed to put the most annoying Windows 11 UI choices back in your hands. 

Built using **Go** and the **Fyne** toolkit, it offers a native-feeling interface to modify registry settings safely and instantly.

## ✨ Features

### ⚙️ System Customization
* **Dark Mode:** Toggle system-wide dark/light themes instantly.
* **Classic Context Menu:** Restore the Windows 10 style right-click menu (no more "Show more options" clicks).
* **Custom Selection Color:** Use the advanced color picker to change the default Windows highlight and tracking colors.

### 📂 Desktop & Explorer
* **Shortcut Arrows:** Clean up your desktop by removing the arrow overlay from icons.
* **Recycle Bin Management:** Hide or show the Recycle Bin icon with a single check.
* **File Visibility:** Quickly toggle hidden files and file extensions without digging through folder options.

### 📍 Taskbar Personalization
* **Alignment Control:** Easily switch between Centered and Left-aligned icons.
* **Search Modes:** Choose exactly how search appears—Hidden, Icon only, or the full Search Box.
* **Task View Toggle:** Remove the Task View button to save space.
* **Seconds in Clock:** Enable high-precision time by showing seconds in the system tray.

---

## 🚀 Getting Started

1. Download the latest `WinShaper.exe` from the **[Releases](https://github.com/sh1neez/WinShaper/releases)** page.
2. Run the application (Administrator privileges are required to modify registry keys).
3. Select your preferred tweaks.
4. Click **Apply Changes** — the app will automatically restart `explorer.exe` to refresh your UI.

## 🛠️ Build from Source

If you want to compile it yourself, make sure you have **Go 1.20+** and a C compiler (required by Fyne for graphics).

```bash
# Clone the repo
git clone [https://github.com/sh1neez/WinShaper.git](https://github.com/sh1neez/WinShaper.git)

# Enter the directory
cd WinShaper

# Install Go dependencies
go mod download

# Run the app
go run .