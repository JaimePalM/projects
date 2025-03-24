import pdfplumber
import pandas as pd

# Path to your PDF
pdf_path = 'MIR_RPA_2024.pdf'

search_str1 = 'LAURA'  # Replace with your first search string
search_str2 = 'LEÓN'  # Replace with your second search string

with pdfplumber.open(pdf_path) as pdf:
    for page_num, page in enumerate(pdf.pages, start=1):
        tables = page.extract_tables()
        if tables:
            for table in tables:
                for row in table:
                    # Check if both search strings are in the same row (anywhere in the row)
                    row_text = ' '.join(cell if cell is not None else '' for cell in row)
                    if search_str1 in row_text and search_str2 in row_text:
                        print(f"✅ Found on page {page_num}: {row}")