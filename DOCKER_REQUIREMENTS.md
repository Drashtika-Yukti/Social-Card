# Docker Requirements

Because the Social Forge engine relies on a headless Chromium instance to render absolute pixel-perfect CSS templates dynamically, several underlying system dependencies must be properly orchestrated in the runtime environment.

## 1. System Packages (APT)

The environment must install the following shared objects (so) and fonts to execute `chromedp` successfully:

- `chromium` - The core browser binary itself.
- `fonts-liberation` - Crucial open-source fonts so text renders correctly across environments.
- `libnss3` - Network Security Services required by Chrome.
- `libatk-bridge2.0-0` - Accessibility toolkit bridge.
- `libxcomposite1`, `libxrandr2`, `libgbm1` - Display compositing and buffering libraries (even though we run headless).
- `libasound2` - Audio libraries required by the Chrome runtime.
- `curl` - (Added) Required for querying the `/health` endpoint directly inside the container.

## 2. Docker Architecture

### Build Stage (`golang:1.26-bullseye`)
A modern Go version (`1.26`) must be used to match the module configuration. No C-libraries (`CGO_ENABLED=0`) are used in the native Go build, ensuring maximum portability.

### Runtime Stage (`debian:bullseye-slim`)
Provides the minimal OS footprint required to apt-install the Chromium dependencies. 

## 3. Container Configurations

- **Network Interface**: Binds natively on port `9090`.
- **Environment Variables**:
  - `PORT`: (Required) Overrides default binding port.
  - `API_KEY`: (Required) Passed globally inside the container to authorize inbound requests.
  - `ENV`: Should be set to `production` for security.
