// routerConfig.ts

import { Routes } from '@angular/router';
import { CreateAgencyComponent } from './components/Agency/create/create.component';
import { EditAgencyComponent } from './components/Agency/edit/edit.component';
import { IndexAgencyComponent } from './components/Agency/index/index.component';
import { CreateTeamComponent } from './components/Team/create/create.component';
import { EditTeamComponent } from './components/Team/edit/edit.component';
import { IndexTeamComponent } from './components/Team/index/index.component';
import { CreateUserComponent } from './components/User/create/create.component';
import { EditUserComponent } from './components/User/edit/edit.component';
import { IndexUserComponent } from './components/User/index/index.component';
import { CreateAdvertiserComponent } from './components/Advertiser/create/create.component';
import { EditAdvertiserComponent } from './components/Advertiser/edit/edit.component';
import { IndexAdvertiserComponent } from './components/Advertiser/index/index.component';
import { CreateBillingProfileComponent } from './components/BillingProfile/create/create.component';
import { EditBillingProfileComponent } from './components/BillingProfile/edit/edit.component';
import { IndexBillingProfileComponent } from './components/BillingProfile/index/index.component';
import { CreatePaymentMethodComponent } from './components/PaymentMethod/create/create.component';
import { EditPaymentMethodComponent } from './components/PaymentMethod/edit/edit.component';
import { IndexPaymentMethodComponent } from './components/PaymentMethod/index/index.component';
import { CreateAdAccountComponent } from './components/AdAccount/create/create.component';
import { EditAdAccountComponent } from './components/AdAccount/edit/edit.component';
import { IndexAdAccountComponent } from './components/AdAccount/index/index.component';
import { CreateDSPComponent } from './components/DSP/create/create.component';
import { EditDSPComponent } from './components/DSP/edit/edit.component';
import { IndexDSPComponent } from './components/DSP/index/index.component';
import { CreateCampaignComponent } from './components/Campaign/create/create.component';
import { EditCampaignComponent } from './components/Campaign/edit/edit.component';
import { IndexCampaignComponent } from './components/Campaign/index/index.component';
import { CreateKPIComponent } from './components/KPI/create/create.component';
import { EditKPIComponent } from './components/KPI/edit/edit.component';
import { IndexKPIComponent } from './components/KPI/index/index.component';
import { CreateAudienceSegmentComponent } from './components/AudienceSegment/create/create.component';
import { EditAudienceSegmentComponent } from './components/AudienceSegment/edit/edit.component';
import { IndexAudienceSegmentComponent } from './components/AudienceSegment/index/index.component';
import { CreateDataProviderComponent } from './components/DataProvider/create/create.component';
import { EditDataProviderComponent } from './components/DataProvider/edit/edit.component';
import { IndexDataProviderComponent } from './components/DataProvider/index/index.component';
import { CreateLineItemComponent } from './components/LineItem/create/create.component';
import { EditLineItemComponent } from './components/LineItem/edit/edit.component';
import { IndexLineItemComponent } from './components/LineItem/index/index.component';
import { CreateTargetingProfileComponent } from './components/TargetingProfile/create/create.component';
import { EditTargetingProfileComponent } from './components/TargetingProfile/edit/edit.component';
import { IndexTargetingProfileComponent } from './components/TargetingProfile/index/index.component';
import { CreateDeviceCriterionComponent } from './components/DeviceCriterion/create/create.component';
import { EditDeviceCriterionComponent } from './components/DeviceCriterion/edit/edit.component';
import { IndexDeviceCriterionComponent } from './components/DeviceCriterion/index/index.component';
import { CreateBrandSafetyPolicyComponent } from './components/BrandSafetyPolicy/create/create.component';
import { EditBrandSafetyPolicyComponent } from './components/BrandSafetyPolicy/edit/edit.component';
import { IndexBrandSafetyPolicyComponent } from './components/BrandSafetyPolicy/index/index.component';
import { CreateContentCategoryComponent } from './components/ContentCategory/create/create.component';
import { EditContentCategoryComponent } from './components/ContentCategory/edit/edit.component';
import { IndexContentCategoryComponent } from './components/ContentCategory/index/index.component';
import { CreatePublisherComponent } from './components/Publisher/create/create.component';
import { EditPublisherComponent } from './components/Publisher/edit/edit.component';
import { IndexPublisherComponent } from './components/Publisher/index/index.component';
import { CreateInventorySourceComponent } from './components/InventorySource/create/create.component';
import { EditInventorySourceComponent } from './components/InventorySource/edit/edit.component';
import { IndexInventorySourceComponent } from './components/InventorySource/index/index.component';
import { CreateAdSlotComponent } from './components/AdSlot/create/create.component';
import { EditAdSlotComponent } from './components/AdSlot/edit/edit.component';
import { IndexAdSlotComponent } from './components/AdSlot/index/index.component';
import { CreateDealComponent } from './components/Deal/create/create.component';
import { EditDealComponent } from './components/Deal/edit/edit.component';
import { IndexDealComponent } from './components/Deal/index/index.component';
import { CreatePlacementComponent } from './components/Placement/create/create.component';
import { EditPlacementComponent } from './components/Placement/edit/edit.component';
import { IndexPlacementComponent } from './components/Placement/index/index.component';
import { CreateCreativeAssetComponent } from './components/CreativeAsset/create/create.component';
import { EditCreativeAssetComponent } from './components/CreativeAsset/edit/edit.component';
import { IndexCreativeAssetComponent } from './components/CreativeAsset/index/index.component';
import { CreateCreativeFileComponent } from './components/CreativeFile/create/create.component';
import { EditCreativeFileComponent } from './components/CreativeFile/edit/edit.component';
import { IndexCreativeFileComponent } from './components/CreativeFile/index/index.component';
import { CreateCreativeVariationComponent } from './components/CreativeVariation/create/create.component';
import { EditCreativeVariationComponent } from './components/CreativeVariation/edit/edit.component';
import { IndexCreativeVariationComponent } from './components/CreativeVariation/index/index.component';
import { CreateCreativeApprovalComponent } from './components/CreativeApproval/create/create.component';
import { EditCreativeApprovalComponent } from './components/CreativeApproval/edit/edit.component';
import { IndexCreativeApprovalComponent } from './components/CreativeApproval/index/index.component';
import { CreateTrackingPixelComponent } from './components/TrackingPixel/create/create.component';
import { EditTrackingPixelComponent } from './components/TrackingPixel/edit/edit.component';
import { IndexTrackingPixelComponent } from './components/TrackingPixel/index/index.component';
import { CreateConversionEventComponent } from './components/ConversionEvent/create/create.component';
import { EditConversionEventComponent } from './components/ConversionEvent/edit/edit.component';
import { IndexConversionEventComponent } from './components/ConversionEvent/index/index.component';
import { CreatePerformanceMetricComponent } from './components/PerformanceMetric/create/create.component';
import { EditPerformanceMetricComponent } from './components/PerformanceMetric/edit/edit.component';
import { IndexPerformanceMetricComponent } from './components/PerformanceMetric/index/index.component';
import { CreateReportComponent } from './components/Report/create/create.component';
import { EditReportComponent } from './components/Report/edit/edit.component';
import { IndexReportComponent } from './components/Report/index/index.component';
import { CreateInsertionOrderComponent } from './components/InsertionOrder/create/create.component';
import { EditInsertionOrderComponent } from './components/InsertionOrder/edit/edit.component';
import { IndexInsertionOrderComponent } from './components/InsertionOrder/index/index.component';
import { CreateRateCardComponent } from './components/RateCard/create/create.component';
import { EditRateCardComponent } from './components/RateCard/edit/edit.component';
import { IndexRateCardComponent } from './components/RateCard/index/index.component';
import { CreateRateComponent } from './components/Rate/create/create.component';
import { EditRateComponent } from './components/Rate/edit/edit.component';
import { IndexRateComponent } from './components/Rate/index/index.component';
import { CreateExperimentComponent } from './components/Experiment/create/create.component';
import { EditExperimentComponent } from './components/Experiment/edit/edit.component';
import { IndexExperimentComponent } from './components/Experiment/index/index.component';
import { CreateExperimentVariantComponent } from './components/ExperimentVariant/create/create.component';
import { EditExperimentVariantComponent } from './components/ExperimentVariant/edit/edit.component';
import { IndexExperimentVariantComponent } from './components/ExperimentVariant/index/index.component';
import { CreateGeoRegionComponent } from './components/GeoRegion/create/create.component';
import { EditGeoRegionComponent } from './components/GeoRegion/edit/edit.component';
import { IndexGeoRegionComponent } from './components/GeoRegion/index/index.component';

export const AgencyRoutes: Routes = [
  { path: 'createAgency',
    component: CreateAgencyComponent
  },
  {
    path: 'editAgency/:id',
    component: EditAgencyComponent
  },
  { path: 'indexAgency',
    component: IndexAgencyComponent
  }
];
export const TeamRoutes: Routes = [
  { path: 'createTeam',
    component: CreateTeamComponent
  },
  {
    path: 'editTeam/:id',
    component: EditTeamComponent
  },
  { path: 'indexTeam',
    component: IndexTeamComponent
  }
];
export const UserRoutes: Routes = [
  { path: 'createUser',
    component: CreateUserComponent
  },
  {
    path: 'editUser/:id',
    component: EditUserComponent
  },
  { path: 'indexUser',
    component: IndexUserComponent
  }
];
export const AdvertiserRoutes: Routes = [
  { path: 'createAdvertiser',
    component: CreateAdvertiserComponent
  },
  {
    path: 'editAdvertiser/:id',
    component: EditAdvertiserComponent
  },
  { path: 'indexAdvertiser',
    component: IndexAdvertiserComponent
  }
];
export const BillingProfileRoutes: Routes = [
  { path: 'createBillingProfile',
    component: CreateBillingProfileComponent
  },
  {
    path: 'editBillingProfile/:id',
    component: EditBillingProfileComponent
  },
  { path: 'indexBillingProfile',
    component: IndexBillingProfileComponent
  }
];
export const PaymentMethodRoutes: Routes = [
  { path: 'createPaymentMethod',
    component: CreatePaymentMethodComponent
  },
  {
    path: 'editPaymentMethod/:id',
    component: EditPaymentMethodComponent
  },
  { path: 'indexPaymentMethod',
    component: IndexPaymentMethodComponent
  }
];
export const AdAccountRoutes: Routes = [
  { path: 'createAdAccount',
    component: CreateAdAccountComponent
  },
  {
    path: 'editAdAccount/:id',
    component: EditAdAccountComponent
  },
  { path: 'indexAdAccount',
    component: IndexAdAccountComponent
  }
];
export const DSPRoutes: Routes = [
  { path: 'createDSP',
    component: CreateDSPComponent
  },
  {
    path: 'editDSP/:id',
    component: EditDSPComponent
  },
  { path: 'indexDSP',
    component: IndexDSPComponent
  }
];
export const CampaignRoutes: Routes = [
  { path: 'createCampaign',
    component: CreateCampaignComponent
  },
  {
    path: 'editCampaign/:id',
    component: EditCampaignComponent
  },
  { path: 'indexCampaign',
    component: IndexCampaignComponent
  }
];
export const KPIRoutes: Routes = [
  { path: 'createKPI',
    component: CreateKPIComponent
  },
  {
    path: 'editKPI/:id',
    component: EditKPIComponent
  },
  { path: 'indexKPI',
    component: IndexKPIComponent
  }
];
export const AudienceSegmentRoutes: Routes = [
  { path: 'createAudienceSegment',
    component: CreateAudienceSegmentComponent
  },
  {
    path: 'editAudienceSegment/:id',
    component: EditAudienceSegmentComponent
  },
  { path: 'indexAudienceSegment',
    component: IndexAudienceSegmentComponent
  }
];
export const DataProviderRoutes: Routes = [
  { path: 'createDataProvider',
    component: CreateDataProviderComponent
  },
  {
    path: 'editDataProvider/:id',
    component: EditDataProviderComponent
  },
  { path: 'indexDataProvider',
    component: IndexDataProviderComponent
  }
];
export const LineItemRoutes: Routes = [
  { path: 'createLineItem',
    component: CreateLineItemComponent
  },
  {
    path: 'editLineItem/:id',
    component: EditLineItemComponent
  },
  { path: 'indexLineItem',
    component: IndexLineItemComponent
  }
];
export const TargetingProfileRoutes: Routes = [
  { path: 'createTargetingProfile',
    component: CreateTargetingProfileComponent
  },
  {
    path: 'editTargetingProfile/:id',
    component: EditTargetingProfileComponent
  },
  { path: 'indexTargetingProfile',
    component: IndexTargetingProfileComponent
  }
];
export const DeviceCriterionRoutes: Routes = [
  { path: 'createDeviceCriterion',
    component: CreateDeviceCriterionComponent
  },
  {
    path: 'editDeviceCriterion/:id',
    component: EditDeviceCriterionComponent
  },
  { path: 'indexDeviceCriterion',
    component: IndexDeviceCriterionComponent
  }
];
export const BrandSafetyPolicyRoutes: Routes = [
  { path: 'createBrandSafetyPolicy',
    component: CreateBrandSafetyPolicyComponent
  },
  {
    path: 'editBrandSafetyPolicy/:id',
    component: EditBrandSafetyPolicyComponent
  },
  { path: 'indexBrandSafetyPolicy',
    component: IndexBrandSafetyPolicyComponent
  }
];
export const ContentCategoryRoutes: Routes = [
  { path: 'createContentCategory',
    component: CreateContentCategoryComponent
  },
  {
    path: 'editContentCategory/:id',
    component: EditContentCategoryComponent
  },
  { path: 'indexContentCategory',
    component: IndexContentCategoryComponent
  }
];
export const PublisherRoutes: Routes = [
  { path: 'createPublisher',
    component: CreatePublisherComponent
  },
  {
    path: 'editPublisher/:id',
    component: EditPublisherComponent
  },
  { path: 'indexPublisher',
    component: IndexPublisherComponent
  }
];
export const InventorySourceRoutes: Routes = [
  { path: 'createInventorySource',
    component: CreateInventorySourceComponent
  },
  {
    path: 'editInventorySource/:id',
    component: EditInventorySourceComponent
  },
  { path: 'indexInventorySource',
    component: IndexInventorySourceComponent
  }
];
export const AdSlotRoutes: Routes = [
  { path: 'createAdSlot',
    component: CreateAdSlotComponent
  },
  {
    path: 'editAdSlot/:id',
    component: EditAdSlotComponent
  },
  { path: 'indexAdSlot',
    component: IndexAdSlotComponent
  }
];
export const DealRoutes: Routes = [
  { path: 'createDeal',
    component: CreateDealComponent
  },
  {
    path: 'editDeal/:id',
    component: EditDealComponent
  },
  { path: 'indexDeal',
    component: IndexDealComponent
  }
];
export const PlacementRoutes: Routes = [
  { path: 'createPlacement',
    component: CreatePlacementComponent
  },
  {
    path: 'editPlacement/:id',
    component: EditPlacementComponent
  },
  { path: 'indexPlacement',
    component: IndexPlacementComponent
  }
];
export const CreativeAssetRoutes: Routes = [
  { path: 'createCreativeAsset',
    component: CreateCreativeAssetComponent
  },
  {
    path: 'editCreativeAsset/:id',
    component: EditCreativeAssetComponent
  },
  { path: 'indexCreativeAsset',
    component: IndexCreativeAssetComponent
  }
];
export const CreativeFileRoutes: Routes = [
  { path: 'createCreativeFile',
    component: CreateCreativeFileComponent
  },
  {
    path: 'editCreativeFile/:id',
    component: EditCreativeFileComponent
  },
  { path: 'indexCreativeFile',
    component: IndexCreativeFileComponent
  }
];
export const CreativeVariationRoutes: Routes = [
  { path: 'createCreativeVariation',
    component: CreateCreativeVariationComponent
  },
  {
    path: 'editCreativeVariation/:id',
    component: EditCreativeVariationComponent
  },
  { path: 'indexCreativeVariation',
    component: IndexCreativeVariationComponent
  }
];
export const CreativeApprovalRoutes: Routes = [
  { path: 'createCreativeApproval',
    component: CreateCreativeApprovalComponent
  },
  {
    path: 'editCreativeApproval/:id',
    component: EditCreativeApprovalComponent
  },
  { path: 'indexCreativeApproval',
    component: IndexCreativeApprovalComponent
  }
];
export const TrackingPixelRoutes: Routes = [
  { path: 'createTrackingPixel',
    component: CreateTrackingPixelComponent
  },
  {
    path: 'editTrackingPixel/:id',
    component: EditTrackingPixelComponent
  },
  { path: 'indexTrackingPixel',
    component: IndexTrackingPixelComponent
  }
];
export const ConversionEventRoutes: Routes = [
  { path: 'createConversionEvent',
    component: CreateConversionEventComponent
  },
  {
    path: 'editConversionEvent/:id',
    component: EditConversionEventComponent
  },
  { path: 'indexConversionEvent',
    component: IndexConversionEventComponent
  }
];
export const PerformanceMetricRoutes: Routes = [
  { path: 'createPerformanceMetric',
    component: CreatePerformanceMetricComponent
  },
  {
    path: 'editPerformanceMetric/:id',
    component: EditPerformanceMetricComponent
  },
  { path: 'indexPerformanceMetric',
    component: IndexPerformanceMetricComponent
  }
];
export const ReportRoutes: Routes = [
  { path: 'createReport',
    component: CreateReportComponent
  },
  {
    path: 'editReport/:id',
    component: EditReportComponent
  },
  { path: 'indexReport',
    component: IndexReportComponent
  }
];
export const InsertionOrderRoutes: Routes = [
  { path: 'createInsertionOrder',
    component: CreateInsertionOrderComponent
  },
  {
    path: 'editInsertionOrder/:id',
    component: EditInsertionOrderComponent
  },
  { path: 'indexInsertionOrder',
    component: IndexInsertionOrderComponent
  }
];
export const RateCardRoutes: Routes = [
  { path: 'createRateCard',
    component: CreateRateCardComponent
  },
  {
    path: 'editRateCard/:id',
    component: EditRateCardComponent
  },
  { path: 'indexRateCard',
    component: IndexRateCardComponent
  }
];
export const RateRoutes: Routes = [
  { path: 'createRate',
    component: CreateRateComponent
  },
  {
    path: 'editRate/:id',
    component: EditRateComponent
  },
  { path: 'indexRate',
    component: IndexRateComponent
  }
];
export const ExperimentRoutes: Routes = [
  { path: 'createExperiment',
    component: CreateExperimentComponent
  },
  {
    path: 'editExperiment/:id',
    component: EditExperimentComponent
  },
  { path: 'indexExperiment',
    component: IndexExperimentComponent
  }
];
export const ExperimentVariantRoutes: Routes = [
  { path: 'createExperimentVariant',
    component: CreateExperimentVariantComponent
  },
  {
    path: 'editExperimentVariant/:id',
    component: EditExperimentVariantComponent
  },
  { path: 'indexExperimentVariant',
    component: IndexExperimentVariantComponent
  }
];
export const GeoRegionRoutes: Routes = [
  { path: 'createGeoRegion',
    component: CreateGeoRegionComponent
  },
  {
    path: 'editGeoRegion/:id',
    component: EditGeoRegionComponent
  },
  { path: 'indexGeoRegion',
    component: IndexGeoRegionComponent
  }
];
