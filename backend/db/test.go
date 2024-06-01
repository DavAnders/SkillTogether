package db

import (
	"context"
	"database/sql"
	"log"
)

func (q *Queries) GetUserIDFromDiscordTokenTest(ctx context.Context, sessionToken string) (string, error) {
    var discordID string
    query := `SELECT discord_id FROM user_sessions WHERE session_token = $1`
    err := q.db.QueryRowContext(ctx, query, sessionToken).Scan(&discordID)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("No Discord ID found for session token %s", sessionToken)
        } else {
            log.Printf("Error fetching Discord ID for session token %s: %v", sessionToken, err)
        }
        return "", err
    }
    log.Printf("Fetched Discord ID %s for session token %s", discordID, sessionToken)
    return discordID, nil
}

