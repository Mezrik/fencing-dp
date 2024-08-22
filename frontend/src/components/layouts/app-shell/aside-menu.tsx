import { Book, LifeBuoy, Settings2, SquareUser, Swords, Circle } from 'lucide-react';
import { ComponentProps, ComponentType, FC } from 'react';
import { Button } from '@/components/ui/button';
import { Tooltip, TooltipContent, TooltipTrigger } from '@/components/ui/tooltip';
import { t } from '@lingui/macro';
import { pathnames } from '@/app/pathnames';
import { cn } from '@/utils/class-names';
import { NavLink } from 'react-router-dom';

type AsideMenuItemProps = {
  name: string;
  to: string;
  icon: ComponentType<ComponentProps<'svg'>>;
};

const menuItems: AsideMenuItemProps[] = [
  {
    name: t`Competitions`,
    to: pathnames.competitionsPath,
    icon: Swords,
  },
  {
    name: t`Documentation`,
    to: '',
    icon: Book,
  },
  {
    name: t`Settings`,
    to: '',
    icon: Settings2,
  },
];

const AsideMenuItem: FC<AsideMenuItemProps> = ({ name, to, icon }) => {
  const Icon = icon;
  return (
    <Tooltip key={name}>
      <TooltipTrigger asChild>
        <NavLink to={to}>
          {({ isActive }) => (
            <Button
              variant="ghost"
              size="icon"
              className={cn('rounded-lg', isActive && 'bg-muted')}
              aria-label={name}
            >
              <Icon className="size-5" />
            </Button>
          )}
        </NavLink>
      </TooltipTrigger>
      <TooltipContent side="right" sideOffset={5}>
        {name}
      </TooltipContent>
    </Tooltip>
  );
};

export const AsideMenu: FC<{ children?: React.ReactNode }> = () => {
  return (
    <aside className="inset-y fixed  left-0 z-20 flex h-full flex-col border-r">
      <div className="border-b p-2">
        <Button variant="outline" size="icon" aria-label="Home">
          <Circle className="size-5 fill-foreground" />
        </Button>
      </div>
      <nav className="grid gap-1 p-2">
        {menuItems.map((props) => (
          <AsideMenuItem key={props.name} {...props} />
        ))}
      </nav>
      <nav className="mt-auto grid gap-1 p-2">
        <AsideMenuItem to="" name={t`Help`} icon={LifeBuoy} />
        <AsideMenuItem to="" name={t`Account`} icon={SquareUser} />
      </nav>
    </aside>
  );
};
