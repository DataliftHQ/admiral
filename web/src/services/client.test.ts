import { type AxiosError } from 'axios'

import client, { errorInterceptor } from './client'
import type { AdmiralError } from './errors'

describe('error interceptor', () => {
  describe('on axios error', () => {
    let err: Promise<AdmiralError>
    beforeAll(() => {
      err = errorInterceptor({
        message: 'Request timeout of 1ms reached',
        isAxiosError: true,
      } as AxiosError)
    })

    it('returns a AdmiralError', async () => {
      await expect(err).rejects.toEqual({
        status: {
          code: 500,
          text: 'Client Error',
        },
        message: 'Request timeout of 1ms reached',
      })
    })
  })

  describe('on auth error', () => {
    let window: Window & typeof globalThis
    const axiosError = {
      response: {
        status: 401,
        statusText: 'Not Authorized',
        data: {
          code: 16,
          message: 'Whoops!',
        },
      },
    } as AxiosError

    beforeAll(() => {
      window = global.window
      global.window = Object.create(window)
      Object.defineProperty(window, 'location', {
        value: {
          href: '/example?foo=bar',
          pathname: '/example',
          search: '?foo=bar',
        },
        writable: true,
      })

      errorInterceptor(axiosError).catch((_) => _)
    })

    afterAll(() => {
      global.window = window
    })

    it('redirects to provided url', () => {
      expect(window.location.href).toBe('/auth/login?redirect_url=%2Fexample%3Ffoo%3Dbar')
    })
  })

  describe('on known error', () => {
    const axiosError = {
      response: {
        status: 404,
        statusText: 'Not Found',
        data: {
          code: 5,
          message: 'Could not find resource',
        },
      },
    } as AxiosError
    let err: Promise<AdmiralError>
    beforeAll(() => {
      err = errorInterceptor(axiosError)
    })

    it('returns a AdmiralError', async () => {
      await expect(err).rejects.toEqual({
        code: 5,
        message: 'Could not find resource',
        status: {
          code: 404,
          text: 'Not Found',
        },
      })
    })
  })

  describe('on known error with details', () => {
    const axiosError = {
      response: {
        status: 400,
        statusText: 'Invalid Argument',
        data: {
          code: 3,
          message: 'Invalid Input',
          details: [
            {
              '@type': 'type.googleapis.com/google.rpc.BadRequest',
              fieldViolations: [
                {
                  field: 'SomeRequest.email_address',
                  description: 'INVALID_EMAIL_ADDRESS',
                },
                {
                  field: 'SomeRequest.username',
                  description: 'INVALID_USER_NAME',
                },
              ],
            },
          ],
        },
      },
    } as AxiosError
    let err: Promise<AdmiralError>
    beforeAll(() => {
      err = errorInterceptor(axiosError)
    })

    it('returns a DataliftError', async () => {
      await expect(err).rejects.toEqual({
        code: 3,
        message: 'Invalid Input',
        status: {
          code: 400,
          text: 'Invalid Argument',
        },
        details: [
          {
            _type: 'type.googleapis.com/google.rpc.BadRequest',
            fieldViolations: [
              {
                field: 'SomeRequest.email_address',
                description: 'INVALID_EMAIL_ADDRESS',
              },
              {
                field: 'SomeRequest.username',
                description: 'INVALID_USER_NAME',
              },
            ],
          },
        ],
      })
    })
  })

  describe('on unknown error', () => {
    let window: Window & typeof globalThis
    let err: Promise<AdmiralError>
    beforeAll(() => {
      window = global.window
      global.window = Object.create(window)
      Object.defineProperty(window, 'location', {
        value: {
          href: '/example?foo=bar',
          pathname: '/example',
          search: '?foo=bar',
        },
        writable: true,
      })

      err = errorInterceptor({
        isAxiosError: false,
        message: 'Unauthorized to perform action',
        response: {
          status: 401,
          statusText: 'Unauthenticated',
          data: {},
        },
        name: 'foobar',
      } as AxiosError)
    })

    afterAll(() => {
      global.window = window
    })

    it('returns a DataliftError', async () => {
      await expect(err).rejects.toEqual({
        data: {},
        message: 'Unauthorized to perform action',
        status: {
          code: 401,
          text: 'Unauthenticated',
        },
      })
    })
  })
})

describe('axios client', () => {
  it('treats status codes >= 400 as error', () => {
    expect(client.defaults.validateStatus?.(400)).toBe(false)
  })

  it('treats status codes >= 500 as error', () => {
    expect(client.defaults.validateStatus?.(500)).toBe(false)
  })

  it('treats status codes < 400 as success', () => {
    expect(client.defaults.validateStatus?.(399)).toBe(true)
  })
})
