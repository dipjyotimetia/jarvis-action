# jarvis-action

```yaml
name: Merge checks

on:
  schedule:
    - cron: '15 * * * *'
  pull_request:
    types:
      - opened

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Run pr merge checks
      uses: dipjyotimetia/jarvis@main
      with:
        PAT: ${{ secrets.PAT }}
        OWNER: "dipjyotimetia"
        GITHUB_REPOSITORY: "jarvis"
```

