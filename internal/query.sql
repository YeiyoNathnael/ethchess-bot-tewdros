
-- name: CreateUser :exec
INSERT INTO users (
    telegram_id,
    lichess_username,
    is_whitelisted,
    created_at
) VALUES (?, ?, 0, ?);
-- name: GetUser :one
SELECT * FROM users
WHERE telegram_id = ?;
-- name: UserExists :one
SELECT EXISTS(
    SELECT 1 FROM users WHERE telegram_id = ?
);
-- name: IsPresident :one
SELECT EXISTS(
    SELECT 1 FROM presidents WHERE telegram_id = ?
);
-- name: WhitelistUser :exec
UPDATE users
SET is_whitelisted = 1,
    updated_at = ?
WHERE telegram_id = ?;
-- name: RemoveWhitelist :exec
UPDATE users
SET is_whitelisted = 0,
    updated_at = ?
WHERE telegram_id = ?;
-- name: CreateOrganizer :exec
INSERT INTO organizers (
    telegram_id,
    name,
    logo_url,
    created_at
)
SELECT users.telegram_id, ?, ?, ?
FROM users
WHERE users.telegram_id = ?
AND is_whitelisted = 1;
-- name: IsOrganizer :one
SELECT EXISTS(
    SELECT 1 FROM organizers WHERE telegram_id = ?
);
-- name: CreateDraft :exec
INSERT INTO tournament_drafts (
    draft_uuid,
    tournament_class,
    requested_by,
    name,
    format,
    time_control,
    start_datetime,
    end_datetime,
    timezone,
    capacity,
    visibility,
    prizes,
    organizer_telegram_id,
    rationale,
    status,
    created_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 'draft', ?);
-- name: GetDraft :one
SELECT * FROM tournament_drafts
WHERE draft_uuid = ?;
-- name: SubmitDraft :exec
UPDATE tournament_drafts
SET status = 'submitted',
    updated_at = ?
WHERE draft_uuid = ?
AND requested_by = ?;
-- name: ApproveDraft :exec
UPDATE tournament_drafts
SET status = 'approved',
    updated_at = ?
WHERE draft_uuid = ?
AND status = 'submitted';
-- name: PublishDraft :exec
UPDATE tournament_drafts
SET status = 'published',
    updated_at = ?
WHERE draft_uuid = ?
AND status = 'approved';
-- name: InsertAuditLog :exec
INSERT INTO audit_log (
    entity_type,
    entity_id,
    action,
    performed_by,
    created_at
) VALUES (?, ?, ?, ?, ?);
