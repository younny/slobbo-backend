package operations

import (
	"net/http"

	m "github.com/younny/slobbo-backend/src/middleware"
	"github.com/younny/slobbo-backend/src/types"

	"github.com/go-chi/render"
)

func (server *Server) GetWorkshops(w http.ResponseWriter, r *http.Request) {
	pageID := r.Context().Value(m.PageIDKey)
	if err := render.Render(w, r, server.DB.GetWorkshops(pageID.(int))); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func GetWorkshop(w http.ResponseWriter, r *http.Request) {
	workshop := r.Context().Value(m.WorkshopCtxKey).(*types.Workshop)

	if err := render.Render(w, r, workshop); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func (server *Server) CreateWorkshop(w http.ResponseWriter, r *http.Request) {
	workshopRequest := &types.WorkshopRequest{}

	if err := render.Bind(r, workshopRequest); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := workshopRequest.Validate(); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	newWorkshop := types.Workshop{
		Name:        workshopRequest.Name,
		Description: workshopRequest.Description,
		Organiser:   workshopRequest.Organiser,
		Location:    workshopRequest.Location,
		Duration:    workshopRequest.Duration,
		Capacity:    workshopRequest.Capacity,
		Price:       workshopRequest.Price,
		Thumbnail:   workshopRequest.Thumbnail,
	}

	if err := server.DB.CreateWorkshop(&newWorkshop); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := render.Render(w, r, &newWorkshop); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}

}

func (server *Server) UpdateWorkshop(w http.ResponseWriter, r *http.Request) {
	workshop := r.Context().Value(m.WorkshopCtxKey).(*types.Workshop)

	if err := render.Bind(r, workshop); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := server.DB.UpdateWorkshop(workshop); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}

	if err := render.Render(w, r, workshop); err != nil {
		_ = render.Render(w, r, types.ErrRender(err))
		return
	}
}

func (server *Server) DeleteWorkshop(w http.ResponseWriter, r *http.Request) {
	workshop := r.Context().Value(m.WorkshopCtxKey).(*types.Workshop)

	if err := server.DB.DeleteWorkshop(workshop.ID); err != nil {
		_ = render.Render(w, r, types.ErrInvalidRequst(err))
		return
	}
}
