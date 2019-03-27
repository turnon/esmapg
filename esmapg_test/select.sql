SELECT order_no, created_at
	, (
		SELECT row_to_json(t)
		FROM (
			SELECT name, mobile
				, (
					SELECT json_agg(t)
					FROM (
						SELECT path
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
			SELECT name
			FROM stores
			WHERE orders.store_id = stores.id
		) t
	) AS store
	, (
		SELECT row_to_json(t)
		FROM (
			SELECT text
			FROM contracts
			WHERE orders.id = contracts.order_id
		) t
	) AS contract
	, (
		SELECT json_agg(t)
		FROM (
			SELECT barcode
				, (
					SELECT row_to_json(t)
					FROM (
						SELECT code
						FROM goods
						WHERE line_items.good_id = goods.id
					) t
				) AS good
				, (
					SELECT json_agg(t)
					FROM (
						SELECT amount
							, (
								SELECT row_to_json(t)
								FROM (
									SELECT code
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