import React from 'react';
import './App.css';
import './pages/page.css';
import 'bootstrap/dist/css/bootstrap.min.css';

import Home from './pages/Home';
import Login from './pages/Login';
import NotFound from './pages/NotFound';

import {
  BrowserRouter as Router,
  Route,
  NavLink
} from 'react-router-dom';

import { CSSTransition } from 'react-transition-group'
import { Navbar, Nav } from 'react-bootstrap'

function App() {
  const routes = [
    { path: '/app', name: 'Home', Component: Home, is_navlink: true },
    { path: '/app/login', name: 'Login', Component: Login, is_navlink: true},
    { path: '/app/404', name: 'NotFound', Component: NotFound, is_navlink: false}
  ];

  return (
    <div className="App">
      <Router>
        <>
          <Navbar bg="light">
            <Nav className="mx-auto">
              {routes.filter(route => route.is_navlink).map(route => (
                <Nav.Link
                  key={route.path}
                  as={NavLink}
                  to={route.path}
                  activeClassName="active"
                  exact
                >
                  {route.name}
                </Nav.Link>
              ))}
            </Nav>
          </Navbar>
          
          <div>
            {routes.map(({ path, Component }) => (
              <Route key={path} exact path={path}>
                {({ match }) => (
                  <CSSTransition
                    in={match != null}
                    timeout={300}
                    classNames='page'
                    unmountOnExit
                  >
                    <div className='page'>
                      <Component />
                    </div>
                  </CSSTransition>
                )}
              </Route>
            ))}
          </div>
        </>
      </Router>
    </div>
  );
}

export default App;