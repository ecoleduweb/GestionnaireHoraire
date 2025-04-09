import type { Activity } from 'lucide-svelte';
import * as yup from 'yup';

const ActivityValidationSchema = yup.object().shape({
  // Validation des informations de l'activité
  name: yup
    .string()
    .max(50, "Le nom de l'activité ne doit pas dépasser 50 caractères")
    .required("Le nom de l'activité est requis"),
  description: yup
    .string()
    .max(100, "La description de l'activité ne doit pas dépasser 100 caractères"),

  // Validation des information utilisateur, projet et catégorie
  userId: yup
    .number() 
    .typeError('Veuillez sélectionner un utilisateur'),
  projectId: yup
    .number()
    .min(1, 'Veuillez sélectionner un projet')
    .typeError('Veuillez sélectionner un projet')
    .required('Veuillez sélectionner un projet'),
  categoryId: yup
    .number()
    .min(1, 'Veuillez sélectionner une catégorie')
    .typeError('Veuillez sélectionner une catégorie')
    .required('Veuillez sélectionner une catégorie'),
});

export default ActivityValidationSchema;
