"""
    Work in progress until I figure out what fields I need for each table
    and how to generate fake data for them.
"""
import psycopg2
from faker import Faker
import random
import json
from decouple import Config
import os

# Load environment variables from .env file
config = Config(os.path.join(os.path.dirname(__file__), '../.env'))

# Create a Faker instance
fake = Faker()

# Database credentials from the .env file
db_name = config('DB_NAME')
db_user = config('DB_USER')
db_password = config('DB_PASSWORD')
db_host = config('DB_HOST')
db_port = config('DB_PORT')

# Connect to the database
connection = psycopg2.connect(
    dbname=db_name,
    user=db_user,
    password=db_password,
    host=db_host,
    port=db_port
)

cursor = connection.cursor()

# Function to insert fake data into the tables
def insert_fake_data():
    # Insert fake products
    for _ in range(10):
        name = fake.company()
        image_url = [fake.image_url() for _ in range(3)]
        description = fake.text()
        meta = json.dumps({"key": fake.word()})
        cursor.execute(
            "INSERT INTO products (name, image_url, description, meta) VALUES (%s, %s, %s, %s)",
            (name, image_url, description, meta)
        )
    
    connection.commit()

    # Get product IDs
    cursor.execute("SELECT id FROM products")
    product_ids = [item[0] for item in cursor.fetchall()]

    # Insert fake variants
    for product_id in product_ids:
        for _ in range(3):
            name = fake.word()
            sku = fake.bothify(text='???-####', letters='ABCDEFGHIJKLMNOPQRSTUVWXYZ')
            ean = fake.ean13()
            color = fake.color_name()
            size = fake.random_element(elements=('S', 'M', 'L', 'XL'))
            price = random.randint(1000, 9999)
            original_price = price + random.randint(100, 999)
            stock = random.randint(0, 100)
            meta = json.dumps({"key": fake.word()})
            cursor.execute(
                "INSERT INTO variants (product_id, name, sku, ean, color, size, price, original_price, stock, meta) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)",
                (product_id, name, sku, ean, color, size, price, original_price, stock, meta)
            )

    connection.commit()

    # TODO: Continue this pattern for other tables as needed

# Call the function to insert fake data
insert_fake_data()

# Close the database connection
cursor.close()
connection.close()
