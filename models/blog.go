package models

import (
	"FurballCommunity_backend/config/database"
	"time"
)

type Blog struct {
	BlogID      uint      `gorm:"primary_key" json:"blog_id"`
	UserID      uint      `json:"user_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Like        uint      `json:"like"`
	PublishTime time.Time `json:"publish_time" gorm:"default:CURRENT_TIMESTAMP"`
	User        User      `gorm:"foreign_key:UserID"`
	//BannerList  []string  `json:"banner_list"`
}

// BelongsTo 在Blog模型中定义BelongsTo方法，表示blog属于一个user
func (blog *Blog) BelongsTo() interface{} {
	return &User{}
}

func CreateBlog(blog *Blog) (err error) {
	err = database.DB.Create(&blog).Error
	return
}

func GetBlogList() (blogList []*Blog, err error) {
	if err := database.DB.Select("blog_id", "user_id", "title", "publish_time").Find(&blogList).Error; err != nil {
		return nil, err
	}
	return blogList, nil
}

func GetBlogListOfUser(userID uint) (blogList []*Blog, err error) {
	if err := database.DB.Where("user_id = ?", userID).Find(&blogList).Error; err != nil {
		return nil, err
	}
	return blogList, nil
}

func GetBlogInfo(blogID uint) (blog *Blog, err error) {
	blog = new(Blog)
	if err = database.DB.Where("blog_id = ?", blogID).First(blog).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateBlog(blog *Blog) (err error) {
	err = database.DB.Model(&blog).Updates(map[string]interface{}{
		"content": blog.Content,
		"Title":   blog.Title,
	}).Error
	return
}
