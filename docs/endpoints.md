# TikTok API endpoint reference

54 unique GET routes from the published Postman collection. Configure `license_key` on the client. Query parameters below use the same names as the API; check Postman for required/optional rules.

| SDK method | GET path | Query parameters | Purpose |
| --- | --- | --- | --- |
| `GetNoWatermarkUrl` | `/getNoWatermarkUrl` | `video_url`, `region`, `cache_timeout` | Get non watermarked video url |
| `GetNoWatermarkUrlById` | `/getNoWatermarkUrlByID` | `video_id`, `region`, `cache_timeout` | Get non watermarked video url by video ID |
| `GetVideoById` | `/getVideoByID` | `video_id`, `region`, `cache_timeout` | Get video info by ID |
| `GetVideoByUrl` | `/getVideoByUrl` | `video_url`, `region`, `cache_timeout` | Get video info by URL |
| `GetUnshortenedUrl` | `/getUnshortenedUrl` | `video_url`, `cache_timeout` | Convert short TikTok video URL to full TikTok video URL |
| `GetCommentsByUrl` | `/getCommentsByUrl` | `video_url`, `region`, `cursor`, `count`, `cache_timeout` | Get comments list by video URL |
| `GetCommentsById` | `/getCommentsByID` | `video_id`, `region`, `cursor`, `count`, `cache_timeout` | Get comments list by video ID |
| `GetCommentsReplies` | `/getCommentsReplies` | `video_id`, `comment_id`, `cursor`, `count`, `cache_timeout` | Get comments replies |
| `SearchUser` | `/searchUser` | `keyword`, `cursor`, `count`, `region`, `follower_count`, `profile_type`, `other_pref`, `cache_timeout` | Search for user by keyword |
| `SearchVideo` | `/searchVideo` | `keyword`, `cursor`, `count`, `cache_timeout`, `date_posted`, `sort_by`, `region` | Search for video by keyword |
| `SearchMusic` | `/searchMusic` | `keyword`, `cursor`, `count`, `cache_timeout`, `filter_by`, `sort_type`, `region` | Search for music by keyword |
| `SearchLive` | `/searchLive` | `keyword`, `offset`, `count`, `cache_timeout`, `region` | Search for live by keyword |
| `SearchHashtag` | `/searchHashtag` | `keyword`, `cursor`, `count`, `region`, `cache_timeout` | Search for hashtag by keyword |
| `SearchLocation` | `/searchLocation` | `keyword`, `cursor`, `count`, `region`, `cache_timeout` | Search for location by keyword |
| `SearchTop` | `/searchTop` | `keyword`, `cursor`, `count`, `region`, `sort_type`, `publish_time`, `cache_timeout` | Top search by text query |
| `SearchInsights` | `/searchInsights` | `keyword`, `cursor`, `count`, `region`, `cache_timeout` | Search insights by text query |
| `SearchCommentSuggestedUsers` | `/searchCommentSuggestedUsers` | `keyword`, `count`, `region`, `cache_timeout` | Search comment suggested users |
| `GetUserInfo` | `/getUserInfo` | `query`, `region`, `cache_timeout` | Get user info by profile URL or nickname |
| `GetUserInfoById` | `/getUserInfoByID` | `user_id`, `region`, `cache_timeout` | Get user info by profile ID |
| `GetBasicUserInfo` | `/getBasicUserInfo` | `query`, `region`, `cache_timeout` | Get user info by username |
| `GetUserFeed` | `/getUserFeed` | `query`, `max_cursor`, `count`, `region`, `with_pinned_posts`, `sort_type`, `cache_timeout` | Get user feed by profile URL or nickname |
| `GetUserFeedById` | `/getUserFeedByID` | `user_id`, `max_cursor`, `count`, `region`, `with_pinned_posts`, `sort_type`, `cache_timeout` | Get user feed by profile ID |
| `GetUserId` | `/getUserID` | `query`, `region`, `cache_timeout` | Get user ID by profile URL or nickname |
| `GetUsernameById` | `/getUsernameByID` | `user_id`, `region`, `cache_timeout` | Get username by TikTok user ID |
| `GetTwitterUsername` | `/getTwitterUsername` | `query`, `region`, `cache_timeout` | Get Twitter username by profile URL, nickname or TikTok user ID |
| `GetInstagramUsername` | `/getInstagramUsername` | `query`, `region`, `cache_timeout` | Get Instagram username by profile URL, nickname or TikTok user ID |
| `GetYoutubeChannel` | `/getYoutubeChannel` | `query`, `region`, `cache_timeout` | Get YouTube channel info by profile URL, nickname or TikTok user ID |
| `GetFollowers` | `/getFollowers` | `query`, `count`, `min_time`, `cache_timeout` | Get user followers by profile URL or nickname |
| `GetFollowersById` | `/getFollowersByID` | `user_id`, `count`, `min_time`, `cache_timeout` | Get user followers by profile ID |
| `GetFollowing` | `/getFollowing` | `query`, `count`, `min_time`, `cache_timeout` | Get user following by profile URL or nickname |
| `GetFollowingById` | `/getFollowingByID` | `user_id`, `count`, `min_time`, `cache_timeout` | Get user following by profile ID |
| `GetLiveEventInfoById` | `/getLiveEventInfoByID` | `event_id`, `cache_timeout` | Get live event info by ID |
| `GetHashtagInfo` | `/getHashtagInfo` | `query`, `region`, `cache_timeout` | Get hashtag info by URL or tag name |
| `GetHashtagFeed` | `/getHashtagFeed` | `query`, `cursor`, `count`, `region`, `cache_timeout` | Get hashtag feed by URL or tag name |
| `GetHashtagFeedById` | `/getHashtagFeedByID` | `id`, `cursor`, `count`, `region`, `cache_timeout` | Get hashtag feed by ID |
| `GetLocationInfoById` | `/getLocationInfoByID` | `location_id`, `cache_timeout` | Get location info by ID |
| `GetLocationFeedById` | `/getLocationFeedByID` | `location_id`, `count`, `cursor`, `region`, `cache_timeout` | Get videos by location ID |
| `GetMusicInfo` | `/getMusicInfo` | `query`, `region`, `cache_timeout` | Get music info by URL or ID |
| `GetMusicFeed` | `/getMusicFeed` | `query`, `cursor`, `count`, `cache_timeout`, `region` | Get music feed by URL or ID |
| `GetMusicCharts` | `/getMusicCharts` | `chart_type`, `cache_timeout`, `region` | Get music charts |
| `GetTrendingCategoriesList` | `/getTrendingCategoriesList` | `region`, `offset`, `count`, `cache_timeout` | Get trending categories list |
| `GetTrendingFeed` | `/getTrendingFeed` | `region`, `pull_type`, `cache_timeout` | Get trending feed |
| `GetStickersById` | `/getStickersByID` | `sticker_ids`, `cache_timeout` | Get effect information by IDs |
| `GetStickerById` | `/getStickerByID` | `sticker_id`, `cache_timeout` | Get effect information by ID |
| `GetStickerFeedById` | `/getStickerFeedByID` | `sticker_id`, `count`, `cursor`, `region`, `cache_timeout` | Get videos by effect ID |
| `GetUserStories` | `/getUserStories` | `user_id`, `count`, `min_cursor`, `region`, `cache_timeout` | Get user stories |
| `GetStoryInfoById` | `/getStoryInfoByID` | `story_id`, `region`, `cache_timeout` | Get story info by ID |
| `GetUserPlaylists` | `/getUserPlaylists` | `user_id`, `cursor`, `cache_timeout` | Get user playlists |
| `GetUserPlaylistById` | `/getUserPlaylistByID` | `user_id`, `playlist_id`, `cache_timeout` | Get user playlist by ID |
| `GetVideoByPlaylistId` | `/getVideoByPlaylistID` | `user_id`, `playlist_id`, `max_cursor`, `cache_timeout` | Get video by playlist ID |
| `GetAvailableRankingList` | `/getAvailableRankingList` | `room_id`, `user_id`, `cache_timeout` | Get available ranking list |
| `GetRankingInfo` | `/getRankingInfo` | `rank_type`, `room_id`, `user_id`, `cache_timeout` | Get ranking info |
| `GetLiveInfo` | `/getLiveInfo` | `room_id`, `user_id`, `cache_timeout` | Get live info |
| `GetLiveFeed` | `/getLiveFeed` | `draw_room_id`, `max_time`, `region`, `cache_timeout` | Get live feed |
