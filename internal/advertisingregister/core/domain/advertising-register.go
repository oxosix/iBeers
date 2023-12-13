package domain

import (
	"net/http"
	"time"

	"github.com/d90ares/gloisp-web/internal/advertisingregister/core/dto"
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

type Advertising struct {
	ID          ID           `json:"id"`
	Title       string       `json:"name"`
	Description string       `json:"description"`
	Category    CategoryType `json:"category"`
	Logo        LogoPath     `json:"logo"`
	Banner      BannerPath   `json:"banner"`
	Phone1      string       `json:"phone1"`
	Phone2      string       `json:"phone2"`
	Email       string       `json:"email"`
	Site        string       `json:"site"`
	WhatsApp    string       `json:"whatsapp"`
	Facebook    string       `json:"facebook"`
	Instagram   string       `json:"instagram"`
	Twitter     string       `json:"twitter"`
	CreatedAt   time.Time    `json:"created_at"`
}

type LogoPath struct {
	ID        ID     `json:"imageid"`
	ImagePath string `json:"imagepath"`
}

type BannerPath struct {
	ID        string `json:"imageid"`
	ImagePath string `json:"imagepath"`
}

type CategoryType int

const (
	academias = iota + 1
	acessorios
	brinquedos
	acougue
	administracao
	advocacia
	alimentacao
	animais
	autopecas
	bar
	bebidas
	beleza
	borracharia
	comunicacao
	contrucao
	contabilidade
	curriculos
	educacao
	eletronicos
	esportes
	eventos
	farmacia
	funilaria
	gasagua
	grafica
	igreja
	imobiliaria
	industria
	informatica
	lanchonete
	limpeza
	loterica
	manutencao
	mecanina
	mercado
	metalurgica
	moda
	moradia
	moveis
	naturais
	odontologia
	ofertaempregos
	padaria
	papelaria
	pizzaria
	restaurante
	revestimentos
	saude
	seguranca
	serralheria
	servicos
	som
	transportes
	variedades
	veiculos
	vestuario
	vidracaria
	outros
)

func (c CategoryType) String() string {
	switch c {
	case academias:
		return "Academias"
	case acessorios:
		return "Acessórios"
	case brinquedos:
		return "brinquedos"
	case acougue:
		return "acougue"
	case administracao:
		return "administracao"
	case advocacia:
		return "advocacia"
	case alimentacao:
		return "alimentacao"
	case animais:
		return "animais"
	case autopecas:
		return "autopecas"
	case bar:
		return "bar"
	case bebidas:
		return "bebidas"
	case beleza:
		return "beleza"
	case borracharia:
		return "borracharia"
	case comunicacao:
		return "comunicacao"
	case contrucao:
		return "contrucao"
	case contabilidade:
		return "contabilidade"
	case curriculos:
		return "curriculos"
	case educacao:
		return "educacao"
	case eletronicos:
		return "eletronicos"
	case esportes:
		return "esportes"
	case eventos:
		return "eventos"
	case farmacia:
		return "farmacia"
	case funilaria:
		return "funilaria"
	case gasagua:
		return "gas e agua"
	case grafica:
		return "grafica"
	case igreja:
		return "igreja"
	case imobiliaria:
		return "imobiliaria"
	case industria:
		return "industria"
	case informatica:
		return "informatica"
	case lanchonete:
		return "lanchonete"
	case limpeza:
		return "limpeza"
	case loterica:
		return "loterica"
	case manutencao:
		return "manutencao"
	case mecanina:
		return "mecanina"
	case mercado:
		return "mercado"
	case metalurgica:
		return "metalurgica"
	case moda:
		return "moda"
	case moradia:
		return "moradia"
	case moveis:
		return "moveis"
	case naturais:
		return "naturais"
	case odontologia:
		return "odontologia"
	case ofertaempregos:
		return "oferta de empregos"
	case padaria:
		return "padaria"
	case papelaria:
		return "papelaria"
	case pizzaria:
		return "pizzaria"
	case restaurante:
		return "restaurante"
	case revestimentos:
		return "revestimentos"
	case saude:
		return "saude"
	case seguranca:
		return "seguranca"
	case serralheria:
		return "serralheria"
	case servicos:
		return "servicos"
	case som:
		return "som"
	case transportes:
		return "transportes"
	case variedades:
		return "variedades"
	case veiculos:
		return "veiculos"
	case vestuario:
		return "vestuario"
	case vidracaria:
		return "vidracaria"
	case outros:
		return "outros"
	}
	return "Desconhecido"
}

type AdvertisingService interface {
	Create(rw http.ResponseWriter, r *http.Request)
	Fetch(rw http.ResponseWriter, r *http.Request)
}

type AdvertisingUseCase interface {
	Create(advertisingRequest *dto.AdvertisingRequest) (*Advertising, error)
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination[[]Advertising], error)
}

type AdvertisingRepository interface {
	Create(advertisingRequest *dto.AdvertisingRequest) (*Advertising, error)
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination[[]Advertising], error)
}
