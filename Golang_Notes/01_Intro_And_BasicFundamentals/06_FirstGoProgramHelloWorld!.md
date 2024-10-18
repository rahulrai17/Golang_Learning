## First Go Program "Hello World! "

### 1. Writing and Running Your First Go Program in **Terminal** (Command Prompt/PowerShell)

#### Step 1: Set Up Your Go Workspace

By default, Go uses the `GOPATH` environment variable to define the workspace, typically located at `C:\Users\<YourUsername>\go`. You can structure your project under this directory.

Create a new directory for your project:

1. Open **Command Prompt** or **PowerShell**.
2. Navigate to your `GOPATH`:
   ```bash
   cd %USERPROFILE%\go\src
   ```
3. Create a folder for your project (e.g., `helloworld`):
   ```bash
   mkdir helloworld
   cd helloworld
   ```

#### Step 2: Write Your First Go Program

1. Create a Go file called `main.go`:

   ```bash
   notepad main.go
   ```

   This will open Notepad. Add the following Go code:

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, World from Go!")
   }
   ```

2. Save the file and close Notepad.

#### Step 3: Run Your Go Program

1. Back in the terminal, use the `go run` command to compile and execute the program:

   ```bash
   go run main.go
   ```

2. You should see the following output in the terminal:
   ```
   Hello, World from Go!
   ```

#### Step 4: Compile and Run as a Binary

You can also compile the program to create an executable `.exe` file.

1. Run the following command to compile the program:

   ```bash
   go build
   ```

2. You should see a `helloworld.exe` file generated in the directory. Run the executable:

   ```bash
   .\helloworld.exe
   ```

3. This will output:
   ```
   Hello, World from Go!
   ```

---

### 2. Writing and Running Your First Go Program in **VS Code**

#### Step 1: Set Up VS Code for Go

1. **Install VS Code**: If you don't have VS Code installed, download it from [here](https://code.visualstudio.com/Download).
2. **Install the Go Extension**: Open VS Code, and install the Go extension by searching for **"Go"** in the Extensions view (`Ctrl+Shift+X`) or by navigating to [Go Extension for VS Code](https://marketplace.visualstudio.com/items?itemName=golang.Go).

#### Step 2: Create Your Go Project

1. Open **VS Code**.
2. Click **File** > **Open Folder** and navigate to your Go workspace: `C:\Users\<YourUsername>\go\src`.
3. Create and open a new folder (e.g., `helloworld`).

4. In VS Code, create a new file called `main.go` by clicking **New File** and naming it `main.go`.

#### Step 3: Write the Go Code

1. In the `main.go` file, write the following Go code:

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, World from VS Code!")
   }
   ```

2. Save the file (`Ctrl+S`).

#### Step 4: Running the Go Program in VS Code

1. **Open the Integrated Terminal**: You can open the terminal within VS Code by selecting **Terminal** > **New Terminal** from the top menu or using the shortcut `Ctrl+` (backtick).
2. Navigate to your project directory in the terminal if not already there:

   ```bash
   cd %USERPROFILE%\go\src\helloworld
   ```

3. Run the program using the following command:

   ```bash
   go run main.go
   ```

4. You should see the following output in the terminal inside VS Code:
   ```
   Hello, World from VS Code!
   ```

#### Step 5: Compile the Program in VS Code

1. To compile your Go program into an executable, run:

   ```bash
   go build
   ```

2. Run the executable from the terminal:

   ```bash
   .\helloworld.exe
   ```

3. The output will be:
   ```
   Hello, World from VS Code!
   ```

---

### Summary

#### Using Terminal:

1. Write the Go program in `main.go`.
2. Run it using `go run main.go` or compile it using `go build` and then execute the binary.

#### Using VS Code:

1. Set up your Go workspace, write the code in `main.go`, and use the integrated terminal to run `go run` or `go build`.
2. VS Code helps with linting, code completion, and project management, making Go development easier and more efficient.

With these steps, you're ready to build Go programs in both the terminal and VS Code!
