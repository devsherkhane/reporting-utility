

# Student Reporting Utility 
A modular backend utility engineered for high-throughput student record management and automated document generation. This project emphasizes Separation of Concerns and Type Safety, leveraging Go’s concurrency and performance capabilities.

## Architecture & Design Patterns
   * Unlike standard CRUD apps, this utility is built using the  Service-Repository Pattern to decouple business logic from data persistence.

   * Repository Layer: Handles direct MySQL interactions, ensuring that the database schema is abstracted from the rest of the application.

   * Service Layer: Contains the core business rules, such as calculating report metrics and preparing data for the PDF engine.

   * Modular Reporting: A dedicated PDF generation package built on gopdf that produces granular profile cards and comprehensive tabular summaries.

## Features
   * Custom PDF Engine: Engineered a custom layout logic to transform structured SQL data into print-ready PDF profile cards.

   * Session-Based Security: Implemented secure user authentication using gorilla/sessions to restrict access to sensitive student records.

   * Stateless Routing: Utilized gorilla/mux for clean, RESTful API endpoints and dynamic route matching.

   * Environment Agnostic: Configuration is managed via a config.yaml file, allowing seamless transitions between development and production environments.
 
---
