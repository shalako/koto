import React from 'react'
import Typography from '@material-ui/core/Typography'
import milhousePicture from './../../../assets/images/milhouse.gif'
import { ContainerStyled } from './styles'
import { ImageDemo } from './styles'
export const Slide1 = () => {
  return (
    <ContainerStyled maxWidth="md">
      <Typography variant="h3" gutterBottom>Welcome to Koto</Typography>

      <Typography variant="h5" gutterBottom>
      Like most social networks, koto starts with friends. And from the look of it, you don't have any.
      </Typography>
      <ImageDemo src={milhousePicture} alt="no friends" />

      <Typography variant="h5" gutterBottom>
      Swipe right to learn more ➡️
      </Typography>
    </ContainerStyled>
  )
}