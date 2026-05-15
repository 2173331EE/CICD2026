<script lang="ts">
  import { createAddClassHandler } from "$lib/boatClasses/addClass";
  import { loadBoatClassesFromApi } from "$lib/api";
  import type { BoatClass } from "$lib/boatClasses/Class/types";
  import ClassWindowDialog from "$lib/components/ClassWindowDialog.svelte";
  import { createDeleteClassHandler } from "$lib/boatClasses/deleteClass";

  let showClassWindow = $state(false);
  let isLoadingClasses = $state(false);
  let isMutatingClasses = $state(false);
  let selectedClassName = $state<string | null>(null);
  let classMessage = $state("Click Class to load data from Go backend.");

  let boatClasses = $state<BoatClass[]>([]);

  async function loadBoatClasses() {
    isLoadingClasses = true;
    selectedClassName = null;
    classMessage = "Loading classes from Go backend...";

    try {
      const data = await loadBoatClassesFromApi();
      boatClasses = data;
      classMessage = `${data.length} classes loaded from backend.`;
    } catch (error) {
      classMessage = "Backend unreachable. Start Go server on port 8080.";
      console.error("Failed to load classes:", error);
    } finally {
      isLoadingClasses = false;
    }
  }

  async function openClassWindow() {
    showClassWindow = true;
    await loadBoatClasses();
  }

  function closeClassWindow() {
    showClassWindow = false;
  }

  function selectClass(name: string) {
    selectedClassName = name;
    classMessage = `Selected class: ${name}`;
  }

  const addClass = createAddClassHandler({
    getBoatClasses: () => boatClasses,
    setBoatClasses: (value) => {
      boatClasses = value;
    },
    getIsLoadingClasses: () => isLoadingClasses,
    getIsMutatingClasses: () => isMutatingClasses,
    setIsMutatingClasses: (value) => {
      isMutatingClasses = value;
    },
    setSelectedClassName: (value) => {
      selectedClassName = value;
    },
    setClassMessage: (value) => {
      classMessage = value;
    },
  });

  const deleteSelectedClass = createDeleteClassHandler({
    getBoatClasses: () => boatClasses,
    setBoatClasses: (value) => {
      boatClasses = value;
    },
    getIsLoadingClasses: () => isLoadingClasses,
    getIsMutatingClasses: () => isMutatingClasses,
    setIsMutatingClasses: (value) => {
      isMutatingClasses = value;
    },
    getSelectedClassName: () => selectedClassName,
    setSelectedClassName: (value) => {
      selectedClassName = value;
    },
    setClassMessage: (value) => {
      classMessage = value;
    },
  });
</script>

<main class="screen">
  <section class="window" aria-label="Yacht Racing Results">
    <header class="title-bar">
      <h1>Yacht Racing Results + changement Dan</h1>
      <div class="window-controls" aria-hidden="true">
        <span class="control">&#9633;</span>
        <span class="control">&#10005;</span>
      </div>
    </header>

    <div class="content" aria-label="results-area"></div>

    <footer class="actions">
      <div class="left-actions">
        <button type="button" onclick={openClassWindow}>Class</button>
        <button type="button">Boat</button>
        <button type="button">Series</button>
        <button type="button">Race</button>
        <button type="button">Entry</button>
      </div>
      <button type="button">Cancel</button>
    </footer>
  </section>
</main>

<ClassWindowDialog
  {showClassWindow}
  {boatClasses}
  {selectedClassName}
  {classMessage}
  {isLoadingClasses}
  {isMutatingClasses}
  {closeClassWindow}
  {selectClass}
  {addClass}
  {deleteSelectedClass}
/>
