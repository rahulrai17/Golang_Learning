## Install Go and set up the Go environment on Linux.

### Step 1: Download Go

1. **Visit the official Go website**:

   - Go to [https://go.dev/dl/](https://go.dev/dl/).

2. **Download the Go tarball**:
   - Download the latest Go tarball for Linux (`.tar.gz`). You can download it directly using the terminal. For example, to download version `go1.21.1` (you should replace it with the latest version):
     ```bash
     wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz
     ```

### Step 2: Install Go

1. **Extract the tarball**:

   - Extract the tarball to `/usr/local`, which is the recommended Go installation directory:
     ```bash
     sudo tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz
     ```

2. **Verify the installation**:
   - Check if Go is properly installed by checking the version. Add the Go binary to your `PATH` temporarily by running:
     ```bash
     /usr/local/go/bin/go version
     ```
   - You should see output similar to:
     ```
     go version go1.21.1 linux/amd64
     ```

### Step 3: Set Up Go Environment Variables

Go requires setting two important environment variables:

1. **GOPATH**: This is your workspace directory, where Go will store your code, packages, and binaries. It’s usually set to something like `$HOME/go`.
2. **GOROOT**: This is the directory where Go is installed. It’s automatically set to `/usr/local/go`, so you don't need to modify it manually unless you install Go somewhere else.

#### Steps to set the environment variables:

1. **Open the shell configuration file**:

   - If you use `bash`, edit the `.bashrc` file:
     ```bash
     nano ~/.bashrc
     ```
   - For `zsh` users, edit the `.zshrc` file:
     ```bash
     nano ~/.zshrc
     ```

2. **Add Go to your PATH and set GOPATH**:
   - Add the following lines at the end of the file:
     ```bash
     # Set Go environment variables
     export PATH=$PATH:/usr/local/go/bin
     export GOPATH=$HOME/go
     export PATH=$PATH:$GOPATH/bin
     ```
3. **Save and close the file**:

   - Press `CTRL + X` to exit, `Y` to confirm changes, and then `Enter` to save the file.

4. **Reload the shell configuration**:
   - Apply the changes by running:
     ```bash
     source ~/.bashrc
     ```

### Step 4: Verify the Go Environment

1. **Check the environment variables**:

   - Run the following command to check if the `GOROOT` and `GOPATH` are set correctly:
     ```bash
     go env
     ```
   - You should see the `GOROOT` as `/usr/local/go` and `GOPATH` as `$HOME/go`.

2. **Test the Go installation**:
   - Run a simple Go program to ensure everything is working.

### Step 5: Write and Run a Simple Go Program

1. **Create a workspace directory**:

   - Make a directory for your Go code:
     ```bash
     mkdir -p $GOPATH/src/hello
     ```

2. **Write a simple Go program**:

   - Create a file named `main.go` in the `hello` directory:
     ```bash
     nano $GOPATH/src/hello/main.go
     ```
   - Add the following code:

     ```go
     package main

     import "fmt"

     func main() {
         fmt.Println("Hello, Go!")
     }
     ```

3. **Run the Go program**:
   - Navigate to the directory and run the program:
     ```bash
     cd $GOPATH/src/hello
     go run main.go
     ```
   - You should see the output:
     ```
     Hello, Go!
     ```

### Step 6: Optional - Install Go from Package Manager (Alternative Method)

If you prefer using a package manager (like `apt` on Ubuntu), you can install Go with the following command:

```bash
sudo apt update
sudo apt install golang-go
```

However, this may not always give you the latest version of Go. It’s generally recommended to download the latest release directly from the official website as shown earlier.

---

You now have Go installed and your development environment set up on Linux!
