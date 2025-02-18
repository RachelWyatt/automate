import { HttpErrorResponse } from '@angular/common/http';
import { Action } from '@ngrx/store';
import { Client } from './client.model';

export enum ClientActionTypes {
  GET_ALL = 'CLIENTS::GET_ALL',
  GET_ALL_SUCCESS = 'CLIENTS::GET_ALL::SUCCESS',
  GET_ALL_FAILURE = 'CLIENTS::GET_ALL::FAILURE',
  GET = 'CLIENTS::GET',
  GET_SUCCESS = 'CLIENTS::GET::SUCCESS',
  GET_FAILURE = 'CLIENTS::GET::FAILURE',
  DELETE          = 'CLIENTS::DELETE',
  DELETE_SUCCESS  = 'CLIENTS::DELETE::SUCCESS',
  DELETE_FAILURE  = 'CLIENTS::DELETE::FAILURE',
  CREATE = 'CLIENTS::CREATE',
  CREATE_SUCCESS = 'CLIENTS::CREATE::SUCCESS',
  CREATE_FAILURE = 'CLIENTS::CREATE::FAILURE'
}

export interface CreateClientSuccessPayload {
  name: string;
  client_key: {
    name: string,
    public_key: string,
    expiration_date: string,
    private_key: string
  };
}

export interface ClientsSuccessPayload {
  clients: Client[];
  total: number;
}

export interface ClientsPayload {
  clientName: string;
  org_id: string;
  page: number;
  per_page: number;
  server_id: string;
}

export class GetClients implements Action {
  readonly type = ClientActionTypes.GET_ALL;
  constructor(public payload: ClientsPayload) { }
}

export class GetClientsSuccess implements Action {
  readonly type = ClientActionTypes.GET_ALL_SUCCESS;
  constructor(public payload: ClientsSuccessPayload) { }
}

export class GetClientsFailure implements Action {
  readonly type = ClientActionTypes.GET_ALL_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}


export class GetClient implements Action {
  readonly type = ClientActionTypes.GET;
  constructor(public payload: { server_id: string, org_id: string, name: string }) { }
}

export class GetClientSuccess implements Action {
  readonly type = ClientActionTypes.GET_SUCCESS;
  constructor(public payload: Client) { }
}

export class GetClientFailure implements Action {
  readonly type = ClientActionTypes.GET_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export interface CreateClientPayload {
  name: string;
  validator: boolean;
  org_id: string;
  server_id: string;
  create_key: boolean;
}

export class CreateClient implements Action {
  readonly type = ClientActionTypes.CREATE;
  constructor(public payload: CreateClientPayload) { }
}

export class CreateClientSuccess implements Action {
  readonly type = ClientActionTypes.CREATE_SUCCESS;
  constructor(public payload: CreateClientSuccessPayload) { }
}

export class CreateClientFailure implements Action {
  readonly type = ClientActionTypes.CREATE_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export class DeleteClient implements Action {
  readonly type = ClientActionTypes.DELETE;
  constructor(public payload: { server_id: string, org_id: string, name: string }) { }
}

export class DeleteClientSuccess implements Action {
  readonly type = ClientActionTypes.DELETE_SUCCESS;
  constructor(public payload: { name: string }) { }
}

export class DeleteClientFailure implements Action {
  readonly type = ClientActionTypes.DELETE_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export type ClientActions =
  | GetClients
  | GetClientsSuccess
  | GetClientsFailure
  | GetClient
  | GetClientSuccess
  | GetClientFailure
  | CreateClient
  | CreateClientSuccess
  | CreateClientFailure
  | DeleteClient
  | DeleteClientSuccess
  | DeleteClientFailure;
