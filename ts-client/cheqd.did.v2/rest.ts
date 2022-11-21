/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

export interface V2DidDoc {
  context?: string[];
  id?: string;
  controller?: string[];
  verification_method?: V2VerificationMethod[];
  authentication?: string[];
  assertion_method?: string[];
  capability_invocation?: string[];
  capability_delegation?: string[];
  key_agreement?: string[];
  service?: V2Service[];
  also_known_as?: string[];
}

export interface V2DidDocWithMetadata {
  did_doc?: V2DidDoc;
  metadata?: V2Metadata;
}

export interface V2Metadata {
  created?: string;
  updated?: string;
  deactivated?: boolean;
  version_id?: string;
  next_version_id?: string;
  previous_version_id?: string;
}

export interface V2MsgCreateDidDocPayload {
  context?: string[];
  id?: string;
  controller?: string[];
  verification_method?: V2VerificationMethod[];
  authentication?: string[];
  assertion_method?: string[];
  capability_invocation?: string[];
  capability_delegation?: string[];
  key_agreement?: string[];
  also_known_as?: string[];
  service?: V2Service[];
}

export interface V2MsgCreateDidDocResponse {
  value?: V2DidDocWithMetadata;
}

export interface V2MsgDeactivateDidDocPayload {
  id?: string;
  version_id?: string;
}

export interface V2MsgDeactivateDidDocResponse {
  value?: V2DidDocWithMetadata;
}

export interface V2MsgUpdateDidDocPayload {
  context?: string[];
  id?: string;
  controller?: string[];
  verification_method?: V2VerificationMethod[];
  authentication?: string[];
  assertion_method?: string[];
  capability_invocation?: string[];
  capability_delegation?: string[];
  key_agreement?: string[];
  also_known_as?: string[];
  service?: V2Service[];
  version_id?: string;
}

export interface V2MsgUpdateDidDocResponse {
  value?: V2DidDocWithMetadata;
}

export interface V2QueryGetDidDocResponse {
  value?: V2DidDocWithMetadata;
}

export interface V2Service {
  id?: string;
  type?: string;
  service_endpoint?: string[];
}

export interface V2SignInfo {
  verification_method_id?: string;

  /** @format byte */
  signature?: string;
}

export interface V2VerificationMethod {
  id?: string;
  type?: string;
  controller?: string;
  verification_material?: string;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title cheqd/did/v2/diddoc.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryDidDoc
   * @request GET:/cheqd/did/v2/diddoc/{id}
   */
  queryDidDoc = (id: string, params: RequestParams = {}) =>
    this.request<V2QueryGetDidDocResponse, RpcStatus>({
      path: `/cheqd/did/v2/diddoc/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
