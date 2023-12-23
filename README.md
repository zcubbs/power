# Project Power

Project Power is a dynamic project generation system designed to create project structures based on user-defined specifications. It utilizes a flexible blueprint system, allowing for extensible and customizable project configurations.

<center>
<img src="docs/assets/power_logo_white.png" />
</center>

## Features

- **Dynamic Blueprint System**: Supports various technology stacks with the ability to add new ones through plugins.
- **YAML Spec Files**: Each blueprint can define its configuration options using YAML spec files, enabling dynamic user input handling.
- **Plugin Support**: Extend functionality with custom blueprints by adding plugins.
- **gRPC and RESTful API Support**: Built to support both gRPC and RESTful APIs, ensuring compatibility with a wide range of client applications.

## Project Structure

- `cmd/`: Entry points for the application or services.
- `proto/`: Protocol Buffer files defining gRPC services.
- `blueprint/`: Contains blueprints for different technologies and the core logic for blueprint handling.
    - `golang/`, `springboot/`, etc.: Directories for each technology stack.
    - `blueprint.go`: Manages blueprint registration and component generation.
    - `plugin.go`: Handles dynamic plugin loading.
- `designer/`: Handles the logic of generating project blueprints based on specifications.
    - `designer.go`: Core file for orchestrating project generation.
- `tests/`: Test scripts and data.
- `.github/`: Configuration files for GitHub Actions and workflows.

## Getting Started

TODO

## Blueprint System

The Blueprint System is a core component of Project Power, designed to offer a flexible and dynamic approach to project generation. It allows for the definition and registration of various project "blueprints," each representing a different technology stack or project configuration.

### How It Works

- **Blueprint Registration**: Blueprints are registered in the system with a unique identifier. Each blueprint consists of two key parts:
    - **Generator**: A Go implementation that defines how the project structure for that blueprint will be generated based on given specifications.
    - **YAML Spec File**: A YAML file that outlines the configuration options available for that blueprint. This includes parameters like choice of database, use of specific libraries, and other customizable settings.

- **Dynamic Blueprint Loading**: Project Power supports dynamic loading of blueprints through plugins. This enables the system to extend its capabilities without modifying the core codebase.

## Plugin System

The Plugin System in Project Power enhances the flexibility and extensibility of the blueprint system. It allows developers to add new blueprints dynamically as plugins, without needing to alter the core application code.

### Understanding the Plugin System

- **Dynamic Loading**: The system can dynamically load custom blueprints at runtime. These blueprints are compiled as Go plugins (`.so` files) and can be added to the system simply by placing them in a designated directory.
- **YAML Specification**: Each plugin must have an associated YAML specification file detailing its configuration options, similar to built-in blueprints.

### Creating a Plugin

To create a new plugin for the system, follow these steps:

1. **Implement the Blueprint Generator**:
    - Implement a Go file that adheres to the `ComponentGenerator` interface, specifically the `Generate` method.
    - The implementation should define how to generate the project structure based on the provided specs.

2. **Create a YAML Spec File**:
    - Write a YAML file that details the configuration options for your plugin, structured according to the `BlueprintSpec` struct.

3. **Compile as a Go Plugin**:
    - Compile your blueprint implementation into a Go plugin (`.so` file) using the Go compiler.
    - The command generally looks like: `go build -buildmode=plugin -o yourplugin.so yourplugin.go`.

### Adding a Plugin to the System

1. **Place the Plugin and YAML Spec**:
    - Place the compiled `.so` file and its corresponding YAML spec file in the designated plugin directory of the system.

2. **System Auto-Loading**:
    - On startup, the system will automatically load and register all plugins found in the plugin directory.
    - Each plugin's YAML spec file will also be loaded, allowing the system to understand the configuration options it provides.

### Example Plugin Structure

wip...
