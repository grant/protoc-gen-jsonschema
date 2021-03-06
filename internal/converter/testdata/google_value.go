package testdata

const GoogleValue = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "arg": {
            "oneOf": [
                {
                    "type": "array"
                },
                {
                    "type": "boolean"
                },
                {
                    "type": "number"
                },
                {
                    "type": "object"
                },
                {
                    "type": "string"
                }
            ]
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
