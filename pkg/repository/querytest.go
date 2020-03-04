
	package repository
	
	
	
	// tableName := "user_content"
	// joinTableName := "content"
	// userIDString := userID.String()

	// userContents := []response.WireUserContent{}

	// selectStatments := []repository.Select{
	// 	repository.Select{
	// 		TableName:  tableName,
	// 		ColumnName: "id"},
	// 	repository.Select{
	// 		TableName:  joinTableName,
	// 		ColumnName: "id AS content_id"},
	// 	repository.Select{
	// 		TableName:  joinTableName,
	// 		ColumnName: "title"},
	// 	repository.Select{
	// 		TableName:  joinTableName,
	// 		ColumnName: "description"},
	// 	repository.Select{
	// 		TableName:  joinTableName,
	// 		ColumnName: "image_url"},
	// 	repository.Select{
	// 		TableName:  joinTableName,
	// 		ColumnName: "content_type"},
	// 	repository.Select{
	// 		TableName:  joinTableName,
	// 		ColumnName: "domain"},
	// 	repository.Select{
	// 		TableName:  tableName,
	// 		ColumnName: "sent_by"},
	// 	repository.Select{
	// 		TableName:  tableName,
	// 		ColumnName: "created_at"}}

	// whereStatements := []repository.Where{}
	// orderStatements := []repository.Order{}
	// joinStatements := []repository.Join{
	// 	repository.Join{
	// 		TableName:      tableName,
	// 		ColumnName:     "content_id",
	// 		JoinTableName:  "content",
	// 		JoinColumnName: "id"}}

	// whereStatements = append(whereStatements,
	// 	repository.Where{
	// 		TableName:  tableName,
	// 		ColumnName: "user_id",
	// 		Argument:   &userIDString})

	// if senderType != nil {
	// 	where := repository.Where{
	// 		TableName:  tableName,
	// 		ColumnName: "sender_type",
	// 		Argument:   senderType}
	// 	whereStatements = append(whereStatements, where)
	// }

	// if contentType != nil {
	// 	where := repository.Where{
	// 		TableName:  tableName,
	// 		ColumnName: "content_type",
	// 		Argument:   contentType}
	// 	whereStatements = append(whereStatements, where)
	// }

	// if archived != nil && *archived == true {
	// 	where := repository.Where{
	// 		TableName:  tableName,
	// 		ColumnName: "archived_at IS NOT NULL",
	// 		Argument:   nil}
	// 	whereStatements = append(whereStatements, where)
	// }

	// query := repository.Query{
	// 	TableName:        tableName,
	// 	SelectStatements: selectStatments,
	// 	WhereStatements:  whereStatements,
	// 	OrderStatements:  orderStatements,
	// 	JoinStatements:   joinStatements}

	// rows, err := DB.Query(query.GenerateQuery(), query.GenerateArguments()...)