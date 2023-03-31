# porter 🏗
a package manager... 
![image](https://user-images.githubusercontent.com/38886930/229208345-8452f001-18ba-4246-b080-8bfe9ad2d306.png)

# Technical Overview of Porter
Security and privacy are critical considerations in the design of the Porter application. The application provides a secure and reliable way for users to manage their application installations and configurations, including the ability to work with private repositories.

Porter allows users to add private repositories as crates, enabling them to manage sensitive application data securely. When adding a private repository as a crate, users are prompted to provide their Git credentials, which are securely stored in Porter's configuration files using industry-standard encryption mechanisms.

Additionally, Porter employs strict security measures to protect user data and ensure the integrity of the installation process. For instance, during the installation process, Porter verifies the authenticity and integrity of the downloaded application files to prevent tampering and ensure that the application is safe to install.

Overall, the Porter application provides robust security and privacy features that enable users to manage their application installations and configurations with confidence, even when working with private repositories.

## Crates
Crates are the building blocks of Porter. They are Git repositories with YAML configuration files that specify how to install and configure an application. A user can add a crate to their configuration by providing the Git repository URL.

## Container
A container is a reference to an organization that contains multiple repositories. Porter will scan all repositories in the specified organization for crates that can be added to the user's configuration.

## Install and Uninstall
Porter provides the ability to install or uninstall an application that is specified in the YAML configuration file. The installation process involves parsing the YAML configuration file to determine the necessary steps for installing the application. The uninstallation process will reverse the installation steps and remove the application from the system.

# Examples
To add a crate to a user's configuration, the following command can be used:
```bash
porter add crate github.com/justjordant/cli-ops
```

To install an application that has been added to the user's configuration:
```bash
porter install justjordant/cli-ops
```

To uninstall an application that has been installed using Porter:
```bash
porter uninstall justjordant/cli-ops
```
