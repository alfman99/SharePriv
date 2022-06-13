SELECT t_Invi.propietario as "usuario que ha invitado a la persona con mas archivos subidos"
FROM practica.usuario as t_Usuario 
JOIN practica.invitacion_registro as t_Invi 
ON codigo=t_Usuario.invitacion_registro_codigo 
WHERE username=(
  SELECT a_Publico.propietario 
  FROM practica.archivo_grupo as a_Grupo 
  JOIN practica.archivo_publico as a_Publico 
  ON a_Grupo.propietario = a_Publico.propietario 
  GROUP BY a_Publico.propietario 
  ORDER BY count(*) 
  DESC LIMIT 1
);

Nested Loop  (cost=5005.65..5021.69 rows=1 width=11) (actual time=46.838..46.844 rows=1 loops=1)
   InitPlan 1 (returns $0)
     ->  Limit  (cost=5005.08..5005.08 rows=1 width=19) (actual time=46.754..46.758 rows=1 loops=1)
           ->  Sort  (cost=5005.08..5026.23 rows=8458 width=19) (actual time=46.752..46.755 rows=1 loops=1)
                 Sort Key: (count(*)) DESC
                 Sort Method: top-N heapsort  Memory: 25kB
                 ->  HashAggregate  (cost=4878.21..4962.79 rows=8458 width=19) (actual time=44.117..45.700 rows=6673 loops=1)
                       Group Key: a_publico.propietario
                       ->  Hash Join  (cost=2044.46..4651.14 rows=45415 width=11) (actual time=13.053..33.144 rows=40817 loops=1)
                             Hash Cond: ((a_grupo.propietario)::text = (a_publico.propietario)::text)
                             ->  Seq Scan on archivo_grupo a_grupo  (cost=0.00..2012.28 rows=16028 width=11) (actual time=0.016..6.948 rows=16028 loops=1)
                             ->  Hash  (cost=1744.76..1744.76 rows=23976 width=11) (actual time=12.891..12.892 rows=23976 loops=1)
                                   Buckets: 32768  Batches: 1  Memory Usage: 1275kB
                                   ->  Seq Scan on archivo_publico a_publico  (cost=0.00..1744.76 rows=23976 width=11) (actual time=0.010..7.506 rows=23976 loops=1)
   ->  Index Scan using usuario_pkey on usuario t_usuario  (cost=0.29..8.30 rows=1 width=17) (actual time=46.808..46.809 rows=1 loops=1)
         Index Cond: ((username)::text = ($0)::text)
   ->  Index Scan using invitacion_registro_pkey on invitacion_registro t_invi  (cost=0.29..8.30 rows=1 width=28) (actual time=0.025..0.025 rows=1 loops=1)
         Index Cond: ((codigo)::text = (t_usuario.invitacion_registro_codigo)::text)