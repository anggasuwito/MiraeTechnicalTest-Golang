package r_order

func (r OrderRepo) CountTotalOrderRepo() (total int, err error) {
	err = r.DB.
		QueryRow(`
		SELECT 
		    COUNT(*) 
		FROM orders
		`).
		Scan(&total)
	if err != nil {
		return total, err
	}
	return total, err
}
