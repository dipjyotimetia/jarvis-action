# jarvis

```yaml
name: Merge checks

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:

  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Run pr merge checks
      uses: dipjyotimetia/jarvis@main
      with:
        PAT: ${{ secrets.GITHUB_TOKEN }}
        OWNER: ""
        GITHUB_REPOSITORY: ""
```
