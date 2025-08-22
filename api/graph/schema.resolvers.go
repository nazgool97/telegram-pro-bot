package graph


import (
"context"
"encoding/json"


dbmodel "probot/api/internal/model"
"probot/api/graph/generated"
"probot/api/graph/model"
)


// Mutation возвращает реализацию MutationResolver
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }


type mutationResolver struct{ *Resolver }


func (m *mutationResolver) SaveFlow(ctx context.Context, input model.SaveFlowInput) (bool, error) {
// Преобразуем input.Nodes в JSON для jsonb
nodes := make([]map[string]any, 0, len(input.Nodes))
for _, n := range input.Nodes {
nodes = append(nodes, map[string]any{
"id": n.ID,
"type": n.Type,
"data": n.Data,
"position": map[string]float64{
"x": n.Position.X,
"y": n.Position.Y,
},
})
}
b, err := json.Marshal(nodes)
if err != nil { return false, err }


flow := &dbmodel.Flow{ ID: input.ID, Name: input.Name, Nodes: b }
if err := m.DB.Save(flow).Error; err != nil { return false, err }
return true, nil
}