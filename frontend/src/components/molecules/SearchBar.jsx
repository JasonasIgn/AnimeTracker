import React from "react";

export const SearchBar = ({ suggestions, onChange }) => (
  <div className="search-bar">
    <input
      type="text"
      placeholder="search"
      className={`${suggestions ? "active" : ""}`}
      onChange={e => {
        onChange(e.target.value);
      }}
    />
    <ul className={"suggestions-list"}>
      {suggestions.map((element, index) => {
        console.log(element);
        return <li key={index}>{element}</li>;
      })}
    </ul>
  </div>
);
