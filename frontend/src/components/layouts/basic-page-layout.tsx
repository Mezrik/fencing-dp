import { FC } from 'react';
import { Head } from '../meta';
import { cn } from '@/utils/class-names';

export const BasicPageLayout: FC<{
  children: React.ReactNode;
  title: string;
  subtitle?: string;
  actions?: React.ReactNode;
}> = ({ children, title, subtitle, actions }) => {
  return (
    <>
      <Head title={title} />
      <div className="py-6">
        <div
          className={cn(
            'flex mx-auto max-w-7xl px-4 sm:px-6 md:px-8',
            actions ? 'justify-between' : 'justify-start',
          )}
        >
          <div>
            {subtitle && (
              <h3 className="text-md text-primary uppercase tracking-wider">{subtitle}</h3>
            )}
            <h1 className="text-2xl font-semibold text-gray-900">{title}</h1>
          </div>
          {actions}
        </div>
        <div className="mx-auto max-w-7xl px-4 py-6 sm:px-6 md:px-8">{children}</div>
      </div>
    </>
  );
};
