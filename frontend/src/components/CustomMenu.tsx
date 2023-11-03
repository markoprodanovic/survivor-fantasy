// CustomMenu.tsx
import * as React from "react";
import { Menu } from "react-admin";
import HomeIcon from "@mui/icons-material/Home";

const CustomMenu = (props: any) => {
  return (
    <Menu {...props}>
      <Menu.DashboardItem primaryText="Home" leftIcon={<HomeIcon />} />
      <Menu.ResourceItem name="tribes" />
      <Menu.ResourceItem name="players" />
      <Menu.ResourceItem name="episodes" />
      <Menu.ResourceItem name="users" />
    </Menu>
  );
};

export default CustomMenu;
