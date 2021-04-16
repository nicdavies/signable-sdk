# Signable SDK
A CLI tool to aid integration with the Signable API V1. It acts as a full-mock of the Signable V1 API!

## Features

- Authentication
- Test all V1 API endpoints locally.
- Receive webhooks.
- Upload PDF files.
- Supports success and error responses.

## Getting Started

### Commands

#### serve
Description: starts the sdk.

Flags: `--without-webhooks`

### Config
// TODO


## TODO
- [ ] Build `serve` command (webserver component)
- [ ] Build authentication handler (Bearer tokens)
- [ ] Build JSON response stubs
- [ ] Build `config` command to handle custom config (json format), and allow user to specify path to config file
- [ ] Accept input flags
- [ ] Build up `branding` endpoints
- [ ] Build up `contact` endpoints
- [ ] Build up `envelope` endpoints
- [ ] Build up `settings` endpoints
- [ ] Build up `team` endpoints
- [ ] Build up `template` endpoints
- [ ] Build up `user` endpoints
- [ ] Build up `webhook` endpoints
- [ ] Build Webhook triggers
- [ ] Build output logging to CLI (similar to NGROK cli)

## Dreams
- [ ] Parameterised JSON response stubs, that parse and replace
      with the user's input data.
- [ ] Config option to set the webserver URL.
- [ ] Parameterize timestamps in JSON response stubs.
- [ ] Parameterize `href` urls in JSON response stubs (use the configured webserver url & port).
  
