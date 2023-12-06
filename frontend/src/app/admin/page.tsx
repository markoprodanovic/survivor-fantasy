import { NextPage } from "next";
import dynamic from "next/dynamic";

const DynamicAdminApp = dynamic(() => import("../../components/AdminApp"), {
  ssr: false,
  loading: () => <p>Loading ...</p>,
});

const Admin: NextPage = () => <DynamicAdminApp />;

export default Admin;
