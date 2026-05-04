package chesspuzzle

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Tibz-Dankan/BiTE/internal/models"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

// csvPath is relative to the server/ directory where the seeder is run from.
// Run with: go run cmd/seed/main.go (from inside server/)
// const csvPath = "internal/data/lichess_puzzles_filtered.csv"

func ChessPuzzleSeeder() {
	// ── 1. Load environment ──────────────────────────────────────────────────
	env := os.Getenv("GO_ENV")

	if env == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file")
		}
		log.Println("Loaded .env file")
	}

	log.Println("GO_ENV:", env)

	// ── 2. Connect to database (mirrors db.go DSN logic exactly) ────────────
	var dsn string
	var csvPath string
	var err error

	switch env {
	case "development":
		dsn = os.Getenv("BiTE_DEV_DSN")
		csvPath, err = filepath.Abs("./internal/data/lichess_puzzles_filtered.csv")
		if err != nil {
			log.Fatalf("Error getting absolute path for CSV file: %v", err)
		}
	case "production":
		dsn = os.Getenv("BiTE_PROD_DSN")
		// csvPath, err = filepath.Abs("./internal/data/lichess_puzzles_filtered.csv")
		// if err != nil {
		// 	log.Fatalf("Error getting absolute path for CSV file: %v", err)
		// }
		csvPath = "/app/server/internal/data/lichess_puzzles_filtered.csv"
	default:
		log.Fatalf("Unrecognized GO_ENV: '%s' — set GO_ENV to 'development' or 'production'", env)
	}

	if dsn == "" {
		log.Fatalf("DSN environment variable is empty for GO_ENV=%s", env)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")

	// ── 3. Open the CSV file ─────────────────────────────────────────────────
	f, err := os.Open(csvPath)
	if err != nil {

		log.Fatalf(
			"Cannot open CSV file at '%s': %v\n"+
				"Make sure you are running this command from inside the server/ directory:\n"+
				"  cd server && go run cmd/seed/main.go",
			csvPath, err,
		)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true

	// Skip header row
	if _, err := reader.Read(); err != nil {
		log.Fatalf("Cannot read CSV header: %v", err)
	}

	// ── 4. Stream CSV and insert in batches ──────────────────────────────────
	const batchSize = 1000

	var batch []models.ChessPuzzle
	inserted := 0
	skipped := 0
	startTime := time.Now()

	fmt.Printf("\nSeeding chess_puzzles from: %s\n", csvPath)
	// fmt.Println("Progress printed every 10,000 rows...\n")
	fmt.Println("Progress printed every 10,000 rows...")

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Skipping malformed row: %v", err)
			skipped++
			continue
		}

		// Guard: minimum 9 columns required
		if len(row) < 9 {
			skipped++
			continue
		}

		// ── Parse columns ────────────────────────────────────────────────────
		// 0=PuzzleId 1=FEN 2=Moves 3=Rating 4=RatingDeviation
		// 5=Popularity 6=NbPlays 7=Themes 8=GameUrl 9=OpeningTags

		puzzleID := strings.TrimSpace(row[0])
		fen := strings.TrimSpace(row[1])
		moves := strings.TrimSpace(row[2])

		rating, err := strconv.Atoi(strings.TrimSpace(row[3]))
		if err != nil {
			skipped++
			continue
		}

		rd, err := strconv.Atoi(strings.TrimSpace(row[4]))
		if err != nil {
			skipped++
			continue
		}

		popularity, _ := strconv.ParseInt(strings.TrimSpace(row[5]), 10, 16)
		nbPlays, _ := strconv.Atoi(strings.TrimSpace(row[6]))

		themes := pq.StringArray{}
		if strings.TrimSpace(row[7]) != "" {
			themes = pq.StringArray(strings.Fields(row[7]))
		}

		gameUrl := strings.TrimSpace(row[8])

		openingTags := pq.StringArray{}
		if len(row) > 9 && strings.TrimSpace(row[9]) != "" {
			openingTags = pq.StringArray(strings.Fields(row[9]))
		}

		// Derive player color from the FEN active color field:
		// FEN field 2 = whose turn BEFORE the trigger move.
		// The trigger move flips the turn → that is the player's color.
		color := "w"
		fenParts := strings.Fields(fen)
		if len(fenParts) >= 2 && fenParts[1] == "w" {
			color = "b"
		}

		batch = append(batch, models.ChessPuzzle{
			ID:              puzzleID,
			FEN:             fen,
			Moves:           moves,
			Rating:          rating,
			RatingDeviation: rd,
			Volatility:      0.09,
			Popularity:      int16(popularity),
			NbPlays:         nbPlays,
			Themes:          themes,
			GameUrl:         gameUrl,
			OpeningTags:     openingTags,
			Color:           color,
		})

		// ── Flush when batch is full ─────────────────────────────────────────
		if len(batch) >= batchSize {
			if err := flushBatch(db, batch); err != nil {
				log.Printf("Batch insert error: %v", err)
			}
			inserted += len(batch)
			batch = batch[:0]

			if inserted%10000 == 0 {
				elapsed := time.Since(startTime).Round(time.Second)
				rate := int(float64(inserted) / time.Since(startTime).Seconds())
				fmt.Printf("  Inserted: %7d | Skipped: %4d | Elapsed: %s | ~%d rows/sec\n",
					inserted, skipped, elapsed, rate)
			}
		}
	}

	// ── Flush any remaining rows ─────────────────────────────────────────────
	if len(batch) > 0 {
		if err := flushBatch(db, batch); err != nil {
			log.Printf("Final batch insert error: %v", err)
		}
		inserted += len(batch)
	}

	// ── 5. Create indexes after bulk insert ──────────────────────────────────
	// Indexes are always created after the bulk insert — building them on an
	// already-populated table is significantly faster than maintaining them
	// row-by-row during the insert.
	fmt.Println("\nCreating indexes on chess_puzzles...")
	createIndexes(db)

	// ── 6. Final summary ─────────────────────────────────────────────────────
	fmt.Printf("\n✓ Seeding complete\n")
	fmt.Printf("  Inserted : %d\n", inserted)
	fmt.Printf("  Skipped  : %d\n", skipped)
	fmt.Printf("  Duration : %s\n", time.Since(startTime).Round(time.Second))
}

// flushBatch inserts a slice of ChessPuzzle records.
// ON CONFLICT (id) DO NOTHING ensures the seeder is safe to re-run —
// existing puzzles are silently skipped rather than causing an error.
func flushBatch(db *gorm.DB, batch []models.ChessPuzzle) error {
	return db.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		}).
		CreateInBatches(batch, len(batch)).
		Error
}

// createIndexes adds performance indexes to chess_puzzles after the bulk insert.
func createIndexes(db *gorm.DB) {
	indexes := []struct {
		name string
		sql  string
	}{
		{
			"idx_chess_puzzles_rating",
			`CREATE INDEX IF NOT EXISTS idx_chess_puzzles_rating
			    ON chess_puzzles (rating)`,
		},
		{
			"idx_chess_puzzles_nb_plays",
			`CREATE INDEX IF NOT EXISTS idx_chess_puzzles_nb_plays
			    ON chess_puzzles ("nbPlays" DESC)`,
		},
		{
			"idx_chess_puzzles_popularity",
			`CREATE INDEX IF NOT EXISTS idx_chess_puzzles_popularity
			    ON chess_puzzles (popularity DESC)`,
		},
		{
			"idx_chess_puzzles_color",
			`CREATE INDEX IF NOT EXISTS idx_chess_puzzles_color
			    ON chess_puzzles (color)`,
		},
		{
			"idx_chess_puzzles_themes",
			`CREATE INDEX IF NOT EXISTS idx_chess_puzzles_themes
			    ON chess_puzzles USING GIN (themes)`,
		},
		{
			"idx_chess_puzzles_opening_tags",
			`CREATE INDEX IF NOT EXISTS idx_chess_puzzles_opening_tags
			    ON chess_puzzles USING GIN ("openingTags")`,
		},
	}

	for _, idx := range indexes {
		if err := db.Exec(idx.sql).Error; err != nil {
			log.Printf("  ✗ %s: %v", idx.name, err)
		} else {
			fmt.Printf("  ✓ %s\n", idx.name)
		}
	}
}

// func init() {
// 	go func() {
// 		time.Sleep(2 * time.Minute)
// 		// time.Sleep(15 * time.Second)
// 		env := os.Getenv("GO_ENV")
// 		if env == "production" {
// 			ChessPuzzleSeeder()
// 		}
// 	}()
// }
