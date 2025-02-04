document.addEventListener('DOMContentLoaded', function() {
    const sectionsContainer = document.getElementById('sections-container');
    const addSectionBtn = document.getElementById('add-section');
    const saveSectionBtn = document.getElementById('save-section');
    const sectionTitle = document.getElementById('section-title');
    const sectionContent = document.getElementById('section-content');

    let sections = [];
    let currentSection = null;

    // Configurar drag and drop
    sectionsContainer.addEventListener('dragover', e => {
        e.preventDefault();
        const draggingItem = document.querySelector('.dragging');
        const siblings = [...sectionsContainer.querySelectorAll('.section-item:not(.dragging)')];
        
        const nextSibling = siblings.find(sibling => {
            const box = sibling.getBoundingClientRect();
            return e.clientY <= box.top + box.height / 2;
        });

        if (nextSibling) {
            sectionsContainer.insertBefore(draggingItem, nextSibling);
        } else {
            sectionsContainer.appendChild(draggingItem);
        }
    });

    // Adicionar nova seção
    addSectionBtn.addEventListener('click', () => {
        sectionTitle.value = '';
        sectionContent.value = '';
        currentSection = null;
    });

    // Salvar seção
    saveSectionBtn.addEventListener('click', async () => {
        const title = sectionTitle.value.trim();
        const content = sectionContent.value.trim();
        
        if (!title || !content) {
            alert('Por favor, preencha título e conteúdo');
            return;
        }

        try {
            const response = await fetch('/api/sections', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ title, content }),
            });

            if (!response.ok) throw new Error('Erro ao salvar seção');

            const section = await response.json();
            addSectionToList(section);
            
            sectionTitle.value = '';
            sectionContent.value = '';
            currentSection = null;
        } catch (error) {
            console.error('Erro:', error);
            alert('Erro ao salvar seção');
        }
    });

    function addSectionToList(section) {
        const sectionElement = document.createElement('div');
        sectionElement.className = 'section-item';
        sectionElement.draggable = true;
        sectionElement.dataset.id = section.ID;
        sectionElement.innerHTML = `
            <h3>${section.Title}</h3>
            <div class="section-preview">${section.Content.substring(0, 100)}...</div>
        `;

        sectionElement.addEventListener('dragstart', () => {
            sectionElement.classList.add('dragging');
        });

        sectionElement.addEventListener('dragend', () => {
            sectionElement.classList.remove('dragging');
            updateSectionsOrder();
        });

        sectionElement.addEventListener('click', () => {
            currentSection = section;
            sectionTitle.value = section.Title;
            sectionContent.value = section.Content;
        });

        sectionsContainer.appendChild(sectionElement);
    }

    async function updateSectionsOrder() {
        const sectionElements = document.querySelectorAll('.section-item');
        const newOrder = Array.from(sectionElements).map(el => el.dataset.id);

        try {
            const response = await fetch('/api/sections/reorder', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ order: newOrder }),
            });

            if (!response.ok) throw new Error('Erro ao reordenar seções');
        } catch (error) {
            console.error('Erro:', error);
            alert('Erro ao reordenar seções');
        }
    }

    // Carregar seções existentes
    async function loadSections() {
        try {
            const response = await fetch('/api/sections');
            if (!response.ok) throw new Error('Erro ao carregar seções');
            
            sections = await response.json();
            sections.forEach(addSectionToList);
        } catch (error) {
            console.error('Erro:', error);
            alert('Erro ao carregar seções');
        }
    }

    loadSections();
}); 