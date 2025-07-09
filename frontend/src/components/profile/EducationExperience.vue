<script setup>
import { ref, computed } from 'vue'
import { useProfileStore } from '@/stores/profile'
import {
  Button, Form, FormGroup, FormRow, TextInput, TextareaInput, DateInput, Modal
} from '@/components/ui'
import { ProfileEntryCard } from '@/components/common'
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

  // Removed await profile.loadProfile() to prevent duplicate entries
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

      <ProfileEntryCard v-for="entry in profile.education_experiences" :key="entry.id">
        <template #title>
          <strong>{{ entry.institution }}</strong>
          <span v-if="entry.major"> - {{ entry.major }}</span>
          <span v-if="entry.minor"> (Minor: {{ entry.minor }})</span>
          <span v-if="entry.area_of_study"> - {{ entry.area_of_study }}</span>
        </template>

        <template #subtitle>
          <span v-if="entry.location"> at {{ entry.location }}</span>
        </template>

        <template #metadata>
          <span v-if="entry.start_date || entry.end_date">{{ entry.start_date }} - {{ entry.end_date || (entry.completion_date ? entry.completion_date : '') }}</span>
        </template>

        <template #description>
          <p v-if="entry.description">{{ entry.description }}</p>
        </template>

        <template #details>
          <div v-if="entry.gpa || entry.credits_earned" class="info-box">
            <span v-if="entry.gpa"><strong>GPA:</strong> {{ entry.gpa }}</span>
            <span v-if="entry.credits_earned"><strong>Credits Earned:</strong> {{ entry.credits_earned }}</span>
            <span v-if="entry.credit_type"><strong>Credit Type:</strong> {{ entry.credit_type }}</span>
          </div>
          <div v-if="entry.honors_recognition && entry.honors_recognition.length > 0" class="success-info-box">
            <strong>Honors & Recognition:</strong>
            <ul>
              <li v-for="honor in entry.honors_recognition" :key="honor">{{ honor }}</li>
            </ul>
          </div>
        </template>

        <template #actions>
          <Button @click="startEditEducationEntry(entry)" variant="primary" size="small">Edit</Button>
          <Button @click="removeEducationEntry(entry.id)" variant="danger" size="small">Delete</Button>
        </template>
      </ProfileEntryCard>

      <div v-if="profile.education_experiences.length === 0" class="empty-education-experience">
        <p>No education experience entries yet. Click "Add Education Experience" to get started.</p>
      </div>
    </div>

    <div v-else>
      <ProfileEntryCard v-for="entry in profile.education_experiences" :key="entry.id">
        <template #title>
          <strong>{{ entry.institution }}</strong>
          <span v-if="entry.major"> - {{ entry.major }}</span>
          <span v-if="entry.minor"> (Minor: {{ entry.minor }})</span>
          <span v-if="entry.area_of_study"> - {{ entry.area_of_study }}</span>
        </template>

        <template #subtitle>
          <span v-if="entry.location"> at {{ entry.location }}</span>
        </template>

        <template #metadata>
          <span v-if="entry.start_date || entry.end_date">{{ entry.start_date }} - {{ entry.end_date || (entry.completion_date ? entry.completion_date : '') }}</span>
        </template>

        <template #description>
          <p v-if="entry.description">{{ entry.description }}</p>
        </template>

        <template #details>
          <div v-if="entry.gpa || entry.credits_earned" class="info-box">
            <span v-if="entry.gpa"><strong>GPA:</strong> {{ entry.gpa }}</span>
            <span v-if="entry.credits_earned"><strong>Credits Earned:</strong> {{ entry.credits_earned }}</span>
            <span v-if="entry.credit_type"><strong>Credit Type:</strong> {{ entry.credit_type }}</span>
          </div>
          <div v-if="entry.honors_recognition && entry.honors_recognition.length > 0" class="success-info-box">
            <strong>Honors & Recognition:</strong>
            <ul>
              <li v-for="honor in entry.honors_recognition" :key="honor">{{ honor }}</li>
            </ul>
          </div>
        </template>
      </ProfileEntryCard>
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


</style>
