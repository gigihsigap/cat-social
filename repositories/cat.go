package repository

import (
	model "cat-social/models"
	"cat-social/models/dto/response"
	"cat-social/models/enum"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

type CatRepository interface {
	FindAll(filterParams map[string]interface{}) ([]response.CatResponse, error)
	FindByUserID(i int) (model.Cat, error)
	FindByID(catID string) (model.Cat, error)
	FindByIDAndUserID(catID string, userID int) (model.Cat, error)
	Create(cat model.Cat) (response.CreateCatResponse, error)
	Update(cat model.Cat) (model.Cat, error)
	Delete(catID string, userID int) error
	UpdateHasMatch(id int, isHasMatch bool) (int, error)
}
type catRepository struct {
	db *pgx.Conn
}

func NewCatRepository(db *pgx.Conn) *catRepository {
	return &catRepository{db}
}

func (r *catRepository) FindAll(filterParams map[string]interface{}) ([]response.CatResponse, error) {
	query := "SELECT id, name, race, sex, age_in_month, description, image_urls FROM cats WHERE 1=1"
	var args []interface{}
	argIndex := 1

	if catID, ok := filterParams["id"].(string); ok && catID != "" {
		query += fmt.Sprintf(" AND id = $%d", argIndex)
		args = append(args, catID)
		argIndex++
	}

	if race, ok := filterParams["race"].(enum.Race); ok && race != "" {
		query += fmt.Sprintf(" AND race = '%s'", race)
	}

	if sexStr, ok := filterParams["sex"].(string); ok && sexStr != "" {
		sex := enum.Sex(sexStr)
		query += fmt.Sprintf(" AND sex = '%s'", sex)
	}

	if hasMatched, ok := filterParams["hasMatched"].(bool); ok {
		query += fmt.Sprintf(" AND has_matched = %t", hasMatched)
	}

	if ageInMonth, ok := filterParams["ageInMonth"].(string); ok && ageInMonth != "" {
		var comparison string
		var age int
		if strings.HasPrefix(ageInMonth, ">") {
			comparison = ">"
			age, _ = strconv.Atoi(strings.TrimPrefix(ageInMonth, ">"))
		} else if strings.HasPrefix(ageInMonth, "<") {
			comparison = "<"
			age, _ = strconv.Atoi(strings.TrimPrefix(ageInMonth, "<"))
		} else if strings.HasPrefix(ageInMonth, "=") {
			comparison = "="
			age, _ = strconv.Atoi(strings.TrimPrefix(ageInMonth, "="))
		} else {
			return nil, errors.New("Invalid comparison operator")
		}
		query += fmt.Sprintf(" AND age_in_month %s %d", comparison, age)
	}

	if owned, ok := filterParams["owned"].(bool); ok {
		userId := filterParams["userID"]
		if owned {
			query += fmt.Sprintf(" AND user_id = %d", userId)
		} else {
			query += fmt.Sprintf(" AND user_id != %d", userId)
		}
	}

	if search, ok := filterParams["search"].(string); ok && search != "" {
		query += fmt.Sprintf(" AND name ILIKE '%%%s%%'", search)
	}

	if limit, ok := filterParams["limit"].(int); ok && limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, limit)
		argIndex++
	}

	if offset, ok := filterParams["offset"].(int); ok && offset >= 0 {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, offset)
		argIndex++
	}
	fmt.Println(query)
	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []response.CatResponse
	for rows.Next() {
		var cat model.Cat
		err := rows.Scan(&cat.ID, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.ImageUrls)
		if err != nil {
			return nil, err
		}
		catResponse := response.CatResponse{
			ID:          cat.ID,
			Name:        cat.Name,
			Race:        cat.Race,
			Sex:         cat.Sex,
			AgeInMonth:  cat.AgeInMonth,
			ImageURLs:   cat.ImageUrls,
			Description: cat.Description,
			HasMatched:  cat.HasMatched,
		}
		cats = append(cats, catResponse)
	}
	return cats, nil
}

func (r *catRepository) FindByID(catID string) (model.Cat, error) {
	var cat model.Cat
	err := r.db.QueryRow(context.Background(), "SELECT id, name, race, sex, age_in_month, has_matched, description, image_urls FROM cats WHERE id = $1", catID).Scan(&cat.ID, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.HasMatched, &cat.Description, &cat.ImageUrls)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Cat{}, nil // Kucing tidak ditemukan, tidak ada error
		}
		return model.Cat{}, err // Error lainnya
	}
	return cat, nil
}

func (r *catRepository) FindByUserID(i int) (model.Cat, error) {
	var cat model.Cat
	err := r.db.QueryRow(context.Background(), "SELECT id, name, race, sex, age_in_month, description, image_urls FROM cats WHERE user_id = $1", i).Scan(&cat.ID, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.ImageUrls)
	fmt.Println(err)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Cat{}, nil // Kucing tidak ditemukan, tidak ada error
		}
		return model.Cat{}, err // Error lainnya
	}
	return cat, nil
}

func (r *catRepository) FindByIDAndUserID(catID string, userID int) (model.Cat, error) {
	var cat model.Cat
	err := r.db.QueryRow(context.Background(), "SELECT id, name, race, sex, age_in_month, description, image_urls FROM cats WHERE user_id = $1 and id = $2", userID, catID).Scan(&cat.ID, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.ImageUrls)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Cat{}, nil // Kucing tidak ditemukan, tidak ada error
		}
		return model.Cat{}, err // Error lainnya
	}
	return cat, nil
}

func (r *catRepository) Create(cat model.Cat) (response.CreateCatResponse, error) {
	var id string
	var createdAt time.Time
	err := r.db.QueryRow(context.Background(), "INSERT INTO cats (name, race, sex, age_in_month, description, image_urls, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at", cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description, cat.ImageUrls, cat.UserID).Scan(&id, &createdAt)
	if err != nil {
		return response.CreateCatResponse{}, err
	}

	// Konversi waktu pembuatan ke format ISO 8601
	createdAtISO8601 := createdAt.Format(time.RFC3339)

	// Buat respons yang akan dikirimkan kembali
	response := response.CreateCatResponse{
		ID:        id,
		CreatedAt: createdAtISO8601,
	}

	return response, nil
}

func (r *catRepository) Update(cat model.Cat) (model.Cat, error) {
	_, err := r.db.Exec(context.Background(), "UPDATE cats SET name = $1, race = $2, sex = $3, age_in_month = $4, description = $5, image_urls = $6 WHERE id = $7", cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description, cat.ImageUrls, cat.ID)
	if err != nil {
		return model.Cat{}, err
	}
	fmt.Println("cat updeted")
	return cat, nil
}

func (r *catRepository) UpdateHasMatch(id int, isHasMatch bool) (int, error) {
	_, err := r.db.Exec(context.Background(), "UPDATE cats SET has_matched = $1 WHERE id = $2", isHasMatch, id)
	if err != nil {
		return id, err
	}
	fmt.Println("cat updated")
	return id, nil
}

func (r *catRepository) Delete(catID string, userID int) error {
	_, err := r.db.Exec(context.Background(), "DELETE FROM cats WHERE id = $1 and user_id = $2", catID, userID)
	if err != nil {
		return err
	}
	return nil
}
