import React, {useState, useEffect} from 'react'

function App() {
  const [favoriteColor, setFavoriteColor] = useState("red");
  const [favoriteNumber, setFavoriteNumber] = useState("ten");

  useEffect(() => {
    setFavoriteNumber("eleven");
  }, [favoriteNumber])
  
  return (
      <div className="container">
        <div className="row">
          <div className="col">
            <h2 className="pt-4">Hook Test</h2>
            <hr />
            <p>Your favorite color is {favoriteColor}, and your favorite number is {favoriteNumber}.</p>
            <p>
              <button className="btn btn-primary" onClick={() => setFavoriteColor("green")}>
                Change color
              </button>
            </p>
          </div>
        </div>
      </div>
  );
}

export default App;
