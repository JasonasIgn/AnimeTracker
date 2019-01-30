import React, { Component } from "react";
import { SearchBar } from "../../molecules/SearchBar";

export class LandingPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      fullList: ["asda", "aaaaa", "bbbbb", "cccc", "dddd", "kkkkk"],
      filteredList: [],
      searchQuery: ""
    };
  }

  onSearchQueryChange = searchQuery => {
    const regex = new RegExp(`^${searchQuery}`, "i");
    this.setState({
      filteredList: searchQuery
        ? this.state.fullList.sort().filter(v => regex.test(v))
        : []
    });
  };

  render() {
    return (
      <div className="landing-page">
        <div className="search-bar-container">
          <SearchBar
            suggestions={this.state.filteredList}
            onChange={this.onSearchQueryChange}
          />
        </div>
      </div>
    );
  }
}
