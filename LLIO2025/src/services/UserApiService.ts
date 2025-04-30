import type { User, UserInfo } from '../Models/index';
import { GET } from '../ts/server';

interface UsersResponse {
  users: User[];
}

const getUserInfo = async (): Promise<UserInfo> => {
  try {
    const response = await GET<UserInfo>('/user/me');
    return response;
  } catch (error) {
    console.error("Erreur lors de la récupération des informations de l'utilisateur:", error);
    throw error;
  }
};

const getAllUsers = async (): Promise<User[]> => {
  try {
    const response = await GET<UsersResponse>("/users");
    return response.users;
  } catch (error) {
    console.error("Erreur lors de la récupération des utilisateurs:", error);
    throw new Error("Échec de la récupération des utilisateurs. Veuillez réessayer.");
  }
};

const getManagerUsers = async (): Promise<User[]> => {
  try {
    const response = await GET<User[]>("/users?role=1");
    return response;
  } catch (error) {
    console.error("Erreur lors de la récupération des managers:", error);
    throw new Error("Échec de la récupération des managers. Veuillez réessayer.");
  }
};

export const UserApiService = {
  getAllUsers,
  getManagerUsers,
  getUserInfo,
};
