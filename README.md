Project: https://roadmap.sh/projects/expense-tracker

## Installation

Cloning project:
```
https://github.com/SALutHere/RoadmapSh-ExpenseTracker.git
```

Installing dependencies (in project's directory):
```
go mod tidy
```

Building (in project's directory):
```
go build
```

Getting started:
```
expense-tracker.exe --help
```

```
Program for expenses management

Usage:
  expense-tracker [flags]
  expense-tracker [command]

Available Commands:
  add         Add a new expense
  clear       Clears your expenses list
  completion  Generate the autocompletion script for the specified shell
  delete      Delete an existing expense
  help        Help about any command
  list        List all the expenses
  summary     Get sum of expenses by specified filters
  update      Update a new expense

Flags:
  -h, --help   help for expense-tracker

Use "expense-tracker [command] --help" for more information about a command.
```

## Usage expample
### Listing

```
./expense-tracker list
```

```
ID                                      Description     Amount  Category        Date
68e666de-5d8e-4a4c-b46e-377774f1b353    Description-1   $10.00  Category-1      19-06-2025 23:41:44
```

### Adding a new expense

```
./expense-tracker add -a 10 -c "Category-1" -d "Description-1"
```

```
A new expense was added successfully! ID: 642b36f8-089d-4e78-8bda-0a90f4dda0b1
```

### Updating a new expense

```
./expense-tracker update -i 642b36f8-089d-4e78-8bda-0a90f4dda0b1 -a 20 -c "Category-2" -d "Description-2"
```

```
An expense was updated successfully! ID: 642b36f8-089d-4e78-8bda-0a90f4dda0b1
```

### Deleting an existent expense

```
./expense-tracker delete -i 642b36f8-089d-4e78-8bda-0a90f4dda0b1
```

```
Expense was deleted successfully! ID: 642b36f8-089d-4e78-8bda-0a90f4dda0b1
```

### Getting summary
#### Total

```
./expense-tracker summary 
```

```
Total expenses: $10.00
```

#### With filters
```
./expense-tracker summary -d 19 -m 6 -y 2025 -c "Category-1"
```

```
Total expenses: $10.00
```

