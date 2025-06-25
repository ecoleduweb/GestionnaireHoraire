import { createForm } from 'felte';
import type { ProjectBase } from '../Models';
import * as yup from 'yup';

const schema = yup.object().shape({
  name: yup
    .string()
    .max(50, 'Le nom du projet ne doit pas dépasser 50 caractères')
    .nullable(),
  managerId: yup
    .number()
    .typeError('Veuillez sélectionner un responsable')
    .min(1, 'Veuillez sélectionner un responsable'),
  uniqueId: yup
    .string()
    .required('L\'identifiant unique est requis')
    .max(20, 'L\'identifiant unique ne doit pas dépasser 20 caractères'),
  estimatedHours: yup
      .number()
      .typeError('Veuillez entrer un nombre pour les heures estimées')
      .min(0, 'Les heures estimées ne peuvent pas être négatives')
      .nullable(),
});

export const validateProjectForm = (handleSubmit: (values) => void, project: ProjectBase) => {
  return createForm({
    initialValues: { ...project },
    validate: async (values) => {
      try {
        await schema.validate(values, { abortEarly: false });
        return {};
      } catch (err) {
        const errors = {};
        err.inner.forEach((value) => {
          errors[value.path] = value.message;
        });

        return errors;
      }
    },
    onSubmit: handleSubmit,
  });
}
