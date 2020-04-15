import React, { useState } from 'react';
import ContentPreview from './ContentPreview';
import ContentBottomSheet from './ContentBottomSheet';

function HomeContentPreview(props) {

    const [optionsActive, setOptionsActive] = useState(false)

    let content = props.content

    return (
        <>
            <ContentPreview
                content={content}
                linkActive={!optionsActive}
                onDotPress={() => setOptionsActive(true)}
                onTagPress={tag => props.navigation.navigate('Tag Screen', { tag: tag })}
            />
            <ContentBottomSheet
                type='archive'
                optionsActive={optionsActive}
                setOptionsActive={value => setOptionsActive(value)}
                onForward={() => props.onForward()}
            />
        </>
    );
}

export default HomeContentPreview