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
                        {index < props.content.length - 1 ? <Divider key={props.index} /> : null}
                    </View>
                ))
            }
        </ScrollView>
    );
}

export default ContentGroup