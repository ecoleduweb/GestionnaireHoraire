import type { UserInfo } from '../Models/index';
import { GET } from '../ts/server';

const getUserInfo = async (): Promise<UserInfo> => {
  try {
    const response = await GET<UserInfo>('/user/me');
    return response;
  } catch (error) {
    console.error("Erreur lors de la récupération des informations de l'utilisateur:", error);
    throw error;
  }
};

export const UserApiService = {
  getAllUsers,
  getManagerUsers,
  getUserInfo,
  getUsers,
  updateUserRole,
};
