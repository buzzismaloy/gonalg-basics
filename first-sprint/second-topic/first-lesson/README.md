# First lesson

```mermaid

graph LR
    A["GO DATA TYPES"]

    A --> B["Basic Types"]
    A --> C["Composite Types"]
    A --> D["References"]

    %% Basic Types
    B --> B1["Boolean"]
    B --> B2["Numeric"]
    B --> B3["String"]

    B1 --> B1a["bool"]

    B2 --> B2b["Integers"]
    B2 --> B2c["Unsigned Integers"]
    B2 --> B2d["Floating Point"]
    B2 --> B2e["Complex Numbers"]

    B2b --> B2b1["int, int8, int16, int32, int64"]
    B2c --> B2c1["uint, uint8, uint16, uint32, uint64, uintptr"]
    B2d --> B2d1["float32, float64"]
    B2e --> B2e1["complex64, complex128"]

    B3 --> B3a["string"]
    B3 --> B3b["rune"]

    %% Composite Types
    C --> C1["Array"]
    C --> C2["Slice"]
    C --> C3["Map"]
    C --> C4["Struct"]

    C1 --> C1a["Fixed length"]
    C2 --> C2a["Dynamic size"]
    C3 --> C3a["Key-value storage"]
    C4 --> C4a["Custom data grouping"]
