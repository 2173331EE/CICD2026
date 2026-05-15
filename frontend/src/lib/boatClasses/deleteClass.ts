import type { BoatClass } from './Class/types';

type DeleteClassDependencies = {
  getBoatClasses: () => BoatClass[];
  setBoatClasses: (value: BoatClass[]) => void;
  getIsLoadingClasses: () => boolean;
  getIsMutatingClasses: () => boolean;
  setIsMutatingClasses: (value: boolean) => void;
  getSelectedClassName: () => string | null;
  setSelectedClassName: (value: string | null) => void;
  setClassMessage: (value: string) => void;
};

export function createDeleteClassHandler(deps: DeleteClassDependencies) {
  return async function deleteSelectedClass() {
    if (deps.getIsLoadingClasses() || deps.getIsMutatingClasses()) return;

    const selectedClassName = deps.getSelectedClassName();
    if (!selectedClassName) {
      deps.setClassMessage('Select a class before deleting.');
      return;
    }

    const nameToDelete = selectedClassName;
    const confirmed = window.confirm(`Delete class ${nameToDelete}?`);
    if (!confirmed) {
      deps.setClassMessage('Delete cancelled.');
      return;
    }

    const currentClasses = deps.getBoatClasses();
    deps.setBoatClasses(currentClasses.filter((item) => item.name !== nameToDelete));
    deps.setSelectedClassName(null);
    deps.setClassMessage(`Class ${nameToDelete} deleted locally. Syncing with backend...`);

    deps.setIsMutatingClasses(true);
    try {
      const response = await fetch(`http://localhost:8080/api/classes/${encodeURIComponent(nameToDelete)}`, {
        method: 'DELETE'
      });

      if (!response.ok) {
        throw new Error(`Request failed with status ${response.status}`);
      }

      deps.setClassMessage(`Class ${nameToDelete} deleted locally and on backend.`);
    } catch (error) {
      deps.setClassMessage(`Class ${nameToDelete} deleted locally, but backend sync failed.`);
      console.error('Failed to sync class deletion:', error);
    } finally {
      deps.setIsMutatingClasses(false);
    }
  };
}
