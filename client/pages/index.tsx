import type { NextPage } from 'next'
import React, { useEffect, useState } from 'react'
import { Card, Col, Container, Row } from 'react-bootstrap'

const Home: NextPage = () => {

  const [projects, setProjects] = useState<any[]>([])

  useEffect(() => {
    fetch("http://localhost:8080/api/repositories")
  }, [])

  return (
    <div>
      <Container>
        <Row>
          <Col>
            <h1>Savannah</h1>
          </Col>
        </Row>
        <Row>
          {projects.map((p, i) => <Card key={i}>
            <Card.Body>
              <Card.Title>{p.name}</Card.Title>
              <Card.Link href="#">Sync</Card.Link>
            </Card.Body>
          </Card>)}
        </Row>
      </Container>
    </div>)
}

export default Home
