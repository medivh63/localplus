package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
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

func (r *SQLiteRepository) GetQuestionByID(id string) (*models.Question, error) {
	var q models.Question
	if err := r.db.Where("question_id = ?", id).First(&q).Error; err != nil {
		return nil, err
	}
	return &q, nil
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

// GetAllQuestionIDs 从数据获取所有questionID
func (r *SQLiteRepository) GetAllQuestionIDs() ([]string, error) {
	var ids []string
	if err := r.db.Model(&models.Question{}).Select("question_id").Find(&ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}

// GetAnsweredQuestionIDs 根据quizID查询数据库已经答过的题目ID
func (r *SQLiteRepository) GetAnsweredQuestionIDs(quizID string) ([]string, error) {
	var ids []string
	if err := r.db.Model(&models.Quiz{}).Where("quiz_id = ?", quizID).Select("question_id").Find(&ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}
