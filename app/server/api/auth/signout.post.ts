export default defineEventHandler(async (event) => {
  try {
    setCookie(event, 'auth_token', '')
    setCookie(event, 'modalStore', '')
    return {
      message: 'Signout successful!',
    }
  }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  catch (error: any) {
    throw createError({
      message: `Failed to signout: ${error.message}`,
      statusCode: error.statusCode ? error.statusCode : 500,
    })
  }
})
