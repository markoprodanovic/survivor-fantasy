import {
  List,
  Datagrid,
  TextField,
  Create,
  SimpleForm,
  NumberInput,
  DateField,
  DateInput,
  ArrayInput,
  SimpleFormIterator,
  ReferenceInput,
  SelectInput,
  DateTimeInput,
} from "react-admin";

export const EpisodesList = (props: any) => (
  <List {...props}>
    <Datagrid bulkActionButtons={false}>
      <TextField source="episode_number" />
      <DateField source="episode_date" />
    </Datagrid>
  </List>
);

export const EpisodesCreate = (props: any) => (
  <Create {...props}>
    <SimpleForm>
      <NumberInput source="episode_number" />
      <DateTimeInput source="episode_date" />
      <ArrayInput source="points">
        <SimpleFormIterator>
          <ReferenceInput label="Player" source="castId" reference="players">
            <SelectInput optionText="first_name" />
          </ReferenceInput>
          <NumberInput source="points" />
        </SimpleFormIterator>
      </ArrayInput>
    </SimpleForm>
  </Create>
);
