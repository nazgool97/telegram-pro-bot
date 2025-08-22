import React, { useCallback } from 'react'
import ReactFlow, { Controls, Background, addEdge, useNodesState, useEdgesState, Connection } from 'reactflow'
import 'reactflow/dist/style.css'
import axios from 'axios'


const TextNode = ({ data }: any) => <div style={{ padding: 8, border: '1px solid #ddd', borderRadius: 8 }}>{data.text || 'Текст'}</div>
const ButtonsNode = () => <div style={{ padding: 8, border: '1px solid #ddd', borderRadius: 8 }}>Кнопки</div>
const PaymentNode = ({ data }: any) => <div style={{ padding: 8, border: '1px solid #ddd', borderRadius: 8 }}>Оплата {data.amount}₽</div>


const nodeTypes = { text: TextNode, buttons: ButtonsNode, payment: PaymentNode }


export default function App() {
const [nodes, setNodes, onNodesChange] = useNodesState([])
const [edges, setEdges, onEdgesChange] = useEdgesState([])
const onConnect = useCallback((params: Connection) => setEdges((eds) => addEdge(params, eds)), [])


const addNode = (type: string) => {
const id = `${type}_${Date.now()}`
const text = type === 'text' ? window.prompt('Текст сообщения', '') : undefined
setNodes((nds: any[]) => [
...nds,
{
id,
type,
data: { text, amount: type === 'payment' ? 100 : undefined },
position: { x: 100 + nds.length * 50, y: 100 },
},
])
}


const save = async () => {
const payload = {
id: 'start',
name: 'Конструктор',
nodes: nodes.map((n: any) => ({
id: n.id,
type: n.type,
data: JSON.stringify(n.data),
position: n.position,
})),
}


// ВАЖНО: GraphQL variables, а не вставка JSON в текст запроса
await axios.post('http://localhost:8080/query', {
query: 'mutation Save($input: SaveFlowInput!){ saveFlow(input: $input) }',
variables: { input: payload },
})


alert('✅ Сохранено!')
}


return (
<div style={{ height: '100vh' }}>
<div style={{ margin: 10, display: 'flex', gap: 8 }}>
<button onClick={() => addNode('text')}>+ Текст</button>
<button onClick={() => addNode('buttons')}>+ Кнопки</button>
}