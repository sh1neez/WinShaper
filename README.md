# 🛠️ WinShaper

<p align="center">
  <img src="icon.svg" width="128" height="128" alt="WinShaper Logo">
</p>

<p align="center">
  <strong>Shape Your Windows</strong><br>
  A lightweight, modern Windows 11 tweaker built with Go. Personalize your OS with a Mica-inspired UI.
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
**WinShaper** is a fast, safe, and open-source tool designed to fine-tune Windows 11. Forget bloated tweakers with outdated designs. WinShaper offers a sleek, modern interface powered by the **Fyne** toolkit and a robust set of features under the hood.

## ✨ Powerful Features

### ⚙️ System UI
* **Dark Mode:** Toggle system-wide dark theme with one click.
* **Classic Context Menu:** Restore the familiar Windows 10 right-click menu.
* **Accent Colors:** Instantly change Windows selection and accent colors.

### 👁️ Desktop Clean
* **Shortcut Arrows:** Remove the arrow overlay from desktop shortcuts.
* **Recycle Bin:** Easily hide the Recycle Bin from your desktop.

### 🎨 Taskbar & Explorer
* **Taskbar Clock:** Show seconds in the system tray clock.
* **File Extensions:** Quickly toggle visibility of file extensions.
* **Hidden Files:** Show or hide hidden system files with ease.

---

## 🚀 Quick Start

1. Go to the **[Releases](https://github.com/sh1neez/WinShaper/releases)** section.
2. Download the latest `WinShaper.exe`.
3. Run as Administrator.
4. Select your desired tweaks and click **Apply Changes**.

## 🛠️ Build from source

To build the project yourself, you will need **Go 1.20+** and the **Fyne toolkit** installed.

```bash
# Clone the repository
git clone [https://github.com/sh1neez/WinShaper.git](https://github.com/sh1neez/WinShaper.git)

# Navigate to the folder
cd WinShaper

# Install Fyne and dependencies
go mod download

# Run the project
go run main.go