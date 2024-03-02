package persistence

import (
	"context"
	"online-script/internal/core/domain"

	"github.com/jackc/pgx/v5"
)

// RoleRepository implements RoleRepository interface using pgx.
type RoleRepository struct {
	db *pgx.Conn // or *pgx.ConnPool if you're using a connection pool
}

// NewRoleRepository creates a new instance of RoleRepository.
func NewRoleRepository(db *pgx.Conn) domain.RoleRepository {
	return &RoleRepository{db: db}
}

// ListRoles retrieves a list of roles from the database with pagination.
// The roles data are stored via demo data, and the access to those is read only.
// For a role record to be modified it needs to be accessed through db.
func (r *RoleRepository) ListRoles(ctx context.Context) ([]domain.SecRole, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM sec_role ORDER BY role_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []domain.SecRole
	for rows.Next() {
		var role domain.SecRole
		err := rows.Scan(&role.RoleID, &role.RoleName)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}
