<script setup>
import { ref, computed } from 'vue'
import { useProfileStore } from '@/stores/profile'
import {
  Button, Form, FormGroup, FormRow, TextInput, TextareaInput, DateInput, Modal
} from '@/components/ui'
import { ProfileEntryCard } from '@/components/common'
import { WorkExperience } from '@/models'

const props = defineProps({
  editable: {
    type: Boolean,
    default: false
  }
})

const profile = useProfileStore()

// Work Experience Entry Management
const newWorkEntry = ref(new WorkExperience())
const editingWorkId = ref(null)
const showWorkModal = ref(false)

// Work Experience Functions
function resetNewWorkEntry() {
  newWorkEntry.value = new WorkExperience()
}

function startAddWorkEntry() {
  editingWorkId.value = null
  resetNewWorkEntry()
  showWorkModal.value = true
}

function startEditWorkEntry(entry) {
  editingWorkId.value = entry.id
  newWorkEntry.value = { ...entry }
  showWorkModal.value = true
}

async function saveWorkEntry() {
  if (!newWorkEntry.value.job_title_start || !newWorkEntry.value.employer) return

  if (editingWorkId.value) {
    profile.editWorkExperience(editingWorkId.value, { ...newWorkEntry.value })
  } else {
    profile.addWorkExperience({ ...newWorkEntry.value, id: Date.now().toString() })
  }

  await profile.loadProfile()
  editingWorkId.value = null
  resetNewWorkEntry()
  showWorkModal.value = false
}

function cancelWorkEntry() {
  resetNewWorkEntry()
  editingWorkId.value = null
  showWorkModal.value = false
}

async function removeWorkEntry(id) {
  profile.removeWorkExperience(id)
  await profile.loadProfile()
}
</script>

<template>
  <div class="profile-section">
    <h2>Work Experience</h2>
    <hr class="divider" />

    <div v-if="!editable && profile.work_experiences.length === 0">
      <p>No work experience entries yet.</p>
    </div>

    <div v-if="editable" class="work-experience-edit">
      <div class="work-experience-header">
        <Button @click="startAddWorkEntry" variant="primary">Add Work Experience</Button>
      </div>

      <ProfileEntryCard v-for="entry in profile.work_experiences" :key="entry.id">
        <template #title>
          <strong>{{ entry.job_title_start }}</strong>
          <span v-if="entry.job_title_end && entry.job_title_end !== entry.job_title_start"> - {{ entry.job_title_end }}</span>
        </template>
        
        <template #subtitle>
          <span v-if="entry.employer"> at <em>{{ entry.employer }}</em></span>
          <span v-if="entry.company_slug"> ({{ entry.company_slug }})</span>
        </template>
        
        <template #metadata>
          <span v-if="entry.start_date || entry.end_date">{{ entry.start_date }} - {{ entry.end_date || (entry.is_current ? 'Present' : '') }}</span>
        </template>
        
        <template #description>
          <p v-if="entry.responsibilities">{{ entry.responsibilities }}</p>
          <p v-if="entry.comments">{{ entry.comments }}</p>
          <p v-if="entry.reason_for_leaving"><strong>Reason for leaving:</strong> {{ entry.reason_for_leaving }}</p>
        </template>
        
        <template #details>
          <div v-if="entry.supervisor_name || entry.supervisor_email || entry.supervisor_phone || entry.supervisor_title" class="info-box">
            <strong>Supervisor:</strong>
            <span v-if="entry.supervisor_name">{{ entry.supervisor_name }}</span>
            <span v-if="entry.supervisor_title"> ({{ entry.supervisor_title }})</span>
            <span v-if="entry.supervisor_email"> - {{ entry.supervisor_email }}</span>
            <span v-if="entry.supervisor_phone"> - {{ entry.supervisor_phone }}</span>
          </div>
          <p v-if="entry.can_contact"><strong>Can contact:</strong> Yes</p>
        </template>
        
        <template #actions>
          <Button @click="startEditWorkEntry(entry)" variant="primary" size="small">Edit</Button>
          <Button @click="removeWorkEntry(entry.id)" variant="danger" size="small">Delete</Button>
        </template>
      </ProfileEntryCard>

      <div v-if="profile.work_experiences.length === 0" class="empty-work-experience">
        <p>No work experience entries yet. Click "Add Work Experience" to get started.</p>
      </div>
    </div>

    <div v-else>
      <ProfileEntryCard v-for="entry in profile.work_experiences" :key="entry.id">
        <template #title>
          <strong>{{ entry.job_title_start }}</strong>
          <span v-if="entry.job_title_end && entry.job_title_end !== entry.job_title_start"> - {{ entry.job_title_end }}</span>
        </template>
        
        <template #subtitle>
          <span v-if="entry.employer"> at <em>{{ entry.employer }}</em></span>
          <span v-if="entry.company_slug"> ({{ entry.company_slug }})</span>
        </template>
        
        <template #metadata>
          <span v-if="entry.start_date || entry.end_date">{{ entry.start_date }} - {{ entry.end_date || (entry.is_current ? 'Present' : '') }}</span>
        </template>
        
        <template #description>
          <p v-if="entry.responsibilities">{{ entry.responsibilities }}</p>
          <p v-if="entry.comments">{{ entry.comments }}</p>
          <p v-if="entry.reason_for_leaving"><strong>Reason for leaving:</strong> {{ entry.reason_for_leaving }}</p>
        </template>
        
        <template #details>
          <div v-if="entry.supervisor_name || entry.supervisor_email || entry.supervisor_phone || entry.supervisor_title" class="info-box">
            <strong>Supervisor:</strong>
            <span v-if="entry.supervisor_name">{{ entry.supervisor_name }}</span>
            <span v-if="entry.supervisor_title"> ({{ entry.supervisor_title }})</span>
            <span v-if="entry.supervisor_email"> - {{ entry.supervisor_email }}</span>
            <span v-if="entry.supervisor_phone"> - {{ entry.supervisor_phone }}</span>
          </div>
          <p v-if="entry.can_contact"><strong>Can contact:</strong> Yes</p>
        </template>
      </ProfileEntryCard>
    </div>

    <!-- Work Experience Modal -->
    <Modal v-model="showWorkModal">
      <template #title>
        {{ editingWorkId ? 'Edit' : 'Add' }} Work Experience
      </template>

      <Form>
        <FormRow>
          <FormGroup style="flex-grow: 1;" label="Job Title (Start)" required>
            <TextInput v-model="newWorkEntry.job_title_start" placeholder="e.g., Software Engineer" />
          </FormGroup>
          <FormGroup style="flex-grow: 1;" label="Job Title (End)">
            <TextInput v-model="newWorkEntry.job_title_end" placeholder="e.g., Senior Software Engineer" />
          </FormGroup>
        </FormRow>

        <FormRow>
          <FormGroup style="flex-grow: 1;" label="Employer" required>
            <TextInput v-model="newWorkEntry.employer" placeholder="e.g., Tech Corp" />
          </FormGroup>
          <FormGroup style="flex-grow: 1;" label="Company Slug">
            <TextInput v-model="newWorkEntry.company_slug" placeholder="e.g., tech-corp" />
          </FormGroup>
        </FormRow>

        <FormRow>
          <FormGroup style="flex-grow: 1;" label="Start Date">
            <DateInput v-model="newWorkEntry.start_date" placeholder="Start Date" />
          </FormGroup>
          <FormGroup style="flex-grow: 1;" label="End Date">
            <DateInput v-model="newWorkEntry.end_date" placeholder="End Date" />
          </FormGroup>
        </FormRow>

        <FormGroup label="Is Current Position">
          <input type="checkbox" v-model="newWorkEntry.is_current" />
        </FormGroup>

        <FormGroup label="Responsibilities">
          <TextareaInput v-model="newWorkEntry.responsibilities" placeholder="Describe your responsibilities and achievements" />
        </FormGroup>

        <FormGroup label="Comments">
          <TextareaInput v-model="newWorkEntry.comments" placeholder="Additional comments" />
        </FormGroup>

        <FormGroup label="Reason for Leaving">
          <TextInput v-model="newWorkEntry.reason_for_leaving" placeholder="e.g., Career advancement, company restructuring" />
        </FormGroup>

        <FormGroup label="Can Contact">
          <input type="checkbox" v-model="newWorkEntry.can_contact" />
        </FormGroup>

        <FormRow>
          <FormGroup style="flex-grow: 1;" label="Supervisor Name">
            <TextInput v-model="newWorkEntry.supervisor_name" placeholder="e.g., John Smith" />
          </FormGroup>
          <FormGroup style="flex-grow: 1;" label="Supervisor Title">
            <TextInput v-model="newWorkEntry.supervisor_title" placeholder="e.g., Engineering Manager" />
          </FormGroup>
        </FormRow>

        <FormRow>
          <FormGroup style="flex-grow: 1;" label="Supervisor Email">
            <TextInput v-model="newWorkEntry.supervisor_email" placeholder="e.g., john.smith@company.com" />
          </FormGroup>
          <FormGroup style="flex-grow: 1;" label="Supervisor Phone">
            <TextInput v-model="newWorkEntry.supervisor_phone" placeholder="e.g., +1-555-123-4567" />
          </FormGroup>
        </FormRow>
      </Form>

      <template #footer>
        <Button @click="saveWorkEntry" variant="primary">
          {{ editingWorkId ? 'Update' : 'Add' }} Work Experience
        </Button>
        <Button @click="cancelWorkEntry" variant="secondary">Cancel</Button>
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

  .work-experience-edit {
    margin-top: 1rem;
  }

  .work-experience-header {
    display: flex;
    justify-content: flex-end;
    margin-bottom: 1rem;
  }





  .empty-work-experience {
    text-align: center;
    padding: 2rem;
    color: var(--color-text-subtle);
    background: var(--color-surface-hover);
    border-radius: 8px;
    border: 1px dashed var(--color-border);
  }

  .empty-work-experience p {
    margin: 0;
    font-size: 1rem;
  }




</style>
