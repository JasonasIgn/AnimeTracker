import React from "react";
import config from "../../config";

export const SearchBar = ({ suggestions, onChange }) => (
  <div className="search-bar">
    <input
      type="text"
      placeholder="search"
      className={`${suggestions.length > 0 ? "active" : ""}`}
      onChange={e => {
        onChange(e.target.value);
      }}
    />
    <ul
      className={`suggestions-list ${
        suggestions.length === 0 ? "disabled" : ""
      }`}
    >
      {suggestions.map((element, index) => {
        if (index < config.utils.maxSuggestionsInList)
          return <li key={index}>{element.Title}</li>;
      })}
    </ul>
  </div>
);
