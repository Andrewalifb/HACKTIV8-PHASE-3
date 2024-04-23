package controller

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Andrewalifb/HACKTIV8-PHASE-3/employee-rest-api-mongodb/config"
	"github.com/Andrewalifb/HACKTIV8-PHASE-3/employee-rest-api-mongodb/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateEmployee(c echo.Context) error {
	collectionName := os.Getenv("EMPLOYEE_COLLECTION")
	collection, err := config.ConnectionDatabase(context.Background(), collectionName)
	if err != nil {
		return err
	}

	e := new(model.Employee)
	err = c.Bind(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	result, err := collection.InsertOne(context.Background(), e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func GetAllEmployee(c echo.Context) error {
	collectionName := os.Getenv("EMPLOYEE_COLLECTION")
	collection, err := config.ConnectionDatabase(context.Background(), collectionName)
	if err != nil {
		return err
	}

	var datas []model.Employee


	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}


	for cursor.Next(context.Background()) {
		var data model.Employee
		if err := cursor.Decode(&data); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}


		datas = append(datas, data)
	}

	if err := cursor.Err(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, datas)
}

func GetEmployeeById(c echo.Context) error {
	collectionName := os.Getenv("EMPLOYEE_COLLECTION")
	collection, err := config.ConnectionDatabase(context.Background(), collectionName)
	if err != nil {
		return err
	}


	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	var data model.Employee


	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&data) 
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, data)
}

func UpdateDataEmployee(c echo.Context) error {
	collectionName := os.Getenv("EMPLOYEE_COLLECTION")
	collection, err := config.ConnectionDatabase(context.Background(), collectionName)
	if err != nil {
		return err
	}

	// untuk mengubah tipe data string pada id menjadi Primitive.ObjectID -> WAJIB
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	e := new(model.Employee)
	err = c.Bind(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}


	result, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": e},
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func DeleteDataPerson(c echo.Context) error {
	collectionName := os.Getenv("EMPLOYEE_COLLECTION")
	collection, err := config.ConnectionDatabase(context.Background(), collectionName)
	if err != nil {
		return err
	}


	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}


	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}


	if result.DeletedCount == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Data Not Found")
	}

	return c.JSON(http.StatusCreated, result)
}

func GetPagging(c echo.Context) error {
	collectionName := os.Getenv("EMPLOYEE_COLLECTION")
	collection, err := config.ConnectionDatabase(context.Background(), collectionName)
	if err != nil {
		return err
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid limit parameter")
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid page parameter")
	}

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64((page - 1) * limit))

	var employees []model.Employee
	cursor, err := collection.Find(context.Background(), bson.D{{}}, findOptions)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for cursor.Next(context.Background()) {
		var employee model.Employee
		if err := cursor.Decode(&employee); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		employees = append(employees, employee)
	}

	return c.JSON(http.StatusOK, employees)
}



func GetSorting(c echo.Context) error {
	collectionName := os.Getenv("EMPLOYEE_COLLECTION")
	collection, err := config.ConnectionDatabase(context.Background(), collectionName)
	if err != nil {
		return err
	}

	
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "firstname", Value: -1}}) // -1 sort form z - a

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}


	var employees []model.Employee
  if err = cursor.All(context.Background(), &employees); err != nil {
    return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
  }

	return c.JSON(http.StatusOK, employees)
}