import React, { useState } from 'react';
import ContentPreview from './ContentPreview';
import ContentBottomSheet from './ContentBottomSheet';

function HomeContentPreview(props) {
    const [optionsActive, setOptionsActive] = useState(false)

    return (
        <>
            <ContentPreview
                linkActive={!optionsActive}
                onDotPress={() => setOptionsActive(true)}
                onTagPress={tag => props.navigation.navigate('Tag Screen', { tag: tag })}
                {...props}
            />
            <ContentBottomSheet
                type='archive'
                optionsActive={optionsActive}
                setOptionsActive={value => setOptionsActive(value)}
                {...props}
            />
        </>
    );
}

export default HomeContentPreview