package persistence

import (
	"context"
	"online-script/internal/core/domain"

	"github.com/jackc/pgx/v5"
)

// RoleRepository implements RoleRepositorysitory interface using pgx.
type RoleRepository struct {
	db *pgx.Conn // or *pgx.ConnPool if you're using a connection pool
}

// NewRoleRepository creates a new instance of RoleRepository.
func NewRoleRepository(db *pgx.Conn) domain.RoleRepository {
	return &RoleRepository{db: db}
}

// ListRoles retrieves a list of roles from the database with pagination.
func (r *RoleRepository) ListRoles(ctx context.Context, params domain.ListRolesParams) ([]domain.SecRole, error) {
	rows, err := r.db.Query(ctx, "SELECT role_id, role_name FROM sec_role ORDER BY role_id LIMIT $1 OFFSET $2", params.Limit, params.Offset)
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
