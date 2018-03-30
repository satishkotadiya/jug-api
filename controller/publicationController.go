package controller

import (
	"strconv"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"jug-api/model"
	"jug-api/dao/daoMongo"
)

func (app *App) SalvarPublication(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	publ := model.Publication{}

	if err := json.NewDecoder(request.Body).Decode(&publ); err != nil {
		respondWithMessage(response, 400, "Publicação Inválida")
		return
	}

	dao := daoMongo.PublicationDaoMongo{}
	err := dao.Salvar(publ)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Salvar Publicação")
	} else {
		respondWithMessage(response, 200, "Publicação Salva com Sucesso")
	}
}

func (app *App) AtualizarPublication(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	publ := model.Publication{}

	if err := json.NewDecoder(request.Body).Decode(&publ); err != nil {
		respondWithMessage(response, 400, "Publicação Inválida")
	}

	dao := daoMongo.PublicationDaoMongo{}
	err := dao.Atualizar(publ)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Atualizar Publicação")
	} else {
		respondWithMessage(response, 200, "Publicação Atualizada com Sucesso")
	}
}

func (app *App) RemoverPublication(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	publ := model.Publication{}

	if err := json.NewDecoder(request.Body).Decode(&publ); err != nil {
		respondWithMessage(response, 400, "Publicação Inválida")
	}

	dao := daoMongo.PublicationDaoMongo{}
	err := dao.Remover(publ)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Remover Publicação")
	} else {
		respondWithMessage(response, 200, "Publicação Removida com Sucesso")
	}
}

func (app *App) ListarPublications(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	dao := daoMongo.PublicationDaoMongo{}
	publs, err := dao.Listar()

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Publicações")
	} else if len(publs) == 0 {
		respondWithMessage(response, 204, "Não Há Publicações Salvas")
	} else {
		respondWithJSON(response, 200, publs)
	}
}

func (app *App) GetPublById(response http.ResponseWriter, request *http.Request)  {
	defer request.Body.Close()

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil || id <= 0 {
		respondWithMessage(response, 400, "ID Inválido")
	}

	publ := model.Publication{}
	dao := daoMongo.PublicationDaoMongo{}

	err = dao.GetPublById(id, publ)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Publicação")
	} else if publ.Conteudo == "" {
		respondWithMessage(response, 204, "Publicação Não Foi Econtrada")
	} else {
		respondWithJSON(response, 200, publ)
	}
}

func (app *App) GetPublsByTec(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	vars := mux.Vars(request)
	tecnology := vars["tecnologia"]

	if tecnology == "" || len(tecnology) == 0 {
		respondWithMessage(response, 400, "Tecnologia Inválida")
	}

	dao := daoMongo.PublicationDaoMongo{}
	publs, err := dao.GetPublsByTec(tecnology)

	if err != nil {
		respondWithMessage(response, 500, "Erro ao Recuperar Publicação")
	} else if len(publs) == 0 {
		respondWithMessage(response, 204, "Publicação Não Foi Econtrada")
	} else {
		respondWithJSON(response, 200, publs)
	}
}