package enums

// MediaTypeEnum is a media type classification.
type MediaTypeEnum string

const (
	MediaTypeEnumUnknown           MediaTypeEnum = "unknown"
	MediaTypeEnumImage             MediaTypeEnum = "image"
	MediaTypeEnumVideo             MediaTypeEnum = "video"
	MediaTypeEnumSound             MediaTypeEnum = "sound"
	MediaTypeEnumMediaWithDuration MediaTypeEnum = "mediaWithDuration"
	MediaTypeEnumMedia             MediaTypeEnum = "media"
)

// ConversionStatusEnum is a conversion/processing status.
type ConversionStatusEnum string

const (
	ConversionStatusEnumCompleted  ConversionStatusEnum = "completed"
	ConversionStatusEnumProcessing ConversionStatusEnum = "processing"
	ConversionStatusEnumFailed     ConversionStatusEnum = "failed"
)

// DocumentChangeType is a document change type in gallery context.
type DocumentChangeType string

const (
	DocumentChangeTypeCreate DocumentChangeType = "create"
	DocumentChangeTypeUpdate DocumentChangeType = "update"
	DocumentChangeTypeDelete DocumentChangeType = "delete"
)

// GalleryDropDownEnumeration is a dropdown menu action for gallery items.
type GalleryDropDownEnumeration string

const (
	GalleryDropDownEnumerationDownload GalleryDropDownEnumeration = "download"
	GalleryDropDownEnumerationDelete   GalleryDropDownEnumeration = "delete"
	GalleryDropDownEnumerationReport   GalleryDropDownEnumeration = "report"
	GalleryDropDownEnumerationEditTags GalleryDropDownEnumeration = "editTags"
	GalleryDropDownEnumerationViewInfo GalleryDropDownEnumeration = "viewInfo"
)

// ImageSizeEnum is an image size preset.
type ImageSizeEnum string

const (
	ImageSizeEnumOriginal ImageSizeEnum = "original"
	ImageSizeEnumMax200   ImageSizeEnum = "max200"
	ImageSizeEnumMax400   ImageSizeEnum = "max400"
)

// MediaCellType is a cell type in media grid.
type MediaCellType string

const (
	MediaCellTypeTagCell    MediaCellType = "tagCell"
	MediaCellTypeTaggedCell MediaCellType = "taggedCell"
	MediaCellTypeNonTagCell MediaCellType = "nonTagCell"
	MediaCellTypeEmptyCell  MediaCellType = "emptyCell"
)

// MediaSelectManyAction is a batch action on selected media.
type MediaSelectManyAction string

const (
	MediaSelectManyActionAddTags     MediaSelectManyAction = "addTags"
	MediaSelectManyActionDownload    MediaSelectManyAction = "download"
	MediaSelectManyActionDelete      MediaSelectManyAction = "delete"
	MediaSelectManyActionEditInfo    MediaSelectManyAction = "editInfo"
	MediaSelectManyActionRotateRight MediaSelectManyAction = "rotateRight"
)

// RotatingEnum is a rotation angle for images.
type RotatingEnum string

const (
	RotatingEnumRotating0   RotatingEnum = "rotating0"
	RotatingEnumRotating90  RotatingEnum = "rotating90"
	RotatingEnumRotating180 RotatingEnum = "rotating180"
	RotatingEnumRotating270 RotatingEnum = "rotating270"
)

// ThumbnailSizeEnum is a thumbnail size preset.
type ThumbnailSizeEnum string

const (
	ThumbnailSizeEnumXS   ThumbnailSizeEnum = "xS"
	ThumbnailSizeEnumS    ThumbnailSizeEnum = "s"
	ThumbnailSizeEnumM    ThumbnailSizeEnum = "m"
	ThumbnailSizeEnumL    ThumbnailSizeEnum = "l"
	ThumbnailSizeEnumFull ThumbnailSizeEnum = "full"
)
