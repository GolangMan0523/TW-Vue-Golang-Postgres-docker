import axios from '../../utils/axios'

export async function resetPassword(
  email: string,
): Promise<void> {
  try {
    return axios.post('/auth/reset', {email})
  } catch (error) {
    throw error
  }
}