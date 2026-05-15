<script lang="ts">
  import type { BoatClass } from '$lib/boatClasses/Class/types';

  let {
    showClassWindow,
    boatClasses,
    selectedClassName,
    classMessage,
    isLoadingClasses,
    isMutatingClasses,
    closeClassWindow,
    selectClass,
    addClass,
    deleteSelectedClass
  } = $props<{
    showClassWindow: boolean;
    boatClasses: BoatClass[];
    selectedClassName: string | null;
    classMessage: string;
    isLoadingClasses: boolean;
    isMutatingClasses: boolean;
    closeClassWindow: () => void;
    selectClass: (name: string) => void;
    addClass: () => void | Promise<void>;
    deleteSelectedClass: () => void | Promise<void>;
  }>();
</script>

{#if showClassWindow}
  <div class="modal-backdrop" aria-hidden="true" onclick={closeClassWindow}></div>

  <dialog class="class-window" aria-label="Classes of boat" open>
    <header class="class-title-bar">
      <h2>Classes of boat</h2>
      <button type="button" class="dialog-close" aria-label="Close" onclick={closeClassWindow}>
        &#10005;
      </button>
    </header>

    <div class="class-table-wrap">
      <table class="class-table">
        <thead>
          <tr>
            <th>Class name</th>
            <th>H/cap type</th>
            <th>H/cap value</th>
          </tr>
        </thead>
        <tbody>
          {#each boatClasses as boatClass}
            <tr
              class:selected-row={selectedClassName === boatClass.name}
              onclick={() => selectClass(boatClass.name)}
            >
              <td>{boatClass.name}</td>
              <td>{boatClass.handicapType}</td>
              <td>{boatClass.handicapValue}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>

    <div class="message-box">{classMessage}</div>

    <footer class="class-actions">
      <div class="class-left-actions">
        <button type="button" onclick={addClass} disabled={isLoadingClasses || isMutatingClasses}>Add</button>
        <button
          type="button"
          onclick={deleteSelectedClass}
          disabled={isLoadingClasses || isMutatingClasses || !selectedClassName}
        >
          Delete
        </button>
      </div>
      <div class="class-right-actions">
        <button type="button" onclick={() => window.print()} disabled={isLoadingClasses || isMutatingClasses}>
          Print
        </button>
        <button type="button" onclick={closeClassWindow}>Cancel</button>
      </div>
    </footer>
  </dialog>
{/if}
