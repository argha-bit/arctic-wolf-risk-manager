# Greetings
The Arctic-Wolf-Risk-Manager project has been developed to store and retrieve ongoing risks, adhering to the requirements outlined in the assessment file. The project has been structured in a modular fashion, enabling the easy separation of responsibilities across different components, aligning with the principles of **separation of concerns**. This approach not only improves maintainability but also facilitates easier writing of unit tests, ensuring code reliability and scalability.

Additionally, the project manages concurrent access to in-memory data storage by utilizing mutex locks, ensuring data consistency. The design also allows for flexibility, as the data storage adapter can be easily modified to accommodate a new database without requiring significant changes to the overall architecture, minimizing potential impact on other modules.

To enhance portability, a Dockerfile has been included, allowing the application to run seamlessly on any system regardless of the underlying operating system or dependencies.

Please feel free to share any suggestions for improvement. I look forward to your feedback.

Thank you,
### Argha