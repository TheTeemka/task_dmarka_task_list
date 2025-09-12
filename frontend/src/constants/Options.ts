export interface Option {
    id?: number| null;
    name: string;
    color?: string;
}


export const statusOptionsWithColor = [
    { id: null, name: 'None', color: 'var(--color-status-none)' },  // Reset option
    {id: 1, name: "Pending", color: 'var(--color-status-pending)' },      // Mild yellow (similar to Notion's pending)
    {id: 2, name: "In Progress", color: 'var(--color-status-in-progress)' },  // Mild blue (similar to Notion's in progress)
    {id: 3, name: "Completed", color: 'var(--color-status-completed)' },   // Mild green (similar to Notion's completed)
    {id: 4, name: "Cancelled", color: 'var(--color-status-cancelled)' },   // Mild red (similar to Notion's cancelled)
];

export function getStatusOptionByID(id : number) : Option {
    return statusOptionsWithColor.find(option => option.id === id) || statusOptionsWithColor[0]; // Default to gray if not found
}

export const priorityOptionsWithColor = [
    { id: null, name: 'None', color: 'var(--color-priority-none)' },  // Reset option
    {id: 1, name: "Low", color: 'var(--color-priority-low)' },      // Mild green
    {id: 2, name: "Medium", color: 'var(--color-priority-medium)' },   // Mild yellow
    {id: 3, name: "High", color: 'var(--color-priority-high)' },     // Mild red
    {id: 4, name: "Urgent", color: 'var(--color-priority-urgent)' },   // Mild blue
];

export function getPriorityOptionByID(id : number) : Option {
    return priorityOptionsWithColor.find(option => option.id === id) || priorityOptionsWithColor[0]; // Default to gray if not found
}