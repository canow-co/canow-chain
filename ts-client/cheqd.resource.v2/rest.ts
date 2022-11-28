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

export interface Cheqdresourcev2Metadata {
  collection_id?: string;
  id?: string;
  name?: string;
  version?: string;
  resource_type?: string;
  also_known_as?: string[];
  media_type?: string;
  created?: string;

  /** @format byte */
  checksum?: string;
  previous_version_id?: string;
  next_version_id?: string;
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

export interface V2MsgCreateResourcePayload {
  /** @format byte */
  data?: string;
  collection_id?: string;
  id?: string;
  name?: string;
  version?: string;
  resource_type?: string;
  also_known_as?: string[];
}

export interface V2MsgCreateResourceResponse {
  resource?: Cheqdresourcev2Metadata;
}

export interface V2QueryGetCollectionResourcesResponse {
  resources?: Cheqdresourcev2Metadata[];
}

export interface V2QueryGetResourceMetadataResponse {
  resource?: Cheqdresourcev2Metadata;
}

export interface V2QueryGetResourceResponse {
  resource?: V2ResourceWithMetadata;
}

export interface V2Resource {
  /** @format byte */
  data?: string;
}

export interface V2ResourceWithMetadata {
  resource?: V2Resource;
  metadata?: Cheqdresourcev2Metadata;
}

export interface V2SignInfo {
  verification_method_id?: string;

  /** @format byte */
  signature?: string;
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
 * @title cheqd/resource/v2/fee.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryCollectionResources
   * @request GET:/cheqd/resource/v2/collection/{collection_id}/resources
   */
  queryCollectionResources = (collectionId: string, params: RequestParams = {}) =>
    this.request<V2QueryGetCollectionResourcesResponse, RpcStatus>({
      path: `/cheqd/resource/v2/collection/${collectionId}/resources`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryResource
   * @request GET:/cheqd/resource/v2/collection/{collection_id}/resources/{id}
   */
  queryResource = (collectionId: string, id: string, params: RequestParams = {}) =>
    this.request<V2QueryGetResourceResponse, RpcStatus>({
      path: `/cheqd/resource/v2/collection/${collectionId}/resources/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryResourceMetadata
   * @request GET:/cheqd/resource/v2/collection/{collection_id}/resources/{id}/metadata
   */
  queryResourceMetadata = (collectionId: string, id: string, params: RequestParams = {}) =>
    this.request<V2QueryGetResourceMetadataResponse, RpcStatus>({
      path: `/cheqd/resource/v2/collection/${collectionId}/resources/${id}/metadata`,
      method: "GET",
      format: "json",
      ...params,
    });
}
