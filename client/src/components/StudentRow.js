import React, { useRef, useState } from "react";
import { API } from "../config/api";

const StudentRow = (props) => {
  const [formInput, setFormInput] = useState({
    id: props.student.id,
    nim: props.student.nim,
    fullname: props.student.fullname,
    majority: props.student.majority,
    address: props.student.address,
  });
  const [editStatus, setEditStatus] = useState(false);
  const [errors, setErrors] = useState({
    nim: "",
    fullname: "",
    majority: "",
    address: "",
  });
  const dataReset = useRef({});

  const handleInputChange = (e) => {
    setFormInput({ ...formInput, [e.target.name]: e.target.value });
  };

  const handleEditCondition = () => {
    if (editStatus) {
      setFormInput(dataReset.current);
      setEditStatus(!editStatus);
    } else {
      dataReset.current = { ...formInput };
      setEditStatus(!editStatus);
    }
  };

  const handleFormSubmit = (e) => {
    e.preventDefault();
    let errorMsg = {};

    // NIM Validation
    let regex = new RegExp(/^\d{12}/);
    if (formInput.nim.trim() === "") {
      errorMsg.nim = "NIM can't be empty";
    } else if (!regex.test(formInput.nim)) {
      errorMsg.nim = "NIM must be a 12 digit number";
    } else {
      errorMsg.nim = "";
    }

    // Fullname Validation
    errorMsg.fullname =
      formInput.fullname.trim() === "" ? "Fullname can't be empty" : "";

    // Majority Validation
    errorMsg.majority =
      formInput.majority.trim() === "" ? "majority can't be empty" : "";

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

    // If passed the validation
    if (formValid) {
      handleEditStudent(formInput);
      setEditStatus(false);
    }
  };

  const handleEditStudent = async (data) => {
    const response = await API.patch("/students/" + data.id, {
      nim: formInput.nim,
      fullname: formInput.fullname,
      majority: formInput.majority,
      address: formInput.address,
    });
    console.log(response.data);
    if (response.data.status === 200) {
      alert("Student data updated successfully...");
    }
  };

  const handleDeleteStudent = async (id) => {
    let result = window.confirm("Are you sure ?");

    if (result === true) {
      // user choose "OK"
      const response = await API.delete("/students/" + id);
      console.log(response.data);
      if (response.data.status === 200) {
        alert("Student data deleted successfully...");
      }
    } else {
      // user choose "Cancel"
      alert("You cancelled to delete the student");
    }
  };

  return (
    <>
      {editStatus ? (
        <tr>
          <td colSpan="5">
            <form onSubmit={handleFormSubmit}>
              <div className="row">
                <div className="col">
                  <input
                    type="text"
                    name="nim"
                    placeholder="Insert NIM"
                    className="form-control"
                    value={formInput.nim}
                    onChange={handleInputChange}
                    disabled
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
                  <button type="submit" className="btn btn-success me-2">
                    Save
                  </button>
                  <button
                    type="submit"
                    className="btn btn-warning"
                    onClick={handleEditCondition}
                  >
                    Cancel
                  </button>
                </div>
              </div>
            </form>
          </td>
        </tr>
      ) : (
        <tr>
          <td>{formInput.nim}</td>
          <td>{formInput.fullname}</td>
          <td>{formInput.majority}</td>
          <td>{formInput.address}</td>
          <td>
            <button
              className="btn btn-secondary me-2"
              id={formInput.nim}
              onClick={handleEditCondition}
            >
              Edit
            </button>
            <button
              className="btn btn-danger"
              id={formInput.nim}
              onClick={() => {
                handleDeleteStudent(formInput.id);
              }}
            >
              Delete
            </button>
          </td>
        </tr>
      )}
    </>
  );
};

export default StudentRow;
