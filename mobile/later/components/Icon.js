import React from 'react';
import SimpleLineIcons from 'react-native-vector-icons/SimpleLineIcons';
import AntIcons from 'react-native-vector-icons/AntDesign';

function Icon(props) {

    switch (props.type) {
        case 'watch':
            return <SimpleLineIcons name='screen-desktop' {...props} />
        case 'read':
            return <SimpleLineIcons name='eyeglass' {...props} />
        case 'listen':
            return <SimpleLineIcons name='earphones' {...props} />
        case 'add_friend':
            return <AntIcons name='adduser' {...props} />
        case 'search':
            return <AntIcons name='search1' {...props} />
        case 'plus':
            return <AntIcons name='pluscircleo' {...props} />
        case 'back':
            return <AntIcons name='left' {...props} />
        default:
            return null
    }
}

export default Icon