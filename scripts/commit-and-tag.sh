#!/bin/bash

# Цвета для красивого вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Функция для вывода помощи
show_help() {
    echo -e "${BLUE}Usage:${NC}"
    echo "  ./commit-and-tag.sh <commit-message> <tag-name>"
    echo ""
    echo -e "${BLUE}Examples:${NC}"
    echo "  ./commit-and-tag.sh \"add project setup\" \"01-project-setup\""
    echo "  ./commit-and-tag.sh \"implement search endpoint\" \"05-search-endpoint\""
    echo ""
    echo -e "${BLUE}Options:${NC}"
    echo "  -h, --help    Show this help message"
    exit 0
}

# Проверка аргументов
if [ "$1" = "-h" ] || [ "$1" = "--help" ]; then
    show_help
fi

if [ $# -ne 2 ]; then
    echo -e "${RED}Error: Requires exactly 2 arguments${NC}"
    echo ""
    show_help
fi

COMMIT_MESSAGE="$1"
TAG_NAME="$2"
CURRENT_BRANCH=$(git branch --show-current)

# Проверяем, что мы в git репозитории
if ! git rev-parse --git-dir > /dev/null 2>&1; then
    echo -e "${RED}Error: Not in a git repository${NC}"
    exit 1
fi

echo -e "${BLUE}🔍 Current branch:${NC} ${YELLOW}$CURRENT_BRANCH${NC}"
echo -e "${BLUE}📝 Commit message:${NC} $COMMIT_MESSAGE"
echo -e "${BLUE}🏷️  Tag name:${NC} $TAG_NAME"
echo ""

# Показываем статус изменений
echo -e "${BLUE}📊 Changes to be committed:${NC}"
git status --short

if [ -z "$(git status --porcelain)" ]; then
    echo -e "${YELLOW}⚠️  No changes to commit${NC}"
    exit 0
fi

echo ""
read -p "$(echo -e ${YELLOW}Continue? [y/N]:${NC} )" -n 1 -r
echo ""

if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${YELLOW}❌ Cancelled${NC}"
    exit 0
fi

# Проверяем, не существует ли уже такой тег
if git tag -l | grep -q "^$TAG_NAME$"; then
    echo -e "${RED}Error: Tag '$TAG_NAME' already exists${NC}"
    echo -e "${BLUE}Existing tags:${NC}"
    git tag -l | tail -5
    exit 1
fi

echo ""
echo -e "${BLUE}🚀 Starting commit and tag process...${NC}"

# Добавляем все изменения
echo -e "${BLUE}📦 Adding changes...${NC}"
if ! git add .; then
    echo -e "${RED}Error: Failed to add changes${NC}"
    exit 1
fi

# Коммитим
echo -e "${BLUE}💾 Creating commit...${NC}"
if ! git commit -m "$COMMIT_MESSAGE"; then
    echo -e "${RED}Error: Failed to create commit${NC}"
    exit 1
fi

# Создаем тег с аннотацией
echo -e "${BLUE}🏷️  Creating tag...${NC}"
if ! git tag -a "$TAG_NAME" -m "$COMMIT_MESSAGE"; then
    echo -e "${RED}Error: Failed to create tag${NC}"
    exit 1
fi

# Пушим коммит
echo -e "${BLUE}⬆️  Pushing commit...${NC}"
if ! git push origin "$CURRENT_BRANCH"; then
    echo -e "${RED}Error: Failed to push commit${NC}"
    exit 1
fi

# Пушим тег
echo -e "${BLUE}🏷️  Pushing tag...${NC}"
if ! git push origin "$TAG_NAME"; then
    echo -e "${RED}Error: Failed to push tag${NC}"
    exit 1
fi

echo ""
echo -e "${GREEN}✅ Successfully completed!${NC}"
echo -e "${BLUE}📝 Commit:${NC} $(git log -1 --oneline)"
echo -e "${BLUE}🏷️  Tag:${NC} $TAG_NAME"
echo -e "${BLUE}🌿 Branch:${NC} $CURRENT_BRANCH"

# Показываем последние теги для контекста
echo ""
echo -e "${BLUE}📋 Recent tags:${NC}"
git tag -l --sort=-version:refname | head -5
