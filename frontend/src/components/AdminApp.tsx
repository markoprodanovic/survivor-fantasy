"use client";
import { Admin, Resource } from "react-admin";
import simpleRestProvider from "ra-data-simple-rest";
import { TribesList, TribesCreate, TribesEdit, TribesShow } from "./tribes";
import GroupsIcon from "@mui/icons-material/Groups";

const dataProvider = simpleRestProvider("http://localhost:9090/api/v1"); // Adjust the URL to your backend's API endpoint.

const AdminApp = () => {
  return (
    <Admin dataProvider={dataProvider}>
      <Resource
        icon={GroupsIcon}
        name="tribes"
        list={TribesList}
        create={TribesCreate}
        edit={TribesEdit}
        show={TribesShow}
      />
    </Admin>
  );
};

export default AdminApp;
