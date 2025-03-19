import json, sys

data = json.load(sys.stdin)  # Reads JSON directly from stdin
a = float(data['A'])
b = float(data['B'])
#result = f"Sum: {a + b}"
print(a + b)  # Output only the number, no text
