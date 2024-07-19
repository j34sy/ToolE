# ToolE
A light-weight command line tool for simple tasks

## Cobra and Viper
This CLI tool is utilizing SPF13's Go CLI Frameworks/ Libraries Cobra and Viper, as well as Cobra-CLI for easier setup

More info about these tools can be found [here](https://go.dev/solutions/clis), [here](https://github.com/spf13/cobra) and [here](https://github.com/spf13/viper).


## Key-Features

This CLI tool is designed to help you to work more fluently on the command-line.

You need to create a small ToDo? - No Problem

You want to remeber something? - Go ahead

Or you need a subnet? - You'll get it quickly

You need to do some maths? - I got you


More features will be designed in the future.

# Commands

## todos

List the ToDos in the ToDo datafile within the used ToolE directory (you define it with --config), that aren't done yet.

With the flags you can customized the output:
    -v/--verbose: show the descriptions (if set) for each ToDo and show the marking state (done?)
    -a/--all: show all ToDos, include those marked as done

### add

Adds a ToDo to the data file.
A ToDo must have an unique title and can have an optional description. The add command needs at least one argument (the title).

### done

Marks a ToDo as done in the data file.
To do so the command requires one argument that's either the title of the ToDo or its ID shown with the ```bash todos``` command.

### remove

Removes a ToDo from the data file.
To do so the command requires one argument that's either the titleor its ID shown with the ```bash todos``` command.


## subnet

Displays the subent information about a given IPv4 address and its CIDR.
These information contains: Subnet mask, Network address, Broadcast address, Usable host range, total hosts, and usable hosts.


## Full Command List

```bash
ToolE -h  # prints information about anything
ToolE -v  # prints the version of ToolE
ToolE  # not much, look for yourself
ToolE --config /path/to/your/desired/directory  # global flag for any command, specifiing with directory to use
ToolE subnet 10.0.0.1/24  # prints subnet info about any given IPv4 address with CIDR
ToolE todos  # lists current not done ToDos
ToolE todos -v  # includes the descriptions
TooLE todos -a  # includes done ToDos
ToolE todos add MyToDo "Optional Description"  # adds a ToDo with unique title and optional description
ToolE todos done MyToDo | 0  # marks ToDo as done by either title or number in list (of ToolE todos)
ToolE todos remove MyToDo | 0  # removes ToDo from list by either title or number in list (of ToolE todos)
```

