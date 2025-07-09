<script setup>
import { ref, computed } from 'vue'
import { useProfileStore } from '@/stores/profile'
import {
  Button, Form, FormGroup, FormRow, TextInput, TextareaInput, DateInput, Modal
} from '@/components/ui'
import { EducationExperience } from '@/models'

const props = defineProps({
  editable: {
    type: Boolean,
    default: true
  }
})

const profile = useProfileStore()

// Education Experience Entry Management
const newEducationEntry = ref(new EducationExperience())
const editingEducationId = ref(null)
const showEducationModal = ref(false)

// Education Experience Functions
function resetNewEducationEntry() {
  newEducationEntry.value = new EducationExperience()
}

function startAddEducationEntry() {
  editingEducationId.value = null
  resetNewEducationEntry()
  showEducationModal.value = true
}

function startEditEducationEntry(entry) {
  editingEducationId.value = entry.id
  newEducationEntry.value = { ...entry }
  showEducationModal.value = true
}

async function saveEducationEntry() {
  if (!newEducationEntry.value.institution) return

  if (editingEducationId.value) {
    profile.editEducationExperience(editingEducationId.value, { ...newEducationEntry.value })
  } else {
    profile.addEducationExperience({ ...newEducationEntry.value, id: Date.now().toString() })
  }

  await profile.loadProfile()
  editingEducationId.value = null
  resetNewEducationEntry()
  showEducationModal.value = false
}

function cancelEducationEntry() {
  resetNewEducationEntry()
  editingEducationId.value = null
  showEducationModal.value = false
}

async function removeEducationEntry(id) {
  profile.removeEducationExperience(id)
  await profile.loadProfile()
}
</script>

<template>
  <div class="profile-section">
    <h2>Education Experience</h2>
    <hr class="divider" />

    <div v-if="!editable && profile.education_experiences.length === 0">
      <p>No education experience entries yet.</p>
    </div>

    <div v-if="editable" class="education-experience-edit">
      <div class="education-experience-header">
        <Button @click="startAddEducationEntry" variant="primary">Add Education Experience</Button>
      </div>

      <div v-for="entry in profile.education_experiences" :key="entry.id" class="education-entry">
        <div class="education-entry-content">
          <strong>{{ entry.institution }}</strong>
          <span v-if="entry.major"> - {{ entry.major }}</span>
          <span v-if="entry.minor"> (Minor: {{ entry.minor }})</span>
          <span v-if="entry.area_of_study"> - {{ entry.area_of_study }}</span>
          <span v-if="entry.start_date || entry.end_date">{{ entry.start_date }} - {{ entry.end_date || (entry.completion_date ? entry.completion_date : '') }}</span>
          <span v-if="entry.location"> at {{ entry.location }}</span>
          <p v-if="entry.description">{{ entry.description }}</p>
          <div v-if="entry.gpa || entry.credits_earned" class="education-details">
            <span v-if="entry.gpa"><strong>GPA:</strong> {{ entry.gpa }}</span>
            <span v-if="entry.credits_earned"><strong>Credits Earned:</strong> {{ entry.credits_earned }}</span>
            <span v-if="entry.credit_type"><strong>Credit Type:</strong> {{ entry.credit_type }}</span>
          </div>
          <div v-if="entry.honors_recognition && entry.honors_recognition.length > 0" class="honors-recognition">
            <strong>Honors & Recognition:</strong>
            <ul>
              <li v-for="honor in entry.honors_recognition" :key="honor">{{ honor }}</li>
            </ul>
          </div>
        </div>
        <div class="education-entry-actions">
          <Button @click="startEditEducationEntry(entry)" variant="primary" size="small">Edit</Button>
          <Button @click="removeEducationEntry(entry.id)" variant="danger" size="small">Delete</Button>
        </div>
      </div>

      <div v-if="profile.education_experiences.length === 0" class="empty-education-experience">
        <p>No education experience entries yet. Click "Add Education Experience" to get started.</p>
      </div>
    </div>

    <div v-else>
      <div v-for="entry in profile.education_experiences" :key="entry.id" class="education-entry">
        <div class="education-entry-content">
          <strong>{{ entry.institution }}</strong>
          <span v-if="entry.major"> - {{ entry.major }}</span>
          <span v-if="entry.minor"> (Minor: {{ entry.minor }})</span>
          <span v-if="entry.area_of_study"> - {{ entry.area_of_study }}</span>
          <span v-if="entry.start_date || entry.end_date">{{ entry.start_date }} - {{ entry.end_date || (entry.completion_date ? entry.completion_date : '') }}</span>
          <span v-if="entry.location"> at {{ entry.location }}</span>
          <p v-if="entry.description">{{ entry.description }}</p>
          <div v-if="entry.gpa || entry.credits_earned" class="education-details">
            <span v-if="entry.gpa"><strong>GPA:</strong> {{ entry.gpa }}</span>
            <span v-if="entry.credits_earned"><strong>Credits Earned:</strong> {{ entry.credits_earned }}</span>
            <span v-if="entry.credit_type"><strong>Credit Type:</strong> {{ entry.credit_type }}</span>
          </div>
          <div v-if="entry.honors_recognition && entry.honors_recognition.length > 0" class="honors-recognition">
            <strong>Honors & Recognition:</strong>
            <ul>
              <li v-for="honor in entry.honors_recognition" :key="honor">{{ honor }}</li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <!-- Education Experience Modal -->
    <Modal v-model="showEducationModal">
      <template #title>
        {{ editingEducationId ? 'Edit' : 'Add' }} Education Experience
      </template>

      <Form>
        <FormRow>
          <FormGroup style="flex-grow: 1;" label="Institution" required>
            <TextInput v-model="newEducationEntry.institution" placeholder="e.g., University of Technology" />
          </FormGroup>
          <FormGroup style="flex-grow: 1;" label="Location">
            <TextInput v-model="newEducationEntry.location" placeholder="e.g., San Francisco, CA" />
          </FormGroup>
        </FormRow>

        <FormRow>
          <FormGroup style="flex-grow: 1;" label="Major">
            <TextInput v-model="newEducationEntry.major" placeholder="e.g., Computer Science" />
          </FormGroup>
          <FormGroup style="flex-grow: 1;" label="Minor">
            <TextInput v-model="newEducationEntry.minor" placeholder="e.g., Mathematics" />
          </FormGroup>
        </FormRow>

        <FormRow>
          <FormGroup style="flex-grow: 1;" label="Area of Study">
            <TextInput v-model="newEducationEntry.area_of_study" placeholder="e.g., Software Engineering" />
          </FormGroup>
          <FormGroup style="flex-grow: 1;" label="GPA">
            <TextInput v-model="newEducationEntry.gpa" placeholder="e.g., 3.8" />
          </FormGroup>
        </FormRow>

        <FormRow>
          <FormGroup style="flex-grow: 1;" label="Start Date">
            <DateInput v-model="newEducationEntry.start_date" placeholder="Start Date" />
          </FormGroup>
          <FormGroup style="flex-grow: 1;" label="End Date">
            <DateInput v-model="newEducationEntry.end_date" placeholder="End Date" />
          </FormGroup>
        </FormRow>

        <FormRow>
          <FormGroup style="flex-grow: 1;" label="Completion Date">
            <DateInput v-model="newEducationEntry.completion_date" placeholder="Completion Date" />
          </FormGroup>
          <FormGroup style="flex-grow: 1;" label="Credits Earned">
            <TextInput v-model="newEducationEntry.credits_earned" placeholder="e.g., 120" />
          </FormGroup>
        </FormRow>

        <FormGroup label="Credit Type">
          <TextInput v-model="newEducationEntry.credit_type" placeholder="e.g., Semester Hours" />
        </FormGroup>

        <FormGroup label="Description">
          <TextareaInput v-model="newEducationEntry.description" placeholder="Describe your education experience, achievements, and relevant coursework" />
        </FormGroup>

        <FormGroup label="Honors & Recognition">
          <TextareaInput
            :model-value="Array.isArray(newEducationEntry.honors_recognition) ? newEducationEntry.honors_recognition.join('\n') : newEducationEntry.honors_recognition"
            @update:model-value="newEducationEntry.honors_recognition = $event.split('\n').filter(line => line.trim())"
            placeholder="List any honors, awards, or recognition received (one per line)"
          />
        </FormGroup>
      </Form>

      <template #footer>
        <Button @click="saveEducationEntry" variant="primary">
          {{ editingEducationId ? 'Update' : 'Add' }} Education Experience
        </Button>
        <Button @click="cancelEducationEntry" variant="secondary">Cancel</Button>
      </template>
    </Modal>
  </div>
</template>

<style scoped>
  .profile-section {
    max-width: 960px;
    margin-bottom: 1.5rem;
    padding: 1rem;
    background: var(--color-background-elevated);
    border-radius: var(--border-radius-medium);
    border: 1px solid var(--color-border);
  }

  .education-experience-edit {
    margin-top: 1rem;
  }

  .education-experience-header {
    display: flex;
    justify-content: flex-end;
    margin-bottom: 1rem;
  }

  .education-entry {
    border: 1.5px solid var(--color-border);
    border-radius: var(--border-radius-medium);
    padding: 1.25rem 1rem 1rem 1rem;
    margin-bottom: 1.25rem;
    background: var(--color-surface);
    box-shadow: 0 1px 4px var(--color-shadow);
    transition: background 0.2s, border 0.2s;
  }

  .education-entry strong {
    color: var(--color-primary-600);
    font-size: 1.1rem;
    font-weight: 700;
  }

  .education-entry em {
    color: var(--color-primary-500);
    font-style: normal;
    font-weight: 600;
  }

  .education-entry span {
    color: var(--color-text-muted);
    font-size: 0.95rem;
    display: block;
    margin-bottom: 0.5rem;
  }

  .education-entry p {
    color: var(--color-text);
    margin: 0.5rem 0 0 0;
    font-size: 1rem;
  }

  .education-entry-actions {
    margin-top: 0.75rem;
    display: flex;
    gap: 0.5rem;
  }

  .empty-education-experience {
    text-align: center;
    padding: 2rem;
    color: var(--color-text-subtle);
    background: var(--color-surface-hover);
    border-radius: 8px;
    border: 1px dashed var(--color-border);
  }

  .empty-education-experience p {
    margin: 0;
    font-size: 1rem;
  }

  .education-details {
    margin-top: 0.5rem;
    padding: 0.5rem;
    background: var(--color-surface-hover);
    border-radius: var(--border-radius-small);
    border-left: 3px solid var(--color-primary-500);
  }

  .education-details span {
    display: inline-block;
    margin-right: 1rem;
    margin-bottom: 0.25rem;
  }

  .honors-recognition {
    margin-top: 0.5rem;
    padding: 0.5rem;
    background: var(--color-surface-hover);
    border-radius: var(--border-radius-small);
    border-left: 3px solid var(--color-success-500);
  }

  .honors-recognition ul {
    margin: 0.5rem 0 0 0;
    padding-left: 1.5rem;
  }

  .honors-recognition li {
    margin-bottom: 0.25rem;
  }
</style>
