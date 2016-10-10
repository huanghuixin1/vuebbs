package controllers

type commentController struct {

}

var CommentController = commentController{}

//添加评论
func (this *commentController) Add(aid int,detail string) bool {
	return false;
}