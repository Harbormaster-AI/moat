import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import HomePageComponent from './components/HomePageComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListMerchantComponent from './components/ListMerchantComponent';
import CreateMerchantComponent from './components/CreateMerchantComponent';
import ViewMerchantComponent from './components/ViewMerchantComponent';
import ListChannelComponent from './components/ListChannelComponent';
import CreateChannelComponent from './components/CreateChannelComponent';
import ViewChannelComponent from './components/ViewChannelComponent';
import ListBrandComponent from './components/ListBrandComponent';
import CreateBrandComponent from './components/CreateBrandComponent';
import ViewBrandComponent from './components/ViewBrandComponent';
import ListCatalogComponent from './components/ListCatalogComponent';
import CreateCatalogComponent from './components/CreateCatalogComponent';
import ViewCatalogComponent from './components/ViewCatalogComponent';
import ListCategoryComponent from './components/ListCategoryComponent';
import CreateCategoryComponent from './components/CreateCategoryComponent';
import ViewCategoryComponent from './components/ViewCategoryComponent';
import ListProductComponent from './components/ListProductComponent';
import CreateProductComponent from './components/CreateProductComponent';
import ViewProductComponent from './components/ViewProductComponent';
import ListProductVariantComponent from './components/ListProductVariantComponent';
import CreateProductVariantComponent from './components/CreateProductVariantComponent';
import ViewProductVariantComponent from './components/ViewProductVariantComponent';
import ListProductPricingComponent from './components/ListProductPricingComponent';
import CreateProductPricingComponent from './components/CreateProductPricingComponent';
import ViewProductPricingComponent from './components/ViewProductPricingComponent';
import ListMediaAssetComponent from './components/ListMediaAssetComponent';
import CreateMediaAssetComponent from './components/CreateMediaAssetComponent';
import ViewMediaAssetComponent from './components/ViewMediaAssetComponent';
import ListFulfillmentCenterComponent from './components/ListFulfillmentCenterComponent';
import CreateFulfillmentCenterComponent from './components/CreateFulfillmentCenterComponent';
import ViewFulfillmentCenterComponent from './components/ViewFulfillmentCenterComponent';
import ListInventoryItemComponent from './components/ListInventoryItemComponent';
import CreateInventoryItemComponent from './components/CreateInventoryItemComponent';
import ViewInventoryItemComponent from './components/ViewInventoryItemComponent';
import ListSupplierComponent from './components/ListSupplierComponent';
import CreateSupplierComponent from './components/CreateSupplierComponent';
import ViewSupplierComponent from './components/ViewSupplierComponent';
import ListSellerComponent from './components/ListSellerComponent';
import CreateSellerComponent from './components/CreateSellerComponent';
import ViewSellerComponent from './components/ViewSellerComponent';
import ListCustomerComponent from './components/ListCustomerComponent';
import CreateCustomerComponent from './components/CreateCustomerComponent';
import ViewCustomerComponent from './components/ViewCustomerComponent';
import ListCustomerAddressComponent from './components/ListCustomerAddressComponent';
import CreateCustomerAddressComponent from './components/CreateCustomerAddressComponent';
import ViewCustomerAddressComponent from './components/ViewCustomerAddressComponent';
import ListWishlistComponent from './components/ListWishlistComponent';
import CreateWishlistComponent from './components/CreateWishlistComponent';
import ViewWishlistComponent from './components/ViewWishlistComponent';
import ListWishlistItemComponent from './components/ListWishlistItemComponent';
import CreateWishlistItemComponent from './components/CreateWishlistItemComponent';
import ViewWishlistItemComponent from './components/ViewWishlistItemComponent';
import ListCartComponent from './components/ListCartComponent';
import CreateCartComponent from './components/CreateCartComponent';
import ViewCartComponent from './components/ViewCartComponent';
import ListCartItemComponent from './components/ListCartItemComponent';
import CreateCartItemComponent from './components/CreateCartItemComponent';
import ViewCartItemComponent from './components/ViewCartItemComponent';
import ListOrderComponent from './components/ListOrderComponent';
import CreateOrderComponent from './components/CreateOrderComponent';
import ViewOrderComponent from './components/ViewOrderComponent';
import ListOrderLineComponent from './components/ListOrderLineComponent';
import CreateOrderLineComponent from './components/CreateOrderLineComponent';
import ViewOrderLineComponent from './components/ViewOrderLineComponent';
import ListPaymentComponent from './components/ListPaymentComponent';
import CreatePaymentComponent from './components/CreatePaymentComponent';
import ViewPaymentComponent from './components/ViewPaymentComponent';
import ListRefundComponent from './components/ListRefundComponent';
import CreateRefundComponent from './components/CreateRefundComponent';
import ViewRefundComponent from './components/ViewRefundComponent';
import ListShipmentComponent from './components/ListShipmentComponent';
import CreateShipmentComponent from './components/CreateShipmentComponent';
import ViewShipmentComponent from './components/ViewShipmentComponent';
import ListShipmentItemComponent from './components/ListShipmentItemComponent';
import CreateShipmentItemComponent from './components/CreateShipmentItemComponent';
import ViewShipmentItemComponent from './components/ViewShipmentItemComponent';
import ListReturnRequestComponent from './components/ListReturnRequestComponent';
import CreateReturnRequestComponent from './components/CreateReturnRequestComponent';
import ViewReturnRequestComponent from './components/ViewReturnRequestComponent';
import ListReturnItemComponent from './components/ListReturnItemComponent';
import CreateReturnItemComponent from './components/CreateReturnItemComponent';
import ViewReturnItemComponent from './components/ViewReturnItemComponent';
import ListPromotionComponent from './components/ListPromotionComponent';
import CreatePromotionComponent from './components/CreatePromotionComponent';
import ViewPromotionComponent from './components/ViewPromotionComponent';
import ListCouponComponent from './components/ListCouponComponent';
import CreateCouponComponent from './components/CreateCouponComponent';
import ViewCouponComponent from './components/ViewCouponComponent';
import ListCouponRedemptionComponent from './components/ListCouponRedemptionComponent';
import CreateCouponRedemptionComponent from './components/CreateCouponRedemptionComponent';
import ViewCouponRedemptionComponent from './components/ViewCouponRedemptionComponent';
import ListTaxRuleComponent from './components/ListTaxRuleComponent';
import CreateTaxRuleComponent from './components/CreateTaxRuleComponent';
import ViewTaxRuleComponent from './components/ViewTaxRuleComponent';
import ListShippingMethodComponent from './components/ListShippingMethodComponent';
import CreateShippingMethodComponent from './components/CreateShippingMethodComponent';
import ViewShippingMethodComponent from './components/ViewShippingMethodComponent';
import ListCarrierServiceComponent from './components/ListCarrierServiceComponent';
import CreateCarrierServiceComponent from './components/CreateCarrierServiceComponent';
import ViewCarrierServiceComponent from './components/ViewCarrierServiceComponent';
import ListReviewComponent from './components/ListReviewComponent';
import CreateReviewComponent from './components/CreateReviewComponent';
import ViewReviewComponent from './components/ViewReviewComponent';
import ListSubscriptionComponent from './components/ListSubscriptionComponent';
import CreateSubscriptionComponent from './components/CreateSubscriptionComponent';
import ViewSubscriptionComponent from './components/ViewSubscriptionComponent';
import ListPaymentProviderComponent from './components/ListPaymentProviderComponent';
import CreatePaymentProviderComponent from './components/CreatePaymentProviderComponent';
import ViewPaymentProviderComponent from './components/ViewPaymentProviderComponent';
import ListInvoiceComponent from './components/ListInvoiceComponent';
import CreateInvoiceComponent from './components/CreateInvoiceComponent';
import ViewInvoiceComponent from './components/ViewInvoiceComponent';
import ListGiftCardComponent from './components/ListGiftCardComponent';
import CreateGiftCardComponent from './components/CreateGiftCardComponent';
import ViewGiftCardComponent from './components/ViewGiftCardComponent';
import ListGiftCardRedemptionComponent from './components/ListGiftCardRedemptionComponent';
import CreateGiftCardRedemptionComponent from './components/CreateGiftCardRedemptionComponent';
import ViewGiftCardRedemptionComponent from './components/ViewGiftCardRedemptionComponent';
import ListPayoutComponent from './components/ListPayoutComponent';
import CreatePayoutComponent from './components/CreatePayoutComponent';
import ViewPayoutComponent from './components/ViewPayoutComponent';
function App() {
  return (
    <div>
        <Router>
                <HeaderComponent className="header"/>
                <div className="container">
                    <Switch>
                          <Route path = "/" exact component = {HomePageComponent}></Route>
                            <Route path = "/merchants" component = {ListMerchantComponent}></Route>
                            <Route path = "/add-merchant/:id" component = {CreateMerchantComponent}></Route>
                            <Route path = "/view-merchant/:id" component = {ViewMerchantComponent}></Route>
                          {/* <Route path = "/update-merchant/:id" component = {UpdateMerchantComponent}></Route> */}
                            <Route path = "/channels" component = {ListChannelComponent}></Route>
                            <Route path = "/add-channel/:id" component = {CreateChannelComponent}></Route>
                            <Route path = "/view-channel/:id" component = {ViewChannelComponent}></Route>
                          {/* <Route path = "/update-channel/:id" component = {UpdateChannelComponent}></Route> */}
                            <Route path = "/brands" component = {ListBrandComponent}></Route>
                            <Route path = "/add-brand/:id" component = {CreateBrandComponent}></Route>
                            <Route path = "/view-brand/:id" component = {ViewBrandComponent}></Route>
                          {/* <Route path = "/update-brand/:id" component = {UpdateBrandComponent}></Route> */}
                            <Route path = "/catalogs" component = {ListCatalogComponent}></Route>
                            <Route path = "/add-catalog/:id" component = {CreateCatalogComponent}></Route>
                            <Route path = "/view-catalog/:id" component = {ViewCatalogComponent}></Route>
                          {/* <Route path = "/update-catalog/:id" component = {UpdateCatalogComponent}></Route> */}
                            <Route path = "/categorys" component = {ListCategoryComponent}></Route>
                            <Route path = "/add-category/:id" component = {CreateCategoryComponent}></Route>
                            <Route path = "/view-category/:id" component = {ViewCategoryComponent}></Route>
                          {/* <Route path = "/update-category/:id" component = {UpdateCategoryComponent}></Route> */}
                            <Route path = "/products" component = {ListProductComponent}></Route>
                            <Route path = "/add-product/:id" component = {CreateProductComponent}></Route>
                            <Route path = "/view-product/:id" component = {ViewProductComponent}></Route>
                          {/* <Route path = "/update-product/:id" component = {UpdateProductComponent}></Route> */}
                            <Route path = "/productVariants" component = {ListProductVariantComponent}></Route>
                            <Route path = "/add-productVariant/:id" component = {CreateProductVariantComponent}></Route>
                            <Route path = "/view-productVariant/:id" component = {ViewProductVariantComponent}></Route>
                          {/* <Route path = "/update-productVariant/:id" component = {UpdateProductVariantComponent}></Route> */}
                            <Route path = "/productPricings" component = {ListProductPricingComponent}></Route>
                            <Route path = "/add-productPricing/:id" component = {CreateProductPricingComponent}></Route>
                            <Route path = "/view-productPricing/:id" component = {ViewProductPricingComponent}></Route>
                          {/* <Route path = "/update-productPricing/:id" component = {UpdateProductPricingComponent}></Route> */}
                            <Route path = "/mediaAssets" component = {ListMediaAssetComponent}></Route>
                            <Route path = "/add-mediaAsset/:id" component = {CreateMediaAssetComponent}></Route>
                            <Route path = "/view-mediaAsset/:id" component = {ViewMediaAssetComponent}></Route>
                          {/* <Route path = "/update-mediaAsset/:id" component = {UpdateMediaAssetComponent}></Route> */}
                            <Route path = "/fulfillmentCenters" component = {ListFulfillmentCenterComponent}></Route>
                            <Route path = "/add-fulfillmentCenter/:id" component = {CreateFulfillmentCenterComponent}></Route>
                            <Route path = "/view-fulfillmentCenter/:id" component = {ViewFulfillmentCenterComponent}></Route>
                          {/* <Route path = "/update-fulfillmentCenter/:id" component = {UpdateFulfillmentCenterComponent}></Route> */}
                            <Route path = "/inventoryItems" component = {ListInventoryItemComponent}></Route>
                            <Route path = "/add-inventoryItem/:id" component = {CreateInventoryItemComponent}></Route>
                            <Route path = "/view-inventoryItem/:id" component = {ViewInventoryItemComponent}></Route>
                          {/* <Route path = "/update-inventoryItem/:id" component = {UpdateInventoryItemComponent}></Route> */}
                            <Route path = "/suppliers" component = {ListSupplierComponent}></Route>
                            <Route path = "/add-supplier/:id" component = {CreateSupplierComponent}></Route>
                            <Route path = "/view-supplier/:id" component = {ViewSupplierComponent}></Route>
                          {/* <Route path = "/update-supplier/:id" component = {UpdateSupplierComponent}></Route> */}
                            <Route path = "/sellers" component = {ListSellerComponent}></Route>
                            <Route path = "/add-seller/:id" component = {CreateSellerComponent}></Route>
                            <Route path = "/view-seller/:id" component = {ViewSellerComponent}></Route>
                          {/* <Route path = "/update-seller/:id" component = {UpdateSellerComponent}></Route> */}
                            <Route path = "/customers" component = {ListCustomerComponent}></Route>
                            <Route path = "/add-customer/:id" component = {CreateCustomerComponent}></Route>
                            <Route path = "/view-customer/:id" component = {ViewCustomerComponent}></Route>
                          {/* <Route path = "/update-customer/:id" component = {UpdateCustomerComponent}></Route> */}
                            <Route path = "/customerAddresss" component = {ListCustomerAddressComponent}></Route>
                            <Route path = "/add-customerAddress/:id" component = {CreateCustomerAddressComponent}></Route>
                            <Route path = "/view-customerAddress/:id" component = {ViewCustomerAddressComponent}></Route>
                          {/* <Route path = "/update-customerAddress/:id" component = {UpdateCustomerAddressComponent}></Route> */}
                            <Route path = "/wishlists" component = {ListWishlistComponent}></Route>
                            <Route path = "/add-wishlist/:id" component = {CreateWishlistComponent}></Route>
                            <Route path = "/view-wishlist/:id" component = {ViewWishlistComponent}></Route>
                          {/* <Route path = "/update-wishlist/:id" component = {UpdateWishlistComponent}></Route> */}
                            <Route path = "/wishlistItems" component = {ListWishlistItemComponent}></Route>
                            <Route path = "/add-wishlistItem/:id" component = {CreateWishlistItemComponent}></Route>
                            <Route path = "/view-wishlistItem/:id" component = {ViewWishlistItemComponent}></Route>
                          {/* <Route path = "/update-wishlistItem/:id" component = {UpdateWishlistItemComponent}></Route> */}
                            <Route path = "/carts" component = {ListCartComponent}></Route>
                            <Route path = "/add-cart/:id" component = {CreateCartComponent}></Route>
                            <Route path = "/view-cart/:id" component = {ViewCartComponent}></Route>
                          {/* <Route path = "/update-cart/:id" component = {UpdateCartComponent}></Route> */}
                            <Route path = "/cartItems" component = {ListCartItemComponent}></Route>
                            <Route path = "/add-cartItem/:id" component = {CreateCartItemComponent}></Route>
                            <Route path = "/view-cartItem/:id" component = {ViewCartItemComponent}></Route>
                          {/* <Route path = "/update-cartItem/:id" component = {UpdateCartItemComponent}></Route> */}
                            <Route path = "/orders" component = {ListOrderComponent}></Route>
                            <Route path = "/add-order/:id" component = {CreateOrderComponent}></Route>
                            <Route path = "/view-order/:id" component = {ViewOrderComponent}></Route>
                          {/* <Route path = "/update-order/:id" component = {UpdateOrderComponent}></Route> */}
                            <Route path = "/orderLines" component = {ListOrderLineComponent}></Route>
                            <Route path = "/add-orderLine/:id" component = {CreateOrderLineComponent}></Route>
                            <Route path = "/view-orderLine/:id" component = {ViewOrderLineComponent}></Route>
                          {/* <Route path = "/update-orderLine/:id" component = {UpdateOrderLineComponent}></Route> */}
                            <Route path = "/payments" component = {ListPaymentComponent}></Route>
                            <Route path = "/add-payment/:id" component = {CreatePaymentComponent}></Route>
                            <Route path = "/view-payment/:id" component = {ViewPaymentComponent}></Route>
                          {/* <Route path = "/update-payment/:id" component = {UpdatePaymentComponent}></Route> */}
                            <Route path = "/refunds" component = {ListRefundComponent}></Route>
                            <Route path = "/add-refund/:id" component = {CreateRefundComponent}></Route>
                            <Route path = "/view-refund/:id" component = {ViewRefundComponent}></Route>
                          {/* <Route path = "/update-refund/:id" component = {UpdateRefundComponent}></Route> */}
                            <Route path = "/shipments" component = {ListShipmentComponent}></Route>
                            <Route path = "/add-shipment/:id" component = {CreateShipmentComponent}></Route>
                            <Route path = "/view-shipment/:id" component = {ViewShipmentComponent}></Route>
                          {/* <Route path = "/update-shipment/:id" component = {UpdateShipmentComponent}></Route> */}
                            <Route path = "/shipmentItems" component = {ListShipmentItemComponent}></Route>
                            <Route path = "/add-shipmentItem/:id" component = {CreateShipmentItemComponent}></Route>
                            <Route path = "/view-shipmentItem/:id" component = {ViewShipmentItemComponent}></Route>
                          {/* <Route path = "/update-shipmentItem/:id" component = {UpdateShipmentItemComponent}></Route> */}
                            <Route path = "/returnRequests" component = {ListReturnRequestComponent}></Route>
                            <Route path = "/add-returnRequest/:id" component = {CreateReturnRequestComponent}></Route>
                            <Route path = "/view-returnRequest/:id" component = {ViewReturnRequestComponent}></Route>
                          {/* <Route path = "/update-returnRequest/:id" component = {UpdateReturnRequestComponent}></Route> */}
                            <Route path = "/returnItems" component = {ListReturnItemComponent}></Route>
                            <Route path = "/add-returnItem/:id" component = {CreateReturnItemComponent}></Route>
                            <Route path = "/view-returnItem/:id" component = {ViewReturnItemComponent}></Route>
                          {/* <Route path = "/update-returnItem/:id" component = {UpdateReturnItemComponent}></Route> */}
                            <Route path = "/promotions" component = {ListPromotionComponent}></Route>
                            <Route path = "/add-promotion/:id" component = {CreatePromotionComponent}></Route>
                            <Route path = "/view-promotion/:id" component = {ViewPromotionComponent}></Route>
                          {/* <Route path = "/update-promotion/:id" component = {UpdatePromotionComponent}></Route> */}
                            <Route path = "/coupons" component = {ListCouponComponent}></Route>
                            <Route path = "/add-coupon/:id" component = {CreateCouponComponent}></Route>
                            <Route path = "/view-coupon/:id" component = {ViewCouponComponent}></Route>
                          {/* <Route path = "/update-coupon/:id" component = {UpdateCouponComponent}></Route> */}
                            <Route path = "/couponRedemptions" component = {ListCouponRedemptionComponent}></Route>
                            <Route path = "/add-couponRedemption/:id" component = {CreateCouponRedemptionComponent}></Route>
                            <Route path = "/view-couponRedemption/:id" component = {ViewCouponRedemptionComponent}></Route>
                          {/* <Route path = "/update-couponRedemption/:id" component = {UpdateCouponRedemptionComponent}></Route> */}
                            <Route path = "/taxRules" component = {ListTaxRuleComponent}></Route>
                            <Route path = "/add-taxRule/:id" component = {CreateTaxRuleComponent}></Route>
                            <Route path = "/view-taxRule/:id" component = {ViewTaxRuleComponent}></Route>
                          {/* <Route path = "/update-taxRule/:id" component = {UpdateTaxRuleComponent}></Route> */}
                            <Route path = "/shippingMethods" component = {ListShippingMethodComponent}></Route>
                            <Route path = "/add-shippingMethod/:id" component = {CreateShippingMethodComponent}></Route>
                            <Route path = "/view-shippingMethod/:id" component = {ViewShippingMethodComponent}></Route>
                          {/* <Route path = "/update-shippingMethod/:id" component = {UpdateShippingMethodComponent}></Route> */}
                            <Route path = "/carrierServices" component = {ListCarrierServiceComponent}></Route>
                            <Route path = "/add-carrierService/:id" component = {CreateCarrierServiceComponent}></Route>
                            <Route path = "/view-carrierService/:id" component = {ViewCarrierServiceComponent}></Route>
                          {/* <Route path = "/update-carrierService/:id" component = {UpdateCarrierServiceComponent}></Route> */}
                            <Route path = "/reviews" component = {ListReviewComponent}></Route>
                            <Route path = "/add-review/:id" component = {CreateReviewComponent}></Route>
                            <Route path = "/view-review/:id" component = {ViewReviewComponent}></Route>
                          {/* <Route path = "/update-review/:id" component = {UpdateReviewComponent}></Route> */}
                            <Route path = "/subscriptions" component = {ListSubscriptionComponent}></Route>
                            <Route path = "/add-subscription/:id" component = {CreateSubscriptionComponent}></Route>
                            <Route path = "/view-subscription/:id" component = {ViewSubscriptionComponent}></Route>
                          {/* <Route path = "/update-subscription/:id" component = {UpdateSubscriptionComponent}></Route> */}
                            <Route path = "/paymentProviders" component = {ListPaymentProviderComponent}></Route>
                            <Route path = "/add-paymentProvider/:id" component = {CreatePaymentProviderComponent}></Route>
                            <Route path = "/view-paymentProvider/:id" component = {ViewPaymentProviderComponent}></Route>
                          {/* <Route path = "/update-paymentProvider/:id" component = {UpdatePaymentProviderComponent}></Route> */}
                            <Route path = "/invoices" component = {ListInvoiceComponent}></Route>
                            <Route path = "/add-invoice/:id" component = {CreateInvoiceComponent}></Route>
                            <Route path = "/view-invoice/:id" component = {ViewInvoiceComponent}></Route>
                          {/* <Route path = "/update-invoice/:id" component = {UpdateInvoiceComponent}></Route> */}
                            <Route path = "/giftCards" component = {ListGiftCardComponent}></Route>
                            <Route path = "/add-giftCard/:id" component = {CreateGiftCardComponent}></Route>
                            <Route path = "/view-giftCard/:id" component = {ViewGiftCardComponent}></Route>
                          {/* <Route path = "/update-giftCard/:id" component = {UpdateGiftCardComponent}></Route> */}
                            <Route path = "/giftCardRedemptions" component = {ListGiftCardRedemptionComponent}></Route>
                            <Route path = "/add-giftCardRedemption/:id" component = {CreateGiftCardRedemptionComponent}></Route>
                            <Route path = "/view-giftCardRedemption/:id" component = {ViewGiftCardRedemptionComponent}></Route>
                          {/* <Route path = "/update-giftCardRedemption/:id" component = {UpdateGiftCardRedemptionComponent}></Route> */}
                            <Route path = "/payouts" component = {ListPayoutComponent}></Route>
                            <Route path = "/add-payout/:id" component = {CreatePayoutComponent}></Route>
                            <Route path = "/view-payout/:id" component = {ViewPayoutComponent}></Route>
                          {/* <Route path = "/update-payout/:id" component = {UpdatePayoutComponent}></Route> */}
                    </Switch>
                </div>
              <FooterComponent />
        </Router>
    </div>
    
  );
}

export default App;
