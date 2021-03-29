# Image Validator

This service accepts requests to validate images by URL. It sends the given image to an AI vision service to ensure that the image likely matches some criteria. for example, given an image of a dog, and a criteria that it's 90% certain to actually be a dog, this service returns whether or not the image matches that criteria.

Additionally, this service checks the backing AI system to ensure that the given image does _not_ have any adult or "racy" content. See [Azure content moderator](https://azure.microsoft.com/en-us/services/cognitive-services/content-moderator/) for more information on that.

## Things to Think About

- Rate limiting / backpressure
- A cache to quickly determine whether an image was already checked and, if so, if it's valid or invalid
