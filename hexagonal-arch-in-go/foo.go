func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)

	if status == "" {
		query := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		if err: = d.client.Select(&customers, query); err != nil {
			logger.Error("Error while querying all customers: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	} else {
		query := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		if err: = err = d.client.Select(&customers, query, status); err != nil {
			logger.Error("Error while querying customers with status '" + status : "': " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return customers, nil
 }


 router := mux.NewRouter()
 router.HandleFunc("/v1/api/time", getTime)
 log.Fatal(http.ListenAndServe("localhost:8080", router))
 
 func getTime(w http.ResponseWriter, r *http.Request) {
	 response := make(map[string]string, 0)
	 tz := r.URL.Query().Get("tz")
	 timezones := strings.Split(tz, ",")
 
	 if len(timezones) <= 1 {
		 loc, err := time.LoadLocation(tz)
		 if err != nil {
			 w.WriteHeader(http.StatusNotFound)
			 w.Write([]byte(fmt.Sprintf("invalid timezone %s", tz)))
		 } else {
			 response["current_time"] = time.Now().In(loc).String()
			 w.Header().Add("Content-Type", "application/json")
			 json.NewEncoder(w).Encode(response)
		 }
	  } else {
		 for _, tzdb := range timezones {
			 loc, err := time.LoadLocation(tzdb)
			 if err != nil {
				 w.WriteHeader(http.StatusNotFound)
				 w.Write([]byte(fmt.Sprintf("invalid timezone %s in input", tzdb)))
				 return
			  }
			  now := time.Now().In(loc)
			  response[tzdb] = now.String()
		 }
		 w.Header().Add("Content-Type", "application/json")
		 json.NewEncoder(w).Encode(response)
	  }
 }
 
 // ------------------
 func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomer()
	if err != nil {
	  writeResponse(w, err.Code, err)
	  return
	}
  
	writeResponse(w, http.StatusOK, customers)  
  }
  // -----------------
  func (db CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)

	query := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	if err: = d.client.Select(&customers, query); err != nil {
		logger.Error("Error while querying all customers: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
  }
