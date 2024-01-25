# nxlabels
simple command to wrap "nixos-rebuild rebuild" with given labels and superuser access who compile the configuration file, copy the changes to a folder who can push on a github repository which have the command line arguments as a commit message 

## Prerequisites
golang **1.21.*** or highter 

## Install
modify the *install.sh* 

```bash
    #!/bin/bash
    user="<user-name-on-the-machine>"
    # file to compile
    file="main.go"
    # executable name
    output="nixlbl"
    # destination directory
    dest="/home/$user/.local/bin"
    # Compile go file
    go build -o $output $file

    # Check if compilation was successful
    if [ $? -eq 0 ]; then
        echo "Compilation successful, moving executable to $dest"
        # Move executable to destination directory
        mv $output $dest
    else
        echo "Compilation failed"
    fi

```

to install the program we need to write two simple cmd commands:

```bash
    cd <project-folder>
    
    sh install.sh 
```

## Configuration
this is the configuration file, this file MUST stay inside the root folder of the project with the name "config.yaml" 

```yaml
    path: <path-to-pushable-configuration>
```

if u want custom name or destination or both, feel free to modify the **viper-config.go** in configurations folder.

```golang 
    func ViperConfig() (*Config, error) {
	// Read Config File and unmarshal to struct
        viper.SetConfigName("config")
        viper.SetConfigType("yaml")
        viper.AddConfigPath("$HOME/Desktop/nixlabels")
        viper.AddConfigPath(".")
    ...
}
```
