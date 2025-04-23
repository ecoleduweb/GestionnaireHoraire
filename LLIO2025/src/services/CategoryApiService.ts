// src/services/CategoryApiService.ts
import type { Category } from "../Models/index";
import { GET } from "../ts/server";

interface CategoriesResponse {
  categories: Category[];
}

// Récupérer toutes les catégories
const getAllCategories = async (): Promise<Category[]> => {
  try {
    const response = await GET<CategoriesResponse>("/categories");
    return response.categories;
  } catch (error) {
    console.error("Erreur lors de la récupération des catégories:", error);
    throw error;
  }
};

export const CategoryApiService = {
  getAllCategories
};