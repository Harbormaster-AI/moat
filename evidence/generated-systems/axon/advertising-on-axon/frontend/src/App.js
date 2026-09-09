import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import HomePageComponent from './components/HomePageComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListAgencyComponent from './components/ListAgencyComponent';
import CreateAgencyComponent from './components/CreateAgencyComponent';
import ViewAgencyComponent from './components/ViewAgencyComponent';
import ListTeamComponent from './components/ListTeamComponent';
import CreateTeamComponent from './components/CreateTeamComponent';
import ViewTeamComponent from './components/ViewTeamComponent';
import ListUserComponent from './components/ListUserComponent';
import CreateUserComponent from './components/CreateUserComponent';
import ViewUserComponent from './components/ViewUserComponent';
import ListAdvertiserComponent from './components/ListAdvertiserComponent';
import CreateAdvertiserComponent from './components/CreateAdvertiserComponent';
import ViewAdvertiserComponent from './components/ViewAdvertiserComponent';
import ListBillingProfileComponent from './components/ListBillingProfileComponent';
import CreateBillingProfileComponent from './components/CreateBillingProfileComponent';
import ViewBillingProfileComponent from './components/ViewBillingProfileComponent';
import ListPaymentMethodComponent from './components/ListPaymentMethodComponent';
import CreatePaymentMethodComponent from './components/CreatePaymentMethodComponent';
import ViewPaymentMethodComponent from './components/ViewPaymentMethodComponent';
import ListAdAccountComponent from './components/ListAdAccountComponent';
import CreateAdAccountComponent from './components/CreateAdAccountComponent';
import ViewAdAccountComponent from './components/ViewAdAccountComponent';
import ListDSPComponent from './components/ListDSPComponent';
import CreateDSPComponent from './components/CreateDSPComponent';
import ViewDSPComponent from './components/ViewDSPComponent';
import ListCampaignComponent from './components/ListCampaignComponent';
import CreateCampaignComponent from './components/CreateCampaignComponent';
import ViewCampaignComponent from './components/ViewCampaignComponent';
import ListKPIComponent from './components/ListKPIComponent';
import CreateKPIComponent from './components/CreateKPIComponent';
import ViewKPIComponent from './components/ViewKPIComponent';
import ListAudienceSegmentComponent from './components/ListAudienceSegmentComponent';
import CreateAudienceSegmentComponent from './components/CreateAudienceSegmentComponent';
import ViewAudienceSegmentComponent from './components/ViewAudienceSegmentComponent';
import ListDataProviderComponent from './components/ListDataProviderComponent';
import CreateDataProviderComponent from './components/CreateDataProviderComponent';
import ViewDataProviderComponent from './components/ViewDataProviderComponent';
import ListLineItemComponent from './components/ListLineItemComponent';
import CreateLineItemComponent from './components/CreateLineItemComponent';
import ViewLineItemComponent from './components/ViewLineItemComponent';
import ListTargetingProfileComponent from './components/ListTargetingProfileComponent';
import CreateTargetingProfileComponent from './components/CreateTargetingProfileComponent';
import ViewTargetingProfileComponent from './components/ViewTargetingProfileComponent';
import ListDeviceCriterionComponent from './components/ListDeviceCriterionComponent';
import CreateDeviceCriterionComponent from './components/CreateDeviceCriterionComponent';
import ViewDeviceCriterionComponent from './components/ViewDeviceCriterionComponent';
import ListBrandSafetyPolicyComponent from './components/ListBrandSafetyPolicyComponent';
import CreateBrandSafetyPolicyComponent from './components/CreateBrandSafetyPolicyComponent';
import ViewBrandSafetyPolicyComponent from './components/ViewBrandSafetyPolicyComponent';
import ListContentCategoryComponent from './components/ListContentCategoryComponent';
import CreateContentCategoryComponent from './components/CreateContentCategoryComponent';
import ViewContentCategoryComponent from './components/ViewContentCategoryComponent';
import ListPublisherComponent from './components/ListPublisherComponent';
import CreatePublisherComponent from './components/CreatePublisherComponent';
import ViewPublisherComponent from './components/ViewPublisherComponent';
import ListInventorySourceComponent from './components/ListInventorySourceComponent';
import CreateInventorySourceComponent from './components/CreateInventorySourceComponent';
import ViewInventorySourceComponent from './components/ViewInventorySourceComponent';
import ListAdSlotComponent from './components/ListAdSlotComponent';
import CreateAdSlotComponent from './components/CreateAdSlotComponent';
import ViewAdSlotComponent from './components/ViewAdSlotComponent';
import ListDealComponent from './components/ListDealComponent';
import CreateDealComponent from './components/CreateDealComponent';
import ViewDealComponent from './components/ViewDealComponent';
import ListPlacementComponent from './components/ListPlacementComponent';
import CreatePlacementComponent from './components/CreatePlacementComponent';
import ViewPlacementComponent from './components/ViewPlacementComponent';
import ListCreativeAssetComponent from './components/ListCreativeAssetComponent';
import CreateCreativeAssetComponent from './components/CreateCreativeAssetComponent';
import ViewCreativeAssetComponent from './components/ViewCreativeAssetComponent';
import ListCreativeFileComponent from './components/ListCreativeFileComponent';
import CreateCreativeFileComponent from './components/CreateCreativeFileComponent';
import ViewCreativeFileComponent from './components/ViewCreativeFileComponent';
import ListCreativeVariationComponent from './components/ListCreativeVariationComponent';
import CreateCreativeVariationComponent from './components/CreateCreativeVariationComponent';
import ViewCreativeVariationComponent from './components/ViewCreativeVariationComponent';
import ListCreativeApprovalComponent from './components/ListCreativeApprovalComponent';
import CreateCreativeApprovalComponent from './components/CreateCreativeApprovalComponent';
import ViewCreativeApprovalComponent from './components/ViewCreativeApprovalComponent';
import ListTrackingPixelComponent from './components/ListTrackingPixelComponent';
import CreateTrackingPixelComponent from './components/CreateTrackingPixelComponent';
import ViewTrackingPixelComponent from './components/ViewTrackingPixelComponent';
import ListConversionEventComponent from './components/ListConversionEventComponent';
import CreateConversionEventComponent from './components/CreateConversionEventComponent';
import ViewConversionEventComponent from './components/ViewConversionEventComponent';
import ListPerformanceMetricComponent from './components/ListPerformanceMetricComponent';
import CreatePerformanceMetricComponent from './components/CreatePerformanceMetricComponent';
import ViewPerformanceMetricComponent from './components/ViewPerformanceMetricComponent';
import ListReportComponent from './components/ListReportComponent';
import CreateReportComponent from './components/CreateReportComponent';
import ViewReportComponent from './components/ViewReportComponent';
import ListInsertionOrderComponent from './components/ListInsertionOrderComponent';
import CreateInsertionOrderComponent from './components/CreateInsertionOrderComponent';
import ViewInsertionOrderComponent from './components/ViewInsertionOrderComponent';
import ListRateCardComponent from './components/ListRateCardComponent';
import CreateRateCardComponent from './components/CreateRateCardComponent';
import ViewRateCardComponent from './components/ViewRateCardComponent';
import ListRateComponent from './components/ListRateComponent';
import CreateRateComponent from './components/CreateRateComponent';
import ViewRateComponent from './components/ViewRateComponent';
import ListExperimentComponent from './components/ListExperimentComponent';
import CreateExperimentComponent from './components/CreateExperimentComponent';
import ViewExperimentComponent from './components/ViewExperimentComponent';
import ListExperimentVariantComponent from './components/ListExperimentVariantComponent';
import CreateExperimentVariantComponent from './components/CreateExperimentVariantComponent';
import ViewExperimentVariantComponent from './components/ViewExperimentVariantComponent';
import ListGeoRegionComponent from './components/ListGeoRegionComponent';
import CreateGeoRegionComponent from './components/CreateGeoRegionComponent';
import ViewGeoRegionComponent from './components/ViewGeoRegionComponent';
function App() {
  return (
    <div>
        <Router>
                <HeaderComponent className="header"/>
                <div className="container">
                    <Switch>
                          <Route path = "/" exact component = {HomePageComponent}></Route>
                            <Route path = "/agencys" component = {ListAgencyComponent}></Route>
                            <Route path = "/add-agency/:id" component = {CreateAgencyComponent}></Route>
                            <Route path = "/view-agency/:id" component = {ViewAgencyComponent}></Route>
                          {/* <Route path = "/update-agency/:id" component = {UpdateAgencyComponent}></Route> */}
                            <Route path = "/teams" component = {ListTeamComponent}></Route>
                            <Route path = "/add-team/:id" component = {CreateTeamComponent}></Route>
                            <Route path = "/view-team/:id" component = {ViewTeamComponent}></Route>
                          {/* <Route path = "/update-team/:id" component = {UpdateTeamComponent}></Route> */}
                            <Route path = "/users" component = {ListUserComponent}></Route>
                            <Route path = "/add-user/:id" component = {CreateUserComponent}></Route>
                            <Route path = "/view-user/:id" component = {ViewUserComponent}></Route>
                          {/* <Route path = "/update-user/:id" component = {UpdateUserComponent}></Route> */}
                            <Route path = "/advertisers" component = {ListAdvertiserComponent}></Route>
                            <Route path = "/add-advertiser/:id" component = {CreateAdvertiserComponent}></Route>
                            <Route path = "/view-advertiser/:id" component = {ViewAdvertiserComponent}></Route>
                          {/* <Route path = "/update-advertiser/:id" component = {UpdateAdvertiserComponent}></Route> */}
                            <Route path = "/billingProfiles" component = {ListBillingProfileComponent}></Route>
                            <Route path = "/add-billingProfile/:id" component = {CreateBillingProfileComponent}></Route>
                            <Route path = "/view-billingProfile/:id" component = {ViewBillingProfileComponent}></Route>
                          {/* <Route path = "/update-billingProfile/:id" component = {UpdateBillingProfileComponent}></Route> */}
                            <Route path = "/paymentMethods" component = {ListPaymentMethodComponent}></Route>
                            <Route path = "/add-paymentMethod/:id" component = {CreatePaymentMethodComponent}></Route>
                            <Route path = "/view-paymentMethod/:id" component = {ViewPaymentMethodComponent}></Route>
                          {/* <Route path = "/update-paymentMethod/:id" component = {UpdatePaymentMethodComponent}></Route> */}
                            <Route path = "/adAccounts" component = {ListAdAccountComponent}></Route>
                            <Route path = "/add-adAccount/:id" component = {CreateAdAccountComponent}></Route>
                            <Route path = "/view-adAccount/:id" component = {ViewAdAccountComponent}></Route>
                          {/* <Route path = "/update-adAccount/:id" component = {UpdateAdAccountComponent}></Route> */}
                            <Route path = "/dSPs" component = {ListDSPComponent}></Route>
                            <Route path = "/add-dSP/:id" component = {CreateDSPComponent}></Route>
                            <Route path = "/view-dSP/:id" component = {ViewDSPComponent}></Route>
                          {/* <Route path = "/update-dSP/:id" component = {UpdateDSPComponent}></Route> */}
                            <Route path = "/campaigns" component = {ListCampaignComponent}></Route>
                            <Route path = "/add-campaign/:id" component = {CreateCampaignComponent}></Route>
                            <Route path = "/view-campaign/:id" component = {ViewCampaignComponent}></Route>
                          {/* <Route path = "/update-campaign/:id" component = {UpdateCampaignComponent}></Route> */}
                            <Route path = "/kPIs" component = {ListKPIComponent}></Route>
                            <Route path = "/add-kPI/:id" component = {CreateKPIComponent}></Route>
                            <Route path = "/view-kPI/:id" component = {ViewKPIComponent}></Route>
                          {/* <Route path = "/update-kPI/:id" component = {UpdateKPIComponent}></Route> */}
                            <Route path = "/audienceSegments" component = {ListAudienceSegmentComponent}></Route>
                            <Route path = "/add-audienceSegment/:id" component = {CreateAudienceSegmentComponent}></Route>
                            <Route path = "/view-audienceSegment/:id" component = {ViewAudienceSegmentComponent}></Route>
                          {/* <Route path = "/update-audienceSegment/:id" component = {UpdateAudienceSegmentComponent}></Route> */}
                            <Route path = "/dataProviders" component = {ListDataProviderComponent}></Route>
                            <Route path = "/add-dataProvider/:id" component = {CreateDataProviderComponent}></Route>
                            <Route path = "/view-dataProvider/:id" component = {ViewDataProviderComponent}></Route>
                          {/* <Route path = "/update-dataProvider/:id" component = {UpdateDataProviderComponent}></Route> */}
                            <Route path = "/lineItems" component = {ListLineItemComponent}></Route>
                            <Route path = "/add-lineItem/:id" component = {CreateLineItemComponent}></Route>
                            <Route path = "/view-lineItem/:id" component = {ViewLineItemComponent}></Route>
                          {/* <Route path = "/update-lineItem/:id" component = {UpdateLineItemComponent}></Route> */}
                            <Route path = "/targetingProfiles" component = {ListTargetingProfileComponent}></Route>
                            <Route path = "/add-targetingProfile/:id" component = {CreateTargetingProfileComponent}></Route>
                            <Route path = "/view-targetingProfile/:id" component = {ViewTargetingProfileComponent}></Route>
                          {/* <Route path = "/update-targetingProfile/:id" component = {UpdateTargetingProfileComponent}></Route> */}
                            <Route path = "/deviceCriterions" component = {ListDeviceCriterionComponent}></Route>
                            <Route path = "/add-deviceCriterion/:id" component = {CreateDeviceCriterionComponent}></Route>
                            <Route path = "/view-deviceCriterion/:id" component = {ViewDeviceCriterionComponent}></Route>
                          {/* <Route path = "/update-deviceCriterion/:id" component = {UpdateDeviceCriterionComponent}></Route> */}
                            <Route path = "/brandSafetyPolicys" component = {ListBrandSafetyPolicyComponent}></Route>
                            <Route path = "/add-brandSafetyPolicy/:id" component = {CreateBrandSafetyPolicyComponent}></Route>
                            <Route path = "/view-brandSafetyPolicy/:id" component = {ViewBrandSafetyPolicyComponent}></Route>
                          {/* <Route path = "/update-brandSafetyPolicy/:id" component = {UpdateBrandSafetyPolicyComponent}></Route> */}
                            <Route path = "/contentCategorys" component = {ListContentCategoryComponent}></Route>
                            <Route path = "/add-contentCategory/:id" component = {CreateContentCategoryComponent}></Route>
                            <Route path = "/view-contentCategory/:id" component = {ViewContentCategoryComponent}></Route>
                          {/* <Route path = "/update-contentCategory/:id" component = {UpdateContentCategoryComponent}></Route> */}
                            <Route path = "/publishers" component = {ListPublisherComponent}></Route>
                            <Route path = "/add-publisher/:id" component = {CreatePublisherComponent}></Route>
                            <Route path = "/view-publisher/:id" component = {ViewPublisherComponent}></Route>
                          {/* <Route path = "/update-publisher/:id" component = {UpdatePublisherComponent}></Route> */}
                            <Route path = "/inventorySources" component = {ListInventorySourceComponent}></Route>
                            <Route path = "/add-inventorySource/:id" component = {CreateInventorySourceComponent}></Route>
                            <Route path = "/view-inventorySource/:id" component = {ViewInventorySourceComponent}></Route>
                          {/* <Route path = "/update-inventorySource/:id" component = {UpdateInventorySourceComponent}></Route> */}
                            <Route path = "/adSlots" component = {ListAdSlotComponent}></Route>
                            <Route path = "/add-adSlot/:id" component = {CreateAdSlotComponent}></Route>
                            <Route path = "/view-adSlot/:id" component = {ViewAdSlotComponent}></Route>
                          {/* <Route path = "/update-adSlot/:id" component = {UpdateAdSlotComponent}></Route> */}
                            <Route path = "/deals" component = {ListDealComponent}></Route>
                            <Route path = "/add-deal/:id" component = {CreateDealComponent}></Route>
                            <Route path = "/view-deal/:id" component = {ViewDealComponent}></Route>
                          {/* <Route path = "/update-deal/:id" component = {UpdateDealComponent}></Route> */}
                            <Route path = "/placements" component = {ListPlacementComponent}></Route>
                            <Route path = "/add-placement/:id" component = {CreatePlacementComponent}></Route>
                            <Route path = "/view-placement/:id" component = {ViewPlacementComponent}></Route>
                          {/* <Route path = "/update-placement/:id" component = {UpdatePlacementComponent}></Route> */}
                            <Route path = "/creativeAssets" component = {ListCreativeAssetComponent}></Route>
                            <Route path = "/add-creativeAsset/:id" component = {CreateCreativeAssetComponent}></Route>
                            <Route path = "/view-creativeAsset/:id" component = {ViewCreativeAssetComponent}></Route>
                          {/* <Route path = "/update-creativeAsset/:id" component = {UpdateCreativeAssetComponent}></Route> */}
                            <Route path = "/creativeFiles" component = {ListCreativeFileComponent}></Route>
                            <Route path = "/add-creativeFile/:id" component = {CreateCreativeFileComponent}></Route>
                            <Route path = "/view-creativeFile/:id" component = {ViewCreativeFileComponent}></Route>
                          {/* <Route path = "/update-creativeFile/:id" component = {UpdateCreativeFileComponent}></Route> */}
                            <Route path = "/creativeVariations" component = {ListCreativeVariationComponent}></Route>
                            <Route path = "/add-creativeVariation/:id" component = {CreateCreativeVariationComponent}></Route>
                            <Route path = "/view-creativeVariation/:id" component = {ViewCreativeVariationComponent}></Route>
                          {/* <Route path = "/update-creativeVariation/:id" component = {UpdateCreativeVariationComponent}></Route> */}
                            <Route path = "/creativeApprovals" component = {ListCreativeApprovalComponent}></Route>
                            <Route path = "/add-creativeApproval/:id" component = {CreateCreativeApprovalComponent}></Route>
                            <Route path = "/view-creativeApproval/:id" component = {ViewCreativeApprovalComponent}></Route>
                          {/* <Route path = "/update-creativeApproval/:id" component = {UpdateCreativeApprovalComponent}></Route> */}
                            <Route path = "/trackingPixels" component = {ListTrackingPixelComponent}></Route>
                            <Route path = "/add-trackingPixel/:id" component = {CreateTrackingPixelComponent}></Route>
                            <Route path = "/view-trackingPixel/:id" component = {ViewTrackingPixelComponent}></Route>
                          {/* <Route path = "/update-trackingPixel/:id" component = {UpdateTrackingPixelComponent}></Route> */}
                            <Route path = "/conversionEvents" component = {ListConversionEventComponent}></Route>
                            <Route path = "/add-conversionEvent/:id" component = {CreateConversionEventComponent}></Route>
                            <Route path = "/view-conversionEvent/:id" component = {ViewConversionEventComponent}></Route>
                          {/* <Route path = "/update-conversionEvent/:id" component = {UpdateConversionEventComponent}></Route> */}
                            <Route path = "/performanceMetrics" component = {ListPerformanceMetricComponent}></Route>
                            <Route path = "/add-performanceMetric/:id" component = {CreatePerformanceMetricComponent}></Route>
                            <Route path = "/view-performanceMetric/:id" component = {ViewPerformanceMetricComponent}></Route>
                          {/* <Route path = "/update-performanceMetric/:id" component = {UpdatePerformanceMetricComponent}></Route> */}
                            <Route path = "/reports" component = {ListReportComponent}></Route>
                            <Route path = "/add-report/:id" component = {CreateReportComponent}></Route>
                            <Route path = "/view-report/:id" component = {ViewReportComponent}></Route>
                          {/* <Route path = "/update-report/:id" component = {UpdateReportComponent}></Route> */}
                            <Route path = "/insertionOrders" component = {ListInsertionOrderComponent}></Route>
                            <Route path = "/add-insertionOrder/:id" component = {CreateInsertionOrderComponent}></Route>
                            <Route path = "/view-insertionOrder/:id" component = {ViewInsertionOrderComponent}></Route>
                          {/* <Route path = "/update-insertionOrder/:id" component = {UpdateInsertionOrderComponent}></Route> */}
                            <Route path = "/rateCards" component = {ListRateCardComponent}></Route>
                            <Route path = "/add-rateCard/:id" component = {CreateRateCardComponent}></Route>
                            <Route path = "/view-rateCard/:id" component = {ViewRateCardComponent}></Route>
                          {/* <Route path = "/update-rateCard/:id" component = {UpdateRateCardComponent}></Route> */}
                            <Route path = "/rates" component = {ListRateComponent}></Route>
                            <Route path = "/add-rate/:id" component = {CreateRateComponent}></Route>
                            <Route path = "/view-rate/:id" component = {ViewRateComponent}></Route>
                          {/* <Route path = "/update-rate/:id" component = {UpdateRateComponent}></Route> */}
                            <Route path = "/experiments" component = {ListExperimentComponent}></Route>
                            <Route path = "/add-experiment/:id" component = {CreateExperimentComponent}></Route>
                            <Route path = "/view-experiment/:id" component = {ViewExperimentComponent}></Route>
                          {/* <Route path = "/update-experiment/:id" component = {UpdateExperimentComponent}></Route> */}
                            <Route path = "/experimentVariants" component = {ListExperimentVariantComponent}></Route>
                            <Route path = "/add-experimentVariant/:id" component = {CreateExperimentVariantComponent}></Route>
                            <Route path = "/view-experimentVariant/:id" component = {ViewExperimentVariantComponent}></Route>
                          {/* <Route path = "/update-experimentVariant/:id" component = {UpdateExperimentVariantComponent}></Route> */}
                            <Route path = "/geoRegions" component = {ListGeoRegionComponent}></Route>
                            <Route path = "/add-geoRegion/:id" component = {CreateGeoRegionComponent}></Route>
                            <Route path = "/view-geoRegion/:id" component = {ViewGeoRegionComponent}></Route>
                          {/* <Route path = "/update-geoRegion/:id" component = {UpdateGeoRegionComponent}></Route> */}
                    </Switch>
                </div>
              <FooterComponent />
        </Router>
    </div>
    
  );
}

export default App;
