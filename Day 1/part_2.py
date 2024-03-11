def main():
    book_path = "input.txt"
    text = get_text(book_path)
    line_list = get_line_list(text)
    sum = 0
    for line in line_list:
        sum += get_num_from_string(line)
    print(sum)
    return

def str_to_int(digit: str) -> int:
    conv = {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }
    return conv[digit]

def get_num_from_string(line: str) -> int:
    number = 0
    possible_digits = {
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9"
    }
    first_digit = ""
    last_digit = ""
    for char in line:
        first_digit += char
        for digit in possible_digits:
            if first_digit.find(digit) != -1:
                if first_digit.isalpha():
                    digit=str_to_int(digit)
                number += int(digit)
        if number > 0:
            break   
    for char in line[::-1]:
        last_digit = char + last_digit
        for digit in possible_digits:
            if last_digit.find(digit) != -1:
                if last_digit.isalpha():
                    digit=str_to_int(digit)
                number = number * 10 + int(digit)
        if number > 10:
            break   
    return number

def get_line_list(text: str) -> list:
    lines = text.split('\n')
    return lines

def get_text(file: str) -> str:
    with open(file) as f:
        return f.read()
        
main()