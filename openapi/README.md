# openapi

OpenAPI files to generate stubs and classes for microservices single source of truth.

## Requirements

Install required packages:

```bash
make install
```

## Usage

To generate files in all languages:

```bash
make gen
```

## File Structure

All openapi files are stored in `open-api/` directory. Within that directory, there is a `domains/` directory which contains data types which are not used on any endpoint, and cannot be directly generated as others.

All generated files end up in `gen/golang/` directory.



