package bottle

type Bottle struct {
    ID      int    `json:"id"`
    OwnerID int    `json:"owner_id"`
    Brand   string `json:"brand"`
    Volume  int    `json:"volume"`
}
