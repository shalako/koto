import React, { useEffect, useState } from 'react'
import { PageLayout } from '@view/shared/PageLayout'
import GroupTopBar from './GroupTopBar'
import { Member } from './Member'
import DeleteGroupDialog from './DeleteGroupDialog'
import { connect } from 'react-redux'
import Actions from '@store/actions'
import selectors from '@selectors/index'
import { ApiTypes, StoreTypes } from 'src/types'
import { v4 as uuidv4 } from 'uuid'
import { getGroupAvatarUrl, getGroupCoverUrl } from '@services/avatarUrl'
import { Container } from '@view/shared/styles'
import UserForInvite from './UserForInvite'
import {
  GroupCover,
  GroupMainWrapper,
  LeftSideBar,
  RightSideBar,
  CentralBar,
  AvatarStyled,
  GroupName,
  GroupPublicity,
  GroupDescriptopn,
  BarTitle,
} from './styles'

interface Props {
  groupDetails: ApiTypes.Groups.GroupDetails | null
  friends: ApiTypes.Friends.Friend[] | null

  onGetFriends: () => void
  onGetInvitesToConfirmRequest: () => void
}

const AdminLayoutPrivate: React.FC<Props> = React.memo((props) => {
  const [groupInvites, setGroupInvites] = useState<ApiTypes.Groups.Invite[] | null>(null)
  const [isRequested, setRequested] = useState(false)
  const { 
    groupDetails, 
    onGetInvitesToConfirmRequest, 
    onGetFriends,
    friends,
   } = props

  useEffect(() => {
    if (groupInvites === null && !isRequested) {
      onGetInvitesToConfirmRequest()
      setRequested(true)
    }

    if (groupDetails?.invites?.length && isRequested) {
      setGroupInvites(fixInvitesGroupId())
      setRequested(false)
    }

    if (friends === null) {
      onGetFriends()
    }

  }, [groupInvites, isRequested, friends])

  const fixInvitesGroupId = () => {
    if (!groupDetails?.invites?.length) return []
    
    return groupDetails?.invites?.map(item => {
      item.group_id = groupDetails?.group?.id
      return item
    })
  }

  const filterFriendsForInvite = () => {
    return friends?.filter((item) =>
      !Boolean(
        groupDetails?.members?.some(
          member => member.id === item.user.id
        )
      )
    )
  }

  if (!groupDetails) return null

  const { group, members, status, invites } = groupDetails

  return (
    <PageLayout>
      <GroupCover resource={getGroupCoverUrl(group?.id)}/>
      <GroupTopBar 
        className="desktop-only"
        memberStatus={status}
        membersCounter={members?.length} 
        invitesCounter={invites?.length || 0} 
        groupId={group?.id} 
        isAdminLayout={true}
        isPublic={group?.is_public}
      />
      <Container>
        <GroupMainWrapper>
          <LeftSideBar>
            <AvatarStyled src={getGroupAvatarUrl(group?.id)}/>
            <GroupName>{group?.name}</GroupName>
            <GroupPublicity>{group?.is_public ? 'Public' : 'Private'} group</GroupPublicity>
            <GroupDescriptopn>{group?.description}</GroupDescriptopn>
            <DeleteGroupDialog 
              className="desktop-only"
              groupId={group?.id} 
            />
            <GroupTopBar
              className="mobile-only"
              memberStatus={status}
              membersCounter={members?.length}
              invitesCounter={invites?.length || 0}
              groupId={group?.id}
              isAdminLayout={true}
              isPublic={group?.is_public}
            />
          </LeftSideBar>
          <CentralBar>
            <BarTitle>Members ({members?.length})</BarTitle>
            {Boolean(members?.length) && members.map(item => (
              <Member
                groupId={group?.id}
                isAdminLayout={true}
                key={uuidv4()}
                {...item}
              />
            ))}
            {/* <ViewMoreButton>View more</ViewMoreButton> */}
          </CentralBar>
          <RightSideBar>
            <BarTitle>Invite friends</BarTitle>
            {filterFriendsForInvite()?.map(item => <UserForInvite 
              groupId={group?.id}
              key={uuidv4()} 
              {...item}
              />)}
            {/* <ViewMoreButton>View more</ViewMoreButton> */}
          </RightSideBar>
          <DeleteGroupDialog
            className="mobile-only"
            groupId={group?.id}
          />
        </GroupMainWrapper>
      </Container>
    </PageLayout>
  )
})

type StateProps = Pick<Props, 'groupDetails' | 'friends'>
const mapStateToProps = (state: StoreTypes): StateProps => ({
  groupDetails: selectors.groups.groupDetails(state),
  friends: selectors.friends.friends(state),
})

type DispatchProps = Pick<Props, 'onGetInvitesToConfirmRequest' | 'onGetFriends'>
const mapDispatchToProps = (dispatch): DispatchProps => ({
  onGetInvitesToConfirmRequest: () => dispatch(Actions.groups.getInvitesToConfirmRequest()),
  onGetFriends: () => dispatch(Actions.friends.getFriendsRequest()),
})

export default connect(mapStateToProps, mapDispatchToProps)(AdminLayoutPrivate)
