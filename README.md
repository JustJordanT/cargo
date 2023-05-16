# Cargo 🏗 - (WIP NOT READY FOR USE)
The package manager... 

<img width="627" alt="image" src="https://user-images.githubusercontent.com/38886930/233756499-0e429e9d-9180-4cd5-9754-dd8eb9c4ba8a.png">


# Technical Overview of Cargo
Security and privacy are critical considerations in the design of the Cargo. Cargo provides a secure and reliable way for users to manage their application installations and configurations, including the ability to work with private repositories.

Cargo allows users to add private repositories as crates, enabling them to manage sensitive application data securely. When adding a private repository as a crate, users are prompted to provide their Git credentials, which are securely stored in Cargo's configuration files using industry-standard encryption mechanisms.

Additionally, Cargo employs strict security measures to protect user data and ensure the integrity of the installation process. For instance, during the installation process, Cargo verifies the authenticity and integrity of the downloaded application files to prevent tampering and ensure that the application is safe to install.

Overall, the Cargo application provides robust security and privacy features that enable users to manage their application installations and configurations with confidence, even when working with private repositories.

## Crates
Crates are the building blocks of Cargo. They are Git repositories with YAML configuration files that specify how to install and configure an application. A user can add a crate to their configuration by providing the Git repository URL.

## Container
A container is a reference to an organization that contains multiple repositories. Cargo will scan all repositories in the specified organization for crates that can be added to the user's configuration.

## Install and Uninstall
Cargo provides the ability to install or uninstall an application that is specified in the YAML configuration file. The installation process involves parsing the YAML configuration file to determine the necessary steps for installing the application. The uninstallation process will reverse the installation steps and remove the application from the system.

# Examples
To add a crate to a user's configuration, the following command can be used:
```bash
cargo crate add github.com/justjordant/cli-ops
cargo crate remvoe github.com/justjordant/cli-ops

```
To add a container point to a org
```bash
cargo container add github.com/justjordant
cargo container remove github.com/justjordant
```

To install an application that has been added to the user's configuration:
```bash
cargo crate install justjordant/cli-ops
```

To uninstall an application that has been installed using Cargo:
```bash
cargo uninstall justjordant/cli-ops
```
