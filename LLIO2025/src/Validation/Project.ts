import { createForm } from 'felte';
import type { ProjectBase } from '../Models';
import * as yup from 'yup';

const schema = yup.object().shape({
  name: yup
    .string()
    .required('Le nom du projet est requis')
    .max(50, 'Le nom du projet ne doit pas dépasser 50 caractères'),
  managerId: yup
    .number()
    .typeError('Veuillez sélectionner un responsable')
    .min(1, 'Veuillez sélectionner un responsable'),
  description: yup
    .string()
    .max(20000, 'La description du projet ne doit pas dépasser 20000 caractères')
    .nullable(),
  billable: yup
    .boolean()
    .required('Le statut de facturation est requis'),
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
