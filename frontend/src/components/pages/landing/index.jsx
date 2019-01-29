import React, { Component } from "react";
import { SearchBar } from "../../molecules/SearchBar";

export class LandingPage extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div className="landing-page">
        <div className="search-bar-container">
          <SearchBar />
        </div>
      </div>
    );
  }
}
