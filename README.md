# **Amrita PYQ CLI Tool Documentation**

## **Overview**
The Amrita PYQ CLI Tool is a command-line interface application that simplifies accessing previous year question papers (PYQs) by directly fetching and displaying them in your default browser.

---

## **Prerequisites**
Ensure the following requirements are met before using the application:

- **Network**: You must be connected to the **Amrita WiFi** or use a **VPN** to access the network.  
- **Development Environment**:  
  - **Golang** must be installed on your system.  
  - Ensure a working **Go Compiler** is set up.

---

## **Usage Instructions**

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
![image](https://github.com/user-attachments/assets/862562c1-4db7-4200-82d2-07f58f6fad09)

---

### Step 4: Select an Option
1. Use the menu to choose an option.  
2. The tool will process your request and fetch the desired question paper.

### Output
![image](https://github.com/user-attachments/assets/0eb76815-cfd8-4f48-8c67-bec55bc65f11)

---


### Step 5: View the Question Paper
1. The question paper will automatically open as a **PDF** in your default web browser.

### Output
![image](https://github.com/user-attachments/assets/d926aad6-2adc-41ce-a1b4-c287f1ccab1d)
![image](https://github.com/user-attachments/assets/62e4f6cb-2d38-41ec-9e00-2e1d347114e2)

---
   
### Step 6: Continue or Exit
1. After the PDF is displayed, you will have the option to:
   - **Continue**: Press `1` to go back to the main menu.  
   - **Exit**: Press `0` to close the application.

### Output
![image](https://github.com/user-attachments/assets/cc4cdd7e-f1f4-4926-a00e-00adc098cb66)

---

# **Local Development Environment Setup**

### **Installing Golang**
1. **Download Golang**  
   Visit the official [Golang Downloads Page](https://golang.org/dl/) and select the installer appropriate for your operating system:
   - **Windows**: `.msi` installer  
   - **macOS**: `.pkg` installer  
   - **Linux**: `.tar.gz` archive  

2. **Install Golang**
   - **Windows/macOS**:  
     Run the installer and follow the on-screen instructions. This will automatically configure Golang and set environment variables.  
   - **Linux**:  
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

3. **Verify Installation**  
   Open a terminal and run:  
   ```bash
   go version
   ```  
   If Golang is installed correctly, the installed version will be displayed.

> **Note**  
> If you encounter any issues related to `GOPATH`, please refer to the [GOPATH Troubleshooting Guide](https://go.dev/wiki/SettingGOPATH#windows) for detailed instructions.
