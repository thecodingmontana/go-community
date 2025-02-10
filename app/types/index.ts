import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'

export const sendUniqueCodeSchema = toTypedSchema(
  z.object({
    email: z.string().email('Email is invalid!'),
  }),
)

export const signinFormSchema = toTypedSchema(
  z.object({
    code: z
      .string()
      .max(8, {
        message: 'Code shouldn\'t be more than 8 characters.',
      })
      .min(8, {
        message: 'Code shouldn\'t be less than 8 characters.',
      }),
  }),
)

export interface User {
  id: string
  email: string
  username: string
  emailVerified: boolean
  avatar: string
}

export type ModalType = 'createWorkspace'
  | 'signout' | 'inviteMember' | 'mobileMenu'

export interface ModalData {
  chat?: {
    name: string
    id: string
  }
}

export interface ModalStore {
  type: ModalType | null
  isOpen: boolean
  data: ModalData
}

export const githubUrl = 'https://github.com/thecodingmontana/go-community.git'
export const twitterUrl = 'https://x.com/codewithmontana'

export interface SelectedFile {
  name: string
  fileUrl: string
  size: string
}

export interface SelectedImage {
  name: string
  imageUrl: string
  size: string
}
