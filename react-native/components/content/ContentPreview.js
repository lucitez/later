import React, { useState, useEffect } from 'react';
import { Image } from 'react-native';
import ContentPreviewNoImage from './ContentPreviewNoImage'
import ContentPreviewThumb from './ContentPreviewThumb';

function ContentPreview(props) {

    const [loading, setLoading] = useState(false)
    const [imageAR, setImageAR] = useState(null)

    useEffect(() => {
        if (props.content.imageUrl) {
            Image.getSize(
                props.content.imageUrl,
                (width, height) => {
                    setImageAR(width / height)
                    setLoading(false)
                },
                () => setLoading(false)
            )
        } else {
            setLoading(false)
        }
    })

    if (loading) return null

    if (imageAR) {
        return <ContentPreviewThumb {...props} imageAR={imageAR > 1.5 ? 1.5 : imageAR} />
    } else {
        return <ContentPreviewNoImage {...props} />
    }
}

export default ContentPreview