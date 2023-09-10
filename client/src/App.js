import React from "react";
import { useQuery } from "react-query";
import { API } from "./config/api";
import StudentRow from "./components/StudentRow";
import AddStudent from "./components/AddStudent";

const App = () => {
  const {
    data: students,
    // isLoading: studentsLoading,
    refetch: refetchStudents,
  } = useQuery("studentsCache", async () => {
    try {
      const response = await API.get(`/students`);
      return response.data.data;
    } catch (e) {
      console.log(e);
    }
  });

  return (
    <div className="container">
      <div className="row mt-5">
        <div className="col">
          <h1 className="text-center">Tabel of Students</h1>
          <table className="table mt-3">
            <thead>
              <tr>
                <th>NIM</th>
                <th>Fullname</th>
                <th>Majority</th>
                <th>Address</th>
                <th></th>
              </tr>
            </thead>
            <tbody className="table-group-divider">
              {students?.map((student) => {
                return <StudentRow key={student.nim} student={student} />;
              })}
              <AddStudent refetchStudents={refetchStudents} />
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default App;
