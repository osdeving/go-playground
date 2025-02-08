#!/bin/bash

# Criar diretório base se não existir
mkdir -p chapters

# Para cada arquivo section-X.Y.md
for file in section-*.md; do
    if [[ $file =~ section-([0-9]+)\.([0-9]+)\.md ]]; then
        chapter="${BASH_REMATCH[1]}"
        section="${BASH_REMATCH[2]}"
        
        # Criar diretório do capítulo se não existir
        mkdir -p "chapters/chapter-$chapter"
        
        # Mover e renomear o arquivo
        mv "$file" "chapters/chapter-$chapter/ch$chapter-section-$chapter.$section.md"
        
        echo "Moved $file to chapters/chapter-$chapter/ch$chapter-section-$chapter.$section.md"
    fi
done