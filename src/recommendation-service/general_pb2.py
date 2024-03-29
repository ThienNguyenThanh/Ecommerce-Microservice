# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: general.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rgeneral.proto\x12\x08lofishop\"0\n\x08\x43\x61rtItem\x12\x12\n\nproduct_id\x18\x01 \x01(\t\x12\x10\n\x08quantity\x18\x02 \x01(\x05\"C\n\x0e\x41\x64\x64ItemRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12 \n\x04item\x18\x02 \x01(\x0b\x32\x12.lofishop.CartItem\"#\n\x10\x45mptyCartRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\"!\n\x0eGetCartRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\":\n\x04\x43\x61rt\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12!\n\x05items\x18\x02 \x03(\x0b\x32\x12.lofishop.CartItem\"\x07\n\x05\x45mpty\"B\n\x1aListRecommendationsRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x13\n\x0bproduct_ids\x18\x02 \x03(\t\"2\n\x1bListRecommendationsResponse\x12\x13\n\x0bproduct_ids\x18\x01 \x03(\t\"\x81\x01\n\x07Product\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x0f\n\x07picture\x18\x04 \x01(\t\x12\"\n\tprice_usd\x18\x05 \x01(\x0b\x32\x0f.lofishop.Money\x12\x12\n\ncategories\x18\x06 \x03(\t\";\n\x14ListProductsResponse\x12#\n\x08products\x18\x01 \x03(\x0b\x32\x11.lofishop.Product\"\x1f\n\x11GetProductRequest\x12\n\n\x02id\x18\x01 \x01(\t\"&\n\x15SearchProductsRequest\x12\r\n\x05query\x18\x01 \x01(\t\"<\n\x16SearchProductsResponse\x12\"\n\x07results\x18\x01 \x03(\x0b\x32\x11.lofishop.Product\"X\n\x0fGetQuoteRequest\x12\"\n\x07\x61\x64\x64ress\x18\x01 \x01(\x0b\x32\x11.lofishop.Address\x12!\n\x05items\x18\x02 \x03(\x0b\x32\x12.lofishop.CartItem\"5\n\x10GetQuoteResponse\x12!\n\x08\x63ost_usd\x18\x01 \x01(\x0b\x32\x0f.lofishop.Money\"Y\n\x10ShipOrderRequest\x12\"\n\x07\x61\x64\x64ress\x18\x01 \x01(\x0b\x32\x11.lofishop.Address\x12!\n\x05items\x18\x02 \x03(\x0b\x32\x12.lofishop.CartItem\"(\n\x11ShipOrderResponse\x12\x13\n\x0btracking_id\x18\x01 \x01(\t\"a\n\x07\x41\x64\x64ress\x12\x16\n\x0estreet_address\x18\x01 \x01(\t\x12\x0c\n\x04\x63ity\x18\x02 \x01(\t\x12\r\n\x05state\x18\x03 \x01(\t\x12\x0f\n\x07\x63ountry\x18\x04 \x01(\t\x12\x10\n\x08zip_code\x18\x05 \x01(\x05\"<\n\x05Money\x12\x15\n\rcurrency_code\x18\x01 \x01(\t\x12\r\n\x05units\x18\x02 \x01(\x03\x12\r\n\x05nanos\x18\x03 \x01(\x05\"8\n\x1eGetSupportedCurrenciesResponse\x12\x16\n\x0e\x63urrency_codes\x18\x01 \x03(\t\"K\n\x19\x43urrencyConversionRequest\x12\x1d\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x0f.lofishop.Money\x12\x0f\n\x07to_code\x18\x02 \x01(\t\"\x90\x01\n\x0e\x43reditCardInfo\x12\x1a\n\x12\x63redit_card_number\x18\x01 \x01(\t\x12\x17\n\x0f\x63redit_card_cvv\x18\x02 \x01(\x05\x12#\n\x1b\x63redit_card_expiration_year\x18\x03 \x01(\x05\x12$\n\x1c\x63redit_card_expiration_month\x18\x04 \x01(\x05\"_\n\rChargeRequest\x12\x1f\n\x06\x61mount\x18\x01 \x01(\x0b\x32\x0f.lofishop.Money\x12-\n\x0b\x63redit_card\x18\x02 \x01(\x0b\x32\x18.lofishop.CreditCardInfo\"(\n\x0e\x43hargeResponse\x12\x16\n\x0etransaction_id\x18\x01 \x01(\t\"L\n\tOrderItem\x12 \n\x04item\x18\x01 \x01(\x0b\x32\x12.lofishop.CartItem\x12\x1d\n\x04\x63ost\x18\x02 \x01(\x0b\x32\x0f.lofishop.Money\"\xb6\x01\n\x0bOrderResult\x12\x10\n\x08order_id\x18\x01 \x01(\t\x12\x1c\n\x14shipping_tracking_id\x18\x02 \x01(\t\x12&\n\rshipping_cost\x18\x03 \x01(\x0b\x32\x0f.lofishop.Money\x12+\n\x10shipping_address\x18\x04 \x01(\x0b\x32\x11.lofishop.Address\x12\"\n\x05items\x18\x05 \x03(\x0b\x32\x13.lofishop.OrderItem\"S\n\x1cSendOrderConfirmationRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\x12$\n\x05order\x18\x02 \x01(\x0b\x32\x15.lofishop.OrderResult\"\x9d\x01\n\x11PlaceOrderRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x15\n\ruser_currency\x18\x02 \x01(\t\x12\"\n\x07\x61\x64\x64ress\x18\x03 \x01(\x0b\x32\x11.lofishop.Address\x12\r\n\x05\x65mail\x18\x05 \x01(\t\x12-\n\x0b\x63redit_card\x18\x06 \x01(\x0b\x32\x18.lofishop.CreditCardInfo\":\n\x12PlaceOrderResponse\x12$\n\x05order\x18\x01 \x01(\x0b\x32\x15.lofishop.OrderResult\"!\n\tAdRequest\x12\x14\n\x0c\x63ontext_keys\x18\x01 \x03(\t\"\'\n\nAdResponse\x12\x19\n\x03\x61\x64s\x18\x01 \x03(\x0b\x32\x0c.lofishop.Ad\"(\n\x02\x41\x64\x12\x14\n\x0credirect_url\x18\x01 \x01(\t\x12\x0c\n\x04text\x18\x02 \x01(\t2\xb8\x01\n\x0b\x43\x61rtService\x12\x36\n\x07\x41\x64\x64Item\x12\x18.lofishop.AddItemRequest\x1a\x0f.lofishop.Empty\"\x00\x12\x35\n\x07GetCart\x12\x18.lofishop.GetCartRequest\x1a\x0e.lofishop.Cart\"\x00\x12:\n\tEmptyCart\x12\x1a.lofishop.EmptyCartRequest\x1a\x0f.lofishop.Empty\"\x00\x32}\n\x15RecommendationService\x12\x64\n\x13ListRecommendations\x12$.lofishop.ListRecommendationsRequest\x1a%.lofishop.ListRecommendationsResponse\"\x00\x32\xf1\x01\n\x15ProductCatalogService\x12\x41\n\x0cListProducts\x12\x0f.lofishop.Empty\x1a\x1e.lofishop.ListProductsResponse\"\x00\x12>\n\nGetProduct\x12\x1b.lofishop.GetProductRequest\x1a\x11.lofishop.Product\"\x00\x12U\n\x0eSearchProducts\x12\x1f.lofishop.SearchProductsRequest\x1a .lofishop.SearchProductsResponse\"\x00\x32\x9e\x01\n\x0fShippingService\x12\x43\n\x08GetQuote\x12\x19.lofishop.GetQuoteRequest\x1a\x1a.lofishop.GetQuoteResponse\"\x00\x12\x46\n\tShipOrder\x12\x1a.lofishop.ShipOrderRequest\x1a\x1b.lofishop.ShipOrderResponse\"\x00\x32\xab\x01\n\x0f\x43urrencyService\x12U\n\x16GetSupportedCurrencies\x12\x0f.lofishop.Empty\x1a(.lofishop.GetSupportedCurrenciesResponse\"\x00\x12\x41\n\x07\x43onvert\x12#.lofishop.CurrencyConversionRequest\x1a\x0f.lofishop.Money\"\x00\x32O\n\x0ePaymentService\x12=\n\x06\x43harge\x12\x17.lofishop.ChargeRequest\x1a\x18.lofishop.ChargeResponse\"\x00\x32\x62\n\x0c\x45mailService\x12R\n\x15SendOrderConfirmation\x12&.lofishop.SendOrderConfirmationRequest\x1a\x0f.lofishop.Empty\"\x00\x32\\\n\x0f\x43heckoutService\x12I\n\nPlaceOrder\x12\x1b.lofishop.PlaceOrderRequest\x1a\x1c.lofishop.PlaceOrderResponse\"\x00\x32\x42\n\tAdService\x12\x35\n\x06GetAds\x12\x13.lofishop.AdRequest\x1a\x14.lofishop.AdResponse\"\x00\x62\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'general_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _CARTITEM._serialized_start=27
  _CARTITEM._serialized_end=75
  _ADDITEMREQUEST._serialized_start=77
  _ADDITEMREQUEST._serialized_end=144
  _EMPTYCARTREQUEST._serialized_start=146
  _EMPTYCARTREQUEST._serialized_end=181
  _GETCARTREQUEST._serialized_start=183
  _GETCARTREQUEST._serialized_end=216
  _CART._serialized_start=218
  _CART._serialized_end=276
  _EMPTY._serialized_start=278
  _EMPTY._serialized_end=285
  _LISTRECOMMENDATIONSREQUEST._serialized_start=287
  _LISTRECOMMENDATIONSREQUEST._serialized_end=353
  _LISTRECOMMENDATIONSRESPONSE._serialized_start=355
  _LISTRECOMMENDATIONSRESPONSE._serialized_end=405
  _PRODUCT._serialized_start=408
  _PRODUCT._serialized_end=537
  _LISTPRODUCTSRESPONSE._serialized_start=539
  _LISTPRODUCTSRESPONSE._serialized_end=598
  _GETPRODUCTREQUEST._serialized_start=600
  _GETPRODUCTREQUEST._serialized_end=631
  _SEARCHPRODUCTSREQUEST._serialized_start=633
  _SEARCHPRODUCTSREQUEST._serialized_end=671
  _SEARCHPRODUCTSRESPONSE._serialized_start=673
  _SEARCHPRODUCTSRESPONSE._serialized_end=733
  _GETQUOTEREQUEST._serialized_start=735
  _GETQUOTEREQUEST._serialized_end=823
  _GETQUOTERESPONSE._serialized_start=825
  _GETQUOTERESPONSE._serialized_end=878
  _SHIPORDERREQUEST._serialized_start=880
  _SHIPORDERREQUEST._serialized_end=969
  _SHIPORDERRESPONSE._serialized_start=971
  _SHIPORDERRESPONSE._serialized_end=1011
  _ADDRESS._serialized_start=1013
  _ADDRESS._serialized_end=1110
  _MONEY._serialized_start=1112
  _MONEY._serialized_end=1172
  _GETSUPPORTEDCURRENCIESRESPONSE._serialized_start=1174
  _GETSUPPORTEDCURRENCIESRESPONSE._serialized_end=1230
  _CURRENCYCONVERSIONREQUEST._serialized_start=1232
  _CURRENCYCONVERSIONREQUEST._serialized_end=1307
  _CREDITCARDINFO._serialized_start=1310
  _CREDITCARDINFO._serialized_end=1454
  _CHARGEREQUEST._serialized_start=1456
  _CHARGEREQUEST._serialized_end=1551
  _CHARGERESPONSE._serialized_start=1553
  _CHARGERESPONSE._serialized_end=1593
  _ORDERITEM._serialized_start=1595
  _ORDERITEM._serialized_end=1671
  _ORDERRESULT._serialized_start=1674
  _ORDERRESULT._serialized_end=1856
  _SENDORDERCONFIRMATIONREQUEST._serialized_start=1858
  _SENDORDERCONFIRMATIONREQUEST._serialized_end=1941
  _PLACEORDERREQUEST._serialized_start=1944
  _PLACEORDERREQUEST._serialized_end=2101
  _PLACEORDERRESPONSE._serialized_start=2103
  _PLACEORDERRESPONSE._serialized_end=2161
  _ADREQUEST._serialized_start=2163
  _ADREQUEST._serialized_end=2196
  _ADRESPONSE._serialized_start=2198
  _ADRESPONSE._serialized_end=2237
  _AD._serialized_start=2239
  _AD._serialized_end=2279
  _CARTSERVICE._serialized_start=2282
  _CARTSERVICE._serialized_end=2466
  _RECOMMENDATIONSERVICE._serialized_start=2468
  _RECOMMENDATIONSERVICE._serialized_end=2593
  _PRODUCTCATALOGSERVICE._serialized_start=2596
  _PRODUCTCATALOGSERVICE._serialized_end=2837
  _SHIPPINGSERVICE._serialized_start=2840
  _SHIPPINGSERVICE._serialized_end=2998
  _CURRENCYSERVICE._serialized_start=3001
  _CURRENCYSERVICE._serialized_end=3172
  _PAYMENTSERVICE._serialized_start=3174
  _PAYMENTSERVICE._serialized_end=3253
  _EMAILSERVICE._serialized_start=3255
  _EMAILSERVICE._serialized_end=3353
  _CHECKOUTSERVICE._serialized_start=3355
  _CHECKOUTSERVICE._serialized_end=3447
  _ADSERVICE._serialized_start=3449
  _ADSERVICE._serialized_end=3515
# @@protoc_insertion_point(module_scope)
