import React from 'react';
import { View, ScrollView } from 'react-native';
import HomeContentPreview from './HomeContentPreview';
import SavedContentPreview from './SavedContentPreview';
import { Divider } from '../common'

function ContentGroup(props) {
    return (
        <ScrollView keyboardShouldPersistTaps='handled'>
            {
                props.content.map((contentObj, index) => (
                    <View key={index}>
                        {Preview(props, contentObj)}
                        {index < props.content.length - 1 ? <Divider key={props.index} /> : null}
                    </View>
                ))
            }
        </ScrollView>
    );
}

function Preview(props, content) {
    switch (props.type) {
        case 'home':
            return (
                <HomeContentPreview
                    content={content}
                    onForward={() => props.onForward(content)}
                    onSave={tag => props.onSave(content, tag)}
                />
            )
        case 'save':
            return (
                <SavedContentPreview
                    content={content}
                    onForward={() => props.onForward(content)}
                    onUpdateTag={tag => props.onUpdateTag(content, tag)}
                />
            )
    }
}

export default ContentGroup