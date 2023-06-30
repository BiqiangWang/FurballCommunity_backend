package models

import (
	"FurballCommunity_backend/config/database"
	"strconv"
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

// HasMany 在Order模型中定义HasMany方法，表示一个blog拥有多个blogCmt
func (blog *Blog) HasMany() interface{} {
	return &[]BlogCmt{}
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

func GetBlogLike(blogID uint) (like uint, err error) {
	if err = database.DB.Where("blog_id = ?", blogID).Select("like").First(like).Error; err != nil {
		return like, err
	}
	return
}

func LikeBlog(blog *Blog) (err error) {
	err = database.DB.Model(&blog).Updates(map[string]interface{}{
		"like": blog.Like + 1,
	}).Error
	return
}

func AddToUserLikedList(user *User, blogID uint) (err error) {
	user.LikedBlog = user.LikedBlog + "," + strconv.Itoa(int(blogID))
	err = database.DB.Model(&user).Updates(map[string]interface{}{
		"liked_blog": user.LikedBlog,
	}).Error
	return
}

func UnLikeBlog(blog *Blog) (err error) {
	err = database.DB.Model(&blog).Updates(map[string]interface{}{
		"like": blog.Like - 1,
	}).Error
	return
}

func DeleteBlog(blogID uint) (err error) {
	err = database.DB.Delete(&Blog{}, blogID).Error
	return
}
