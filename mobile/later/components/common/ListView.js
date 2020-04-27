import React from 'react';
import { ScrollView, FlatList } from 'react-native';
import { Divider } from './Divider';

function ListView({ data, renderItem, onRefresh }) {
    return (
        // <ScrollView keyboardShouldPersistTaps='handled'>
        <FlatList
            onRefresh={onRefresh}
            data={data}
            renderItem={renderItem}
            ItemSeparatorComponent={<Divider />}
        />
        // </ScrollView>
    );
}

export default ListView