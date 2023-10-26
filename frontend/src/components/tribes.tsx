import {
  List,
  Datagrid,
  TextField,
  Create,
  SimpleForm,
  TextInput,
  Edit,
  Show,
  SimpleShowLayout,
  SelectInput,
} from "react-admin";

export const TribesList = (props: any) => (
  <List {...props}>
    <Datagrid rowClick="edit">
      <TextField source="id" />
      <TextField source="name" />
      <TextField source="colour" />
    </Datagrid>
  </List>
);

export const TribesCreate = (props: any) => (
  <Create {...props}>
    <SimpleForm>
      <TextInput source="name" />
      <SelectInput
        source="colour"
        choices={[
          { id: "red", name: "Red" },
          { id: "yellow", name: "Yellow" },
          { id: "blue", name: "Blue" },
        ]}
      />
    </SimpleForm>
  </Create>
);

export const TribesEdit = (props: any) => (
  <Edit {...props}>
    <SimpleForm>
      <TextField source="id" disabled />
      <TextInput source="name" />
      <SelectInput
        source="colour"
        choices={[
          { id: "red", name: "Red" },
          { id: "yellow", name: "Yellow" },
          { id: "blue", name: "Blue" },
        ]}
      />
    </SimpleForm>
  </Edit>
);

export const TribesShow = (props: any) => (
  <Show {...props}>
    <SimpleShowLayout>
      <TextField source="id" />
      <TextField source="name" />
      <TextField source="colour" />
    </SimpleShowLayout>
  </Show>
);
