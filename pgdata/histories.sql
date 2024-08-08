CREATE TABLE "histories"(
    "id" SERIAL NOT NULL,
    "user_message" TEXT NOT NULL,
    "bot_response" TEXT NOT NULL,
    "feedback" BOOLEAN NOT NULL,
    "multimedia" BOOLEAN NOT NULL,
    "timestamp" TIMESTAMP(0) WITHOUT TIME ZONE NULL DEFAULT NOW()
);
ALTER TABLE
    "histories" ADD PRIMARY KEY("id");

INSERT INTO "histories" ("id","user_message","bot_response","feedback", "multimedia","timestamp")
VALUES
    (1,'Quiénes somos','Deal Genius se encarga de tomar los mensajes y automatizar las respuestas de articulos provenientes de MercadoLibre que define: Democratizamos el comercio y los servicios financieros para transformar la vida de millones de personas en América Latina.','FALSE','FALSE',NULL),
    (2,'Que puedo consultar','Deal Genius brinda articulos de la pagina https://www.mercadolibre.com.co','FALSE','FALSE',NULL),
    (3,'Que es Alkosto','Lo siento, no podria decirte algo referente diferente a MercadoLibre','FALSE','FALSE',NULL),
    (4,'Como puedo vender productos en MercadoLibre','Si jamás has vendido algo por Internet, Mercado Libre es el lugar perfecto para dar los primeros pasos y garantizarte una buena experiencia. La clave está en aprender el proceso y atender algunos consejos que pueden ayudarte a resolver mejor la operación.

No sólo es súper sencillo sino que la publicación y oferta de tu producto es totalmente gratuita. Sólo pagas un cargo por la venta realizada.

Sigue estos pasos para publicar por primera vez:

1. Ingresa
Ingresa a tu cuenta y haz clic en la opción Vender.

2. Escoge qué deseas publicar
Productos – Vehículos – Inmuebles – Servicios.

Escribe un título para tu publicación. Cuanto más preciso sea el título más fácil será para tus compradores saber que ofreces lo que están buscando.
Selecciona la categoría de tu producto. Tomando el título de la publicación, te sugerimos una categoría para el producto que deseas vender. Verifica que se adapte a lo que estás ofreciendo o escoge otras.

3. Completa la información de tu producto
Ingresa fotos de tu producto: No te olvides que en la venta online una buena imagen es fundamental.
Indica el stock: Indica la cantidad de productos que tienes disponible de cada variante para automatizar tu control de stock.
Completa el código universal de tu producto.
Indica si tu producto tiene variantes: Un producto puede variar por color, material, tamaño, etc. Una variante es cada una de estas opciones de un mismo producto, sin que eso implique un cambio en el precio.
Completa la ficha técnica: Allí podrás cargar las características más importantes del producto como medida, modelo, marca y material.
Revisa tu publicación, ten en cuenta que si infringes derechos de propiedad intelectual podrías ser denunciado por el titular de los derechos o podríamos pausar o dar de baja tu publicación por incumplimiento a nuestras políticas de publicación.

4. Ingresa el precio del producto


5. Escoge el tipo de publicación
Gratuita, Clásica con exposición alta o Premium con exposición máxima. Las publicaciones son gratuitas. ¡Solo pagas cuando realizas una venta!

Pulsa la opción publicar y ¡listo! Tu producto ya estará a la venta.


Puedes entregar tus productos con Mercado Envíos a cualquier parte del país. Ofreciendo envíos gratuitos para el comprador puedes atraer más clientes.

Además, cobrar tus ventas en Mercado Libre es muy sencillo. Puedes ofrecer distintas opciones de pago y hasta facilidades de financiamiento. El dinero de la venta se acredita automáticamente en tu cuenta de Mercado Pago.

Como verás, al publicar tus productos en Mercado Libre puedes brindar un servicio 360: ventas, pagos y envíos, convirtiéndose en una herramienta ágil para potenciar tus ventas de manera práctica y segura.

Referencia: https://vendedores.mercadolibre.com.co/nota/como-publicar-en-mercado-libre-por-primera-vez','FALSE','FALSE',NULL),
    (5,'Cultura: Seis principios orientan nuestras acciones','Nuestro ADN emprendedor es el eje de una empresa cuya cultura prioriza la diversidad, la autonomía y la creatividad.

Trabajamos para que nuestros colaboradores y equipos se sientan protagonistas de su desarrollo mientras crean una experiencia única, centrada en el usuario.

Nuestra estructura, dinámica y abierta a los riesgos, crea un ambiente estimulante y plural, que forma grandes líderes y logra atraer a los mejores talentos de Latinoamérica.','FALSE','FALSE',NULL),
    (6,'Nuestro código de ética: el compromiso de hacer lo correcto','Como empresa líder de la región en industrias claves como el comercio electrónico y los servicios financieros, entendemos que la transparencia es la base de la confianza y la mejor manera de honrar un vínculo cada vez más intenso.
En Mercado Libre creemos que la formalidad y la seguridad jurídica son pilares de la inclusión y el desarrollo. Somos conscientes de nuestro rol económico y social cada vez más relevante y, como ciudadanos corporativos, además de cumplir estrictamente con las leyes vigentes nos esforzamos para evitar acciones u omisiones inadecuadas en nuestras plataformas y promover las mejores prácticas en los productos y servicios digitales que desarrollamos.
En tiempos en que la migración digital se acelera e impulsa un cambio cultural profundo, intensificamos nuestros esfuerzos de transparencia e integridad, a fin de garantizar a nuestros usuarios espacios digitales en los que su seguridad y sus derechos estén custodiados con los más altos estándares internacionales.','FALSE','FALSE',NULL),
    (7,'Mercado Libre presenta su segundo Reporte de Transparencia','Las personas líderes son dueñas. Tienen visión de futuro y no sacrifican los valores a largo plazo por resultados inmediatos. Actúan en nombre de toda la compañía, sin limitarse a su propio equipo. Nunca dicen “ese no es mi trabajo”.','FALSE','FALSE',NULL),
    (8,'Inventar y simplificar','El informe tiene como objetivo dar a conocer los esfuerzos que realiza Mercado Libre para garantizar la seguridad de sus servicios digitales. A los conceptos informados en el primer reporte, se incorporan datos sobre los requerimientos de información de usuarios procesados en cumplimiento de la ley.','FALSE','FALSE',NULL),
    (9,'Nuestros esfuerzos y compromisos contra el lavado de activos y el financiamiento del terrorismo','El lavado de activos es un delito en el cual los beneficios de las actividades delictivas se mueven a través de una serie de transacciones financieras diseñadas para ocultar el verdadero origen de los fondos. En Mercado Libre y Mercado Pago desarrollamos controles, acuerdos de colaboración y procesos de mejora continua para evitar que nuestros productos y servicios sean utilizados para cualquier tipo de actividad ilegal o contraria a nuestros términos y condiciones de uso.','FALSE','FALSE',NULL),
    (10,'Directorio y Gobierno Corporativo: el desafío de agregar valor en forma sostenida','Somos muy claros y transparentes en la definición de normas, principios y procedimientos que regulan la estructura y el funcionamiento de nuestra empresa para multiplicar el impacto positivo de nuestras acciones.','FALSE','FALSE',NULL),
    (11,'Quiero saber los primeros 5 audifonos inalambricos de la pagina','Si claro, te dejo una la siguiente lista:
- https://www.mercadolibre.com.co/audifonos-in-ear-inalambricos-linkon-earspcharge/p/MCO21759292?pdp_filters=item_id:MCO1789249024#is_advertising=true&searchVariation=MCO21759292&position=1&search_layout=stack&type=pad&tracking_id=166f410d-9422-4f6f-9d03-3138022a3981&is_advertising=true&ad_domain=VQCATCORE_LST&ad_position=1&ad_click_id=MDI5NjExOWEtNDczOC00M2FjLWFmMzMtMmRjMjZhNGFjMzcy
- https://www.mercadolibre.com.co/auriculares-jbl-wave-buds-black/p/MCO24541134?pdp_filters=item_id:MCO1397861507#is_advertising=true&searchVariation=MCO24541134&position=2&search_layout=stack&type=pad&tracking_id=1e089f09-65ad-4e60-85c1-25d5e34f641f&is_advertising=true&ad_domain=VQCATCORE_LST&ad_position=2&ad_click_id=NTI2NjA4ZmYtZjAxZi00ZTFkLWFlZTYtZjY3NzI0OGIzNWYz
- https://www.mercadolibre.com.co/audifonos-inalambricos-redmi-buds-4-lite-bluetooth-53-color-blanco/p/MCO25881463#polycard_client=search-nordic&searchVariation=MCO25881463&position=6&search_layout=stack&type=product&tracking_id=1e089f09-65ad-4e60-85c1-25d5e34f641f&wid=MCO2060213896&sid=search
- https://www.mercadolibre.com.co/jbl-headphones-quantum-tws-black-sa-color-negro/p/MCO26031039?pdp_filters=item_id:MCO1795676904#is_advertising=true&searchVariation=MCO26031039&position=4&search_layout=stack&type=pad&tracking_id=1e089f09-65ad-4e60-85c1-25d5e34f641f&is_advertising=true&ad_domain=VQCATCORE_LST&ad_position=4&ad_click_id=YjAwM2U2ZjUtNjQ2ZC00NDRhLWFkYTgtYzVjNDdhZDEzMmUy
- https://www.mercadolibre.com.co/qcy-h3-anc-audifonos-inalambricos-gamer-diadema-hi-res-bk-color-negro/p/MCO28965708?pdp_filters=item_id:MCO2318098402#is_advertising=true&searchVariation=MCO28965708&position=5&search_layout=stack&type=pad&tracking_id=1e089f09-65ad-4e60-85c1-25d5e34f641f&is_advertising=true&ad_domain=VQCATCORE_LST&ad_position=5&ad_click_id=MmRmMmZmYTctYzhjZi00NWI4LThlMWMtNTViNWU3ZjY2MDY4','FALSE','FALSE',NULL),
    (12,'El gaming keyboard mas barato de la pagina','Aqui te dejo esta opción:
- Mini Teclado Inalambrico Plegable Portatil, Bluetooth Color del teclado Blanco: https://www.mercadolibre.com.co/mini-teclado-inalambrico-plegable-portatil-bluetooth-color-del-teclado-blanco/p/MCO21340018?pdp_filters=item_id:MCO1434189525#is_advertising=true&searchVariation=MCO21340018&position=1&search_layout=stack&type=pad&tracking_id=13ea4d91-83e2-4807-b107-606f42cfef08&is_advertising=true&ad_domain=VQCATCORE_LST&ad_position=1&ad_click_id=YjljZDNiMWItNTAzYS00MWVkLWFmYWEtZjA3MDQ3NDU2MmZk','FALSE','FALSE',NULL),
    (13,'Somos','La empresa de tecnología líder en comercio electrónico y soluciones fintech de América Latina. Nuestro propósito es democratizar el comercio y los servicios financieros para transformar la vida de millones de personas en la región.','FALSE','FALSE',NULL),
    (14,'Hacemos','Desarrollamos productos tecnológicos que permiten a millones de usuarios comprar, vender, anunciar, enviar y pagar a través de Internet de forma fácil, segura y eficiente.','FALSE','FALSE',NULL),
    (15,'Innovamos','La tecnología es la herramienta que nos permite desarrollar soluciones escalables, capaces de generar el impacto necesario para impulsar la inclusión y el desarrollo.','FALSE','FALSE',NULL),
    (16,'El gaming keyboard mas caro de la pagina','Aqui te dejo esta opción:
- Smsom Mechanical Keyboard, Wired Gaming Keyboard: https://articulo.mercadolibre.com.co/MCO-2263994438-smsom-mechanical-keyboard-wired-gaming-keyboard-_JM#polycard_client=search-nordic&position=5&search_layout=stack&type=item&tracking_id=7557c533-57b3-4510-9511-a01135992a05','FALSE','FALSE',NULL),
    (17,'Cuidamos','La sustentabilidad es un modo de hacer totalmente integrado a nuestra estrategia de negocio. Para nosotros, generar valor económico, social y ambiental van de la mano.','FALSE','FALSE',NULL),
    (18,'Trabajamos','El ADN emprendedor que nos guía promueve una cultura que prioriza la diversidad, la autonomía y la creatividad, a través de un entorno dinámico y abierto a los riesgos.','FALSE','FALSE',NULL),
    (19,'Crecemos','Somos la primera empresa argentina en ingresar al Nasdaq 100, el selecto grupo de las compañías tecnológicas más importantes de Wall Street.','FALSE','FALSE',NULL),
    (20,'Comunicamos','Somos conscientes de nuestro rol social y económico cada vez más relevante. Compartimos nuestras acciones y noticias corporativas.','FALSE','FALSE',NULL),
    (21,'Cuando consulte por algun producto quiero que me des el link del producto de MercadoLibre Colombia, con este formato:
protocolo: https:// prefijo articulo: articulo. dominio de la pagina: mercadolibre.com. pais colombia: co/ ID del producto: MCO-2376843708-teclado-gaming-inalambrico-dry-studio-black-diamond-75-rgb-_JM#polycard_client=search-nordic&position=6&search_layout=stack&type=item&tracking_id=6e3b3470-ab11-4598-9fda-0bc88f32f16b','Listo, cada vez que consultes un producto nuevo te mandare el link de consulta y el precio actual en precio Colombian (COP) con el formato que me mostraste.','FALSE','FALSE',NULL),
    (22,'Cuales son las categorias para electrodomesticos','Artefactos de Cuidado Personal
Climatización
Cocción
Dispensadores y Purificadores
Lavado
Pequeños Electrodomésticos
Refrigeración
Otros','FALSE','FALSE',NULL),
    (23,'Cuales son las categorias para Consolas y Videojuegos','Accesorios para Consolas
Accesorios para PC Gaming
Consolas
Pinballs y Máquinas de Juego
Repuestos para Consolas
Videojuegos
Otros','FALSE','FALSE',NULL),
    (24,'Cuales son las categorias para Celulares y Telefonos','Accesorios para Celulares
Celulares y Smartphones
Gafas de Realidad Virtual
Radios y Handies
Repuestos de Celulares
Smartwatches y Accesorios
Tarificadores y Cabinas
Telefonía Fija e Inalámbrica
Telefonía IP
Otros','FALSE','FALSE',NULL),
    (25,'Quiero las tres primeras ofertas para Accesorios para Vehículos','Si claro, te dejo una la siguiente lista:
- https://www.mercadolibre.com.co/radio-carro-aiwa-android-11-pantalla-9-tactil-2-din-2gb-32-gb-aw-a802bs/p/MCO29652527?pdp_filters=deal:MCO779366-1&hide_psmb=true#promotion_type=TODAY_PROMOTION&searchVariation=MCO29652527&deal_print_id=287675f5-aecc-429f-a481-abc3d3e754ff&position=8&search_layout=grid&type=product&tracking_id=8f9b89ad-a576-4904-b83a-0e1231391ab0&deal_print_id=e317dfdb-d962-489e-9ac9-c88e4f3636b3&promotion_type=TODAY_PROMOTION
- https://articulo.mercadolibre.com.co/MCO-556641635-maleteros-tomcat-50lts-litros-tipo-shad-givi-envio-gratis-_JM?hide_psmb=true#promotion_type%3DTODAY_PROMOTION%26deal_print_id%3D287675f5-aecc-429f-a481-abc3d3e754ff%26position%3D16%26search_layout%3Dgrid%26type%3Ditem%26tracking_id%3Dcd9828e5-a3e6-4bc7-8d9a-39e601a83177&deal_print_id=968eb570-992e-4609-ac74-5a883152d815&promotion_type=TODAY_PROMOTION
- https://www.mercadolibre.com.co/radio-carro-tactil-7-sistema-android-wifi-gps-usb-camara/p/MCO34858770?pdp_filters=deal:MCO779366-1&hide_psmb=true#promotion_type=TODAY_PROMOTION&searchVariation=MCO34858770&deal_print_id=287675f5-aecc-429f-a481-abc3d3e754ff&position=2&search_layout=grid&type=product&tracking_id=8a931a36-b142-478e-8816-d0f5be6f4da1&deal_print_id=ead69fe0-02ab-4960-9fb0-9bf68fd62091&promotion_type=TODAY_PROMOTION','FALSE','FALSE',NULL);


SELECT setval('histories_id_seq', max(id)) FROM histories;