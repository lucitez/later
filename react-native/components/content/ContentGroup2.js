import React from 'react';
import { View, ScrollView } from 'react-native';
import { Divider } from '../common'
import ContentPreview from './ContentPreview';

export default function ContentGroup({ contents, ...props }) {
    return (
        <ScrollView keyboardShouldPersistTaps='handled'>
            {contents.map((contentObj, index) => (
                <View key={index}>
                    <ContentPreview content={contentObj} {...props} />
                    {index < contents.length - 1 && <Divider />}
                </View>
            ))}
        </ScrollView>
    );
}