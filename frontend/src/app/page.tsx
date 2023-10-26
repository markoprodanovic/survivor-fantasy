import { NextPage } from "next";
import AdminApp from "@/components/AdminApp";
import dynamic from "next/dynamic";

const DynamicAdminApp = dynamic(() => import("../components/AdminApp"), {
  ssr: false,
  loading: () => <p>Loading ...</p>,
});

const Home: NextPage = () => <DynamicAdminApp />;

export default Home;
