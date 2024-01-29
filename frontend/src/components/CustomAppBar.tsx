import React from "react";
import { AppBar, UserMenu, useUserMenu } from "react-admin";
import { MenuItem, ListItemIcon, ListItemText } from "@mui/material";
import SettingsIcon from "@mui/icons-material/Settings";
import Avatar from "@mui/material/Avatar";
import { signOut, useSession } from "next-auth/react";

const CustomLogoutButton = React.forwardRef((props, ref) => {
  const { onClose } = useUserMenu();

  return (
    // FIXME Hacky way to get TS to calm down
    // @ts-ignore
    <MenuItem
      onClick={() => {
        onClose();
        signOut({ callbackUrl: "/" });
      }}
      ref={ref}
      // It's important to pass the props to allow Material UI to manage the keyboard navigation
      {...props}
    >
      <ListItemIcon>
        <SettingsIcon fontSize="small" />
      </ListItemIcon>
      <ListItemText>Sign Out</ListItemText>
    </MenuItem>
  );
});
CustomLogoutButton.displayName = "CustomLogoutButton";

const CustomProfileIcon = ({ image }) => (
  <Avatar
    sx={{
      height: 30,
      width: 30,
    }}
    src={image}
  />
);

const CustomAppBar = () => {
  const { data: session } = useSession();

  return (
    <AppBar
      color="primary"
      position="fixed"
      userMenu={
        <UserMenu icon={<CustomProfileIcon image={session.user.image} />}>
          <CustomLogoutButton />
        </UserMenu>
      }
    />
  );
};

export default CustomAppBar;
