package service

const (
	UserBaseUrl   = "https://api.spotify.com/v1/me"
	SearchBaseUrl = "https://api.spotify.com/v1/search"
	PlayerBaseUrl = "https://api.spotify.com/v1/me/player"
	redirectURI   = "http://localhost:8080/callback"
	state         = "abc123"
	tokenUrl      = "https://accounts.spotify.com/api/token"
	scope         = "user-read-private" +
		"%20user-read-recently-played" +
		"%20user-read-playback-state" +
		"%20user-top-read" +
		"%20app-remote-control" +
		"%20user-modify-playback-state" +
		"%20user-read-currently-playing" +
		"%20user-follow-read" +
		"%20user-read-playback-position" +
		"%20playlist-read-private" +
		"%20user-read-email" +
		"%20user-read-private" +
		"%20user-library-read" +
		"%20playlist-read-collaborative" +
		"%20streaming" +
		"%20user-read-recently-played" +
		"%20user-library-read"
)
