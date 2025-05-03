import { createForm } from 'felte';
import type { Activity } from '../Models';
import * as yup from 'yup';

// Schéma de validation Yup amélioré
const schema = yup.object({
  // Validation des informations de l'activité
  name: yup
    .string()
    .max(50, "Le nom de l'activité ne doit pas dépasser 50 caractères"),
  description: yup
    .string()
    .max(20000, "La description de l'activité ne doit pas dépasser 20000 caractères"),

  // Validation du projet
  projectId: yup
    .number()
    .required('Veuillez sélectionner un projet')
    .typeError('Veuillez sélectionner un projet')
    .min(1, 'Veuillez sélectionner un projet'),
  
  // AVEC VALIDATION de categoryId mais séparée de projectId
  categoryId: yup
    .number()
    .nullable()
    .transform((value) => (value === '' || isNaN(value) ? null : Number(value)))
}).test(
  'category-required-with-project',
  { categoryId: 'Veuillez sélectionner une catégorie' },
  function(values) {
    // Si projectId est présent, categoryId est requis
    if (values.projectId && Number(values.projectId) > 0) {
      return values.categoryId !== null && values.categoryId !== undefined;
    }
    // Sinon, categoryId est optionnel
    return true;
  }
);

// Fonction qui crée un formulaire avec Felte en utilisant le schéma de validation
export const validateActivityForm = (handleSubmit: (values) => void, activity:Activity) => {
  return createForm({
    initialValues: {...activity},
    validate: async (values) => {
      try {        
        // Préparation des valeurs - Conversion explicite
        const valuesToValidate = {
          ...values,
          projectId: values.projectId ? Number(values.projectId) : null,
          categoryId: values.categoryId ? Number(values.categoryId) : null
        };
        
        await schema.validate(valuesToValidate, { abortEarly: false });
        return {};
      } catch(err) {
        const errors: Record<string, string> = {};
        err.inner.forEach(value => {
          errors[value.path] = value.message;
        });
        
        // Si erreur globale, l'assigner à categoryId
        if (err.message && err.message.includes('category-required')) {
          errors.categoryId = 'Veuillez sélectionner une catégorie';
        }
        
        return errors;
      }
    },
    onSubmit: values => {
      // Nettoyer les valeurs avant soumission
      const formattedValues = {
        ...values,
        projectId: Number(values.projectId),
        categoryId: values.categoryId ? Number(values.categoryId) : null
      };
      
      handleSubmit(formattedValues);
    },
  });
};