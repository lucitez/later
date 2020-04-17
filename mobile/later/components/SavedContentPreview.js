import React, { useState } from 'react';
import ContentPreview from './ContentPreview';
import ContentBottomSheet from './ContentBottomSheet';

function SavedContentPreview(props) {
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
                type='save'
                optionsActive={optionsActive}
                setOptionsActive={value => setOptionsActive(value)}
                {...props}
            />
        </>
    );
}

export default SavedContentPreview