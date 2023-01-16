package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//userEndPoints
	rt.router.GET("/users/:user_id", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/users/:user_id/username", rt.wrap(rt.updateUsername))
	rt.router.GET("/result", rt.wrap(rt.fetchUsername))
	rt.router.GET("/users/:user_id/info", rt.wrap(rt.getUserInfo))

	//mediaEndPoints
	rt.router.POST("/media", rt.wrap(rt.postMedia))
	rt.router.GET("/media/:post_id/info", rt.wrap(rt.getMediaMetadata))
	rt.router.GET("/media/:post_id", rt.wrap(rt.getMedia))
	rt.router.DELETE("/media/:post_id", rt.wrap(rt.deleteMedia))

	//FollowersEndPoints
	rt.router.PUT("/users/:user_id/follow/:followed_id", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:user_id/follow/:followed_id", rt.wrap(rt.unfollowUser))

	//BansEndPoints
	rt.router.PUT("/users/:user_id/banned/:ban_id", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:user_id/banned/:ban_id", rt.wrap(rt.unbanUser))

	//LikesEndPoints
	rt.router.PUT("/:post_id/likes/:user_id", rt.wrap(rt.likeMedia))
	rt.router.DELETE("/:post_id/likes/:user_id", rt.wrap(rt.unlikeMedia))
	// TODO change likes/like into likes/:user_id, also consider removing user id from path

	//Comments
	rt.router.POST("/:post_id/comments/comment/:user_id", rt.wrap(rt.addComment))
	rt.router.DELETE("/:post_id/comments/delete/:comment_id", rt.wrap(rt.removeComment))

	//Stream
	rt.router.GET("/feed/:user_id", rt.wrap(rt.getStream))

	//Login
	rt.router.POST("/session", rt.wrap(rt.signIn))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
