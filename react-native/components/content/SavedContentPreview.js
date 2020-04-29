import React, { useState } from 'react';
import ContentPreview from './ContentPreview';
import { ContentBottomSheet } from '../modals';

function SavedContentPreview(props) {
    const [optionsActive, setOptionsActive] = useState(false)

    return (
        <>
            <ContentPreview
                linkActive={!optionsActive}
                onDotPress={() => setOptionsActive(true)}
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