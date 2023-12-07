import {
  List,
  Datagrid,
  TextField,
  Create,
  SimpleForm,
  TextInput,
  Show,
  SimpleShowLayout,
  Edit,
  EditButton,
  ReferenceArrayField,
  SingleFieldList,
  ChipField,
  ImageField,
  // ReferenceArrayInput,
  // AutocompleteArrayInput,
} from "react-admin";

export const UsersList = (props: any) => (
  <List {...props}>
    <Datagrid bulkActionButtons={false}>
      <TextField source="id" />
      <ImageField
        source="image"
        sx={{
          "& img": {
            maxWidth: 35,
            maxHeight: 35,
            objectFit: "contain",
            borderRadius: 10,
          },
        }}
      />
      <TextField source="name" />
      <ReferenceArrayField
        label="Picks"
        reference="players"
        source="player_ids"
      >
        <SingleFieldList>
          <ChipField source="first_name" />
        </SingleFieldList>
      </ReferenceArrayField>
      <TextField source="email" />
      <EditButton />
    </Datagrid>
  </List>
);

export const UsersEdit = (props: any) => (
  <Edit {...props}>
    <SimpleForm>
      <TextInput source="name" />
      <TextInput source="email" />
    </SimpleForm>
  </Edit>
);

export const UsersShow = (props: any) => (
  <Show {...props}>
    <SimpleShowLayout>
      <TextField source="id" />
      <TextField source="name" />
      <TextField source="email" />
    </SimpleShowLayout>
  </Show>
);
