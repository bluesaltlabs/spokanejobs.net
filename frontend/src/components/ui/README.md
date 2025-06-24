# UI Component System

This directory contains reusable, accessible, and composable UI components for forms and layouts.

## Philosophy
- **Reusable**: Configurable via props, emits events for value changes.
- **Accessible**: Keyboard navigation and ARIA attributes.
- **Consistent**: Unified styling and API.

## Components
- `UiButton`
- `UiTextInput`
- `UiTextareaInput`
- `UiSelect`
- `UiMultiSelect`
- `UiDateInput`
- `UiRichTextInput`
- `UiSwitch`
- `UiCheckbox`
- `UiRadio`
- `UiFormGroup`, `UiFormRow`, `UiFormColumn`, `UiForm`

## Usage
Import components as needed:
```vue
<script setup>
import { UiButton, UiTextInput } from '@/components/ui';
</script>
```

See each component file for detailed props and usage examples. 