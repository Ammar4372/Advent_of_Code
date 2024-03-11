def main():
    book_path = "input.txt"
    text = get_text(book_path)
    line_list = get_line_list(text)
    sum = 0
    for line in line_list:
        sum += get_num_from_string(line)
    print(sum)
    return

def get_num_from_string(line: str) -> int:
    number = 0
    digits = []
    for char in line:
        if char.isdigit():
            digits.append(int(char))
            break   
    for char in line[::-1]:
        if char.isdigit():
            digits.append(int(char))
            break 
    number = digits[0] *10 + digits[1]
    return number

def get_line_list(text: str) -> list:
    words = text.split('\n')
    return words

def get_text(file: str) -> str:
    with open(file) as f:
       return f.read()
        
main()