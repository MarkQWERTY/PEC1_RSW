import os
import glob
import re

html_files = glob.glob("*.html")

enlaces_rapidos_new = """<ul class="pie-pagina__lista">
                            <li><a href="maquinaria.html">Maquinaria</a></li>
                            <li><a href="servicios.html">Servicios</a></li>
                            <li><a href="tienda.html">Tienda</a></li>
                            <li><a href="reglas.html">Reglas</a></li>
                            <li><a href="equipo.html">Equipo</a></li>
                            <li><a href="apuntate.html">¡Apúntate!</a></li>
                        </ul>"""

gimnasios_new = """<ul class="pie-pagina__lista">
                            <li><i class="fas fa-map-marker-alt"></i> Seattle</li>
                            <li><i class="fas fa-map-marker-alt"></i> Tel-Aviv</li>
                            <li><i class="fas fa-map-marker-alt"></i> Isla de Epstein</li>
                            <li><i class="fas fa-map-marker-alt"></i> Vallecas</li>
                            <li><i class="fas fa-map-marker-alt"></i> Meco</li>
                            <li><i class="fas fa-map-marker-alt"></i> Berlín</li>
                        </ul>"""

re_enlaces = re.compile(r'<h4>Enlaces Rápidos</h4>\s*<ul class="pie-pagina__lista">.*?</ul>', re.DOTALL)
re_gimnasios = re.compile(r'<h4>Gimnasios</h4>\s*<ul class="pie-pagina__lista">.*?</ul>', re.DOTALL)

for file in html_files:
    with open(file, 'r', encoding='utf-8') as f:
        content = f.read()
    
    content = re_enlaces.sub(f'<h4>Enlaces Rápidos</h4>\n                        {enlaces_rapidos_new}', content)
    content = re_gimnasios.sub(f'<h4>Gimnasios</h4>\n                        {gimnasios_new}', content)
    
    with open(file, 'w', encoding='utf-8') as f:
        f.write(content)

print(f"Updated footers in {len(html_files)} HTML files.")
