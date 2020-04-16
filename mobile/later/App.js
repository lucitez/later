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
import DiscoverScreen from './screens/DiscoverScreen';
import ProfileScreen from './screens/ProfileScreen';
import EditProfileScreen from './screens/EditProfileScreen';

const Tab = createBottomTabNavigator();
const ContentStack = createStackNavigator();
const DiscoverStack = createStackNavigator();
const ShareStack = createStackNavigator();
const ChatStack = createStackNavigator();
const ProfileStack = createStackNavigator();

function CreateContentStack() {
  return (
    <ContentStack.Navigator initialRouteName='Home' headerMode='none'>
      <ContentStack.Screen name='Home' component={ContentScreen} />
      <ContentStack.Screen name='Archive' component={ArchiveScreen} />
      <ContentStack.Screen name='Forward' component={SendShareScreen} />
    </ContentStack.Navigator>
  )
}

function CreateDiscoverStack() {
  return (
    <DiscoverStack.Navigator initialRouteName='Todo' headerMode='none'>
      <DiscoverStack.Screen name='Todo' component={DiscoverScreen} />
    </DiscoverStack.Navigator>
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

function CreateChatStack() {
  return (
    <ChatStack.Navigator initialRouteName='Todo' headerMode='none'>
      <ChatStack.Screen name='Todo' component={DiscoverScreen} />
    </ChatStack.Navigator>
  )
}

function CreateProfileStack() {
  return (
    <ProfileStack.Navigator initialRouteName='Profile' headerMode='none'>
      <ProfileStack.Screen name='Profile' component={ProfileScreen} />
      <ProfileStack.Screen name='Edit' component={EditProfileScreen} />
    </ProfileStack.Navigator>
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
          <Tab.Screen name='Discover' component={CreateDiscoverStack} />
          <Tab.Screen name='Share' component={CreateShareStack} />
          <Tab.Screen name='Chat' component={CreateChatStack} />
          <Tab.Screen name='Profile' component={CreateProfileStack} />
        </Tab.Navigator>
      </NavigationContainer>
    )
  }
}

registerRootComponent(App)