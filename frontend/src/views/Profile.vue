<script setup>
import { ref, onMounted, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useProfileStore } from '@/stores/profile'
import {
  Container, Button, Form, FormGroup, FormRow, TextInput, TextareaInput, DateInput, Switch
} from '@/components/ui'
import { SkeletonText } from '@/components/skeleton'
import { EditIcon } from '@/components/icons'
import { WorkExperience } from '@/models'

const route = useRoute()
const router = useRouter()
const profile = useProfileStore()
const saved = ref(false)

const is_editing = computed(() => route.query.edit === '1')



// Work Experience Entry Management
const newWorkEntry = ref(new WorkExperience())
const editingWorkId = ref(null)
const showWorkForm = ref(false)

// Auto-save functionality
const autoSave = async () => {
  await profile.saveProfile()
  saved.value = true
  setTimeout(() => (saved.value = false), 1500)
}

// Watch for changes in profile fields and auto-save
watch(
  () => [
    profile.personal_information.first_name,
    profile.personal_information.middle_name,
    profile.personal_information.last_name,
    profile.personal_information.email,
    profile.personal_information.phone,
    profile.personal_information.avatar_url
  ],
  () => {
    autoSave() // todo: debounce this.
  },
  { deep: true }
)

onMounted(() => {
  profile.loadProfile()
})

async function updateEditQueryParam(value) {
  const newQuery = { ...route.query }
  if (value) {
    newQuery.edit = '1'
  } else {
    delete newQuery.edit
  }
  await router.replace({ query: newQuery })
}

async function onEditButtonClick() {
  const newValue = !is_editing.value
  await updateEditQueryParam(newValue)
}

// Work Experience Functions
function resetNewWorkEntry() {
  newWorkEntry.value = new WorkExperience()
}

function startAddWorkEntry() {
  editingWorkId.value = null
  resetNewWorkEntry()
  showWorkForm.value = true
}

function startEditWorkEntry(entry) {
  editingWorkId.value = entry.id
  newWorkEntry.value = { ...entry }
  showWorkForm.value = true
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
  showWorkForm.value = false
}

function cancelWorkEntry() {
  resetNewWorkEntry()
  editingWorkId.value = null
  showWorkForm.value = false
}

async function removeWorkEntry(id) {
  profile.removeWorkExperience(id)
  await profile.loadProfile()
}

</script>

<template>
<Container>
  <div class="page-header">
    <h1>Profile</h1>


    <div class="profile-actions">
      <Button @click="onEditButtonClick" variant="primary">
        <EditIcon />
      </Button>
    </div>
  </div>

  <hr class="divider" />

  <!-- todo: add some kind of sidebar to scroll more quickly to each section as they grow? -->


  <!-- Personal Information -->
  <div class="profile-section">
    <h2>Personal Information</h2>
    <hr class="divider" />

    <!-- Avatar -->
    <div class="avatar-preview">
      <div class="avatar">
        <img v-if="profile.personal_information.avatar_url" :src="profile.personal_information.avatar_url" alt="Avatar preview" />
        <span v-else>{{ profile.personal_information.first_name?.charAt(0) || 'U' }}</span>
      </div>
    </div>


    <Form>

      <!-- Avatar URL -->
      <FormGroup style="flex-grow: 1;" label="Avatar URL" v-if="is_editing">
        <TextInput
          v-model="profile.personal_information.avatar_url"
          placeholder="avatar url"
          :editable="is_editing"
          />
      </FormGroup>

      <FormRow>
        <!-- First Name -->
        <FormGroup style="flex-grow: 1;" label="First Name" v-if="profile.personal_information.first_name || is_editing">
          <TextInput
            v-model="profile.personal_information.first_name"
            placeholder="first name"
            :editable="is_editing"
            />
        </FormGroup>

        <!-- Middle Name -->
        <FormGroup style="flex-shrink: 1;" label="Middle Name" v-if="profile.personal_information.middle_name || is_editing">
          <TextInput
            v-model="profile.personal_information.middle_name"
            placeholder="middle name"
            :editable="is_editing"
            />
        </FormGroup>

        <!-- Last Name -->
        <FormGroup style="flex-grow: 1;" label="Last Name" v-if="profile.personal_information.last_name || is_editing">
          <TextInput
            v-model="profile.personal_information.last_name"
            placeholder="last name"
            :editable="is_editing"
            />
        </FormGroup>
      </FormRow>

      <!-- Preferred Name -->
      <FormGroup label="Preferred Name" v-if="profile.personal_information.preferred_name || is_editing">
        <TextInput
          v-model="profile.personal_information.preferred_name"
          placeholder="preferred name"
          :editable="is_editing"
          />
      </FormGroup>

      <FormRow>
        <!-- Email -->
        <FormGroup style="flex-grow: 1;" label="Email" v-if="profile.personal_information.email || is_editing">
          <TextInput
            v-model="profile.personal_information.email"
            placeholder="email"
            :editable="is_editing"
            />
        </FormGroup>

        <!-- Phone -->
        <FormGroup style="flex-grow: 1;" label="Phone" v-if="profile.personal_information.phone || is_editing">
          <TextInput
            v-model="profile.personal_information.phone"
            placeholder="phone"
            :editable="is_editing"
            />
        </FormGroup>
      </FormRow>

        <FormRow>
          <!-- Created At -->
          <FormGroup style="flex-grow: 1;" label="Created">
            <TextInput
              v-model="profile.personal_information.created_at"
              placeholder="yyyy/mm/dd hh:mm:ss"
              :editable="false"
              />
          </FormGroup>


          <!-- Updated At -->
          <FormGroup style="flex-grow: 1;" label="Updated">
            <TextInput
              v-model="profile.personal_information.updated_at"
              placeholder="yyyy/mm/dd hh:mm:ss"
              :editable="false"
            />
          </FormGroup>

        <!-- todo: URLS           (urls)              --> <!-- figure out how to make this an array of values-->
        </FormRow>
      </Form>

</div>


  <!-- Work Experience -->
  <div class="profile-section">
    <h2>Work Experience</h2>
    <hr class="divider" />

    <div v-if="!is_editing && profile.work_experiences.length === 0">
      <p>No work experience entries yet.</p>
    </div>

    <div v-if="is_editing" class="work-experience-edit">
      <div class="work-experience-header">
        <Button @click="startAddWorkEntry" variant="primary">Add Work Experience</Button>
      </div>

      <div v-if="showWorkForm" class="work-form">
        <Form>
          <FormRow>
            <FormGroup style="flex-grow: 1;" label="Job Title (Start)">
              <TextInput v-model="newWorkEntry.job_title_start" placeholder="e.g., Software Engineer" />
            </FormGroup>
            <FormGroup style="flex-grow: 1;" label="Job Title (End)">
              <TextInput v-model="newWorkEntry.job_title_end" placeholder="e.g., Senior Software Engineer" />
            </FormGroup>
          </FormRow>

          <FormRow>
            <FormGroup style="flex-grow: 1;" label="Employer">
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

          <div class="form-action-row">
            <Button @click="saveWorkEntry" variant="primary" class="form-action-btn">
              {{ editingWorkId ? 'Update' : 'Add' }} Work Experience
            </Button>
            <Button @click="cancelWorkEntry" variant="secondary" class="form-action-btn">Cancel</Button>
          </div>
        </Form>
      </div>

      <div v-for="entry in profile.work_experiences" :key="entry.id" class="work-entry">
        <div class="work-entry-content">
          <strong>{{ entry.job_title_start }}</strong>
          <span v-if="entry.job_title_end && entry.job_title_end !== entry.job_title_start"> - {{ entry.job_title_end }}</span>
          <span v-if="entry.employer"> at <em>{{ entry.employer }}</em></span>
          <span v-if="entry.company_slug"> ({{ entry.company_slug }})</span>
          <span v-if="entry.start_date || entry.end_date">{{ entry.start_date }} - {{ entry.end_date || (entry.is_current ? 'Present' : '') }}</span>
          <p v-if="entry.responsibilities">{{ entry.responsibilities }}</p>
          <p v-if="entry.comments">{{ entry.comments }}</p>
          <p v-if="entry.reason_for_leaving"><strong>Reason for leaving:</strong> {{ entry.reason_for_leaving }}</p>
          <div v-if="entry.supervisor_name || entry.supervisor_email || entry.supervisor_phone || entry.supervisor_title" class="supervisor-info">
            <strong>Supervisor:</strong>
            <span v-if="entry.supervisor_name">{{ entry.supervisor_name }}</span>
            <span v-if="entry.supervisor_title"> ({{ entry.supervisor_title }})</span>
            <span v-if="entry.supervisor_email"> - {{ entry.supervisor_email }}</span>
            <span v-if="entry.supervisor_phone"> - {{ entry.supervisor_phone }}</span>
          </div>
          <p v-if="entry.can_contact"><strong>Can contact:</strong> Yes</p>
        </div>
        <div class="work-entry-actions">
          <Button @click="startEditWorkEntry(entry)" variant="primary" size="small">Edit</Button>
          <Button @click="removeWorkEntry(entry.id)" variant="danger" size="small">Delete</Button>
        </div>
      </div>

      <div v-if="profile.work_experiences.length === 0" class="empty-work-experience">
        <p>No work experience entries yet. Click "Add Work Experience" to get started.</p>
      </div>
    </div>

    <div v-else>
      <div v-for="entry in profile.work_experiences" :key="entry.id" class="work-entry">
        <div class="work-entry-content">
          <strong>{{ entry.job_title_start }}</strong>
          <span v-if="entry.job_title_end && entry.job_title_end !== entry.job_title_start"> - {{ entry.job_title_end }}</span>
          <span v-if="entry.employer"> at <em>{{ entry.employer }}</em></span>
          <span v-if="entry.company_slug"> ({{ entry.company_slug }})</span>
          <span v-if="entry.start_date || entry.end_date">{{ entry.start_date }} - {{ entry.end_date || (entry.is_current ? 'Present' : '') }}</span>
          <p v-if="entry.responsibilities">{{ entry.responsibilities }}</p>
          <p v-if="entry.comments">{{ entry.comments }}</p>
          <p v-if="entry.reason_for_leaving"><strong>Reason for leaving:</strong> {{ entry.reason_for_leaving }}</p>
          <div v-if="entry.supervisor_name || entry.supervisor_email || entry.supervisor_phone || entry.supervisor_title" class="supervisor-info">
            <strong>Supervisor:</strong>
            <span v-if="entry.supervisor_name">{{ entry.supervisor_name }}</span>
            <span v-if="entry.supervisor_title"> ({{ entry.supervisor_title }})</span>
            <span v-if="entry.supervisor_email"> - {{ entry.supervisor_email }}</span>
            <span v-if="entry.supervisor_phone"> - {{ entry.supervisor_phone }}</span>
          </div>
          <p v-if="entry.can_contact"><strong>Can contact:</strong> Yes</p>
        </div>
      </div>
    </div>
  </div>


  <!-- Education Experience -->
  <div class="profile-section">
    <h2>Education Experience</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>


  <!-- Licenses & Certifications -->
  <div class="profile-section">
    <h2>Licenses & Certifications</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>


  <!-- Memberships -->
  <div class="profile-section">
    <h2>Memberships</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>


  <!-- Skills -->
  <div class="profile-section">
    <h2>Skills</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>


  <!-- Interests -->
  <div class="profile-section">
    <h2>Interests</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>


  <!-- Projects -->
  <div class="profile-section">
    <h2>Projects</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>


  <!-- References -->
  <div class="profile-section">
    <h2>References</h2>
    <hr class="divider" />

    <SkeletonText style="margin-bottom: 1rem;" />
  </div>

</Container>
</template>

<style scoped>
  .page-header {
    display: flex;
    justify-content: space-between;
    flex-grow: 1;
    margin-bottom: 1rem;
  }

  .profile-actions {
    display: flex;
    justify-content: end;
    padding: 0.5em 0;
  }

  .profile-section {
    max-width: 960px;
    margin-bottom: 1.5rem;
    padding: 1rem;
    background: var(--color-background-elevated);
    border-radius: var(--border-radius-medium);
    border: 1px solid var(--color-border);
  }


  .avatar-preview {
    text-align: center;
    margin-top: 1rem;
    margin-bottom: 1rem;
  }

  .avatar {
    width: 100px;
    height: 100px;
    background: var(--color-surface-hover);
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--color-text-subtle);
    border-radius: var(--border-radius-medium);
    margin: 0 auto;

    img {
      width: 100px;
      height: 100px;
      border-radius: var(--border-radius-small);
      object-fit: cover;
    }
  }

  .work-experience-edit {
    margin-top: 1rem;
  }

  .work-experience-header {
    display: flex;
    justify-content: flex-end;
    margin-bottom: 1rem;
  }

  .work-form {
    margin-bottom: 1.5rem;
    padding: 1rem;
    background: var(--color-surface-hover);
    border-radius: var(--border-radius-small);
    border: 1px solid var(--color-border);
  }

  .work-entry {
    border: 1.5px solid var(--color-border);
    border-radius: var(--border-radius-medium);
    padding: 1.25rem 1rem 1rem 1rem;
    margin-bottom: 1.25rem;
    background: var(--color-surface);
    box-shadow: 0 1px 4px var(--color-shadow);
    transition: background 0.2s, border 0.2s;
  }

  .work-entry strong {
    color: var(--color-primary-600);
    font-size: 1.1rem;
    font-weight: 700;
  }

  .work-entry em {
    color: var(--color-primary-500);
    font-style: normal;
    font-weight: 600;
  }

  .work-entry span {
    color: var(--color-text-muted);
    font-size: 0.95rem;
    display: block;
    margin-bottom: 0.5rem;
  }

  .work-entry p {
    color: var(--color-text);
    margin: 0.5rem 0 0 0;
    font-size: 1rem;
  }

  .work-entry-actions {
    margin-top: 0.75rem;
    display: flex;
    gap: 0.5rem;
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

  .form-action-row {
    display: flex;
    flex-direction: row;
    gap: 1rem;
    margin-top: 1.25rem;
  }

  .form-action-btn {
    flex: 1 1 0;
    min-width: 0;
  }

  .supervisor-info {
    margin-top: 0.5rem;
    padding: 0.5rem;
    background: var(--color-surface-hover);
    border-radius: var(--border-radius-small);
    border-left: 3px solid var(--color-primary-500);
  }

  .supervisor-info strong {
    color: var(--color-text);
    font-weight: 600;
  }
</style>
