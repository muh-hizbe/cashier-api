# Cashier API Documentation

## Overview
This document provides comprehensive documentation for the Cashier API endpoints.

---

## Health Check

### Check Server Status
**Endpoint:** `GET /health`

**Description:** Returns the current status of the server.

**Response:**
- Returns server health status

---

## Products

### Get All Products
**Endpoint:** `GET /products`

**Description:** Retrieves a list of all products.

**Response:**
- Returns an array of product objects

### Create New Product
**Endpoint:** `POST /products`

**Description:** Creates a new product in the system.

**Request Body:**
- Product details (name, price, stock)

**Response:**
- Returns the created product object

### Get Product by ID
**Endpoint:** `GET /products/{id}`

**Description:** Retrieves a specific product by its ID.

**Parameters:**
- `id` (path parameter) - The unique identifier of the product

**Response:**
- Returns the product object

### Update Product
**Endpoint:** `PUT /products/{id}`

**Description:** Updates an existing product.

**Parameters:**
- `id` (path parameter) - The unique identifier of the product

**Request Body:**
- Updated product details (name, price, stock)

**Response:**
- Returns the updated product object

### Delete Product
**Endpoint:** `DELETE /products/{id}`

**Description:** Deletes a product from the system.

**Parameters:**
- `id` (path parameter) - The unique identifier of the product

**Response:**
- Returns deletion confirmation

---

## Categories

### Get All Categories
**Endpoint:** `GET /categories`

**Description:** Retrieves a list of all categories.

**Response:**
- Returns an array of category objects

### Create New Category
**Endpoint:** `POST /categories`

**Description:** Creates a new category in the system.

**Request Body:**
- Category details (name, description)

**Response:**
- Returns the created category object

### Get Category by ID
**Endpoint:** `GET /categories/{id}`

**Description:** Retrieves a specific category by its ID.

**Parameters:**
- `id` (path parameter) - The unique identifier of the category

**Response:**
- Returns the category object

### Update Category
**Endpoint:** `PUT /categories/{id}`

**Description:** Updates an existing category.

**Parameters:**
- `id` (path parameter) - The unique identifier of the category

**Request Body:**
- Updated category details (name, description)

**Response:**
- Returns the updated category object

### Delete Category
**Endpoint:** `DELETE /categories/{id}`

**Description:** Deletes a category from the system.

**Parameters:**
- `id` (path parameter) - The unique identifier of the category

**Response:**
- Returns deletion confirmation
