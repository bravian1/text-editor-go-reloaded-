Go-Reloaded (Text Processing Tool)

This is a command-line tool written in Go for processing text files. It provides various functionalities for manipulating text data, including:

Converting hexadecimal and binary numbers to decimal: Numbers enclosed in parentheses with suffix "(hex)" or "(bin)" are converted to their decimal equivalents.
Uppercasing and lowercasing letters: Words enclosed in parentheses with the suffix "(up)" or "(low)" are converted to uppercase or lowercase, respectively.
Capitalizing letters: The suffix "(cap)" applied to a word capitalizes the first letter. The variation "(cap, N)" capitalizes the first N letters of the preceding word.
Article insertion: The tool inserts the indefinite article "an" before a word if the preceding word ends in "a" or "A" and the following word starts with a vowel.
Punctuation adjustment: Leading punctuation marks are moved to the end of the preceding word. Punctuation marks that are single words are combined with the neighboring words. Special handling is applied for apostrophes.
Usage

The tool is executed from the command line using the following syntax:

go run . <input_file> <output_file>

<input_file>: Path to the text file you want to process.
<output_file>: Path to the file where the processed text will be saved.
Example

Consider a text file with the following content:

This is a F1 (hex) text file 101 (bin) with some data to process. Let's test the CASING (up) and lowercasing (low) . We can also capitalize (cap) the first letter. How about (cap, 2) articles? An example 0 (bin) would be nice.

Running the tool with this file as input:

go run . sample.txt result.txt

The output file (result.txt) will contain the following processed text:

This is a 241 text file 5 with some data to process. Let's test the CASING and lowercasing. We can also Capitalize the first letter. How About articles? An example 0 would be nice.

Dependencies

This code requires the following Go packages which are all standard go packages:

bufio: For reading data from the input file line by line.
log: For logging errors.
os: For opening and closing files.
strconv: For converting between string and numeric data types.
strings: For string manipulation functions.
