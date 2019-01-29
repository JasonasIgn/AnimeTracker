import React, { Component } from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

import { LandingPage } from "./components/pages/landing/index";

class App extends Component {
  render() {
    return (
      <Router>
        <Route path="/" exact component={LandingPage} />
      </Router>
    );
  }
}
export default App;
