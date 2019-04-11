package domain

type IGalleryService interface {

	//All gallerys are public. A user can follow or unfollow a gallery, without permission from the gallery service

	AddItem(gallery GalleryID, item Item, actor Actor) error
	UpdateItem(gallery GalleryID, item Item, actor Actor) error

	Get(gallery []GalleryID) ([]Gallery, error)

	AddContributor(gallery GalleryID, contributor UserID, actor Actor) error
	RemoveContributor(gallery GalleryID, contributor UserID, actor Actor) error
	ChangeOwner(gallery GalleryID, newOwner UserID, actor Actor) error

	DestroyGallery(gallery GalleryID, actor Actor) error
	CreateGallery(actor Actor) (Gallery, error)
}

//gateway to a user's main requests...
type IUserService interface {
	DebugSignIn(id UserID) (Actor, error)
	GetGalleries(actor Actor) ([]GalleryID, error)
	FollowGallery(gallery GalleryID, actor Actor) error
	UnFollowGallery(gallery GalleryID, actor Actor) error
}

//user flow...
//sign in - Auth Service  /  create new user...
//Getgallerys (ids) -User Service
//Foreach gallery, - gallery Service

//some core service that validates Actors...

//need to search for galleries somehow...
//gallerys that contain items...
//gallerys that contain tags...
//Meta-gallery service
//most popular tags, most popular items...
