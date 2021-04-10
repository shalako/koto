import { ApiTypes, CommonTypes } from 'src/types'

export enum Types {
  SET_DIRECT_MSG_POST_TOKEN               = 'SET_DIRECT_MSG_POST_TOKEN',
  SEND_MESSAGE_TO_FRIEND                  = 'SEND_MESSAGE_TO_FRIEND',
  GET_DIRECT_POST_MSG_TOKEN               = 'GET_DIRECT_POST_MSG_TOKEN',
  GET_FRIEND_MSG_API_DATA                 = 'GET_FRIEND_MSG_API_DATA',
  GET_FRIEND_MSG_SUCCESS                  = 'GET_FRIEND_MSG_SUCCESS',

  ADD_FRIEND_TO_ROOM                      = 'ADD_FRIEND_TO_ROOM',
  GET_LAST_MESSAGE_TOKENS_REQUEST         = 'GET_MESSAGE_TOKENS_REQUEST',
  GET_MESSAGE_TOKENS_SUCCESS              = 'GET_MESSAGE_TOKENS_SUCCESS',

  GET_MORE_MESSAGE_REQUEST                = 'GET_MORE_MESSAGE_REQUEST',
  GET_MORE_MESSAGE_SUCCESS                = 'GET_MORE_MESSAGE_SUCCESS',
  GET_MORE_MESSAGE_FAILED                 = 'GET_MORE_MESSAGE_FAILED',

  GET_MESSAGE_TOKENS_FROM_HUB_REQUEST     = 'GET_MESSAGE_TOKENS_FROM_HUB_REQUEST',
  GET_USER_LAST_MESSAGES_FROM_HUB_SUCCESS = 'GET_USER_LAST_MESSAGES_FROM_HUB_SUCCESS',
  GET_MESSAGE_TOKENS_FROM_HUB_FAILED      = 'GET_MESSAGE_TOKENS_FROM_HUB_FAILED',
}

const sendMessageToFriend = (payload: ApiTypes.Feed.PostMessage) => ({
  type: Types.SEND_MESSAGE_TO_FRIEND,
  payload
})

const setDirectMsgPostToken = (payload: CommonTypes.TokenData) => ({
  type: Types.SET_DIRECT_MSG_POST_TOKEN,
  payload
})

const getFriendMsg = (payload: ApiTypes.Messages.GetFriendMsgAPIData) => ({
  type: Types.GET_FRIEND_MSG_API_DATA,
  payload
})

const getFriendMsgSuccess = (payload: CommonTypes.MessageTypes.MessageItemProps[]) => ({
  type: Types.GET_FRIEND_MSG_SUCCESS,
  payload
})

const getDirectPostMsgToken = (payload: string) => ({
  type: Types.GET_DIRECT_POST_MSG_TOKEN,
  payload
})

const addFriendToRoom = (payload: CommonTypes.MessageRoomFriendData) => ({
  type: Types.ADD_FRIEND_TO_ROOM,
  payload
})

const getMessageTokensRequest = () => ({
  type: Types.GET_LAST_MESSAGE_TOKENS_REQUEST,
})

const getMessageTokensSuccess = (payload: CommonTypes.HubTypes.CurrentHub[]) => ({
  type: Types.GET_MESSAGE_TOKENS_SUCCESS,
  payload
})

const getUserLastMessageFromHubSuccess = (payload: {
  hub: string
  usersLastMessage: ApiTypes.Messages.UserMessage[]
}) => ({
  type: Types.GET_USER_LAST_MESSAGES_FROM_HUB_SUCCESS,
  payload
})

const getMessageFromHubFailed = () => ({
  type: Types.GET_MESSAGE_TOKENS_FROM_HUB_FAILED
})

const getMoreMessageRequest = () => ({
  type: Types.GET_MORE_MESSAGE_REQUEST,
})

const getMoreMessageSucces = (payload: CommonTypes.HubTypes.CurrentHub[]) => ({
  type: Types.GET_MORE_MESSAGE_SUCCESS,
  payload
})

const getMoreMessageFailed = () => ({
  type: Types.GET_MORE_MESSAGE_FAILED,
})

export default {
  setDirectMsgPostToken,
  sendMessageToFriend,
  getFriendMsg,
  getFriendMsgSuccess,
  getDirectPostMsgToken,
  addFriendToRoom,
  getMessageTokensRequest,
  getMessageTokensSuccess, 
  getUserLastMessageFromHubSuccess,
  getMessageFromHubFailed,
}
