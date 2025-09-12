export interface Option {
    id?: number| null;
    name: string;
    color?: string;
}


export const statusOptionsWithColor = [
    { id: null, name: 'None', color: '#FFFFFF' },  // Reset option
    {id: 1, name: "Pending", color: '#d6d6d6' },      // Mild yellow (similar to Notion's pending)
    {id: 2, name: "In Progress", color: '#FFF3CD' },  // Mild blue (similar to Notion's in progress)
    {id: 3, name: "Completed", color: '#D4EDDA' },   // Mild green (similar to Notion's completed)
    {id: 4, name: "Cancelled", color: '#F8D7DA' },   // Mild red (similar to Notion's cancelled)
];

export function getStatusOptionByID(id : number) : Option {
    return statusOptionsWithColor.find(option => option.id === id) || statusOptionsWithColor[0]; // Default to gray if not found
}

export const priorityOptionsWithColor = [
    { id: null, name: 'None', color: '#FFFFFF' },  // Reset option
    {id: 1, name: "Low", color: '#D4EDDA' },      // Mild green
    {id: 2, name: "Medium", color: '#FFF3CD' },   // Mild yellow
    {id: 3, name: "High", color: '#D5A1A3' },     // Mild red
    {id: 4, name: "Urgent", color: '#80ADBC' },   // Mild blue
];

export function getPriorityOptionByID(id : number) : Option {
    return priorityOptionsWithColor.find(option => option.id === id) || priorityOptionsWithColor[0]; // Default to gray if not found
}