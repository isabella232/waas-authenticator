/* eslint-disable */
// sequence-waas-authenticator v0.1.0 4547675e59fbdc86616acaf3b39e8f19f897ddbc
// --
// Code generated by webrpc-gen@v0.14.0-dev with typescript generator. DO NOT EDIT.
//
// webrpc-gen -schema=authenticator.ridl -target=typescript -client -out=./clients/authenticator.gen.ts

// WebRPC description and code-gen version
export const WebRPCVersion = "v1"

// Schema version of your RIDL schema
export const WebRPCSchemaVersion = "v0.1.0"

// Schema hash generated from your RIDL schema
export const WebRPCSchemaHash = "4547675e59fbdc86616acaf3b39e8f19f897ddbc"

//
// Types
//


export enum IdentityType {
  None = 'None',
  Guest = 'Guest',
  OIDC = 'OIDC'
}

export interface Version {
  webrpcVersion: string
  schemaVersion: string
  schemaHash: string
  appVersion: string
}

export interface RuntimeStatus {
  healthOK: boolean
  startTime: string
  uptime: number
  ver: string
  branch: string
  commitHash: string
}

export interface RegisterSessionPayload {
  projectId: number
  idToken: string
  sessionAddress: string
  friendlyName: string
  intentJSON: string
}

export interface RefreshSessionPayload {
  sessionId: string
}

export interface DropSessionPayload {
  sessionId: string
  dropSessionId: string
}

export interface ListSessionsPayload {
  sessionId: string
}

export interface GetAddressPayload {
  sessionId: string
}

export interface SendIntentPayload {
  sessionId: string
  intentJson: string
}

export interface Chain {
  id: number
  name: string
  isEnabled: boolean
}

export interface Identity {
  type: IdentityType
  iss: string
  sub: string
  email: string
}

export interface OpenIdProvider {
  iss: string
  azp?: string
  aud?: string
}

export interface Tenant {
  projectId: number
  version: number
  oidcProviders: Array<OpenIdProvider>
  allowedOrigins: Array<string>
  updatedAt: string
}

export interface TenantData {
  proj: number
  prv: string
  addr: string
  salt: string
  seqctx: MiniSequenceContext
  upc: string
  wat: string
  oidc: Array<OpenIdProvider>
  transk: Array<string>
  sessk: Array<string>
  ogns: Array<string>
}

export interface MiniSequenceContext {
  factory: string
  mainModule: string
}

export interface AccountData {
  proj: number
  user: string
  ident: string
  iat: string
}

export interface Session {
  id: string
  address: string
  projectId: number
  userId: string
  identity: Identity
  friendlyName: string
  createdAt: string
  refreshedAt: string
  expiresAt: string
}

export interface SessionData {
  addr: string
  proj: number
  user: string
  ident: string
  iat: string
  exp: string
}

export interface Signature {
  session: string
  signature: string
}

export interface Intent {
  version: string
  packet: string
  signatures: Array<Signature>
}

export interface WaasAuthenticator {
  registerSession(args: RegisterSessionArgs, headers?: object, signal?: AbortSignal): Promise<RegisterSessionReturn>
  listSessions(args: ListSessionsArgs, headers?: object, signal?: AbortSignal): Promise<ListSessionsReturn>
  dropSession(args: DropSessionArgs, headers?: object, signal?: AbortSignal): Promise<DropSessionReturn>
  getAddress(args: GetAddressArgs, headers?: object, signal?: AbortSignal): Promise<GetAddressReturn>
  sendIntent(args: SendIntentArgs, headers?: object, signal?: AbortSignal): Promise<SendIntentReturn>
  chainList(headers?: object, signal?: AbortSignal): Promise<ChainListReturn>
}

export interface RegisterSessionArgs {
  encryptedPayloadKey: string
  payloadCiphertext: string
  payloadSig: string
}

export interface RegisterSessionReturn {
  session: Session
  data: any  
}
export interface ListSessionsArgs {
  encryptedPayloadKey: string
  payloadCiphertext: string
  payloadSig: string
}

export interface ListSessionsReturn {
  sessions: Array<Session>  
}
export interface DropSessionArgs {
  encryptedPayloadKey: string
  payloadCiphertext: string
  payloadSig: string
}

export interface DropSessionReturn {
  ok: boolean  
}
export interface GetAddressArgs {
  encryptedPayloadKey: string
  payloadCiphertext: string
  payloadSig: string
}

export interface GetAddressReturn {
  address: string  
}
export interface SendIntentArgs {
  encryptedPayloadKey: string
  payloadCiphertext: string
  payloadSig: string
}

export interface SendIntentReturn {
  code: string
  data: any  
}
export interface ChainListArgs {
}

export interface ChainListReturn {
  chains: Array<Chain>  
}

export interface WaasAuthenticatorV1 {
  registerSession(args: RegisterSessionArgs, headers?: object, signal?: AbortSignal): Promise<RegisterSessionReturn>
  sendIntent(args: SendIntentArgs, headers?: object, signal?: AbortSignal): Promise<SendIntentReturn>
  chainList(headers?: object, signal?: AbortSignal): Promise<ChainListReturn>
}

export interface RegisterSessionArgs {
  intent: Intent
  friendlyName: string
}

export interface RegisterSessionReturn {
  session: Session
  data: any  
}
export interface SendIntentArgs {
  intent: Intent
}

export interface SendIntentReturn {
  code: string
  data: any  
}
export interface ChainListArgs {
}

export interface ChainListReturn {
  chains: Array<Chain>  
}

export interface WaasAuthenticatorAdmin {
  version(headers?: object, signal?: AbortSignal): Promise<VersionReturn>
  runtimeStatus(headers?: object, signal?: AbortSignal): Promise<RuntimeStatusReturn>
  clock(headers?: object, signal?: AbortSignal): Promise<ClockReturn>
  getTenant(args: GetTenantArgs, headers?: object, signal?: AbortSignal): Promise<GetTenantReturn>
  createTenant(args: CreateTenantArgs, headers?: object, signal?: AbortSignal): Promise<CreateTenantReturn>
  updateTenant(args: UpdateTenantArgs, headers?: object, signal?: AbortSignal): Promise<UpdateTenantReturn>
}

export interface VersionArgs {
}

export interface VersionReturn {
  version: Version  
}
export interface RuntimeStatusArgs {
}

export interface RuntimeStatusReturn {
  status: RuntimeStatus  
}
export interface ClockArgs {
}

export interface ClockReturn {
  serverTime: string  
}
export interface GetTenantArgs {
  projectId: number
}

export interface GetTenantReturn {
  tenant: Tenant  
}
export interface CreateTenantArgs {
  projectId: number
  waasAccessToken: string
  oidcProviders: Array<OpenIdProvider>
  allowedOrigins: Array<string>
}

export interface CreateTenantReturn {
  tenant: Tenant
  upgradeCode: string  
}
export interface UpdateTenantArgs {
  projectId: number
  upgradeCode: string
  oidcProviders: Array<OpenIdProvider>
  allowedOrigins: Array<string>
}

export interface UpdateTenantReturn {
  tenant: Tenant  
}


  
//
// Client
//
export class WaasAuthenticator implements WaasAuthenticator {
  protected hostname: string
  protected fetch: Fetch
  protected path = '/rpc/WaasAuthenticator/'

  constructor(hostname: string, fetch: Fetch) {
    this.hostname = hostname
    this.fetch = (input: RequestInfo, init?: RequestInit) => fetch(input, init)
  }

  private url(name: string): string {
    return this.hostname + this.path + name
  }
  
  registerSession = (args: RegisterSessionArgs, headers?: object, signal?: AbortSignal): Promise<RegisterSessionReturn> => {
    return this.fetch(
      this.url('RegisterSession'),
      createHTTPRequest(args, headers, signal)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          session: <Session>(_data.session),
          data: <any>(_data.data),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
  listSessions = (args: ListSessionsArgs, headers?: object, signal?: AbortSignal): Promise<ListSessionsReturn> => {
    return this.fetch(
      this.url('ListSessions'),
      createHTTPRequest(args, headers, signal)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          sessions: <Array<Session>>(_data.sessions),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
  dropSession = (args: DropSessionArgs, headers?: object, signal?: AbortSignal): Promise<DropSessionReturn> => {
    return this.fetch(
      this.url('DropSession'),
      createHTTPRequest(args, headers, signal)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          ok: <boolean>(_data.ok),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
  getAddress = (args: GetAddressArgs, headers?: object, signal?: AbortSignal): Promise<GetAddressReturn> => {
    return this.fetch(
      this.url('GetAddress'),
      createHTTPRequest(args, headers, signal)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          address: <string>(_data.address),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
  sendIntent = (args: SendIntentArgs, headers?: object, signal?: AbortSignal): Promise<SendIntentReturn> => {
    return this.fetch(
      this.url('SendIntent'),
      createHTTPRequest(args, headers, signal)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          code: <string>(_data.code),
          data: <any>(_data.data),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
  chainList = (headers?: object, signal?: AbortSignal): Promise<ChainListReturn> => {
    return this.fetch(
      this.url('ChainList'),
      createHTTPRequest({}, headers, signal)
      ).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          chains: <Array<Chain>>(_data.chains),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
}
export class WaasAuthenticatorV1 implements WaasAuthenticatorV1 {
  protected hostname: string
  protected fetch: Fetch
  protected path = '/rpc/WaasAuthenticatorV1/'

  constructor(hostname: string, fetch: Fetch) {
    this.hostname = hostname
    this.fetch = (input: RequestInfo, init?: RequestInit) => fetch(input, init)
  }

  private url(name: string): string {
    return this.hostname + this.path + name
  }
  
  registerSession = (args: RegisterSessionArgs, headers?: object, signal?: AbortSignal): Promise<RegisterSessionReturn> => {
    return this.fetch(
      this.url('RegisterSession'),
      createHTTPRequest(args, headers, signal)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          session: <Session>(_data.session),
          data: <any>(_data.data),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
  sendIntent = (args: SendIntentArgs, headers?: object, signal?: AbortSignal): Promise<SendIntentReturn> => {
    return this.fetch(
      this.url('SendIntent'),
      createHTTPRequest(args, headers, signal)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          code: <string>(_data.code),
          data: <any>(_data.data),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
  chainList = (headers?: object, signal?: AbortSignal): Promise<ChainListReturn> => {
    return this.fetch(
      this.url('ChainList'),
      createHTTPRequest({}, headers, signal)
      ).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          chains: <Array<Chain>>(_data.chains),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
}
export class WaasAuthenticatorAdmin implements WaasAuthenticatorAdmin {
  protected hostname: string
  protected fetch: Fetch
  protected path = '/rpc/WaasAuthenticatorAdmin/'

  constructor(hostname: string, fetch: Fetch) {
    this.hostname = hostname
    this.fetch = (input: RequestInfo, init?: RequestInit) => fetch(input, init)
  }

  private url(name: string): string {
    return this.hostname + this.path + name
  }
  
  version = (headers?: object, signal?: AbortSignal): Promise<VersionReturn> => {
    return this.fetch(
      this.url('Version'),
      createHTTPRequest({}, headers, signal)
      ).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          version: <Version>(_data.version),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
  runtimeStatus = (headers?: object, signal?: AbortSignal): Promise<RuntimeStatusReturn> => {
    return this.fetch(
      this.url('RuntimeStatus'),
      createHTTPRequest({}, headers, signal)
      ).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          status: <RuntimeStatus>(_data.status),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
  clock = (headers?: object, signal?: AbortSignal): Promise<ClockReturn> => {
    return this.fetch(
      this.url('Clock'),
      createHTTPRequest({}, headers, signal)
      ).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          serverTime: <string>(_data.serverTime),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
  getTenant = (args: GetTenantArgs, headers?: object, signal?: AbortSignal): Promise<GetTenantReturn> => {
    return this.fetch(
      this.url('GetTenant'),
      createHTTPRequest(args, headers, signal)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          tenant: <Tenant>(_data.tenant),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
  createTenant = (args: CreateTenantArgs, headers?: object, signal?: AbortSignal): Promise<CreateTenantReturn> => {
    return this.fetch(
      this.url('CreateTenant'),
      createHTTPRequest(args, headers, signal)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          tenant: <Tenant>(_data.tenant),
          upgradeCode: <string>(_data.upgradeCode),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
  updateTenant = (args: UpdateTenantArgs, headers?: object, signal?: AbortSignal): Promise<UpdateTenantReturn> => {
    return this.fetch(
      this.url('UpdateTenant'),
      createHTTPRequest(args, headers, signal)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          tenant: <Tenant>(_data.tenant),
        }
      })
    }, (error) => {
      throw WebrpcRequestFailedError.new({ cause: `fetch(): ${error.message || ''}` })
    })
  }
  
}

  const createHTTPRequest = (body: object = {}, headers: object = {}, signal: AbortSignal | null = null): object => {
  return {
    method: 'POST',
    headers: { ...headers, 'Content-Type': 'application/json' },
    body: JSON.stringify(body || {}),
    signal
  }
}

const buildResponse = (res: Response): Promise<any> => {
  return res.text().then(text => {
    let data
    try {
      data = JSON.parse(text)
    } catch(error) {
      let message = ''
      if (error instanceof Error)  {
        message = error.message
      }
      throw WebrpcBadResponseError.new({
        status: res.status,
        cause: `JSON.parse(): ${message}: response text: ${text}`},
      )
    }
    if (!res.ok) {
      const code: number = (typeof data.code === 'number') ? data.code : 0
      throw (webrpcErrorByCode[code] || WebrpcError).new(data)
    }
    return data
  })
}

//
// Errors
//

export class WebrpcError extends Error {
  name: string
  code: number
  message: string
  status: number
  cause?: string

  /** @deprecated Use message instead of msg. Deprecated in webrpc v0.11.0. */
  msg: string

  constructor(name: string, code: number, message: string, status: number, cause?: string) {
    super(message)
    this.name = name || 'WebrpcError'
    this.code = typeof code === 'number' ? code : 0
    this.message = message || `endpoint error ${this.code}`
    this.msg = this.message
    this.status = typeof status === 'number' ? status : 0
    this.cause = cause
    Object.setPrototypeOf(this, WebrpcError.prototype)
  }

  static new(payload: any): WebrpcError {
    return new this(payload.error, payload.code, payload.message || payload.msg, payload.status, payload.cause)
  }
}

// Webrpc errors

export class WebrpcEndpointError extends WebrpcError {
  constructor(
    name: string = 'WebrpcEndpoint',
    code: number = 0,
    message: string = 'endpoint error',
    status: number = 0,
    cause?: string
  ) {
    super(name, code, message, status, cause)
    Object.setPrototypeOf(this, WebrpcEndpointError.prototype)
  }
}

export class WebrpcRequestFailedError extends WebrpcError {
  constructor(
    name: string = 'WebrpcRequestFailed',
    code: number = -1,
    message: string = 'request failed',
    status: number = 0,
    cause?: string
  ) {
    super(name, code, message, status, cause)
    Object.setPrototypeOf(this, WebrpcRequestFailedError.prototype)
  }
}

export class WebrpcBadRouteError extends WebrpcError {
  constructor(
    name: string = 'WebrpcBadRoute',
    code: number = -2,
    message: string = 'bad route',
    status: number = 0,
    cause?: string
  ) {
    super(name, code, message, status, cause)
    Object.setPrototypeOf(this, WebrpcBadRouteError.prototype)
  }
}

export class WebrpcBadMethodError extends WebrpcError {
  constructor(
    name: string = 'WebrpcBadMethod',
    code: number = -3,
    message: string = 'bad method',
    status: number = 0,
    cause?: string
  ) {
    super(name, code, message, status, cause)
    Object.setPrototypeOf(this, WebrpcBadMethodError.prototype)
  }
}

export class WebrpcBadRequestError extends WebrpcError {
  constructor(
    name: string = 'WebrpcBadRequest',
    code: number = -4,
    message: string = 'bad request',
    status: number = 0,
    cause?: string
  ) {
    super(name, code, message, status, cause)
    Object.setPrototypeOf(this, WebrpcBadRequestError.prototype)
  }
}

export class WebrpcBadResponseError extends WebrpcError {
  constructor(
    name: string = 'WebrpcBadResponse',
    code: number = -5,
    message: string = 'bad response',
    status: number = 0,
    cause?: string
  ) {
    super(name, code, message, status, cause)
    Object.setPrototypeOf(this, WebrpcBadResponseError.prototype)
  }
}

export class WebrpcServerPanicError extends WebrpcError {
  constructor(
    name: string = 'WebrpcServerPanic',
    code: number = -6,
    message: string = 'server panic',
    status: number = 0,
    cause?: string
  ) {
    super(name, code, message, status, cause)
    Object.setPrototypeOf(this, WebrpcServerPanicError.prototype)
  }
}

export class WebrpcInternalErrorError extends WebrpcError {
  constructor(
    name: string = 'WebrpcInternalError',
    code: number = -7,
    message: string = 'internal error',
    status: number = 0,
    cause?: string
  ) {
    super(name, code, message, status, cause)
    Object.setPrototypeOf(this, WebrpcInternalErrorError.prototype)
  }
}

export class WebrpcClientDisconnectedError extends WebrpcError {
  constructor(
    name: string = 'WebrpcClientDisconnected',
    code: number = -8,
    message: string = 'client disconnected',
    status: number = 0,
    cause?: string
  ) {
    super(name, code, message, status, cause)
    Object.setPrototypeOf(this, WebrpcClientDisconnectedError.prototype)
  }
}

export class WebrpcStreamLostError extends WebrpcError {
  constructor(
    name: string = 'WebrpcStreamLost',
    code: number = -9,
    message: string = 'stream lost',
    status: number = 0,
    cause?: string
  ) {
    super(name, code, message, status, cause)
    Object.setPrototypeOf(this, WebrpcStreamLostError.prototype)
  }
}

export class WebrpcStreamFinishedError extends WebrpcError {
  constructor(
    name: string = 'WebrpcStreamFinished',
    code: number = -10,
    message: string = 'stream finished',
    status: number = 0,
    cause?: string
  ) {
    super(name, code, message, status, cause)
    Object.setPrototypeOf(this, WebrpcStreamFinishedError.prototype)
  }
}


// Schema errors

export class UnauthorizedError extends WebrpcError {
  constructor(
    name: string = 'Unauthorized',
    code: number = 1000,
    message: string = 'Unauthorized access',
    status: number = 0,
    cause?: string
  ) {
    super(name, code, message, status, cause)
    Object.setPrototypeOf(this, UnauthorizedError.prototype)
  }
}


export enum errors {
  WebrpcEndpoint = 'WebrpcEndpoint',
  WebrpcRequestFailed = 'WebrpcRequestFailed',
  WebrpcBadRoute = 'WebrpcBadRoute',
  WebrpcBadMethod = 'WebrpcBadMethod',
  WebrpcBadRequest = 'WebrpcBadRequest',
  WebrpcBadResponse = 'WebrpcBadResponse',
  WebrpcServerPanic = 'WebrpcServerPanic',
  WebrpcInternalError = 'WebrpcInternalError',
  WebrpcClientDisconnected = 'WebrpcClientDisconnected',
  WebrpcStreamLost = 'WebrpcStreamLost',
  WebrpcStreamFinished = 'WebrpcStreamFinished',
  Unauthorized = 'Unauthorized',
}

const webrpcErrorByCode: { [code: number]: any } = {
  [0]: WebrpcEndpointError,
  [-1]: WebrpcRequestFailedError,
  [-2]: WebrpcBadRouteError,
  [-3]: WebrpcBadMethodError,
  [-4]: WebrpcBadRequestError,
  [-5]: WebrpcBadResponseError,
  [-6]: WebrpcServerPanicError,
  [-7]: WebrpcInternalErrorError,
  [-8]: WebrpcClientDisconnectedError,
  [-9]: WebrpcStreamLostError,
  [-10]: WebrpcStreamFinishedError,
  [1000]: UnauthorizedError,
}

export type Fetch = (input: RequestInfo, init?: RequestInit) => Promise<Response>

