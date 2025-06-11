export const chatColors = [
  '#2a5298', '#f39c12', '#27ae60', '#8e44ad', '#e74c3c', '#16a085',
  '#d35400', '#2980b9', '#c0392b', '#1abc9c', '#9b59b6', '#34495e'
]

export function getColor(name: string) {
  let hash = 0
  for (let i = 0; i < name.length; i++) hash = name.charCodeAt(i) + ((hash << 5) - hash)
  return chatColors[Math.abs(hash) % chatColors.length]
}

export function getInitials(name: string) {
  const parts = name.split(' ')
  if (parts.length === 1) return name.substring(0, 2).toUpperCase()
  return (parts[0][0] + (parts[1]?.[0] || '')).toUpperCase()
}
