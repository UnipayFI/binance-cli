#!/bin/bash

# Binance CLI Installer
# This script automatically downloads and installs the latest release for your platform

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# GitHub repository info
REPO="UnipayFI/binance-cli"
BINARY_NAME="binance-cli"

# Function to print colored output
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to detect platform and architecture
detect_platform() {
    local os=$(uname -s | tr '[:upper:]' '[:lower:]')
    local arch=$(uname -m)
    
    case "$os" in
        linux*)
            PLATFORM="linux"
            ;;
        darwin*)
            PLATFORM="macos"
            ;;
        cygwin*|mingw*|msys*)
            PLATFORM="windows"
            ;;
        *)
            print_error "Unsupported operating system: $os"
            exit 1
            ;;
    esac
    
    case "$arch" in
        x86_64|amd64)
            ARCH="x86_64"
            ;;
        arm64|aarch64)
            ARCH="arm64"
            ;;
        *)
            print_error "Unsupported architecture: $arch"
            exit 1
            ;;
    esac
    
    print_info "Detected platform: $PLATFORM-$ARCH"
}

# Function to get latest release tag
get_latest_release() {
    print_info "Getting latest release information..."
    
    if command -v curl >/dev/null 2>&1; then
        LATEST_TAG=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    elif command -v wget >/dev/null 2>&1; then
        LATEST_TAG=$(wget -qO- "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    else
        print_error "Neither curl nor wget is available. Please install one of them."
        exit 1
    fi
    
    if [ -z "$LATEST_TAG" ]; then
        print_error "Failed to get latest release information"
        exit 1
    fi
    
    print_info "Latest release: $LATEST_TAG"
}

# Function to download and extract
download_and_extract() {
    local filename="$BINARY_NAME-$PLATFORM-$ARCH.tar.gz"
    local download_url="https://github.com/$REPO/releases/download/$LATEST_TAG/$filename"
    
    print_info "Downloading $filename..."
    
    # Download file
    if command -v curl >/dev/null 2>&1; then
        curl -L -o "$filename" "$download_url"
    elif command -v wget >/dev/null 2>&1; then
        wget -O "$filename" "$download_url"
    else
        print_error "Neither curl nor wget is available"
        exit 1
    fi
    
    # Check if download was successful
    if [ ! -f "$filename" ]; then
        print_error "Failed to download $filename"
        exit 1
    fi
    
    print_info "Extracting $filename..."
    
    # Extract archive
    tar -xzf "$filename"
    
    # Clean up archive
    rm "$filename"
    
    # Make binary executable (for Unix-like systems)
    if [ "$PLATFORM" != "windows" ]; then
        chmod +x "$BINARY_NAME"
    fi
    
    print_success "Installation completed!"
    print_info "Binary location: $(pwd)/$BINARY_NAME"
}

# Function to show usage instructions
show_usage() {
    echo
    print_info "Usage instructions:"
    echo "1. Set up your environment variables:"
    echo "   export API_KEY=\"your_api_key\""
    echo "   export API_SECRET=\"your_api_secret\""
    echo
    echo "2. Run the binary:"
    echo "   ./$BINARY_NAME --help"
    echo
    if [ "$PLATFORM" != "windows" ]; then
        echo "3. (Optional) Move to a directory in your PATH:"
        echo "   sudo mv $BINARY_NAME /usr/local/bin/"
    fi
}

# Main execution
main() {
    print_info "Starting Binance CLI installation..."
    
    # Check if we're running as root (warn user)
    if [ "$EUID" -eq 0 ]; then
        print_warning "Running as root. Consider running as a regular user."
    fi
    
    # Detect platform and architecture
    detect_platform
    
    # Get latest release
    get_latest_release
    
    # Download and extract
    download_and_extract
    
    # Show usage instructions
    show_usage
    
    print_success "Binance CLI has been successfully installed!"
}

# Run main function
main "$@"
