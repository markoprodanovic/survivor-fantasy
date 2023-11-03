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
  ReferenceArrayInput,
  AutocompleteArrayInput,
} from "react-admin";

export const UsersList = (props: any) => (
  <List {...props}>
    <Datagrid bulkActionButtons={false}>
      <TextField source="id" />
      <TextField source="first_name" />
      <TextField source="last_name" />
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

export const UsersCreate = (props: any) => (
  <Create {...props}>
    <SimpleForm>
      <TextInput source="first_name" />
      <TextInput source="last_name" />
      <TextInput source="email" />
      <ReferenceArrayInput source="player_ids" reference="players">
        <AutocompleteArrayInput optionText="first_name" label="Players" />
      </ReferenceArrayInput>
    </SimpleForm>
  </Create>
);

export const UsersEdit = (props: any) => (
  <Edit {...props}>
    <SimpleForm>
      <TextInput source="first_name" />
      <TextInput source="last_name" />
      <TextInput source="email" />
    </SimpleForm>
  </Edit>
);

export const UsersShow = (props: any) => (
  <Show {...props}>
    <SimpleShowLayout>
      <TextField source="id" />
      <TextField source="first_name" />
      <TextField source="last_name" />
      <TextField source="email" />
    </SimpleShowLayout>
  </Show>
);
