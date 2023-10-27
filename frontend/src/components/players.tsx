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
  NumberInput,
  ReferenceInput,
  ReferenceField,
  BooleanInput,
  BooleanField,
  EditButton,
  NumberField,
  required,
} from "react-admin";

export const PlayersList = (props: any) => (
  <List {...props}>
    <Datagrid bulkActionButtons={false}>
      <TextField source="id" />
      <TextField source="first_name" />
      <TextField source="last_name" />
      <NumberField source="age" />
      <ReferenceField label="Tribe" source="tribe_id" reference="tribes">
        <TextField source="name" />
      </ReferenceField>
      <BooleanField source="eliminated" />
      <EditButton />
    </Datagrid>
  </List>
);

export const PlayersCreate = (props: any) => (
  <Create {...props}>
    <SimpleForm>
      <TextInput source="first_name" />
      <TextInput source="last_name" />
      <NumberInput source="age" />
      <ReferenceInput
        label="Tribe"
        source="tribe_id"
        reference="tribes"
        validate={[required()]}
      >
        <SelectInput optionText="name" validate={[required()]} />
      </ReferenceInput>
      <BooleanInput source="eliminated" />
    </SimpleForm>
  </Create>
);

export const PlayersEdit = (props: any) => (
  <Edit {...props}>
    <SimpleForm>
      <TextInput source="first_name" />
      <TextInput source="last_name" />
      <NumberInput source="age" />
      <ReferenceInput label="Tribe" source="tribe_id" reference="tribes">
        <SelectInput optionText="name" validate={[required()]} />
      </ReferenceInput>
      <BooleanInput source="eliminated" />
    </SimpleForm>
  </Edit>
);

export const PlayersShow = (props: any) => (
  <Show {...props}>
    <SimpleShowLayout>
      <TextField source="id" />
      <TextField source="first_name" />
      <TextField source="last_name" />
      <NumberField source="age" />
      <ReferenceField label="Tribe" source="tribe_id" reference="tribes">
        <TextField source="name" />
      </ReferenceField>
      <BooleanField source="eliminated" />
    </SimpleShowLayout>
  </Show>
);
