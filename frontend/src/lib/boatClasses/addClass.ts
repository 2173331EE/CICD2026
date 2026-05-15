import type { BoatClass } from './Class/types';

type AddClassDependencies = {
  getBoatClasses: () => BoatClass[];
  setBoatClasses: (value: BoatClass[]) => void;
  getIsLoadingClasses: () => boolean;
  getIsMutatingClasses: () => boolean;
  setIsMutatingClasses: (value: boolean) => void;
  setSelectedClassName: (value: string | null) => void;
  setClassMessage: (value: string) => void;
};

export function createAddClassHandler(deps: AddClassDependencies) {
  return async function addClass() {
    if (deps.getIsLoadingClasses() || deps.getIsMutatingClasses()) return;

    const classNameInput = window.prompt('Class name?', 'New Class');
    if (classNameInput === null) {
      deps.setClassMessage('Add cancelled.');
      return;
    }

    const handicapTypeInput = window.prompt('H/cap type?', 'PY');
    if (handicapTypeInput === null) {
      deps.setClassMessage('Add cancelled.');
      return;
    }

    const handicapValueInput = window.prompt('H/cap value?', '1000');
    if (handicapValueInput === null) {
      deps.setClassMessage('Add cancelled.');
      return;
    }

    const name = classNameInput.trim();
    const handicapType = handicapTypeInput.trim().toUpperCase();
    const handicapValue = Number(handicapValueInput);

    if (!name || !handicapType || !Number.isFinite(handicapValue) || handicapValue <= 0) {
      deps.setClassMessage('Invalid values. Add aborted.');
      return;
    }

    const currentClasses = deps.getBoatClasses();
    const alreadyExists = currentClasses.some((item) => item.name.toLowerCase() === name.toLowerCase());
    if (alreadyExists) {
      deps.setClassMessage(`Class ${name} already exists.`);
      return;
    }

    const newClass: BoatClass = {
      name,
      handicapType,
      handicapValue: Math.round(handicapValue)
    };

    deps.setBoatClasses([...currentClasses, newClass]);
    deps.setSelectedClassName(newClass.name);
    deps.setClassMessage(`Class ${newClass.name} added locally. Syncing with backend...`);

    deps.setIsMutatingClasses(true);
    try {
      const response = await fetch('http://localhost:8080/api/classes', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(newClass)
      });

      if (!response.ok) {
        throw new Error(`Request failed with status ${response.status}`);
      }

      deps.setClassMessage(`Class ${newClass.name} added and saved to backend.`);
    } catch (error) {
      deps.setClassMessage(`Class ${newClass.name} added locally, but backend sync failed.`);
      console.error('Failed to sync class creation:', error);
    } finally {
      deps.setIsMutatingClasses(false);
    }
  };
}
