import { QueryClient, useQueryClient } from '@tanstack/react-query';
import { createBrowserRouter, createHashRouter, RouterProvider } from 'react-router-dom';
import { AppRoot } from './routes/app/root';
import { useMemo } from 'react';

const createRouter = import.meta.env.MODE === 'desktop' ? createHashRouter : createBrowserRouter;

export const createAppRouter = (queryClient: QueryClient) => {
  return createRouter([
    {
      path: '/',
      element: <AppRoot />,
      children: [
        {
          path: 'competitions',
          lazy: async () => {
            const { CompetitionsRoute } = await import('./routes/app/competitions/competitions');
            return { Component: CompetitionsRoute };
          },
        },
      ],
    },
  ]);
};

export const AppRouter = () => {
  const queryClient = useQueryClient();

  const router = useMemo(() => createAppRouter(queryClient), [queryClient]);

  return <RouterProvider router={router} />;
};
