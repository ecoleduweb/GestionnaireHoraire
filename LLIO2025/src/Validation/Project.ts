import { createForm } from "felte";
import type { ProjectBase } from "../Models";
import * as yup from 'yup';

// Schéma de validation
const schema = yup.object({
    name: yup.string().required('Le nom est requis').max(50),
    manager_id: yup.number().required('Le manager est requis').positive(),
    description: yup.string().max(20000),
    billable: yup.boolean(),
    status: yup.number()
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
                err.inner.forEach(value => {
                    errors[value.path] = value.message;
                });

                return errors;
            }
        },
        onSubmit: handleSubmit,
    });
};