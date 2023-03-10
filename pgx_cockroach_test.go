package gorm_vs_korm

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

var pgxCockroach *pgxpool.Pool

// func BenchmarkInsert_Cockroach_Pgx(b *testing.B) {
// 	if done() {
// 		b.SkipNow()
// 		return
// 	}
// 	defer timing()()
// 	b.N = total
// 	ctx := context.Background()
// 	_, err := pgxCockroach.Exec(ctx, `TRUNCATE TABLE `+pgxTableName)
// 	assert.Nil(b, err)
// 	p := pool.New().WithMaxGoroutines(cores)
// 	for z := uint64(1); z <= total; z++ {
// 		z := z
// 		p.Go(func() {
// 			_, err := pgxCockroach.Exec(ctx, `INSERT INTO `+pgxTableName+` (id, content) VALUES ($1, $2)`, z, S.EncodeCB63(z, 0))
// 			assert.Nil(b, err)
// 		})
// 	}
// 	p.Wait()
// }

// func BenchmarkUpdate_Cockroach_Pgx(b *testing.B) {
// 	if done() {
// 		b.SkipNow()
// 		return
// 	}
// 	defer timing()(2)
// 	b.N = total
// 	ctx := context.Background()
// 	p := pool.New().WithMaxGoroutines(cores)
// 	for z := uint64(1); z <= total; z++ {
// 		z := z
// 		p.Go(func() {
// 			_, err := pgxCockroach.Exec(ctx, `UPDATE `+pgxTableName+` SET content = $2 WHERE id = $1`, z, S.EncodeCB63(total+z, 0))
// 			assert.Nil(b, err)
// 			_, err = pgxCockroach.Exec(ctx, `UPDATE `+pgxTableName+` SET content = $2 WHERE id = $1`, z, S.EncodeCB63(z, 0))
// 			assert.Nil(b, err)
// 		})
// 	}
// 	b.N *= 2
// 	p.Wait()
// }

// func BenchmarkGetAllS_Cockroach_Pgx(b *testing.B) {
// 	ctx := context.Background()
// 	p := pool.New().WithMaxGoroutines(cores)
// 	for i := uint64(1); i <= uint64(b.N); i++ {
// 		p.Go(func() {
// 			rows, err := pgxCockroach.Query(ctx, `SELECT * FROM `+pgxTableName+` LIMIT $1`, limit)
// 			assert.Nil(b, err)
// 			defer rows.Close()
// 			res := make([]PgxTestTable, 0, limit)
// 			for rows.Next() {
// 				var row PgxTestTable
// 				err = rows.Scan(&row.Id, &row.Content)
// 				assert.Nil(b, err)
// 				res = append(res, row)
// 			}
// 			assert.Equal(b, len(res), limit)
// 		})
// 	}
// 	p.Wait()
// }

// func BenchmarkGetRowS_Cockroach_Pgx(b *testing.B) {
// 	ctx := context.Background()
// 	p := pool.New().WithMaxGoroutines(cores)
// 	for i := uint64(1); i <= uint64(b.N); i++ {
// 		p.Go(func() {
// 			row := pgxCockroach.QueryRow(ctx, `SELECT * FROM `+pgxTableName+` WHERE content = $1 LIMIT 1`, S.EncodeCB63(1+(i%total), 0))
// 			var row2 PgxTestTable
// 			err := row.Scan(&row2.Id, &row2.Content)
// 			assert.Nil(b, err)
// 		})
// 	}
// 	p.Wait()
// }
