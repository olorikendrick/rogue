package store

import (
    "os/exec"
    "runtime"
)

type Installer string

const (
    Brew   Installer = "brew"
    Winget Installer = "winget"
    Apt    Installer = "apt"
    Pkg    Installer = "pkg"
    Dnf    Installer = "dnf"
    Pacman Installer = "pacman"
)

type Host struct {
    Os           string      `json:"os"`
    Architecture string      `json:"architecture"`
    WordSize     byte        `json:"word_size"`
    Installers   []Installer `json:"installers"`
}

func GetHost() Host {
    arch := runtime.GOARCH
    word := 64
    if ^uintptr(0) == uintptr(^uint32(0)) {
        word = 32
    }

    return Host{
        Os:           runtime.GOOS,
        Architecture: arch,
        WordSize:     byte(word),
        Installers:   getInstallers(),
    }
}

func getInstallers() []Installer {
    candidates := []Installer{Brew, Winget, Apt, Pkg, Dnf, Pacman}
    var available []Installer
    for _, i := range candidates {
        if _, err := exec.LookPath(string(i)); err == nil {
            available = append(available, i)
        }
    }
    return available
}
