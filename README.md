# Farmer's Market Implementation

It's on like Donkey Kong!

## Run

### Docker

```bash
# Build image:
docker build -f build/Dockerfile -t farmers .

# Start server:
docker run -it -p 10010:10010 -e HOST=:10010 farmers

# Create a basket:
curl -X POST localhost:10010/api/basket
# Output:
#
# {"basket":{"id":"bt9pit05st6c7r2p61s0","items":null}}

# Add coffee to basket:
curl -X PATCH localhost:10010/api/basket/bt9pit05st6c7r2p61s0/product/CF1
# Output:
#
# {"basket":{"id":"bt9pit05st6c7r2p61s0","items":[{"code":"CF1","price":11.23,"discounts":null}]}}

# Add more coffee to basket:
curl -X PATCH localhost:10010/api/basket/bt9pit05st6c7r2p61s0/product/CF1
# Output:
#
# {"basket":{"id":"bt9pit05st6c7r2p61s0","items":[{"code":"CF1","price":11.23,"discounts":null},{"code":"CF1","price":11.23,"discounts":[{"code":"BOGO","price":-11.23}]}]}}

# Add still more coffee to basket and print it as plain text:
curl -X PATCH localhost:10010/api/basket/bt9pit05st6c7r2p61s0/product/CF1?format=txt
# Output:
#
# Basket: bt9pit05st6c7r2p61s0
# Item                          Price
# ----                          -----
# CF1                           11.23
# CF1                           11.23
#             BOGO             -11.23
# CF1                           11.23
# -----------------------------------
#                               22.46

#
# Add many more products to the basket !
#

# Get basket and print it as plain text:
curl -X GET localhost:10010/api/basket/bt9pit05st6c7r2p61s0?format=txt
# Output:
#
# Basket: bt9pit05st6c7r2p61s0
# Item                          Price
# ----                          -----
# CF1                           11.23
# CF1                           11.23
#             BOGO             -11.23
# CF1                           11.23
# CF1                           11.23
#             BOGO             -11.23
# CF1                           11.23
# CF1                           11.23
#             BOGO             -11.23
# MK1                            4.75
#             CHMK              -4.75
# MK1                            4.75
# MK1                            4.75
# OM1                            3.69
# OM1                            3.69
# OM1                            3.69
# AP1                            6.00
#             APOM              -3.00
#             APPL              -1.50
# AP1                            6.00
#             APPL              -1.50
# AP1                            6.00
#             APPL              -1.50
# CH1                            3.11
# -----------------------------------
#                               67.87

```
