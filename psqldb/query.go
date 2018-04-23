package psqldb

const DepthOne = `
SELECT count(*) FROM person
WHERE id IN (
 	SELECT person_two_id FROM endorsement
 	WHERE endorsement.person_one_id IN (
  		SELECT id FROM person WHERE name=$1         
 )
);
`

const DepthTwo = `
SELECT count(*) FROM person
WHERE id IN (
	SELECT person_two_id FROM endorsement
	WHERE endorsement.person_one_id IN (
		SELECT id FROM person
		WHERE id IN (
			SELECT person_two_id FROM endorsement
			WHERE endorsement.person_one_id IN (
				SELECT id FROM person WHERE name=$1
			)
		)
	)
);
`

const DepthThree = `
SELECT count(*) FROM person
WHERE id IN (
	SELECT person_two_id FROM endorsement
	WHERE endorsement.person_one_id IN (
		SELECT id FROM person
		WHERE id IN (
			SELECT person_two_id FROM endorsement
			WHERE endorsement.person_one_id IN (
				SELECT id FROM person
				WHERE id IN (
					SELECT person_two_id FROM endorsement
					WHERE endorsement.person_one_id IN (
						SELECT id FROM person WHERE name=$1
					)
				)
			)
		)
	)
);
`
const DepthFour = `
SELECT count(*) FROM person
WHERE id IN (
	SELECT person_two_id FROM endorsement
	WHERE endorsement.person_one_id IN (
		SELECT id FROM person
		WHERE id IN (
			SELECT person_two_id FROM endorsement
			WHERE endorsement.person_one_id IN (
				SELECT id FROM person
				WHERE id IN (
					SELECT person_two_id FROM endorsement
					WHERE endorsement.person_one_id IN (
						SELECT id FROM person
						WHERE id IN (
							SELECT person_two_id FROM endorsement
							WHERE endorsement.person_one_id IN (
								SELECT id FROM person WHERE name=$1
							)
						)
					)
				)
			)
		)
	)
);
`

const DepthFive = `
SELECT count(*) FROM person
WHERE id IN (
	SELECT person_two_id FROM endorsement
	WHERE endorsement.person_one_id IN (
		SELECT id FROM person
		WHERE id IN (
			SELECT person_two_id FROM endorsement
			WHERE endorsement.person_one_id IN (
				SELECT id FROM person
				WHERE id IN (
					SELECT person_two_id FROM endorsement
					WHERE endorsement.person_one_id IN (
						SELECT id FROM person
						WHERE id IN (
							SELECT person_two_id FROM endorsement
							WHERE endorsement.person_one_id IN (
								SELECT id FROM person
								WHERE id IN (
									SELECT person_two_id FROM endorsement
									WHERE endorsement.person_one_id IN (
										SELECT id FROM person WHERE name=$1
									)
								)
							)
						)
					)
				)
			)
		)
	)
);
`
