# Skeleton Components

A collection of reusable Vue components for displaying loading placeholders with animated skeleton effects.

## Components

### SkeletonText
Displays animated placeholder lines for text content.

**Props:**
- `lines` (Number, default: 1) - Number of text lines to display
- `lineHeight` (Number, default: 16) - Height of each line in pixels
- `variant` (String, default: 'default') - Style variant: 'default', 'title', 'subtitle'
- `lastLineWidth` (String, default: null) - Width of the last line (e.g., '60%')
- `animated` (Boolean, default: true) - Whether to show animation

**Usage:**
```vue
<SkeletonText :lines="3" variant="title" />
<SkeletonText :lines="4" :line-height="20" last-line-width="60%" />
```

### SkeletonImage
Displays an animated placeholder for images with an icon.

**Props:**
- `width` (String, default: '100%') - Width of the placeholder
- `height` (String, default: '200px') - Height of the placeholder
- `borderRadius` (String, default: 'var(--border-radius)') - Border radius
- `animated` (Boolean, default: true) - Whether to show animation

**Usage:**
```vue
<SkeletonImage width="300px" height="200px" />
<SkeletonImage width="100%" height="150px" />
```

### SkeletonButton
Displays an animated placeholder for buttons.

**Props:**
- `width` (String, default: '120px') - Width of the button
- `height` (String, default: '40px') - Height of the button
- `borderRadius` (String, default: 'var(--border-radius-small)') - Border radius
- `animated` (Boolean, default: true) - Whether to show animation

**Usage:**
```vue
<SkeletonButton width="100px" height="35px" />
<SkeletonButton width="80px" height="32px" />
```

### SkeletonTableRow
Displays an animated placeholder for table rows.

**Props:**
- `columns` (Number, default: 3) - Number of columns in the table row
- `columnWidths` (Array, default: ['40%', '30%', '30%']) - Widths of each column

**Usage:**
```vue
<SkeletonTableRow :columns="3" :column-widths="['40%', '30%', '30%']" />
<SkeletonTableRow :columns="4" :column-widths="['25%', '25%', '25%', '25%']" />
```

### SkeletonCard
A composite component that combines multiple skeleton elements for card-like layouts.

**Props:**
- `showImage` (Boolean, default: true) - Whether to show image placeholder
- `imageWidth` (String, default: '100%') - Image placeholder width
- `imageHeight` (String, default: '200px') - Image placeholder height
- `showTitle` (Boolean, default: true) - Whether to show title placeholder
- `showDescription` (Boolean, default: true) - Whether to show description placeholder
- `descriptionLines` (Number, default: 3) - Number of description lines
- `lastLineWidth` (String, default: '60%') - Width of last description line
- `showButtons` (Boolean, default: false) - Whether to show button placeholders
- `buttonCount` (Number, default: 2) - Number of button placeholders
- `buttonWidths` (Array, default: ['80px', '100px']) - Widths of button placeholders
- `animated` (Boolean, default: true) - Whether to show animation

**Usage:**
```vue
<SkeletonCard 
  :show-image="true"
  :show-title="true"
  :show-description="true"
  :description-lines="4"
  :show-buttons="true"
  :button-count="2"
/>
```

## Examples

### Company Detail Page Loading State
```vue
<template>
  <div v-if="loading">
    <SkeletonText variant="title" :lines="1" />
    <div class="skeleton-links">
      <SkeletonButton width="100px" height="35px" />
      <SkeletonButton width="80px" height="35px" />
    </div>
    <SkeletonImage width="425px" height="425px" />
    <SkeletonText :lines="4" :line-height="20" />
  </div>
</template>
```

### Company List Page Loading State
```vue
<template>
  <table>
    <thead>
      <tr>
        <th>Name</th>
        <th>Slug</th>
        <th>Website</th>
      </tr>
    </thead>
    <tbody>
      <SkeletonTableRow 
        v-if="loading" 
        v-for="i in 8" 
        :key="`skeleton-${i}`"
        :columns="3"
        :column-widths="['40%', '30%', '30%']"
      />
      <tr v-else v-for="company in companies" :key="company.id">
        <!-- Actual company data -->
      </tr>
    </tbody>
  </table>
</template>
```

### Company Grid Loading State
```vue
<template>
  <div v-if="loading" class="companies-grid">
    <SkeletonCard 
      v-for="i in 6" 
      :key="i"
      :show-image="true"
      :show-title="true"
      :show-description="true"
      :description-lines="3"
      :show-buttons="false"
    />
  </div>
</template>
```

## Styling

All skeleton components use CSS custom properties for consistent theming:
- `--color-background-mute` - Background color for skeleton elements
- `--color-text-muted` - Color for placeholder icons
- `--border-radius` - Default border radius
- `--border-radius-small` - Small border radius for buttons

The components automatically adapt to light/dark themes based on your CSS custom properties. 