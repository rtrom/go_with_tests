package blogposts_test

import (
	blogposts "github.com/quii/learn-go-with-tests/reading-files"
)

func TestNewBlogPosts(t *Testing.T) {
	fs := ftest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted d% posts", len(posts), len(fs))
	}
}
