import {BrowserModule} from '@angular/platform-browser';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatInputModule} from '@angular/material/input';
import {MatDatepickerModule} from '@angular/material/datepicker';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatButtonModule} from '@angular/material/button';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatMomentDateModule} from "@angular/material-moment-adapter";
import {NgModule} from '@angular/core';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import {RouterModule} from '@angular/router';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule} from '@angular/forms';
import {ReactiveFormsModule} from '@angular/forms';
import {AppComponent} from './app.component';
import {MatMenuModule} from '@angular/material/menu';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatSidenavModule} from '@angular/material/sidenav'

import {IndexAgencyComponent} from './components/Agency/index/index.component';
import {CreateAgencyComponent} from './components/Agency/create/create.component';
import {EditAgencyComponent} from './components/Agency/edit/edit.component';
import {IndexTeamComponent} from './components/Team/index/index.component';
import {CreateTeamComponent} from './components/Team/create/create.component';
import {EditTeamComponent} from './components/Team/edit/edit.component';
import {IndexUserComponent} from './components/User/index/index.component';
import {CreateUserComponent} from './components/User/create/create.component';
import {EditUserComponent} from './components/User/edit/edit.component';
import {IndexAdvertiserComponent} from './components/Advertiser/index/index.component';
import {CreateAdvertiserComponent} from './components/Advertiser/create/create.component';
import {EditAdvertiserComponent} from './components/Advertiser/edit/edit.component';
import {IndexBillingProfileComponent} from './components/BillingProfile/index/index.component';
import {CreateBillingProfileComponent} from './components/BillingProfile/create/create.component';
import {EditBillingProfileComponent} from './components/BillingProfile/edit/edit.component';
import {IndexPaymentMethodComponent} from './components/PaymentMethod/index/index.component';
import {CreatePaymentMethodComponent} from './components/PaymentMethod/create/create.component';
import {EditPaymentMethodComponent} from './components/PaymentMethod/edit/edit.component';
import {IndexAdAccountComponent} from './components/AdAccount/index/index.component';
import {CreateAdAccountComponent} from './components/AdAccount/create/create.component';
import {EditAdAccountComponent} from './components/AdAccount/edit/edit.component';
import {IndexDSPComponent} from './components/DSP/index/index.component';
import {CreateDSPComponent} from './components/DSP/create/create.component';
import {EditDSPComponent} from './components/DSP/edit/edit.component';
import {IndexCampaignComponent} from './components/Campaign/index/index.component';
import {CreateCampaignComponent} from './components/Campaign/create/create.component';
import {EditCampaignComponent} from './components/Campaign/edit/edit.component';
import {IndexKPIComponent} from './components/KPI/index/index.component';
import {CreateKPIComponent} from './components/KPI/create/create.component';
import {EditKPIComponent} from './components/KPI/edit/edit.component';
import {IndexAudienceSegmentComponent} from './components/AudienceSegment/index/index.component';
import {CreateAudienceSegmentComponent} from './components/AudienceSegment/create/create.component';
import {EditAudienceSegmentComponent} from './components/AudienceSegment/edit/edit.component';
import {IndexDataProviderComponent} from './components/DataProvider/index/index.component';
import {CreateDataProviderComponent} from './components/DataProvider/create/create.component';
import {EditDataProviderComponent} from './components/DataProvider/edit/edit.component';
import {IndexLineItemComponent} from './components/LineItem/index/index.component';
import {CreateLineItemComponent} from './components/LineItem/create/create.component';
import {EditLineItemComponent} from './components/LineItem/edit/edit.component';
import {IndexTargetingProfileComponent} from './components/TargetingProfile/index/index.component';
import {CreateTargetingProfileComponent} from './components/TargetingProfile/create/create.component';
import {EditTargetingProfileComponent} from './components/TargetingProfile/edit/edit.component';
import {IndexDeviceCriterionComponent} from './components/DeviceCriterion/index/index.component';
import {CreateDeviceCriterionComponent} from './components/DeviceCriterion/create/create.component';
import {EditDeviceCriterionComponent} from './components/DeviceCriterion/edit/edit.component';
import {IndexBrandSafetyPolicyComponent} from './components/BrandSafetyPolicy/index/index.component';
import {CreateBrandSafetyPolicyComponent} from './components/BrandSafetyPolicy/create/create.component';
import {EditBrandSafetyPolicyComponent} from './components/BrandSafetyPolicy/edit/edit.component';
import {IndexContentCategoryComponent} from './components/ContentCategory/index/index.component';
import {CreateContentCategoryComponent} from './components/ContentCategory/create/create.component';
import {EditContentCategoryComponent} from './components/ContentCategory/edit/edit.component';
import {IndexPublisherComponent} from './components/Publisher/index/index.component';
import {CreatePublisherComponent} from './components/Publisher/create/create.component';
import {EditPublisherComponent} from './components/Publisher/edit/edit.component';
import {IndexInventorySourceComponent} from './components/InventorySource/index/index.component';
import {CreateInventorySourceComponent} from './components/InventorySource/create/create.component';
import {EditInventorySourceComponent} from './components/InventorySource/edit/edit.component';
import {IndexAdSlotComponent} from './components/AdSlot/index/index.component';
import {CreateAdSlotComponent} from './components/AdSlot/create/create.component';
import {EditAdSlotComponent} from './components/AdSlot/edit/edit.component';
import {IndexDealComponent} from './components/Deal/index/index.component';
import {CreateDealComponent} from './components/Deal/create/create.component';
import {EditDealComponent} from './components/Deal/edit/edit.component';
import {IndexPlacementComponent} from './components/Placement/index/index.component';
import {CreatePlacementComponent} from './components/Placement/create/create.component';
import {EditPlacementComponent} from './components/Placement/edit/edit.component';
import {IndexCreativeAssetComponent} from './components/CreativeAsset/index/index.component';
import {CreateCreativeAssetComponent} from './components/CreativeAsset/create/create.component';
import {EditCreativeAssetComponent} from './components/CreativeAsset/edit/edit.component';
import {IndexCreativeFileComponent} from './components/CreativeFile/index/index.component';
import {CreateCreativeFileComponent} from './components/CreativeFile/create/create.component';
import {EditCreativeFileComponent} from './components/CreativeFile/edit/edit.component';
import {IndexCreativeVariationComponent} from './components/CreativeVariation/index/index.component';
import {CreateCreativeVariationComponent} from './components/CreativeVariation/create/create.component';
import {EditCreativeVariationComponent} from './components/CreativeVariation/edit/edit.component';
import {IndexCreativeApprovalComponent} from './components/CreativeApproval/index/index.component';
import {CreateCreativeApprovalComponent} from './components/CreativeApproval/create/create.component';
import {EditCreativeApprovalComponent} from './components/CreativeApproval/edit/edit.component';
import {IndexTrackingPixelComponent} from './components/TrackingPixel/index/index.component';
import {CreateTrackingPixelComponent} from './components/TrackingPixel/create/create.component';
import {EditTrackingPixelComponent} from './components/TrackingPixel/edit/edit.component';
import {IndexConversionEventComponent} from './components/ConversionEvent/index/index.component';
import {CreateConversionEventComponent} from './components/ConversionEvent/create/create.component';
import {EditConversionEventComponent} from './components/ConversionEvent/edit/edit.component';
import {IndexPerformanceMetricComponent} from './components/PerformanceMetric/index/index.component';
import {CreatePerformanceMetricComponent} from './components/PerformanceMetric/create/create.component';
import {EditPerformanceMetricComponent} from './components/PerformanceMetric/edit/edit.component';
import {IndexReportComponent} from './components/Report/index/index.component';
import {CreateReportComponent} from './components/Report/create/create.component';
import {EditReportComponent} from './components/Report/edit/edit.component';
import {IndexInsertionOrderComponent} from './components/InsertionOrder/index/index.component';
import {CreateInsertionOrderComponent} from './components/InsertionOrder/create/create.component';
import {EditInsertionOrderComponent} from './components/InsertionOrder/edit/edit.component';
import {IndexRateCardComponent} from './components/RateCard/index/index.component';
import {CreateRateCardComponent} from './components/RateCard/create/create.component';
import {EditRateCardComponent} from './components/RateCard/edit/edit.component';
import {IndexRateComponent} from './components/Rate/index/index.component';
import {CreateRateComponent} from './components/Rate/create/create.component';
import {EditRateComponent} from './components/Rate/edit/edit.component';
import {IndexExperimentComponent} from './components/Experiment/index/index.component';
import {CreateExperimentComponent} from './components/Experiment/create/create.component';
import {EditExperimentComponent} from './components/Experiment/edit/edit.component';
import {IndexExperimentVariantComponent} from './components/ExperimentVariant/index/index.component';
import {CreateExperimentVariantComponent} from './components/ExperimentVariant/create/create.component';
import {EditExperimentVariantComponent} from './components/ExperimentVariant/edit/edit.component';
import {IndexGeoRegionComponent} from './components/GeoRegion/index/index.component';
import {CreateGeoRegionComponent} from './components/GeoRegion/create/create.component';
import {EditGeoRegionComponent} from './components/GeoRegion/edit/edit.component';

import * as appRoutes from './routerConfig';

import {AgencyService} from './services/Agency.service';
import {TeamService} from './services/Team.service';
import {UserService} from './services/User.service';
import {AdvertiserService} from './services/Advertiser.service';
import {BillingProfileService} from './services/BillingProfile.service';
import {PaymentMethodService} from './services/PaymentMethod.service';
import {AdAccountService} from './services/AdAccount.service';
import {DSPService} from './services/DSP.service';
import {CampaignService} from './services/Campaign.service';
import {KPIService} from './services/KPI.service';
import {AudienceSegmentService} from './services/AudienceSegment.service';
import {DataProviderService} from './services/DataProvider.service';
import {LineItemService} from './services/LineItem.service';
import {TargetingProfileService} from './services/TargetingProfile.service';
import {DeviceCriterionService} from './services/DeviceCriterion.service';
import {BrandSafetyPolicyService} from './services/BrandSafetyPolicy.service';
import {ContentCategoryService} from './services/ContentCategory.service';
import {PublisherService} from './services/Publisher.service';
import {InventorySourceService} from './services/InventorySource.service';
import {AdSlotService} from './services/AdSlot.service';
import {DealService} from './services/Deal.service';
import {PlacementService} from './services/Placement.service';
import {CreativeAssetService} from './services/CreativeAsset.service';
import {CreativeFileService} from './services/CreativeFile.service';
import {CreativeVariationService} from './services/CreativeVariation.service';
import {CreativeApprovalService} from './services/CreativeApproval.service';
import {TrackingPixelService} from './services/TrackingPixel.service';
import {ConversionEventService} from './services/ConversionEvent.service';
import {PerformanceMetricService} from './services/PerformanceMetric.service';
import {ReportService} from './services/Report.service';
import {InsertionOrderService} from './services/InsertionOrder.service';
import {RateCardService} from './services/RateCard.service';
import {RateService} from './services/Rate.service';
import {ExperimentService} from './services/Experiment.service';
import {ExperimentVariantService} from './services/ExperimentVariant.service';
import {GeoRegionService} from './services/GeoRegion.service';

@NgModule({
  declarations: [
    IndexAgencyComponent,
    CreateAgencyComponent,
    EditAgencyComponent,
    IndexTeamComponent,
    CreateTeamComponent,
    EditTeamComponent,
    IndexUserComponent,
    CreateUserComponent,
    EditUserComponent,
    IndexAdvertiserComponent,
    CreateAdvertiserComponent,
    EditAdvertiserComponent,
    IndexBillingProfileComponent,
    CreateBillingProfileComponent,
    EditBillingProfileComponent,
    IndexPaymentMethodComponent,
    CreatePaymentMethodComponent,
    EditPaymentMethodComponent,
    IndexAdAccountComponent,
    CreateAdAccountComponent,
    EditAdAccountComponent,
    IndexDSPComponent,
    CreateDSPComponent,
    EditDSPComponent,
    IndexCampaignComponent,
    CreateCampaignComponent,
    EditCampaignComponent,
    IndexKPIComponent,
    CreateKPIComponent,
    EditKPIComponent,
    IndexAudienceSegmentComponent,
    CreateAudienceSegmentComponent,
    EditAudienceSegmentComponent,
    IndexDataProviderComponent,
    CreateDataProviderComponent,
    EditDataProviderComponent,
    IndexLineItemComponent,
    CreateLineItemComponent,
    EditLineItemComponent,
    IndexTargetingProfileComponent,
    CreateTargetingProfileComponent,
    EditTargetingProfileComponent,
    IndexDeviceCriterionComponent,
    CreateDeviceCriterionComponent,
    EditDeviceCriterionComponent,
    IndexBrandSafetyPolicyComponent,
    CreateBrandSafetyPolicyComponent,
    EditBrandSafetyPolicyComponent,
    IndexContentCategoryComponent,
    CreateContentCategoryComponent,
    EditContentCategoryComponent,
    IndexPublisherComponent,
    CreatePublisherComponent,
    EditPublisherComponent,
    IndexInventorySourceComponent,
    CreateInventorySourceComponent,
    EditInventorySourceComponent,
    IndexAdSlotComponent,
    CreateAdSlotComponent,
    EditAdSlotComponent,
    IndexDealComponent,
    CreateDealComponent,
    EditDealComponent,
    IndexPlacementComponent,
    CreatePlacementComponent,
    EditPlacementComponent,
    IndexCreativeAssetComponent,
    CreateCreativeAssetComponent,
    EditCreativeAssetComponent,
    IndexCreativeFileComponent,
    CreateCreativeFileComponent,
    EditCreativeFileComponent,
    IndexCreativeVariationComponent,
    CreateCreativeVariationComponent,
    EditCreativeVariationComponent,
    IndexCreativeApprovalComponent,
    CreateCreativeApprovalComponent,
    EditCreativeApprovalComponent,
    IndexTrackingPixelComponent,
    CreateTrackingPixelComponent,
    EditTrackingPixelComponent,
    IndexConversionEventComponent,
    CreateConversionEventComponent,
    EditConversionEventComponent,
    IndexPerformanceMetricComponent,
    CreatePerformanceMetricComponent,
    EditPerformanceMetricComponent,
    IndexReportComponent,
    CreateReportComponent,
    EditReportComponent,
    IndexInsertionOrderComponent,
    CreateInsertionOrderComponent,
    EditInsertionOrderComponent,
    IndexRateCardComponent,
    CreateRateCardComponent,
    EditRateCardComponent,
    IndexRateComponent,
    CreateRateComponent,
    EditRateComponent,
    IndexExperimentComponent,
    CreateExperimentComponent,
    EditExperimentComponent,
    IndexExperimentVariantComponent,
    CreateExperimentVariantComponent,
    EditExperimentVariantComponent,
    IndexGeoRegionComponent,
    CreateGeoRegionComponent,
    EditGeoRegionComponent,
    AppComponent
  ],
  imports: [

    BrowserModule, 
    NgbModule,
    MatMenuModule,
    MatToolbarModule,
    MatCheckboxModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatDatepickerModule,
	MatMomentDateModule,
    BrowserAnimationsModule,
	HttpClientModule, 
    ReactiveFormsModule,
    FormsModule,
    MatSidenavModule,    
    RouterModule.forRoot(appRoutes.AgencyRoutes), 
    RouterModule.forRoot(appRoutes.TeamRoutes), 
    RouterModule.forRoot(appRoutes.UserRoutes), 
    RouterModule.forRoot(appRoutes.AdvertiserRoutes), 
    RouterModule.forRoot(appRoutes.BillingProfileRoutes), 
    RouterModule.forRoot(appRoutes.PaymentMethodRoutes), 
    RouterModule.forRoot(appRoutes.AdAccountRoutes), 
    RouterModule.forRoot(appRoutes.DSPRoutes), 
    RouterModule.forRoot(appRoutes.CampaignRoutes), 
    RouterModule.forRoot(appRoutes.KPIRoutes), 
    RouterModule.forRoot(appRoutes.AudienceSegmentRoutes), 
    RouterModule.forRoot(appRoutes.DataProviderRoutes), 
    RouterModule.forRoot(appRoutes.LineItemRoutes), 
    RouterModule.forRoot(appRoutes.TargetingProfileRoutes), 
    RouterModule.forRoot(appRoutes.DeviceCriterionRoutes), 
    RouterModule.forRoot(appRoutes.BrandSafetyPolicyRoutes), 
    RouterModule.forRoot(appRoutes.ContentCategoryRoutes), 
    RouterModule.forRoot(appRoutes.PublisherRoutes), 
    RouterModule.forRoot(appRoutes.InventorySourceRoutes), 
    RouterModule.forRoot(appRoutes.AdSlotRoutes), 
    RouterModule.forRoot(appRoutes.DealRoutes), 
    RouterModule.forRoot(appRoutes.PlacementRoutes), 
    RouterModule.forRoot(appRoutes.CreativeAssetRoutes), 
    RouterModule.forRoot(appRoutes.CreativeFileRoutes), 
    RouterModule.forRoot(appRoutes.CreativeVariationRoutes), 
    RouterModule.forRoot(appRoutes.CreativeApprovalRoutes), 
    RouterModule.forRoot(appRoutes.TrackingPixelRoutes), 
    RouterModule.forRoot(appRoutes.ConversionEventRoutes), 
    RouterModule.forRoot(appRoutes.PerformanceMetricRoutes), 
    RouterModule.forRoot(appRoutes.ReportRoutes), 
    RouterModule.forRoot(appRoutes.InsertionOrderRoutes), 
    RouterModule.forRoot(appRoutes.RateCardRoutes), 
    RouterModule.forRoot(appRoutes.RateRoutes), 
    RouterModule.forRoot(appRoutes.ExperimentRoutes), 
    RouterModule.forRoot(appRoutes.ExperimentVariantRoutes), 
    RouterModule.forRoot(appRoutes.GeoRegionRoutes), 
  ],
  providers: [AgencyService,TeamService,UserService,AdvertiserService,BillingProfileService,PaymentMethodService,AdAccountService,DSPService,CampaignService,KPIService,AudienceSegmentService,DataProviderService,LineItemService,TargetingProfileService,DeviceCriterionService,BrandSafetyPolicyService,ContentCategoryService,PublisherService,InventorySourceService,AdSlotService,DealService,PlacementService,CreativeAssetService,CreativeFileService,CreativeVariationService,CreativeApprovalService,TrackingPixelService,ConversionEventService,PerformanceMetricService,ReportService,InsertionOrderService,RateCardService,RateService,ExperimentService,ExperimentVariantService,GeoRegionService],
  bootstrap: [AppComponent]
})
export class AppModule { }
