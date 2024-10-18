## 1) Installing Go and setting up Go environment for windows

To install Go and set up the Go environment on Windows, follow these steps:

### Step 1: Download Go

1. **Visit the official Go website**:
   - Go to [https://go.dev/dl/](https://go.dev/dl/)
2. **Download the installer**:
   - Choose the Windows `.msi` installer. For most users, it's best to pick the latest version and ensure it matches your system architecture (64-bit is most common). Example file: `go1.21.1.windows-amd64.msi`.

### Step 2: Install Go

1. **Run the installer**:

   - Double-click the `.msi` installer you downloaded and follow the prompts.
   - By default, Go will be installed in the directory `C:\Program Files\Go`.

2. **Verify the installation**:
   - Open **Command Prompt** or **PowerShell** and type the following command:
     ```bash
     go version
     ```
   - You should see output similar to:
     ```
     go version go1.21.1 windows/amd64
     ```

### Step 3: Set Up Go Environment Variables

Go uses two important environment variables:

1. **GOPATH**: This is the workspace directory where your Go code, packages, and executables will be stored.
   - It’s often set to something like `C:\Users\<YourUsername>\go`.
2. **GOROOT**: This is the location where Go is installed. It’s automatically set during installation to something like `C:\Program Files\Go`, so you usually don’t need to modify it.

#### Steps to set environment variables:

1. **Open Environment Variables**:

   - Press `Win + R`, type `sysdm.cpl`, and press Enter.
   - In the **System Properties** window, go to the **Advanced** tab and click **Environment Variables**.

2. **Set the GOPATH**:

   - Under **User variables**, click **New**.
   - Set the **Variable name** as `GOPATH`.
   - Set the **Variable value** to your desired directory, e.g., `C:\Users\<YourUsername>\go`.

3. **Update the Path**:

   - Under **System variables**, select the `Path` variable and click **Edit**.
   - Add the following paths (if not already present):
     - `%GOROOT%\bin` (e.g., `C:\Program Files\Go\bin`)
     - `%GOPATH%\bin` (e.g., `C:\Users\<YourUsername>\go\bin`)

4. **Save the changes**:
   - Click **OK** to close all windows.

### Step 4: Test the Go Environment

1. Open **Command Prompt** or **PowerShell**.
2. Run the following command to check if the environment is set up correctly:
   ```bash
   go env
   ```
   - This should display a list of Go environment variables, including `GOROOT` and `GOPATH`.

### Step 5: Write and Run a Simple Go Program

1. **Create a Go workspace directory**:

   - In **Command Prompt**, run:
     ```bash
     mkdir %GOPATH%\src\hello
     ```

2. **Write a simple Go program**:

   - Open your text editor (e.g., Notepad++) and create a file named `main.go` inside the `hello` directory:

     ```go
     package main

     import "fmt"

     func main() {
         fmt.Println("Hello, Go!")
     }
     ```

3. **Build and run the Go program**:
   - In **Command Prompt**, navigate to your Go workspace:
     ```bash
     cd %GOPATH%\src\hello
     ```
   - To run the program, type:
     ```bash
     go run main.go
     ```
   - You should see the output:
     ```
     Hello, Go!
     ```

---

You now have Go installed and a working development environment!
