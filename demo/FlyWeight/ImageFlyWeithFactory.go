package FlyWeight

var imgflyfactory * ImageFlyWeightFactory

type ImageFlyWeightFactory struct{
	maps map[string]*ImageFlyWeight
}

func GetImageFlyWeightFactory() * ImageFlyWeightFactory {
	if imgflyfactory == nil {
		imgflyfactory = &ImageFlyWeightFactory{make(map[string]*ImageFlyWeight)}
	}
	return imgflyfactory
}

func (imf *ImageFlyWeightFactory) Get(filename string) *ImageFlyWeight {
	image := imf.maps[filename]
	if image == nil {
		image = NewImageFlyWeight(filename)
		imf.maps[filename] = image
	}
	return image
}