package models

import (
	"FurballCommunity_backend/config/database"
	"log"
	"time"
)

type BlogCmt struct {
	BlogCmtID   uint      `gorm:"primary_key" json:"blog_cmt_id"`
	UserID      uint      `json:"user_id"`
	BlogID      uint      `json:"blog_id"`
	Content     string    `json:"content"`
	UserName    string    `json:"user_name"`
	Like        uint      `json:"like"`
	CommentTime time.Time `json:"comment_time" gorm:"default:CURRENT_TIMESTAMP"`
}

// BelongsTo 在Pet模型中定义BelongsTo方法，表示blogCmt属于一个blog
func (blogCmt *BlogCmt) BelongsTo() interface{} {
	return &Blog{}
}

func CreateBlogCmt(blogCmt *BlogCmt) (err error) {
	err = database.DB.Create(&blogCmt).Error
	return
}

func GetCmtListOfBlog(blogID uint) (blogCmt []*BlogCmt, err error) {
	log.Printf("GetCmtListOfBlog: blogID=%d\n", blogID)
	if err := database.DB.Where("blog_id = ?", blogID).Find(&blogCmt).Error; err != nil {
		return nil, err
	}
	return blogCmt, nil
}

func DeleteBlogCmt(blogCmtID uint) (err error) {
	err = database.DB.Delete(&BlogCmt{}, blogCmtID).Error
	return
}
