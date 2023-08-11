// internal/omdb.service/omdb_service.go

package omdb_service

import (
	"context"
	"omdb/internal/types"
	pb "omdb/internal/types"
	omdb_api "omdb/omdb.api"
)

type OMDBServer struct {
	pb.UnimplementedOMDBServiceServer
}

func (s *OMDBServer) GetMovieByID(ctx context.Context, req *types.GetMovieByIDRequest) (*types.GetMovieByIDResponse, error) {
	// Assuming GetMovieByID function returns a *types.GetMovieByIDResponse and an error
	return omdb_api.GetMovieByID(req.Id)
}

func (s *OMDBServer) SearchMovies(ctx context.Context, req *types.SearchMoviesRequest) (*types.SearchMoviesResponse, error) {
	// Assuming SearchMovies function returns a *types.SearchMoviesResponse and an error
	return omdb_api.SearchMovies(req.Query, req.Page)
}
