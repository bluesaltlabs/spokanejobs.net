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



</script>

<template>
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
</template>
