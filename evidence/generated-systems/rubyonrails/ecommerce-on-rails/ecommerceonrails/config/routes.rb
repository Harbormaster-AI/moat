Rails.application.routes.draw do
  root "application#health"
  resources :merchants do
    resources :channels
    resources :brands
    resources :fulfillmentcenters
    resources :taxrules
    resources :paymentproviders
    resources :sellers
    resources :promotions
  end
  resources :channels do
    resource :merchant
    resources :catalogs
    resources :promotions
    resources :shippingmethods
    resources :paymentproviders
  end
  resources :brands do
    resource :merchant
    resources :products
  end
  resources :catalogs do
    resource :channel
    resources :categories
  end
  resources :categorys do
    resource :catalog
    resource :parentcategory
    resources :subcategories
    resources :products
  end
  resources :products do
    resource :brand
    resources :categories
    resources :variants
    resources :mediaassets
    resources :reviews
    resource :seller
  end
  resources :productvariants do
    resource :product
    resources :pricing
    resources :inventoryitems
    resources :mediaassets
    resources :subscriptions
    resources :cartitems
    resources :orderlines
    resources :wishlistitems
  end
  resources :productpricings do
    resource :variant
    resource :channel
  end
  resources :mediaassets do
    resource :product
    resource :variant
  end
  resources :fulfillmentcenters do
    resource :merchant
    resources :inventoryitems
    resources :shipments
  end
  resources :inventoryitems do
    resource :variant
    resource :fulfillmentcenter
  end
  resources :suppliers do
    resource :merchant
    resources :products
    resources :fulfillmentcenters
  end
  resources :sellers do
    resource :merchant
    resources :products
    resources :payouts
    resources :orders
  end
  resources :customers do
    resources :addresses
    resources :carts
    resources :orders
    resources :payments
    resources :reviews
    resources :wishlists
    resources :subscriptions
    resources :couponredemptions
    resources :giftcards
  end
  resources :customeraddresss do
    resource :customer
  end
  resources :wishlists do
    resource :customer
    resources :items
  end
  resources :wishlistitems do
    resource :wishlist
    resource :variant
  end
  resources :carts do
    resource :customer
    resource :channel
    resources :items
    resources :appliedpromotions
  end
  resources :cartitems do
    resource :cart
    resource :variant
    resources :appliedpromotions
  end
  resources :orders do
    resource :customer
    resource :channel
    resources :orderlines
    resources :payments
    resources :shipments
    resources :refunds
    resources :appliedpromotions
    resource :seller
    resources :giftcardredemptions
    resources :couponredemptions
    resources :returnrequests
    resource :invoice
  end
  resources :orderlines do
    resource :order
    resource :variant
    resources :appliedpromotions
  end
  resources :payments do
    resource :order
    resource :customer
    resource :paymentprovider
    resources :refunds
  end
  resources :refunds do
    resource :payment
    resource :order
  end
  resources :shipments do
    resource :order
    resources :shipmentitems
    resource :fulfillmentcenter
  end
  resources :shipmentitems do
    resource :shipment
    resource :orderline
  end
  resources :returnrequests do
    resource :order
    resources :items
    resource :refund
    resource :shipment
  end
  resources :returnitems do
    resource :returnrequest
    resource :orderline
  end
  resources :promotions do
    resource :merchant
    resources :channels
    resources :applicableproducts
    resources :applicablecategories
    resources :coupons
  end
  resources :coupons do
    resource :promotion
    resources :redemptions
  end
  resources :couponredemptions do
    resource :coupon
    resource :order
    resource :customer
  end
  resources :taxrules do
    resource :merchant
    resources :channels
  end
  resources :shippingmethods do
    resources :channels
    resource :carrierservice
  end
  resources :carrierservices do
    resources :shippingmethods
  end
  resources :reviews do
    resource :product
    resource :customer
    resource :order
  end
  resources :subscriptions do
    resource :customer
    resource :variant
    resource :paymentprovider
    resource :channel
  end
  resources :paymentproviders do
    resource :merchant
    resources :channels
    resources :payments
    resources :subscriptions
  end
  resources :invoices do
    resource :order
  end
  resources :giftcards do
    resource :customer
    resource :issuedorder
    resources :redemptions
  end
  resources :giftcardredemptions do
    resource :giftcard
    resource :order
  end
  resources :payouts do
    resource :seller
    resources :orders
  end
end
