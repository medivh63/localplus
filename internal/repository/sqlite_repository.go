package repository

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
	"localplus/internal/models"
	"log"
	"os"
	"time"
)

type SQLiteRepository struct {
	db *gorm.DB
}

func NewSQLiteDB(dbPath string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
}

func NewSQLiteRepository(dbPath string) (*SQLiteRepository, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Info,   // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 禁用彩色打印
		},
	)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: newLogger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, err
	}

	// 启用外键约束
	if err := db.Exec("PRAGMA foreign_keys = ON").Error; err != nil {
		return nil, err
	}

	return &SQLiteRepository{db: db}, nil
}

func (r *SQLiteRepository) GetRandomQuestion() (*models.Question, error) {
	var q models.Question
	if err := r.db.Order("RANDOM()").First(&q).Error; err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *SQLiteRepository) SaveQuizAnswer(quiz *models.Quiz) error {
	return r.db.Create(quiz).Error
}

func (r *SQLiteRepository) GetQuizResult(quizID string) ([]models.Quiz, error) {
	var results []models.Quiz
	if err := r.db.Where("quiz_id = ?", quizID).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
