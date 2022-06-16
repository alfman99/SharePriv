import { Tabs } from '@mantine/core';

export interface SelectPageProps {
  component: JSX.Element;
  icon: JSX.Element;
  label: string;
}

const Selector = (props: any) => {
  return (
    <Tabs>
      {props.tabs.map((tab: SelectPageProps, index: number) => {
        return (
          <Tabs.Tab label={tab.label} icon={tab.icon} key={index}>
            {tab.component}
          </Tabs.Tab>
        )
      })}
    </Tabs>
  );
}

export default Selector;