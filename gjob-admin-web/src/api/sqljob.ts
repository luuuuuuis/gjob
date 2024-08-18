import { http } from "@/utils/http";

// params传参
export const getAllSqlJobs = () => {
  return http.request("get", "/api/job/sql/getall");
};
