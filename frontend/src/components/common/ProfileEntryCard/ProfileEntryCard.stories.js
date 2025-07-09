import ProfileEntryCard from './ProfileEntryCard.vue'
import Button from '../../ui/Button/Button.vue'

export default {
  title: 'Common/ProfileEntryCard',
  component: ProfileEntryCard,
  parameters: {
    docs: {
      description: {
        component: 'A reusable card component for displaying profile entries like work experience and education with consistent styling and layout.'
      }
    }
  }
}

// Work Experience Example
const WorkExperienceTemplate = () => ({
  components: { ProfileEntryCard, Button },
  template: `
    <ProfileEntryCard>
      <template #title>
        <strong>Senior Software Engineer</strong>
        <span> - Lead Developer</span>
      </template>
      
      <template #subtitle>
        <span> at <em>Tech Corp</em></span>
        <span> (tech-corp)</span>
      </template>
      
      <template #metadata>
        <span>2020-01-15 - 2023-06-30</span>
      </template>
      
      <template #description>
        <p>Led development of multiple web applications using Vue.js and Node.js. Managed a team of 5 developers and implemented CI/CD pipelines.</p>
        <p>Additional comments about the role and achievements.</p>
        <p><strong>Reason for leaving:</strong> Career advancement opportunity</p>
      </template>
      
      <template #details>
        <div class="info-box">
          <strong>Supervisor:</strong>
          <span>John Smith (Engineering Manager) - john.smith@company.com - +1-555-123-4567</span>
        </div>
        <p><strong>Can contact:</strong> Yes</p>
      </template>
      
      <template #actions>
        <Button variant="primary" size="small">Edit</Button>
        <Button variant="danger" size="small">Delete</Button>
      </template>
    </ProfileEntryCard>
  `
})

export const WorkExperience = WorkExperienceTemplate.bind({})
WorkExperience.parameters = {
  docs: {
    description: {
      story: 'A work experience entry with job title, employer, dates, description, supervisor info, and clean action buttons in a separated section.'
    }
  }
}

// Education Experience Example
const EducationExperienceTemplate = () => ({
  components: { ProfileEntryCard, Button },
  template: `
    <ProfileEntryCard>
      <template #title>
        <strong>University of Technology</strong>
        <span> - Computer Science</span>
        <span> (Minor: Mathematics)</span>
        <span> - Software Engineering</span>
      </template>
      
      <template #subtitle>
        <span> at San Francisco, CA</span>
      </template>
      
      <template #metadata>
        <span>2016-09-01 - 2020-05-15</span>
      </template>
      
      <template #description>
        <p>Studied computer science with focus on software engineering principles, algorithms, and data structures. Completed capstone project on machine learning applications.</p>
      </template>
      
      <template #details>
        <div class="info-box">
          <span><strong>GPA:</strong> 3.8</span>
          <span><strong>Credits Earned:</strong> 120</span>
          <span><strong>Credit Type:</strong> Semester Hours</span>
        </div>
        <div class="success-info-box">
          <strong>Honors & Recognition:</strong>
          <ul>
            <li>Dean's List (2017-2020)</li>
            <li>Computer Science Department Award</li>
            <li>Best Capstone Project</li>
          </ul>
        </div>
      </template>
      
      <template #actions>
        <Button variant="primary" size="small">Edit</Button>
        <Button variant="danger" size="small">Delete</Button>
      </template>
    </ProfileEntryCard>
  `
})

export const EducationExperience = EducationExperienceTemplate.bind({})
EducationExperience.parameters = {
  docs: {
    description: {
      story: 'An education experience entry with institution, degree, dates, description, academic details, honors recognition, and clean action buttons.'
    }
  }
}

// Read-only Example
const ReadOnlyTemplate = () => ({
  components: { ProfileEntryCard },
  template: `
    <ProfileEntryCard>
      <template #title>
        <strong>Software Engineer</strong>
      </template>
      
      <template #subtitle>
        <span> at <em>Startup Inc</em></span>
      </template>
      
      <template #metadata>
        <span>2018-03-01 - Present</span>
      </template>
      
      <template #description>
        <p>Developed full-stack web applications using modern technologies. Collaborated with cross-functional teams to deliver high-quality software solutions.</p>
      </template>
      
      <template #details>
        <div class="info-box">
          <strong>Supervisor:</strong>
          <span>Jane Doe (CTO) - jane.doe@startup.com</span>
        </div>
      </template>
    </ProfileEntryCard>
  `
})

export const ReadOnly = ReadOnlyTemplate.bind({})
ReadOnly.parameters = {
  docs: {
    description: {
      story: 'A read-only profile entry without action buttons, suitable for display-only views.'
    }
  }
}

// Minimal Example
const MinimalTemplate = () => ({
  components: { ProfileEntryCard },
  template: `
    <ProfileEntryCard>
      <template #title>
        <strong>Bachelor of Arts</strong>
      </template>
      
      <template #subtitle>
        <span> at <em>Liberal Arts College</em></span>
      </template>
      
      <template #metadata>
        <span>2015-09-01 - 2019-05-15</span>
      </template>
    </ProfileEntryCard>
  `
})

export const Minimal = MinimalTemplate.bind({})
Minimal.parameters = {
  docs: {
    description: {
      story: 'A minimal profile entry with only basic information, demonstrating the component\'s flexibility.'
    }
  }
}

// Action Section Example
const ActionSectionTemplate = () => ({
  components: { ProfileEntryCard, Button },
  template: `
    <ProfileEntryCard>
      <template #title>
        <strong>Product Manager</strong>
      </template>
      
      <template #subtitle>
        <span> at <em>Innovation Labs</em></span>
      </template>
      
      <template #metadata>
        <span>2021-03-01 - Present</span>
      </template>
      
      <template #description>
        <p>Led product strategy and development for multiple SaaS applications. Managed cross-functional teams and drove user experience improvements.</p>
      </template>
      
      <template #actions>
        <Button variant="primary" size="small">Edit</Button>
        <Button variant="danger" size="small">Delete</Button>
      </template>
    </ProfileEntryCard>
  `
})

export const ActionSection = ActionSectionTemplate.bind({})
ActionSection.parameters = {
  docs: {
    description: {
      story: 'Showcases the clean action section with a border separator and right-aligned buttons for better user experience.'
    }
  }
} 