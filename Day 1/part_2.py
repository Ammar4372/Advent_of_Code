def main():
    book_path = "input.txt"
    text = get_text(book_path)
    line_list = get_line_list(text)
    number_list = []
    for line in line_list:
        number = get_num_from_string(line)
        number_list.append(number)
    sum_of_number_list = sum_list(number_list)
    print(sum_of_number_list)
    return

def sum_list(number_list: list) -> int:
    sum = 0
    for number in number_list:
        sum += number
    return sum

def get_num_from_string(line: str) -> int:
    digit = []
    number = 0
    num_dict = {
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
    spelled_digit = ""
    for letter in line:
        if letter.isdigit():
            digit.append(int(letter))
        else:
            spelled_digit += letter
        if spelled_digit in num_dict:
            digit.append(num_dict[spelled_digit])
            spelled_digit=""
    try :
        number = digit[0]*10 + digit[len(digit)-1]
    except IndexError :
        number = digit[0]*10 + digit[0]
    print(digit,number)
    return number

def get_line_list(text: str) -> list:
    words = text.split('\n')
    return words

def get_text(file: str) -> str:
    with open(file) as f:
       return f.read()
        
main()