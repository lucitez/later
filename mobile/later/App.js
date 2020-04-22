import React from 'react';
import { registerRootComponent } from 'expo';
import { NavigationContainer } from '@react-navigation/native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { createStackNavigator } from '@react-navigation/stack';
import { Icon } from './components/common';
import ContentScreen from './screens/ContentScreen';
import SavedScreen from './screens/SavedScreen';
import { colors } from './assets/colors';
import SharePreviewScreen from './screens/SharePreviewScreen';
import SendShareScreen from './screens/SendShareScreen';
import DiscoverScreen from './screens/DiscoverScreen';
import ProfileScreen from './screens/ProfileScreen';
import EditProfileScreen from './screens/EditProfileScreen';
import LoginScreen from './screens/LoginScreen';
import UserScreen from './screens/UserScreen';

const Tab = createBottomTabNavigator();
const ContentStack = createStackNavigator();
const DiscoverStack = createStackNavigator();
const ShareStack = createStackNavigator();
const ProfileStack = createStackNavigator();

function CreateContentStack() {
  return (
    <ContentStack.Navigator initialRouteName='Home' headerMode='none'>
      <ContentStack.Screen name='Home' component={ContentScreen} />
      <ContentStack.Screen name='Saved' component={SavedScreen} />
      <ContentStack.Screen name='Forward' component={SendShareScreen} />
    </ContentStack.Navigator>
  )
}

function CreateDiscoverStack() {
  return (
    <DiscoverStack.Navigator initialRouteName='Discover' headerMode='none'>
      <DiscoverStack.Screen name='Discover' component={DiscoverScreen} />
      <DiscoverStack.Screen name='User' component={UserScreen} />
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
          initialRouteName='listen'
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
          <Tab.Screen name='listen' component={LoginScreen} />
          <Tab.Screen name='Search' component={CreateDiscoverStack} />
          <Tab.Screen name='Share' component={CreateShareStack} />
          <Tab.Screen name='Profile' component={CreateProfileStack} />
        </Tab.Navigator>
      </NavigationContainer>
    )
  }
}

registerRootComponent(App)