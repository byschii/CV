from playwright.sync_api import sync_playwright
import os

with sync_playwright() as p:
    browser = p.chromium.launch()   

    # open output.html in the browser
    page = browser.new_page()
    page.goto("file://" + os.path.abspath("output.html"))

    # print the page to a PDF file
    page.pdf(path="output.pdf")

    browser.close()


