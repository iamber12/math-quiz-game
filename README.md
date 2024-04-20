# Go Quiz Application

This application conducts a timed quiz using questions from a CSV file, displaying results on completion.

## Features

- Load questions and answers from CSV.
- Timed challenge with customizable duration.
- Command-line based interaction.

## Installation

```bash
git clone https://github.com/iamber12/quiz-game
cd quiz-app
go build
```

## Usage
Execute the quiz application:

```bash
./quiz-app -f path/to/your/csvfile.csv -t timeout_in_seconds
```

### Options

- `-f`: Path to the CSV file with questions (default: `problems.csv`).
- `-t`: Quiz duration in seconds (default: 30).

### CSV Format

The CSV file should have no header and be formatted as:
```
question1,answer1
question2,answer2
```

## Example

To start a 30-second quiz:

```bash
./quiz-app -f my_questions.csv -t 30
```
