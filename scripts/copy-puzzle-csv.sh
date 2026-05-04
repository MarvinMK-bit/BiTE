#!/bin/bash

if [ -f ~/lichess_puzzles_filtered.csv ]; then
  echo "Copying puzzle CSV into app-stack_puzzle-data volume..."
  docker run --rm \
    -v ~/lichess_puzzles_filtered.csv:/src/lichess_puzzles_filtered.csv:ro \
    -v app-stack_puzzle-data:/dest \
    alpine cp /src/lichess_puzzles_filtered.csv /dest/lichess_puzzles_filtered.csv
  echo "Done."
else
  echo "CSV not found at ~/lichess_puzzles_filtered.csv — skipping."
fi