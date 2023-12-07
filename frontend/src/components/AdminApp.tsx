"use client";
import { Admin, Resource, radiantDarkTheme } from "react-admin";
import simpleRestProvider from "ra-data-simple-rest";
import { TribesList, TribesCreate, TribesEdit, TribesShow } from "./tribes";
import {
  PlayersList,
  PlayersCreate,
  PlayersEdit,
  PlayersShow,
} from "./players";
import { EpisodesCreate, EpisodesList } from "./episodes";
import { UsersList, UsersEdit, UsersShow } from "./users";
import { Home } from "./home";
import FlagIcon from "@mui/icons-material/Flag";
import GroupsIcon from "@mui/icons-material/Groups";
import VideocamIcon from "@mui/icons-material/Videocam";
import CustomLayout from "./CustomLayout";

const dataProvider = simpleRestProvider("http://localhost:9090/api/v1"); // Adjust the URL to your backend's API endpoint.

const theme = {
  ...radiantDarkTheme,
  palette: {
    mode: "dark",
  },
  typography: {
    fontFamily: ["Futura", "sans-serif"].join(","),
  },
  components: {
    ...radiantDarkTheme.components,
    MuiTextField: {
      defaultProps: {
        variant: "outlined" as const,
      },
    },
    MuiFormControl: {
      defaultProps: {
        variant: "outlined" as const,
      },
    },
  },
};

const AdminApp = ({ user, ...props }) => {
  return (
    <Admin
      dataProvider={dataProvider}
      theme={theme}
      dashboard={Home}
      layout={(props) => <CustomLayout user={user} {...props} />}
    >
      <Resource
        icon={FlagIcon}
        name="tribes"
        list={TribesList}
        create={TribesCreate}
        edit={TribesEdit}
        show={TribesShow}
      />
      <Resource
        icon={GroupsIcon}
        name="players"
        list={PlayersList}
        create={PlayersCreate}
        edit={PlayersEdit}
        show={PlayersShow}
      />
      <Resource
        icon={VideocamIcon}
        name="episodes"
        list={EpisodesList}
        create={EpisodesCreate}
      />
      <Resource
        icon={GroupsIcon}
        name="users"
        list={UsersList}
        edit={UsersEdit}
        show={UsersShow}
      />
    </Admin>
  );
};

export default AdminApp;
