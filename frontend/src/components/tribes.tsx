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
  EditButton,
} from "react-admin";

export const TribesList = (props: any) => (
  <List {...props}>
    <Datagrid bulkActionButtons={false}>
      <TextField source="id" />
      <TextField source="name" />
      <TextField source="colour" />
      <EditButton />
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
