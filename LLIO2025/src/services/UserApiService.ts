import type { UserInfo } from '../Models/index';
import { GET, PATCH } from '../ts/server';

const getUserInfo = async (): Promise<UserInfo> => {
  try {
    const response = await GET<UserInfo>('/user/me');
    return response;
  } catch (error) {
    console.error("Erreur lors de la récupération des informations de l'utilisateur:", error);
    throw error;
  }
};

const getUsers = async (): Promise<UserInfo[]> => {
  try {
    const response = await GET<UserInfo[]>('/users');
    return response;
  }
  catch (error) {
    console.error("Erreur lors de la récupération des utilisateurs:", error);
    throw error;
  }
};

const updateUserRole = async (userId: number, role: number): Promise<void> => {
  try {
    const response = await PATCH<{ role: number }>(
      `/user/${userId}/role`,
      { role }
    );
    return response;
  } catch (error) {
    console.error("Erreur lors de la mise à jour du rôle de l'utilisateur:", error);
    throw error;
  }
}

export const UserApiService = {
  getUserInfo,
  getUsers,
  updateUserRole,
};
