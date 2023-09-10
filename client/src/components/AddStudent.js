import React, { useState } from "react";
import { useMutation } from "react-query";
import { API } from "../config/api";

const AddStudent = ({ refetchStudents }) => {
  const [formInput, setFormInput] = useState({
    nim: "",
    fullname: "",
    majority: "",
    address: "",
  });

  const [errors, setErrors] = useState({
    nim: "",
    fullname: "",
    majority: "",
    address: "",
  });

  const handleInputChange = (e) => {
    setFormInput({ ...formInput, [e.target.name]: e.target.value });
  };

  const handleAddMahasiswa = useMutation(async () => {
    await API.post("/students", {
      nim: formInput.nim,
      fullname: formInput.fullname,
      majority: formInput.majority,
      address: formInput.address,
    });

    refetchStudents();
  });

  const handleFormSubmit = (e) => {
    e.preventDefault();
    let errorMsg = {};

    // NIM Validation
    if (formInput.nim.trim() === "") {
      errorMsg.nim = "NIM can't be empty";
    } else if (!/^\d{12}/.test(formInput.nim)) {
      errorMsg.nim = "NIM must be a 12 digit number";
    } else {
      errorMsg.nim = "";
    }

    // Fullname Validation
    errorMsg.fullname =
      formInput.fullname.trim() === "" ? "Fullname can't be empty" : "";

    // Majority Validation
    errorMsg.majority =
      formInput.majority.trim() === "" ? "Majority can't be empty" : "";

    // Address Validation
    errorMsg.address =
      formInput.address.trim() === "" ? "Address can't be empty" : "";

    // Update State Error
    setErrors(errorMsg);

    // Check all validation
    let formValid = true;
    for (let inputName in errorMsg) {
      errorMsg[inputName] !== "" && (formValid = false);
    }

    // Check all validation
    if (formValid) {
      console.log("hit post");
      // add new student
      handleAddMahasiswa.mutate();

      // Empty the error form
      setFormInput({
        nim: "",
        fullname: "",
        majority: "",
        address: "",
      });
    }
  };

  return (
    <tr>
      <td colSpan="5">
        <form onSubmit={handleFormSubmit}>
          <div className="row g-3">
            <div className="col">
              <input
                type="text"
                name="nim"
                placeholder="Insert NIM"
                className="form-control"
                value={formInput.nim}
                onChange={handleInputChange}
              />
              {errors && <small>{errors.nim}</small>}
            </div>
            <div className="col">
              <input
                type="text"
                name="fullname"
                placeholder="Insert Fullname"
                className="form-control"
                value={formInput.fullname}
                onChange={handleInputChange}
              />
              {errors && <small>{errors.fullname}</small>}
            </div>
            <div className="col">
              <input
                type="text"
                name="majority"
                placeholder="Insert Majority"
                className="form-control"
                value={formInput.majority}
                onChange={handleInputChange}
              />
              {errors && <small>{errors.majority}</small>}
            </div>
            <div className="col">
              <input
                type="text"
                name="address"
                placeholder="Insert Address"
                className="form-control"
                value={formInput.address}
                onChange={handleInputChange}
              />
              {errors && <small>{errors.address}</small>}
            </div>
            <div className="col">
              <button type="submit" className="btn btn-primary">
                Add
              </button>
            </div>
          </div>
        </form>
      </td>
    </tr>
  );
};

export default AddStudent;
