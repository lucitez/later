import React from 'react';
import { View, ScrollView } from 'react-native';
import ContentPreview from './ContentPreview'
import Divider from './Divider'

function ContentGroup(props) {

    return (
        <ScrollView>
            {
                props.content.map((contentObj, index) => (
                    <View key={index}>
                        <ContentPreview content={contentObj} key={contentObj.id} />
                        <RenderDivider shouldRender={index != props.content.length - 1} index={index} />
                    </View>
                ))
            }
        </ScrollView>
    );
}

function RenderDivider(props) {
    if (props.shouldRender) {
        return <Divider key={props.index} />
    } else {
        return null
    }
}

export default ContentGroup