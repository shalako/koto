import styled from 'styled-components'
import Container from '@material-ui/core/Container'
import Paper from '@material-ui/core/Paper'
import TextareaAutosize from 'react-autosize-textarea'
import Button from '@material-ui/core/Button'
import Link from '@material-ui/core/Link'
import IconButton from '@material-ui/core/IconButton'

export const ContainerStyled = styled(Container)`
  && {
    padding-top: 100px;
  }
`

export const PaperStyled = styled(Paper)`
  && {
    padding: 20px 20px;
    margin-bottom: 15px;
  }
`

export const CommentsWrapepr = styled.div`
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-end;
`

export const CommentWrapper = styled(Paper)`
  && {
    background: #EEEEEE;
    padding: 10px 20px;
    margin-bottom: 10px;
    width: calc(100% - 80px);
  }
`

export const CreateWrapper = styled.div`
  display: flex;
`

export const EditorWrapper = styled.div`
  width: calc(100% - 60px);
  margin-left: 20px;
`

export const TextareaAutosizeStyled = styled(TextareaAutosize)`
  outline: none;
  resize:none;
  border: none;
  border-bottom: 1px solid #000;
  width: 100%;
  min-height: 40px;
  font-family: 'Arial';
  font-size: 14px;
  line-height: 1.5;
  margin-top: 5px;
  margin-bottom: 5px;
  background: transparent;
`

export const ButtonSend = styled(Button)`
  && {
    float: right;
  }
`

export const TextareaTitle = styled.p`
  font-size: 14px;
  margin-top: -5px;
  color: grey;
  transition: 0.3s;
  opacity: 1;

  &.active {
    opacity: 0;
  }
`

export const MessageHeader = styled.header`
  display: flex;
  justify-content: space-between;
`

export const UserInfo = styled.div`
  display: flex;
  align-items: center;
`

export const UserNameWrapper = styled.div`
  margin-left: 20px;
  
`

export const UserName = styled(Link)`
  cursor: pointer;
`

export const MessageDate = styled.div`
  font-size: 12px;
`

export const ButtonsWrapper = styled.div`
  display: flex;
`

export const MessageContent = styled.div`
  font-size: 14px;
  margin-top: 10px;
  padding-left: 60px;
`

export const EditMessageWrapper = styled.div`
  padding-left: 60px;
  
  &:after {
    content: '';
    clear: both;
    display: block;
  }
`

export const CroppedText = styled.p`
  white-space: nowrap;
  overflow: hidden; 
  text-overflow: ellipsis; 
`

export const MessageSticky = styled.div`
  /* position: sticky; */
  /* bottom: 0; */
  /* z-index: 100; */
`

export const EmptyMessage = styled.div`
  text-align: center;
  font-weight: bold;
`

export const UpButton = styled(IconButton)`
  position: fixed;
  right: 20px;
  bottom: 20px;
  z-index: 1000;
`

export const CommentsLinkWrapper = styled.div`
  text-align: right;
`

export const CommentsLink = styled(Link)`
  && {
    cursor: pointer;
    color: #1976d2;
  }
`