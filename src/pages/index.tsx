import React from 'react'
import Head from 'next/head'

import Coding4uLogo from '../assets/coding4u.svg'

import { Container } from '../styles/pages/home'

const Home: React.FC = () => {
  return (
    <Container>
      <Head>
        <title>Homepage</title>
      </Head>

      <Coding4uLogo />
      <h1>Initial Interface</h1>
      <p>develop by patrick</p>
    </Container>
  )
}

export default Home
