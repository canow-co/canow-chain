# Canow chain

Canow ledger

## Development Notes

### Absent IDE Features in Integration Tests Sources

IDE features may be unavailable in the [integration tests](./tests/integration) sources because these files are marked with a custom build tag (`integration`) and so the language server does not build them by default.

If you have faced this issue, set up the language server to pass `integration` build tag to the build system.

If you are using **VS Code** IDE with **gopls** language server, add the following entry to the workspace settings:
```json
    "gopls": {
        "build.buildFlags": [
            "-tags=integration"
        ]
    }
```
