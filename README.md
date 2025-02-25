![Static Badge](https://img.shields.io/badge/ampyq-darkred?logoColor=white)
![Go Badge](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![GitHub last updated (main)](https://img.shields.io/github/last-commit/cse-25/amrita_pyq/main)

# Amrita PYQ CLI Tool 
<a href="https://cse-25.github.io/amrita_pyq/"> <img src="https://img.shields.io/badge/Explore%20Docs-Amrita%20PYQ-blueviolet?style=for-the-badge&logo=rocket" alt="Amrita PYQ Documentation"/> </a>
<div align="center">
  <img src="https://github.com/user-attachments/assets/44822d21-0e1d-4b3e-baa3-3450b5cc14fc" alt="amritalogo" width="450" height="175">
</div>

## Overview
The Amrita PYQ CLI Tool is a command-line interface application that simplifies accessing previous year question papers (PYQs) by directly fetching and displaying them in your default browser.

---

## Prerequisites
Ensure the following requirements are met before using the application:

- **Network**: You must be connected to the **Amrita WiFi** or use a **VPN** to access the network.  
- **Development Environment**:  
  - **Golang** must be installed on your system.  
  - Ensure a working **Go Compiler** is set up.
  - **golangci** must be installed for testing (optional) using the command:
    ```go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest```

---

## Usage Instructions

### Step 1: Connect to Amrita WiFi or VPN
Before using the tool, ensure that you are connected to **Amrita WiFi** or using a **VPN** to access the network. This is mandatory for fetching the previous year question papers.

### Step 2: Clone the Repository
```bash
git clone https://github.com/CSE-25/amrita_pyq
```

---

### Step 3: Run the Application
1. Open the `main.go` file in your preferred code editor (eg:VS Code).  
2. Execute the application:
   ```bash
   go run main.go
   ```
3. A menu will appear with options to choose from.

### Output
![main](https://github.com/user-attachments/assets/148324a7-38fa-48d3-b53a-a292ce928254)

---

### Step 4: Select an Option
1. Use the menu to choose an option.  
2. The tool will process your request and fetch the desired question paper.

### Output
![image](https://github.com/user-attachments/assets/f0a3051c-ff9a-4f7a-84b3-716682ab4ff6)

![image](https://github.com/user-attachments/assets/b4db183c-47e4-4630-8516-07fa016b7131)

---


### Step 5: View the Question Paper
1. **Filter**: Use the `/` key to apply a filter and narrow down the list of PDFs.  
2. The question paper will automatically open as a **PDF** in your default web browser.

### Output
![image](https://github.com/user-attachments/assets/453959d2-10c5-40b0-b981-38791c96c2ac)

![image](https://github.com/user-attachments/assets/dc8e80ce-73ae-4033-901a-d6ccbd387b2f)

![image](https://github.com/user-attachments/assets/b904263f-97c0-439f-9aa5-3c4f477683b6)

The File is opened in PDF format in your default browser(in this case, Microsoft Edge) 

---
   
### Step 6: Continue or Exit
After the PDF is displayed, you will have the following options:
   - **Go Back to Main Menu**: Select `Back to Main Menu` to return to the main menu.  
   - **Exit**: Use the `Quit` option in the menu to close the application.

### Output
![image](https://github.com/user-attachments/assets/7ec80f02-e128-41b8-9907-6ac687e45c17)

---

# Local Development Environment Setup
<p align="center">
  <img src="https://github.com/user-attachments/assets/7c2a4198-cfd9-4451-89a0-44637136e1f4" alt="golang" width="300" height="200">
</p>

## Installing Golang
1. Download Golang
   Visit the official [Golang Downloads Page](https://golang.org/dl/) and select the installer appropriate for your operating system:
   - Windows: `.msi` installer  
   - macOS: `.pkg` installer  
   - Linux: `.tar.gz` archive  

2. Install Golang
   - Windows/macOS:  
     Run the installer and follow the on-screen instructions. This will automatically configure Golang and set environment variables.  
   - Linux:  
     Extract the archive and move the files to `/usr/local`:  
     ```bash
     sudo tar -C /usr/local -xzf go<version>.linux-amd64.tar.gz
     ```  
     Add Golang to your `PATH` by appending this line to your shell's configuration file (`~/.bashrc`, `~/.zshrc`, etc.):  
     ```bash
     export PATH=$PATH:/usr/local/go/bin
     ```  
     Save the changes and reload your shell:  
     ```bash
     source ~/.bashrc
     ```

3. Verify Installation
   Open a terminal and run:  
   ```bash
   go version
   ```  
   If Golang is installed correctly, the installed version will be displayed.

> **Note**  
> If you encounter any issues related to `GOPATH`, please refer to the [GOPATH Troubleshooting Guide](https://go.dev/wiki/SettingGOPATH#windows) for detailed instructions.

---

## Run Unit Test
To run all the unit tests for this project use the following command:
```bash
go test ./... -v
```

You will be able to see the output and status for each unit test available.

To run a unit test specific to a package, go to the package location and run:
```bash
go test -run "PACKAGE_NAME"
```

## Run Lint check
```bash
golangci-lint run
```
**Note:** This command will return an output only when it finds an error. Otherwise, no visible change occurs.

---

## Developers

- [Abhinav Ramakrishnan](https://github.com/Abhinav-ark)
- [Ashwin Narayanan S](https://ashrockzzz2003.github.io/portfolio)
- [Naik Mubashir](https://github.com/naikmubashir)
- [Jayadev D](https://github.com/FLASH2332)
- [Nandagopal R Nair](https://github.com/Nandgopal-R)
- [Shyam Sundar Raju](https://github.com/Shyam-Sundar-Raju)
- [Visruth Thayyil Vijind](https://github.com/beppvis)

---

## Documentation
- [Ashwin V A](https://github.com/WinterSun23)
- [R.D.Tarun](https://github.com/RD-Tarun)
- [Shyam Sundar Raju](https://github.com/Shyam-Sundar-Raju)