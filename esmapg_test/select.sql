SELECT id, order_no, created_at
	, (
		SELECT row_to_json(t)
		FROM (
			SELECT id, name, mobile
				, (
					SELECT json_agg(t)
					FROM (
						SELECT id, path
						FROM addresses
						WHERE members.id = addresses.member_id
					) t
				) AS addresses
			FROM members
			WHERE orders.member_id = members.id
		) t
	) AS member
	, (
		SELECT row_to_json(t)
		FROM (
			SELECT name, id
			FROM stores
			WHERE orders.store_id = stores.id
		) t
	) AS store
	, (
		SELECT row_to_json(t)
		FROM (
			SELECT id, text
			FROM contracts
			WHERE orders.id = contracts.order_id
		) t
	) AS contract
	, (
		SELECT json_agg(t)
		FROM (
			SELECT barcode, id
				, (
					SELECT row_to_json(t)
					FROM (
						SELECT id, code
						FROM goods
						WHERE line_items.good_id = goods.id
					) t
				) AS good
				, (
					SELECT json_agg(t)
					FROM (
						SELECT id, amount
							, (
								SELECT row_to_json(t)
								FROM (
									SELECT id, code
									FROM promotions
									WHERE discounts.promotion_id = promotions.id
								) t
							) AS promotion
						FROM discounts
						WHERE line_items.id = discounts.line_item_id
					) t
				) AS discounts
			FROM line_items
			WHERE orders.id = line_items.order_id
		) t
	) AS line_items
FROM orders