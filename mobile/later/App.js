import React from 'react';
import { registerRootComponent } from 'expo';
import { NavigationContainer } from '@react-navigation/native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { createStackNavigator } from '@react-navigation/stack';
import Icon from './components/Icon';
import ContentScreen from './screens/ContentScreen';
import ArchiveScreen from './screens/ArchiveScreen';
import FriendScreen from './screens/FriendScreen';
import AddFriendScreen from './screens/AddFriendScreen';
import { colors } from './assets/colors';
import SharePreviewScreen from './screens/SharePreviewScreen';
import SendShareScreen from './screens/SendShareScreen';

const Tab = createBottomTabNavigator();
const ContentStack = createStackNavigator();
const FriendStack = createStackNavigator();
const ShareStack = createStackNavigator();

function CreateFriendStack() {
  return (
    <FriendStack.Navigator initialRouteName='Friends' headerMode='none'>
      <FriendStack.Screen name='Friends' component={FriendScreen} />
      <FriendStack.Screen name='Test' component={AddFriendScreen} />
    </FriendStack.Navigator>
  )
}

function CreateShareStack() {
  return (
    <ShareStack.Navigator initialRouteName='Share' headerMode='none'>
      <ShareStack.Screen name='Share' component={SharePreviewScreen} />
      <ShareStack.Screen name='Send Share' component={SendShareScreen} />
    </ShareStack.Navigator>
  )
}

function CreateContentStack() {
  return (
    <ContentStack.Navigator initialRouteName='Home' headerMode='none'>
      <ContentStack.Screen name='Home' component={ContentScreen} />
      <ContentStack.Screen name='Archive' component={ArchiveScreen} />
      <ContentStack.Screen name='Forward' component={SendShareScreen} />
    </ContentStack.Navigator>
  )
}

class App extends React.Component {
  render() {
    return (
      <NavigationContainer>
        <Tab.Navigator
          screenOptions={({ route }) => ({
            tabBarIcon: ({ _, color, size }) => (
              <Icon type={route.name.toLowerCase()} size={size} color={color} />
            )
          })}
          tabBarOptions={{
            activeTintColor: colors.primary,
            inactiveTintColor: 'gray',
          }}
        >
          <Tab.Screen name='Home' component={CreateContentStack} />
          <Tab.Screen name='Share' component={CreateShareStack} />
          <Tab.Screen name='Friends' component={CreateFriendStack} />
        </Tab.Navigator>
      </NavigationContainer>
    )
  }
}

registerRootComponent(App)