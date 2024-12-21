package service

import "github.com/TuringProblem/purepower/internal/model"

func GetProductByID(id int) (*model.Product, error) {
    product := &model.Product {
        ID: id,
        Name: "Example",
        Price: 49.99,
        Image: "/static/images/{addimagehere}"
    }
    return product, nil
}
