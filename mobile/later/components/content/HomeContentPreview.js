import React, { useState } from 'react';
import ContentPreview from './ContentPreview';
import { ContentBottomSheet } from '../modals';

function HomeContentPreview(props) {

    const [optionsActive, setOptionsActive] = useState(false)

    let content = props.content

    return (
        <>
            <ContentPreview
                content={content}
                linkActive={!optionsActive}
                onDotPress={() => {
                    setOptionsActive(true)
                }}
            />
            <ContentBottomSheet
                type='home'
                optionsActive={optionsActive}
                setOptionsActive={value => setOptionsActive(value)}
                {...props}
            />
        </>
    );
}

export default HomeContentPreview