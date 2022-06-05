FROM gitpod/workspace-full

RUN sudo apt-get update && \
    sudo apt-get install gettext libncurses5 libxkbcommon0 libtinfo5 -y

USER gitpod

# subo 0.5.2
RUN brew tap suborbital/subo && \
    brew install subo && \
    brew install httpie && \
    brew install bat && \
    brew install exa && \
    brew install hey

# ------------------------------------
# Install TinyGo
# ------------------------------------
ARG TINYGO_VERSION="0.23.0"
RUN wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo_${TINYGO_VERSION}_amd64.deb
RUN sudo dpkg -i tinygo_${TINYGO_VERSION}_amd64.deb
RUN rm tinygo_${TINYGO_VERSION}_amd64.deb

# ------------------------------------
# Install Rust support
# ------------------------------------
RUN rustup toolchain uninstall stable-x86_64-unknown-linux-gnu && \
    curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh -s -- -y && \
    curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh && \
    rustup target add wasm32-wasi

# ------------------------------------
# Install Sat (Suborbital)
# ------------------------------------
# "sat_version":"v0.1.4"
RUN git clone --depth=1 https://github.com/suborbital/sat.git && \
    cd sat && \
    go build -o .bin/sat -tags netgo,wasmtime . && \
    sudo cp .bin/sat /usr/local/bin/sat && \
    cd .. && \
    rm -rf sat

