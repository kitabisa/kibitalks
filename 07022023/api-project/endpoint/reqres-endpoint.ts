/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/no-var-requires */
import supertest from "supertest";
import env from "dotenv";

env.config();
const request = supertest(process.env.REQRES_URL);

export const createUser = (body: any) => request.post(`/users`).send(body);

export const getAllUsers = (param: any) => request.get(`/users`).query(param);

export const editUser = (path: any, body: any) =>
  request.put(`/users/${path}`).send(body);

export const deleteUser = (path: any) => request.del(`/users/${path}`);
