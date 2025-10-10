# GitHub Actions Workflows

This directory contains automated workflows for the Go-Learning repository.

## Workflows

### 1. Auto Compile Go Code (`go-compile.yml`)

**Purpose:** Automatically compiles all Go files in the repository to ensure code quality and catch compilation errors early.

**Trigger:**
- On `push` when any `.go` file changes
- On `pull_request` when any `.go` file changes

**What it does:**
1. Checks out the repository
2. Sets up Go (version 1.24)
3. Discovers all subdirectories containing `.go` files
4. For each subdirectory:
   - If it contains a `go.mod` file, compiles as a Go module
   - Otherwise, compiles each `.go` file individually
5. Reports build status with clear success/failure indicators
6. Provides a summary of total files compiled
7. Fails the workflow if any compilation error occurs

**Example Output:**
```
🔍 Discovering subfolders with Go files...

📁 Found Go files in: 01-basics
  Building individual Go files...
    Building hello.go...
    ✅ Build successful for hello.go

📊 Summary: Compiled 1 file(s)/module(s)
✅ All builds successful
```

### 2. Daily Topic Issue Creation (`daily-topic-issue.yml`)

**Purpose:** Automatically creates GitHub issues for tracking learning progress on topics.

**Trigger:**
- Scheduled: Daily at 6:00 AM UTC (via cron)
- Manual: Via `workflow_dispatch` (Actions tab → Run workflow button)

**What it does:**
1. Checks out the repository with full git history
2. Detects the most recently added topic folder:
   - Primary method: Uses git log to find the most recently committed folder
   - Fallback method: Uses the highest numbered folder (e.g., `13-databases` > `01-basics`)
3. Extracts a human-readable display name from the folder
   - Example: `01-basics` → `Basics`, `02-control-flow` → `Control Flow`
4. Creates a GitHub issue with:
   - Title: "Learning Progress: {Topic Name}"
   - Label: `daily-topic`
   - Body: Structured template with instructions for tracking progress

**Permissions Required:**
- `issues: write` - To create new issues
- `contents: read` - To read the repository structure

**Issue Template:**
The created issue includes:
- Topic folder reference
- Purpose and tracking instructions
- Markdown template for daily updates
- Auto-generated notice

## Testing the Workflows

### Testing Auto Compile Go Code

1. Create or modify a `.go` file in any topic folder
2. Commit and push the changes (or create a PR)
3. Go to the **Actions** tab to see the workflow run
4. The workflow should compile all Go files and report success/failure

Example test:
```bash
# Create a valid Go file - should pass
echo 'package main
import "fmt"
func main() { fmt.Println("Test") }' > 01-basics/test.go

git add 01-basics/test.go
git commit -m "Test compilation workflow"
git push
```

### Testing Daily Topic Issue Creation

**Option 1: Manual Trigger**
1. Go to the **Actions** tab in GitHub
2. Select "Daily Topic Issue Creation" workflow
3. Click "Run workflow" button
4. Select the branch and click "Run workflow"
5. Check the **Issues** tab for the newly created issue

**Option 2: Wait for Scheduled Run**
- The workflow automatically runs daily at 6:00 AM UTC
- Check the **Issues** tab the next morning for new issues

## Local Testing

You can test the workflow logic locally before pushing:

### Test Compilation Script
```bash
cd /path/to/Go-Learning

# Run the compilation logic
for dir in */; do
  folder="${dir%/}"
  if [[ "$folder" == .* ]]; then continue; fi
  if ls "$folder"/*.go 2>/dev/null | grep -q ".go"; then
    echo "Found Go files in: $folder"
    cd "$folder"
    if [ -f "go.mod" ]; then
      go build .
    else
      for gofile in *.go; do
        if [ -f "$gofile" ]; then
          go build "$gofile"
          rm -f "${gofile%.go}"
        fi
      done
    fi
    cd ..
  fi
done
```

### Test Topic Detection Script
```bash
cd /path/to/Go-Learning

# Find the most recent topic folder
folders=$(ls -d */ 2>/dev/null | grep -E '^[0-9]{2}-' | sort -V)
latest_folder=$(echo "$folders" | tail -1 | sed 's/\/$//')
echo "Latest folder: $latest_folder"

# Extract display name
display_name=$(echo "$latest_folder" | sed 's/^[0-9]*-//' | sed 's/-/ /g')
echo "Display name: $display_name"
```

## Troubleshooting

### Compilation Workflow

**Issue:** Workflow fails with "go: cannot find main module"
- **Solution:** This happens when trying to build a package-only file. The workflow handles this by building individual files with `go build filename.go`

**Issue:** Binary files appear in the repository
- **Solution:** The workflow automatically cleans up binaries after compilation. The `.gitignore` is also configured to ignore common binary names.

### Issue Creation Workflow

**Issue:** No topic folder detected
- **Solution:** Ensure your topic folders follow the naming pattern `XX-topic-name` (e.g., `01-basics`, `02-control-flow`)

**Issue:** Wrong folder detected as "latest"
- **Solution:** The workflow uses git history first, then falls back to folder names. Ensure you commit new folders to git.

**Issue:** Issue creation fails
- **Solution:** Check that the workflow has proper permissions (`issues: write`). This should be configured by default.

## Modifying the Workflows

### Change the Go Version
Edit `go-compile.yml`:
```yaml
- name: Set up Go
  uses: actions/setup-go@v5
  with:
    go-version: '1.24'  # Change this version
```

### Change the Schedule Time
Edit `daily-topic-issue.yml`:
```yaml
on:
  schedule:
    - cron: '0 6 * * *'  # Change this cron expression
```

Cron format: `minute hour day month weekday`
- `0 6 * * *` = 6:00 AM UTC daily
- `0 12 * * *` = 12:00 PM UTC daily
- `0 0 * * 1` = Midnight UTC every Monday

### Customize Issue Template
Edit the `--body` content in `daily-topic-issue.yml` to change the issue template format.

## GitHub Actions Best Practices

✅ **DO:**
- Review workflow runs regularly in the Actions tab
- Fix compilation errors promptly when workflows fail
- Use manual triggers for testing before relying on scheduled runs
- Keep workflows simple and focused on one task

❌ **DON'T:**
- Commit and push broken `.go` files (the workflow will fail)
- Modify workflow files without testing locally first
- Ignore workflow failures (they indicate real issues)

## Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Workflow Syntax](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions)
- [GitHub CLI Documentation](https://cli.github.com/manual/)
- [Go GitHub Actions](https://github.com/actions/setup-go)
