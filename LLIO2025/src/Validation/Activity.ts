import { createForm } from 'felte';
import * as yup from 'yup';


// Schéma de validation Yup
const schema = yup.object().shape({
  // Validation des informations de l'activité
  name: yup
    .string()
    .max(50, "Le nom de l'activité ne doit pas dépasser 50 caractères")
    .required("Le nom de l'activité est requis"),
  description: yup
    .string()
    .max(20000, "La description de l'activité ne doit pas dépasser 20000 caractères"),

  // Validation des informations projet et catégorie
  projectId: yup
    .number()
    .typeError('Veuillez sélectionner un projet')
    .min(1, 'Veuillez sélectionner un projet'),
  categoryId: yup
    .number()
    .typeError('Veuillez sélectionner une catégorie')
    .min(1, 'Veuillez sélectionner une catégorie'),
});

// Fonction qui crée un formulaire avec Felte en utilisant le schéma de validation
export const validateActivityForm = (handleSubmit: (values) => void) => {
  return createForm({
    validate: async (values) => {
      try {
        await schema.validate(values, { abortEarly: false });
        return {};
      } catch(err) {
        const errors = {};
        err.inner.forEach(value => {
          errors[value.path] = value.message;
        });
        
        return errors;
      }
    },
    onSubmit: handleSubmit,
  });
};