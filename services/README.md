# services

This directory contains all backend services. File structure supports services on any programming language. A Dockerfile image definition should be present for each service.

Since all services are written in Golang, no specific considerations are for the time being (like unit tests, mocks, etc). If additional programming languages are used in the future, we should adapt Makefile/Codebuild accordingly.
