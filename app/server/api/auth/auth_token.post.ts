export default defineEventHandler(async (event) => {
  try {
    const { auth_token, expires_at } = await readBody(event)

    if (typeof auth_token !== 'string' || typeof expires_at !== 'number') {
      throw createError({
        statusMessage: 'Invalid credentials!',
        statusCode: 400,
      })
    }

    const expires = new Date(expires_at * 1000)

    setCookie(event, 'auth_token', auth_token, {
      expires,
      path: '/',
      secure: process.env.NODE_ENV === 'production',
      sameSite: 'strict',
      maxAge: Math.floor((expires.getTime() - Date.now()) / 1000),
    })

    return {
      message: 'Authentication successful!',
    }
  }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  catch (error: any) {
    throw createError({
      message: `Failed to signin: ${error.message}`,
      statusCode: error.statusCode ? error.statusCode : 500,
    })
  }
})
