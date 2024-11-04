# gateway


## Diagrams

```mermaid
flowchart TD
    subgraph AWS [AWS]
        SQS[SQS Queue]
        Lambda[Lambda Function]
        ECS[ECS Service]
        
        SQS -->|Trigger| Lambda
        Lambda -->|Invoke| ECS
    end
    
    subgraph Database
        DB[(Database)]
    end

    ECS -->|Store/Fetch Data| DB
```

## Code Generation

This repository defines its models in openapi format.

Generated code should not be checked in the repository, as it should be generated on CI and check interfaces are implemented accordingly.


