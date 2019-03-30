SELECT created_at, id, order_no
	, (
		SELECT row_to_json(t)
		FROM (
			SELECT id, mobile, name
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
			SELECT id, name
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
						SELECT code, id
						FROM goods
						WHERE line_items.good_id = goods.id
					) t
				) AS good
				, (
					SELECT json_agg(t)
					FROM (
						SELECT amount, id
							, (
								SELECT row_to_json(t)
								FROM (
									SELECT code, id
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