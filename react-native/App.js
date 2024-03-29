import React, { useEffect, useState, useMemo } from 'react'
import { AsyncStorage } from 'react-native'
import { Provider } from 'react-redux'
import store from './store'
import { registerRootComponent } from 'expo'
import { NavigationContainer } from '@react-navigation/native'
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs'
import { createStackNavigator } from '@react-navigation/stack'
import { Icon } from './components/common'
import {
  ContentScreen,
  SharePreviewScreen,
  SendShareScreen,
  DiscoverScreen,
  ProfileScreen,
  EditProfileScreen,
  UserScreen,
  SplashScreen,
  ChatDisplayScreen
} from './screens/index'
import {
  SignupScreen,
  LoginScreen,
  SMSCodeScreen,
  CreatePasswordScreen,
} from './screens/onboarding/index'
import { colors } from './assets/colors'
import * as actions from './actions'
import { AuthContext } from './context'
import { refreshTokens } from './util/auth'
import { Notifications } from 'expo';
import * as Permissions from 'expo-permissions';
import Constants from 'expo-constants';
import Network from './util/Network'

const ApplicationStack = createStackNavigator()
const ApplicationTabs = createBottomTabNavigator()
const ContentStack = createStackNavigator()
const DiscoverStack = createStackNavigator()
const ShareStack = createStackNavigator()
const ProfileStack = createStackNavigator()

const HomeScreen = ({ navigation }) => <ContentScreen kind='home' navigation={navigation} />
const SavedScreen = ({ navigation }) => <ContentScreen kind='saved' navigation={navigation} />
const ByTagScreen = ({ navigation, route }) => <ContentScreen kind='byTag' navigation={navigation} route={route} />

function CreateContentStack() {
  return (
    <ContentStack.Navigator initialRouteName='Home' headerMode='none'>
      <ContentStack.Screen name='Home' component={HomeScreen} />
      <ContentStack.Screen name='Saved' component={SavedScreen} />
      <ContentStack.Screen name='Tag' component={ByTagScreen} />
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
      <ProfileStack.Screen name='Chat' component={ChatDisplayScreen} />
    </ProfileStack.Navigator>
  )
}

function CreateApplicationTabs() {
  return (
    <ApplicationTabs.Navigator
      screenOptions={({ route }) => ({
        tabBarIcon: ({ _, color, size }) => (
          <Icon type={route.name.toLowerCase()} size={size} color={color} />
        )
      })}
      tabBarOptions={{
        activeTintColor: colors.primary,
        inactiveTintColor: 'gray',
        showLabel: false,
      }}
    >
      <ApplicationTabs.Screen name='Home' component={CreateContentStack} />
      <ApplicationTabs.Screen name='Search' component={CreateDiscoverStack} />
      <ApplicationTabs.Screen name='Share' component={CreateShareStack} />
      <ApplicationTabs.Screen name='Profile' component={CreateProfileStack} />
    </ApplicationTabs.Navigator>
  )
}

function App() {
  const [isLoading, setIsLoading] = useState(true)
  const [isSignedIn, setIsSignedIn] = useState(false)

  const registerTokens = async () => {
    let refreshToken = await AsyncStorage.getItem('refresh_token')

    if (refreshToken) {
      await store.dispatch(actions.setTokens({ accessToken: null, refreshToken }))
      refreshTokens()
        .then(() => {
          setIsSignedIn(true)
          setIsLoading(false)
        })
        .catch(err => {
          console.log(err)
          setIsLoading(false)

        })
    } else {
      setIsLoading(false)
    }
  }

  const registerPushNotifications = async () => {
    try {
      const { status: existingStatus } = await Permissions.getAsync(Permissions.NOTIFICATIONS);
      let finalStatus = existingStatus;

      if (existingStatus !== 'granted') {
        const { status } = await Permissions.askAsync(Permissions.NOTIFICATIONS);
        finalStatus = status;
      }

      if (finalStatus !== 'granted') {
        return
      }

      const fullToken = await Notifications.getExpoPushTokenAsync()
      const exponentPushTokenPattern = /^ExponentPushToken\[(\S+)\]$/
      const matches = fullToken.match(exponentPushTokenPattern)
      if (matches.length != 2) {
        return
      }

      const token = matches[1]

      Network.PUT('/users/update/expo-token', { token })
    } catch (e) {
      console.warn(e)
    }

    if (Platform.OS === 'android') {
      Notifications.createChannelAndroidAsync('default', {
        name: 'default',
        sound: true,
        priority: 'max',
        vibrate: [0, 250, 250, 250],
      });
    }
  }

  useEffect(() => {
    const init = async () => {
      await registerTokens()
      // await registerPushNotifications()
    }

    init()
  }, [])

  // todo put notification check after signing in
  const authContext = useMemo(
    () => ({
      signIn: () => {
        setIsSignedIn(true)
        registerPushNotifications()
      },
      signOut: () => { setIsSignedIn(false) },
    }),
    []
  );

  return (
    <AuthContext.Provider value={authContext}>
      <Provider store={store}>
        <NavigationContainer>
          <ApplicationStack.Navigator headerMode='none'>
            {
              isLoading ? (
                <ApplicationStack.Screen name='Splash' component={SplashScreen} />
              ) : isSignedIn ? (
                <ApplicationStack.Screen name='Home' component={CreateApplicationTabs} />
              ) : (
                    <>
                      <ApplicationStack.Screen name='Sign Up' component={SignupScreen} />
                      <ApplicationStack.Screen name='Login' component={LoginScreen} />
                      <ApplicationStack.Screen name='SMS' component={SMSCodeScreen} />
                      <ApplicationStack.Screen name='Password' component={CreatePasswordScreen} />
                    </>
                  )
            }
          </ApplicationStack.Navigator>
        </NavigationContainer>
      </Provider>
    </AuthContext.Provider>
  )
}

registerRootComponent(App)