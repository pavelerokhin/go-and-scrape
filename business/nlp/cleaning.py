import re

ebala_regex = r"(?i)данное сообщение .+ функции иностранного агента\.?"

def spoil_ebala(text):
    return re.sub(ebala_regex, "", text)