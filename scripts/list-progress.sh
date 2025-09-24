#!/bin/bash

# Цвета
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${BLUE}📈 Project Progress:${NC}"
echo ""

# Показываем все теги в хронологическом порядке
echo -e "${BLUE}🏷️  Tags (chronological):${NC}"
git for-each-ref --sort=creatordate --format '%(creatordate:short) %(refname:short) - %(contents:subject)' refs/tags | nl -w2 -s'. '

echo ""

# Показываем текущий статус
CURRENT_BRANCH=$(git branch --show-current)
TOTAL_COMMITS=$(git rev-list --count HEAD)
LAST_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "No tags yet")

echo -e "${BLUE}📊 Current Status:${NC}"
echo -e "  Branch: ${YELLOW}$CURRENT_BRANCH${NC}"
echo -e "  Total commits: ${YELLOW}$TOTAL_COMMITS${NC}"
echo -e "  Last tag: ${YELLOW}$LAST_TAG${NC}"

# Показываем коммиты с последнего тега
if [ "$LAST_TAG" != "No tags yet" ]; then
    COMMITS_SINCE_TAG=$(git rev-list --count $LAST_TAG..HEAD)
    echo -e "  Commits since last tag: ${YELLOW}$COMMITS_SINCE_TAG${NC}"
    
    if [ $COMMITS_SINCE_TAG -gt 0 ]; then
        echo ""
        echo -e "${BLUE}📝 Changes since last tag:${NC}"
        git log $LAST_TAG..HEAD --oneline --no-merges | head -10
    fi
fi
