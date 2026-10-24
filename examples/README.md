# TikTok go request examples

Every file below is a complete single-request CLI program with environment-based authentication, all documented query options, timeout configuration, error handling and full response output. Sample defaults come from Postman; replace account names, IDs and URLs for your use case. `--help` requires no API key and makes no network request.

Set `JSONSCRAPER_TIKTOK_LICENSE_KEY` and install the SDK as described in the main README. Each execution makes one live request and may consume quota. None of these programs loops through a feed automatically.

## Get non watermarked video url

[Open the complete example](endpoints/get_no_watermark_url.go) · `GET /getNoWatermarkUrl`

```sh
go run examples/endpoints/get_no_watermark_url.go --help
```

## Get non watermarked video url by video ID

[Open the complete example](endpoints/get_no_watermark_url_by_id.go) · `GET /getNoWatermarkUrlByID`

```sh
go run examples/endpoints/get_no_watermark_url_by_id.go --help
```

## Get video info by ID

[Open the complete example](endpoints/get_video_by_id.go) · `GET /getVideoByID`

```sh
go run examples/endpoints/get_video_by_id.go --help
```

## Get video info by URL

[Open the complete example](endpoints/get_video_by_url.go) · `GET /getVideoByUrl`

```sh
go run examples/endpoints/get_video_by_url.go --help
```

## Convert short TikTok video URL to full TikTok video URL

[Open the complete example](endpoints/get_unshortened_url.go) · `GET /getUnshortenedUrl`

```sh
go run examples/endpoints/get_unshortened_url.go --help
```

## Get comments list by video URL

[Open the complete example](endpoints/get_comments_by_url.go) · `GET /getCommentsByUrl`

```sh
go run examples/endpoints/get_comments_by_url.go --help
```

## Get comments list by video ID

[Open the complete example](endpoints/get_comments_by_id.go) · `GET /getCommentsByID`

```sh
go run examples/endpoints/get_comments_by_id.go --help
```

## Get comments replies

[Open the complete example](endpoints/get_comments_replies.go) · `GET /getCommentsReplies`

```sh
go run examples/endpoints/get_comments_replies.go --help
```

## Search for user by keyword

[Open the complete example](endpoints/search_user.go) · `GET /searchUser`

```sh
go run examples/endpoints/search_user.go --help
```

## Search for video by keyword

[Open the complete example](endpoints/search_video.go) · `GET /searchVideo`

```sh
go run examples/endpoints/search_video.go --help
```

## Search for music by keyword

[Open the complete example](endpoints/search_music.go) · `GET /searchMusic`

```sh
go run examples/endpoints/search_music.go --help
```

## Search for live by keyword

[Open the complete example](endpoints/search_live.go) · `GET /searchLive`

```sh
go run examples/endpoints/search_live.go --help
```

## Search for hashtag by keyword

[Open the complete example](endpoints/search_hashtag.go) · `GET /searchHashtag`

```sh
go run examples/endpoints/search_hashtag.go --help
```

## Search for location by keyword

[Open the complete example](endpoints/search_location.go) · `GET /searchLocation`

```sh
go run examples/endpoints/search_location.go --help
```

## Top search by text query

[Open the complete example](endpoints/search_top.go) · `GET /searchTop`

```sh
go run examples/endpoints/search_top.go --help
```

## Search insights by text query

[Open the complete example](endpoints/search_insights.go) · `GET /searchInsights`

```sh
go run examples/endpoints/search_insights.go --help
```

## Search comment suggested users

[Open the complete example](endpoints/search_comment_suggested_users.go) · `GET /searchCommentSuggestedUsers`

```sh
go run examples/endpoints/search_comment_suggested_users.go --help
```

## Get user info by profile URL or nickname

[Open the complete example](endpoints/get_user_info.go) · `GET /getUserInfo`

```sh
go run examples/endpoints/get_user_info.go --help
```

## Get user info by profile ID

[Open the complete example](endpoints/get_user_info_by_id.go) · `GET /getUserInfoByID`

```sh
go run examples/endpoints/get_user_info_by_id.go --help
```

## Get user info by username

[Open the complete example](endpoints/get_basic_user_info.go) · `GET /getBasicUserInfo`

```sh
go run examples/endpoints/get_basic_user_info.go --help
```

## Get user feed by profile URL or nickname

[Open the complete example](endpoints/get_user_feed.go) · `GET /getUserFeed`

```sh
go run examples/endpoints/get_user_feed.go --help
```

## Get user feed by profile ID

[Open the complete example](endpoints/get_user_feed_by_id.go) · `GET /getUserFeedByID`

```sh
go run examples/endpoints/get_user_feed_by_id.go --help
```

## Get user ID by profile URL or nickname

[Open the complete example](endpoints/get_user_id.go) · `GET /getUserID`

```sh
go run examples/endpoints/get_user_id.go --help
```

## Get username by TikTok user ID

[Open the complete example](endpoints/get_username_by_id.go) · `GET /getUsernameByID`

```sh
go run examples/endpoints/get_username_by_id.go --help
```

## Get Twitter username by profile URL, nickname or TikTok user ID

[Open the complete example](endpoints/get_twitter_username.go) · `GET /getTwitterUsername`

```sh
go run examples/endpoints/get_twitter_username.go --help
```

## Get Instagram username by profile URL, nickname or TikTok user ID

[Open the complete example](endpoints/get_instagram_username.go) · `GET /getInstagramUsername`

```sh
go run examples/endpoints/get_instagram_username.go --help
```

## Get YouTube channel info by profile URL, nickname or TikTok user ID

[Open the complete example](endpoints/get_youtube_channel.go) · `GET /getYoutubeChannel`

```sh
go run examples/endpoints/get_youtube_channel.go --help
```

## Get user followers by profile URL or nickname

[Open the complete example](endpoints/get_followers.go) · `GET /getFollowers`

```sh
go run examples/endpoints/get_followers.go --help
```

## Get user followers by profile ID

[Open the complete example](endpoints/get_followers_by_id.go) · `GET /getFollowersByID`

```sh
go run examples/endpoints/get_followers_by_id.go --help
```

## Get user following by profile URL or nickname

[Open the complete example](endpoints/get_following.go) · `GET /getFollowing`

```sh
go run examples/endpoints/get_following.go --help
```

## Get user following by profile ID

[Open the complete example](endpoints/get_following_by_id.go) · `GET /getFollowingByID`

```sh
go run examples/endpoints/get_following_by_id.go --help
```

## Get live event info by ID

[Open the complete example](endpoints/get_live_event_info_by_id.go) · `GET /getLiveEventInfoByID`

```sh
go run examples/endpoints/get_live_event_info_by_id.go --help
```

## Get hashtag info by URL or tag name

[Open the complete example](endpoints/get_hashtag_info.go) · `GET /getHashtagInfo`

```sh
go run examples/endpoints/get_hashtag_info.go --help
```

## Get hashtag feed by URL or tag name

[Open the complete example](endpoints/get_hashtag_feed.go) · `GET /getHashtagFeed`

```sh
go run examples/endpoints/get_hashtag_feed.go --help
```

## Get hashtag feed by ID

[Open the complete example](endpoints/get_hashtag_feed_by_id.go) · `GET /getHashtagFeedByID`

```sh
go run examples/endpoints/get_hashtag_feed_by_id.go --help
```

## Get location info by ID

[Open the complete example](endpoints/get_location_info_by_id.go) · `GET /getLocationInfoByID`

```sh
go run examples/endpoints/get_location_info_by_id.go --help
```

## Get videos by location ID

[Open the complete example](endpoints/get_location_feed_by_id.go) · `GET /getLocationFeedByID`

```sh
go run examples/endpoints/get_location_feed_by_id.go --help
```

## Get music info by URL or ID

[Open the complete example](endpoints/get_music_info.go) · `GET /getMusicInfo`

```sh
go run examples/endpoints/get_music_info.go --help
```

## Get music feed by URL or ID

[Open the complete example](endpoints/get_music_feed.go) · `GET /getMusicFeed`

```sh
go run examples/endpoints/get_music_feed.go --help
```

## Get music charts

[Open the complete example](endpoints/get_music_charts.go) · `GET /getMusicCharts`

```sh
go run examples/endpoints/get_music_charts.go --help
```

## Get trending categories list

[Open the complete example](endpoints/get_trending_categories_list.go) · `GET /getTrendingCategoriesList`

```sh
go run examples/endpoints/get_trending_categories_list.go --help
```

## Get trending feed

[Open the complete example](endpoints/get_trending_feed.go) · `GET /getTrendingFeed`

```sh
go run examples/endpoints/get_trending_feed.go --help
```

## Get effect information by IDs

[Open the complete example](endpoints/get_stickers_by_id.go) · `GET /getStickersByID`

```sh
go run examples/endpoints/get_stickers_by_id.go --help
```

## Get effect information by ID

[Open the complete example](endpoints/get_sticker_by_id.go) · `GET /getStickerByID`

```sh
go run examples/endpoints/get_sticker_by_id.go --help
```

## Get videos by effect ID

[Open the complete example](endpoints/get_sticker_feed_by_id.go) · `GET /getStickerFeedByID`

```sh
go run examples/endpoints/get_sticker_feed_by_id.go --help
```

## Get user stories

[Open the complete example](endpoints/get_user_stories.go) · `GET /getUserStories`

```sh
go run examples/endpoints/get_user_stories.go --help
```

## Get story info by ID

[Open the complete example](endpoints/get_story_info_by_id.go) · `GET /getStoryInfoByID`

```sh
go run examples/endpoints/get_story_info_by_id.go --help
```

## Get user playlists

[Open the complete example](endpoints/get_user_playlists.go) · `GET /getUserPlaylists`

```sh
go run examples/endpoints/get_user_playlists.go --help
```

## Get user playlist by ID

[Open the complete example](endpoints/get_user_playlist_by_id.go) · `GET /getUserPlaylistByID`

```sh
go run examples/endpoints/get_user_playlist_by_id.go --help
```

## Get video by playlist ID

[Open the complete example](endpoints/get_video_by_playlist_id.go) · `GET /getVideoByPlaylistID`

```sh
go run examples/endpoints/get_video_by_playlist_id.go --help
```

## Get available ranking list

[Open the complete example](endpoints/get_available_ranking_list.go) · `GET /getAvailableRankingList`

```sh
go run examples/endpoints/get_available_ranking_list.go --help
```

## Get ranking info

[Open the complete example](endpoints/get_ranking_info.go) · `GET /getRankingInfo`

```sh
go run examples/endpoints/get_ranking_info.go --help
```

## Get live info

[Open the complete example](endpoints/get_live_info.go) · `GET /getLiveInfo`

```sh
go run examples/endpoints/get_live_info.go --help
```

## Get live feed

[Open the complete example](endpoints/get_live_feed.go) · `GET /getLiveFeed`

```sh
go run examples/endpoints/get_live_feed.go --help
```

