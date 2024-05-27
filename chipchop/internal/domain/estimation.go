package domain

type Location struct {
    Address string
    City    string
    State   string
    Country string
    ZipCode string
}

type Estimation struct {
    ID         int64
    ProductName string
    Price      float64
    Location   Location
    UserID     int64
}
