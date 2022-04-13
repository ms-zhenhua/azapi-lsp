# Terraform AzApi Provider Language Server

Experimental version of [terraform-provider-azapi](https://github.com/Azure/terraform-provider-azapi) language server.

## What is LSP

Read more about the Language Server Protocol at https://microsoft.github.io/language-server-protocol/

## Introduction

This project only supports completion/hover/diagnostics for `terraform-provider-azapi`,
not targeting support all language features for `HCL` or `Terraform`. To get the best user experience, 
it's recommended to use it with language server for `Terraform`.

## Features

- Completion of `azapi` resources
- Completion of allowed azure resource types when input `type` in `azapi` resources
- Completion of allowed azure resource properties when input `body` in `azapi` resources, limitation: it only works when use `jsonencode` function to build the JSON
- Better completion for discriminated object
- Completion for all required properties
- Show hint when hover on `azapi` resources
- Show diagnostics for properties defined inside `body`

## Installation

1. Clone this project to local
2. Run `go install` under the project folder.

## Usage

The most reasonable way you will interact with the language server
is through a client represented by an IDE, or a plugin of an IDE.

VSCode extension: [azapi-vscode](https://github.com/ms-henglu/azapi-vscode)

## Credits

- [terraform-ls](https://github.com/hashicorp/terraform-ls): This project is developed based on `terraform-ls`: framework, testing, file sync feature, etc.

