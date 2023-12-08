param(
    [string]$sourceDir = ".",
    [string]$targetDirs = "."
)

foreach ($dir in $targetDirs.Split(' ')) {
    if (-not (Test-Path $dir)) {
        New-Item -ItemType Directory -Path $dir | Out-Null
    }
    # Print the path of the directory being copied to
    Write-Host "Copying files to $dir"
    Copy-Item -Path "${sourceDir}/*" -Destination $dir -Recurse
}
