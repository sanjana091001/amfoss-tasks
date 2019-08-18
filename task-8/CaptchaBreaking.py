import pytesseract
from PIL import Image
value=Image.open("/home/ubuntu/Downloads/untitled.png")
text= pytesseract.image_to_string(value)
print(text)
