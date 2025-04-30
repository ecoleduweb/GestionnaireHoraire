import { createForm } from "felte";
import type { CreateProject } from "../Models";
import * as yup from 'yup';

// Schéma de validation Yup
const schema = yup.object().shape({
    // Validation des informations du projet
    name: yup
        .string()
        .max(50, "Le nom du projet ne doit pas dépasser 50 caractères"),
    description: yup
        .string()
        .max(20000, "La description du projet ne doit pas dépasser 20000 caractères"),

    // Validation des informations projet et catégorie
    status: yup
        .number()
        .typeError('Veuillez séléctionner un status valide')
        .min(0, 'Veuillez séléctionner un status valide'),
    billable: yup
        .bool()
        .typeError('Vous devez spécifier si le projet est facturable ou non.')
});

export const validateProjectForm = (handleSubmit: (values) => void, project: CreateProject) => {
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