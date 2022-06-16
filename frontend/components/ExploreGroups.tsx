import { Select, Tabs } from "@mantine/core"
import { useState } from "react"
import { Files, InfoCircle } from "tabler-icons-react";
import { GroupData } from "../contexts/AuthContext"
import DisplayGroupArchivos from "./DisplayGroupArchivos";
import DisplayGroupInfo from "./DisplayGroupInfo";

const ExploreGroups = (props: { groups: GroupData[]; }) => {

  const { groups } = props;

  const [group, setGroup] = useState<GroupData>(groups[0])

  const mapGroupsToSelect = () => {
    return groups.map(group => {
      return {
        value: group.Id,
        label: group.Nombre
      }
    })
  }
  
  const handleChangeGroupSelected = (event: any) => {
    const grupoSelec = groups.find(group => group.Id === event)
    if (!grupoSelec) {
      return;
    }
    setGroup(grupoSelec)
  }

  return (
    <>
      <Select
        label={`Grupos a los que perteneces (${groups.length} grupos)`}
        searchable={true}
        nothingFound={'Nada encontrado'}
        placeholder={'Selecciona un grupo'}
        onChange={(event) => handleChangeGroupSelected(event)}
        defaultValue={groups[0].Id}
        data={mapGroupsToSelect()} />

      <Tabs style={{ marginTop: '1em' }}>
        <Tabs.Tab label={'Info'} icon={<InfoCircle size={20} />}>
          <DisplayGroupInfo group={group} />
        </Tabs.Tab>
        <Tabs.Tab label={'Archivos'} icon={<Files size={20} />}>
          <DisplayGroupArchivos group={group} />
        </Tabs.Tab>
      </Tabs>
    </>
  )
}

export default ExploreGroups