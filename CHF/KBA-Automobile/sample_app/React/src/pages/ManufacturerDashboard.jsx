import React, { useState } from "react";

const ManufacturerDashboard = () => {
  const [carData, setCarData] = useState({
    carId: "",
    make: "",
    model: "",
    color: "",
    dateOfManufacture: "",
    manufacturerName: "",
  });

  const [queryCarId, setQueryCarId] = useState("");
  const [carDetails, setCarDetails] = useState(null);

  // Handle form inputs
  const handleChange = (e) => {
    setCarData({ ...carData, [e.target.name]: e.target.value });
  };

  // Add Car API Call
  const addData = async (e) => {
    e.preventDefault();

    const { carId, make, model, color, dateOfManufacture, manufacturerName } = carData;

    if (
      !carId ||
      !make ||
      !model ||
      !color ||
      !dateOfManufacture ||
      !manufacturerName
    ) {
      return alert("Please enter all data properly.");
    }

    try {
      const response = await fetch("http://localhost:3001/api/car", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(carData),
      });
      
      if (!response.ok) {
        const errorText = await response.text();
        console.error("Server error:", errorText);
        alert("Error creating car.");
        return;
      }
      
      const data = await response.json();
      console.log("Car created:", data);
      alert("Car Created Successfully!");
      setCarData({
        carId: "",
        make: "",
        model: "",
        color: "",
        dateOfManufacture: "",
        manufacturerName: "",
      });
    } catch (err) {
      alert("Error creating car.");
      console.log(err);
    }
  };

  // Query Car by ID API Call
  const readData = async (e) => {
    e.preventDefault();

    if (!queryCarId) {
      return alert("Please enter a valid Car ID.");
    }

    try {
      const response = await fetch(`http://localhost:3001/api/car/${queryCarId}`);
      const responseData = await response.json();
      console.log("Queried car:", responseData);
      setCarDetails(responseData);
    } catch (err) {
      alert("Error fetching car.");
      console.log(err);
    }
  };

  // Get All Cars (optional API call — demo static data here)
  const getAllCars = async () => {
    try {
      const response = await fetch("http://localhost:3001/api/car");
      const allCars = await response.json();
      console.log("All cars:", allCars);
      setCarDetails(allCars); // Or handle as needed
    } catch (err) {
      alert("Error fetching all cars.");
      console.log(err);
    }
  };

  return (
    <>
      <div className="max-w-2xl mx-auto p-6 bg-white rounded-2xl shadow-lg">
        <h1 className="text-3xl font-bold text-center mb-8 text-blue-700">
          Manufacturer Dashboard
        </h1>

        <h2 className="text-2xl font-semibold mb-4 text-gray-800">Create Car</h2>
        <form onSubmit={addData} className="space-y-4">
          {[
            { label: "Car ID", name: "carId" },
            { label: "Make", name: "make" },
            { label: "Model", name: "model" },
            { label: "Color", name: "color" },
            { label: "Date of Manufacture", name: "dateOfManufacture", type: "date" },
            { label: "Manufacturer Name", name: "manufacturerName" },
          ].map((field, index) => (
            <div key={index} className="flex flex-col">
              <label className="font-medium text-gray-700 mb-1">{field.label}</label>
              <input
                type={field.type || "text"}
                name={field.name}
                value={carData[field.name]}
                onChange={handleChange}
                required
                className="border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
          ))}

          <div className="flex space-x-3 mt-4">
            <button
              type="submit"
              className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
            >
              Create Car
            </button>
            <button
              type="button"
              onClick={() =>
                setCarData({
                  carId: "",
                  make: "",
                  model: "",
                  color: "",
                  dateOfManufacture: "",
                  manufacturerName: "",
                })
              }
              className="bg-gray-500 text-white px-4 py-2 rounded-lg hover:bg-gray-600 transition"
            >
              Clear
            </button>
          </div>
        </form>
      </div>

      {/* Query Car Section */}
      <div className="max-w-lg mx-auto p-6 mt-12 bg-white rounded-2xl shadow-lg">
        <h2 className="text-2xl font-semibold mb-4 text-gray-800 text-center">
          Query Car
        </h2>

        <form onSubmit={readData} className="space-y-4">
          <div className="flex flex-col">
            <label className="font-medium text-gray-700 mb-1">Enter Car ID</label>
            <input
              type="text"
              value={queryCarId}
              onChange={(e) => setQueryCarId(e.target.value)}
              className="border w-full border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div className="flex space-x-3 mt-4">
            <button
              type="submit"
              className="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition"
            >
              Query Car
            </button>
            <button
              type="button"
              onClick={() => setQueryCarId("")}
              className="bg-gray-500 text-white px-4 py-2 rounded-lg hover:bg-gray-600 transition"
            >
              Clear
            </button>
          </div>
        </form>

        {/* Get All Cars Button */}
        <div className="flex justify-center mt-6">
          <button
            onClick={getAllCars}
            className="bg-indigo-600 text-white px-6 py-2 rounded-lg hover:bg-indigo-700 transition"
          >
            Get All Cars
          </button>
        </div>
      </div>

      {/* Car Details Display Section */}
      {carDetails && (
        <div className="max-w-lg mx-auto p-6 mt-8 bg-gray-100 rounded-2xl shadow">
          <h3 className="text-xl font-semibold text-gray-800 mb-4">Car Details</h3>
          <div className="space-y-2 text-gray-700">
            <p><strong>Car ID:</strong> {carDetails.carId}</p>
            <p><strong>Make:</strong> {carDetails.make}</p>
            <p><strong>Model:</strong> {carDetails.model}</p>
            <p><strong>Color:</strong> {carDetails.color}</p>
            <p><strong>Date of Manufacture:</strong> {carDetails.dateOfManufacture}</p>
            <p><strong>Manufacturer Name:</strong> {carDetails.manufacturerName}</p>
          </div>
        </div>
      )}
    </>
  );
};

export default ManufacturerDashboard;
