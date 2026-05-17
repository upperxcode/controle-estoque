#!/bin/bash

# ==============================================================================
# Script Utilitário: EstoqueMaster
# Gerencia ambiente de Dev, Build, Testes e Limpeza
# ==============================================================================

# Cores e Estilos
GREEN='\033[0;32m'
BG_GREEN='\033[42;37m'
BLUE='\033[0;34m'
BG_BLUE='\033[44;37m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'
BOLD='\033[1m'
NC='\033[0m' # No Color

# Diretório raiz do projeto (onde este script está /..)
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

function header {
    echo -e "${CYAN}================================================================${NC}"
    echo -e "${BOLD}${MAGENTA}   📦 ESTOQUE MASTER - Painel de Controle${NC}"
    echo -e "${CYAN}================================================================${NC}"
}

function show_menu {
    header
    echo -e "${BOLD}Escolha uma opção:${NC}"
    echo -e "  ${YELLOW}1)${NC} 🚀 Iniciar modo ${BOLD}DEV${NC} (Backend + Frontend)"
    echo -e "  ${YELLOW}2)${NC} 🏗️  Realizar ${BOLD}BUILD${NC} de Produção"
    echo -e "  ${YELLOW}3)${NC} 🧪 Rodar todos os ${BOLD}TESTES${NC}"
    echo -e "  ${YELLOW}4)${NC} 🧹 ${BOLD}LIMPAR${NC} ambiente (DB, Binários, Dist)"
    echo -e "  ${YELLOW}5)${NC} 📥 ${BOLD}INSTALAR${NC} dependências"
    echo -e "  ${YELLOW}6)${NC} 📋 ${BOLD}COPIAR${NC} frontend para static"
    echo -e "  ${YELLOW}7)${NC} 🖥️  Apenas ${BOLD}FRONTEND${NC}"
    echo -e "  ${YELLOW}8)${NC} ⚙️  Apenas ${BOLD}BACKEND${NC}"
    echo -e "  ${YELLOW}0)${NC} 🚪 Sair"
    echo -e ""
    read -p "Opção: " choice
    case $choice in
        1) run_dev ;;
        2) run_build ;;
        3) run_test ;;
        4) run_clean ;;
        5) run_install ;;
        6) run_copy ;;
        7) run_frontend ;;
        8) run_backend ;;
        0) exit 0 ;;
        *) echo -e "${RED}Opção inválida!${NC}"; sleep 1; show_menu ;;
    esac
}

function run_install {
    echo -e "\n${BG_BLUE} 📥 INSTALANDO DEPENDÊNCIAS ${NC}"
    echo -e "${BLUE}-> Backend...${NC}"
    (cd "$ROOT_DIR/backend" && go mod tidy)
    echo -e "${BLUE}-> Frontend...${NC}"
    (cd "$ROOT_DIR/frontend" && npm install)
    echo -e "${GREEN}✔ Dependências instaladas com sucesso!${NC}\n"
}

function run_dev {
    echo -e "\n${BG_GREEN} 🚀 INICIANDO AMBIENTE DEV (Pressione Ctrl+C para parar) ${NC}"

    # Função para limpar processos ao sair
    cleanup() {
        echo -e "\n${YELLOW}Finalizando processos em background...${NC}"
        kill $BACKEND_PID 2>/dev/null
        exit
    }
    trap cleanup SIGINT SIGTERM

    # Rodar backend em background (subshell)
    (cd "$ROOT_DIR/backend" && go run cmd/server/main.go) &
    BACKEND_PID=$!

    # Rodar frontend (subshell)
    (cd "$ROOT_DIR/frontend" && npm run dev)

    cleanup
}

function run_build {
    echo -e "\n${BG_BLUE} 🏗️  INICIANDO BUILD DE PRODUÇÃO ${NC}"
    echo -e "${BLUE}-> Compilando Backend Go...${NC}"
    (cd "$ROOT_DIR/backend" && go build -o server cmd/server/main.go)
    echo -e "${BLUE}-> Gerando Build Vue.js...${NC}"
    (cd "$ROOT_DIR/frontend" && npm run build)
    echo -e "${GREEN}✔ Build finalizado com sucesso!${NC}\n"
}

function run_test {
    echo -e "\n${BG_BLUE} 🧪 EXECUTANDO SUÍTE DE TESTES ${NC}"
    echo -e "${YELLOW}[BACKEND]${NC}"
    (cd "$ROOT_DIR/backend" && go test -v ./...)
    echo -e "\n${YELLOW}[FRONTEND]${NC}"
    (cd "$ROOT_DIR/frontend" && npm test)
}

function run_clean {
    echo -e "\n${RED}${BOLD}🧹 LIMPANDO PROJETO...${NC}"
    rm -f "$ROOT_DIR/backend/server"
    rm -f "$ROOT_DIR/backend/estoque.db"
    rm -rf "$ROOT_DIR/frontend/dist"
    echo -e "${GREEN}✔ Limpeza concluída!${NC}\n"
}

function run_frontend {
    echo -e "\n${CYAN}🖥️  Iniciando apenas Frontend...${NC}"
    (cd "$ROOT_DIR/frontend" && npm run dev)
}

function run_backend {
    echo -e "\n${CYAN}⚙️  Iniciando apenas Backend...${NC}"
    (cd "$ROOT_DIR/backend" && go run cmd/server/main.go)
}

function run_copy {
    echo -e "\n${MAGENTA}📋 Copiando arquivos estáticos...${NC}"
    mkdir -p "$ROOT_DIR/backend/static"
    if [ -d "$ROOT_DIR/frontend/dist" ]; then
        cp -r "$ROOT_DIR/frontend/dist"/* "$ROOT_DIR/backend/static/"
        echo -e "${GREEN}✔ Arquivos copiados para backend/static!${NC}\n"
    else
        echo -e "${RED}✘ Erro: Pasta frontend/dist não encontrada. Rode o build primeiro.${NC}\n"
    fi
}

# Lógica Principal
if [ -z "$1" ]; then
    show_menu
else
    case $1 in
        "install") run_install ;;
        "dev") run_dev ;;
        "build") run_build ;;
        "test") run_test ;;
        "clean") run_clean ;;
        "frontend") run_frontend ;;
        "backend") run_backend ;;
        "copy") run_copy ;;
        *) echo -e "${RED}Comando desconhecido: $1${NC}"; exit 1 ;;
    esac
fi
