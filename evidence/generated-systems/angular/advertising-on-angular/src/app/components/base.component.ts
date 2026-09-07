import { HttpClient } from '@angular/common/http';
import * as enumTypes from '../models/EnumTypes';

import {AgencyService} from '../services/Agency.service';
import {TeamService} from '../services/Team.service';
import {UserService} from '../services/User.service';
import {AdvertiserService} from '../services/Advertiser.service';
import {BillingProfileService} from '../services/BillingProfile.service';
import {PaymentMethodService} from '../services/PaymentMethod.service';
import {AdAccountService} from '../services/AdAccount.service';
import {DSPService} from '../services/DSP.service';
import {CampaignService} from '../services/Campaign.service';
import {KPIService} from '../services/KPI.service';
import {AudienceSegmentService} from '../services/AudienceSegment.service';
import {DataProviderService} from '../services/DataProvider.service';
import {LineItemService} from '../services/LineItem.service';
import {TargetingProfileService} from '../services/TargetingProfile.service';
import {DeviceCriterionService} from '../services/DeviceCriterion.service';
import {BrandSafetyPolicyService} from '../services/BrandSafetyPolicy.service';
import {ContentCategoryService} from '../services/ContentCategory.service';
import {PublisherService} from '../services/Publisher.service';
import {InventorySourceService} from '../services/InventorySource.service';
import {AdSlotService} from '../services/AdSlot.service';
import {DealService} from '../services/Deal.service';
import {PlacementService} from '../services/Placement.service';
import {CreativeAssetService} from '../services/CreativeAsset.service';
import {CreativeFileService} from '../services/CreativeFile.service';
import {CreativeVariationService} from '../services/CreativeVariation.service';
import {CreativeApprovalService} from '../services/CreativeApproval.service';
import {TrackingPixelService} from '../services/TrackingPixel.service';
import {ConversionEventService} from '../services/ConversionEvent.service';
import {PerformanceMetricService} from '../services/PerformanceMetric.service';
import {ReportService} from '../services/Report.service';
import {InsertionOrderService} from '../services/InsertionOrder.service';
import {RateCardService} from '../services/RateCard.service';
import {RateService} from '../services/Rate.service';
import {ExperimentService} from '../services/Experiment.service';
import {ExperimentVariantService} from '../services/ExperimentVariant.service';
import {GeoRegionService} from '../services/GeoRegion.service';

import { Directive } from '@angular/core';

/**
 Base class of all Components.
 For convenience, contains all enums and entity lists
 **/

@Directive()
export class BaseComponent {

    constructor (private http: HttpClient) {}

// enum instances
    CampaignStatuss = Object.keys(enumTypes.CampaignStatus);
    LineItemStatuss = Object.keys(enumTypes.LineItemStatus);
    ObjectiveTypes = Object.keys(enumTypes.ObjectiveType);
    PricingModels = Object.keys(enumTypes.PricingModel);
    BidStrategyTypes = Object.keys(enumTypes.BidStrategyType);
    PacingTypes = Object.keys(enumTypes.PacingType);
    FrequencyScopes = Object.keys(enumTypes.FrequencyScope);
    FrequencyPeriods = Object.keys(enumTypes.FrequencyPeriod);
    ChannelTypes = Object.keys(enumTypes.ChannelType);
    AdFormats = Object.keys(enumTypes.AdFormat);
    CreativeTypes = Object.keys(enumTypes.CreativeType);
    DeviceTypes = Object.keys(enumTypes.DeviceType);
    PlatformTypes = Object.keys(enumTypes.PlatformType);
    TargetingOperators = Object.keys(enumTypes.TargetingOperator);
    DealTypes = Object.keys(enumTypes.DealType);
    PublisherTypes = Object.keys(enumTypes.PublisherType);
    DataProviderTypes = Object.keys(enumTypes.DataProviderType);
    BrandSafetyLevels = Object.keys(enumTypes.BrandSafetyLevel);
    ContentRatings = Object.keys(enumTypes.ContentRating);
    MetricTypes = Object.keys(enumTypes.MetricType);
    AttributionModels = Object.keys(enumTypes.AttributionModel);
    IOStatuss = Object.keys(enumTypes.IOStatus);
    ReportTypes = Object.keys(enumTypes.ReportType);
    PixelTypes = Object.keys(enumTypes.PixelType);
    ConversionEventTypes = Object.keys(enumTypes.ConversionEventType);
    GeoRegionTypes = Object.keys(enumTypes.GeoRegionType);
    AccountRoles = Object.keys(enumTypes.AccountRole);
    PaymentTermss = Object.keys(enumTypes.PaymentTerms);
    PaymentMethodTypes = Object.keys(enumTypes.PaymentMethodType);
    InventoryTypes = Object.keys(enumTypes.InventoryType);
    CreativeApprovalStatuss = Object.keys(enumTypes.CreativeApprovalStatus);
    ExperimentStatuss = Object.keys(enumTypes.ExperimentStatus);

// all collection instances
    agencys : any;
    teams : any;
    users : any;
    advertisers : any;
    billingProfiles : any;
    paymentMethods : any;
    adAccounts : any;
    dSPs : any;
    campaigns : any;
    kPIs : any;
    audienceSegments : any;
    dataProviders : any;
    lineItems : any;
    targetingProfiles : any;
    deviceCriterions : any;
    brandSafetyPolicys : any;
    contentCategorys : any;
    publishers : any;
    inventorySources : any;
    adSlots : any;
    deals : any;
    placements : any;
    creativeAssets : any;
    creativeFiles : any;
    creativeVariations : any;
    creativeApprovals : any;
    trackingPixels : any;
    conversionEvents : any;
    performanceMetrics : any;
    reports : any;
    insertionOrders : any;
    rateCards : any;
    rates : any;
    experiments : any;
    experimentVariants : any;
    geoRegions : any;
  
// initialization  
    ngOnInit() {
    }

    initAgencyList() {
        if ( this.agencys == null ) {
            new AgencyService(this.http).getAgencys().subscribe(res => {
                this.agencys = res;
            });
        }
    }
    
    initTeamList() {
        if ( this.teams == null ) {
            new TeamService(this.http).getTeams().subscribe(res => {
                this.teams = res;
            });
        }
    }
    
    initUserList() {
        if ( this.users == null ) {
            new UserService(this.http).getUsers().subscribe(res => {
                this.users = res;
            });
        }
    }
    
    initAdvertiserList() {
        if ( this.advertisers == null ) {
            new AdvertiserService(this.http).getAdvertisers().subscribe(res => {
                this.advertisers = res;
            });
        }
    }
    
    initBillingProfileList() {
        if ( this.billingProfiles == null ) {
            new BillingProfileService(this.http).getBillingProfiles().subscribe(res => {
                this.billingProfiles = res;
            });
        }
    }
    
    initPaymentMethodList() {
        if ( this.paymentMethods == null ) {
            new PaymentMethodService(this.http).getPaymentMethods().subscribe(res => {
                this.paymentMethods = res;
            });
        }
    }
    
    initAdAccountList() {
        if ( this.adAccounts == null ) {
            new AdAccountService(this.http).getAdAccounts().subscribe(res => {
                this.adAccounts = res;
            });
        }
    }
    
    initDSPList() {
        if ( this.dSPs == null ) {
            new DSPService(this.http).getDSPs().subscribe(res => {
                this.dSPs = res;
            });
        }
    }
    
    initCampaignList() {
        if ( this.campaigns == null ) {
            new CampaignService(this.http).getCampaigns().subscribe(res => {
                this.campaigns = res;
            });
        }
    }
    
    initKPIList() {
        if ( this.kPIs == null ) {
            new KPIService(this.http).getKPIs().subscribe(res => {
                this.kPIs = res;
            });
        }
    }
    
    initAudienceSegmentList() {
        if ( this.audienceSegments == null ) {
            new AudienceSegmentService(this.http).getAudienceSegments().subscribe(res => {
                this.audienceSegments = res;
            });
        }
    }
    
    initDataProviderList() {
        if ( this.dataProviders == null ) {
            new DataProviderService(this.http).getDataProviders().subscribe(res => {
                this.dataProviders = res;
            });
        }
    }
    
    initLineItemList() {
        if ( this.lineItems == null ) {
            new LineItemService(this.http).getLineItems().subscribe(res => {
                this.lineItems = res;
            });
        }
    }
    
    initTargetingProfileList() {
        if ( this.targetingProfiles == null ) {
            new TargetingProfileService(this.http).getTargetingProfiles().subscribe(res => {
                this.targetingProfiles = res;
            });
        }
    }
    
    initDeviceCriterionList() {
        if ( this.deviceCriterions == null ) {
            new DeviceCriterionService(this.http).getDeviceCriterions().subscribe(res => {
                this.deviceCriterions = res;
            });
        }
    }
    
    initBrandSafetyPolicyList() {
        if ( this.brandSafetyPolicys == null ) {
            new BrandSafetyPolicyService(this.http).getBrandSafetyPolicys().subscribe(res => {
                this.brandSafetyPolicys = res;
            });
        }
    }
    
    initContentCategoryList() {
        if ( this.contentCategorys == null ) {
            new ContentCategoryService(this.http).getContentCategorys().subscribe(res => {
                this.contentCategorys = res;
            });
        }
    }
    
    initPublisherList() {
        if ( this.publishers == null ) {
            new PublisherService(this.http).getPublishers().subscribe(res => {
                this.publishers = res;
            });
        }
    }
    
    initInventorySourceList() {
        if ( this.inventorySources == null ) {
            new InventorySourceService(this.http).getInventorySources().subscribe(res => {
                this.inventorySources = res;
            });
        }
    }
    
    initAdSlotList() {
        if ( this.adSlots == null ) {
            new AdSlotService(this.http).getAdSlots().subscribe(res => {
                this.adSlots = res;
            });
        }
    }
    
    initDealList() {
        if ( this.deals == null ) {
            new DealService(this.http).getDeals().subscribe(res => {
                this.deals = res;
            });
        }
    }
    
    initPlacementList() {
        if ( this.placements == null ) {
            new PlacementService(this.http).getPlacements().subscribe(res => {
                this.placements = res;
            });
        }
    }
    
    initCreativeAssetList() {
        if ( this.creativeAssets == null ) {
            new CreativeAssetService(this.http).getCreativeAssets().subscribe(res => {
                this.creativeAssets = res;
            });
        }
    }
    
    initCreativeFileList() {
        if ( this.creativeFiles == null ) {
            new CreativeFileService(this.http).getCreativeFiles().subscribe(res => {
                this.creativeFiles = res;
            });
        }
    }
    
    initCreativeVariationList() {
        if ( this.creativeVariations == null ) {
            new CreativeVariationService(this.http).getCreativeVariations().subscribe(res => {
                this.creativeVariations = res;
            });
        }
    }
    
    initCreativeApprovalList() {
        if ( this.creativeApprovals == null ) {
            new CreativeApprovalService(this.http).getCreativeApprovals().subscribe(res => {
                this.creativeApprovals = res;
            });
        }
    }
    
    initTrackingPixelList() {
        if ( this.trackingPixels == null ) {
            new TrackingPixelService(this.http).getTrackingPixels().subscribe(res => {
                this.trackingPixels = res;
            });
        }
    }
    
    initConversionEventList() {
        if ( this.conversionEvents == null ) {
            new ConversionEventService(this.http).getConversionEvents().subscribe(res => {
                this.conversionEvents = res;
            });
        }
    }
    
    initPerformanceMetricList() {
        if ( this.performanceMetrics == null ) {
            new PerformanceMetricService(this.http).getPerformanceMetrics().subscribe(res => {
                this.performanceMetrics = res;
            });
        }
    }
    
    initReportList() {
        if ( this.reports == null ) {
            new ReportService(this.http).getReports().subscribe(res => {
                this.reports = res;
            });
        }
    }
    
    initInsertionOrderList() {
        if ( this.insertionOrders == null ) {
            new InsertionOrderService(this.http).getInsertionOrders().subscribe(res => {
                this.insertionOrders = res;
            });
        }
    }
    
    initRateCardList() {
        if ( this.rateCards == null ) {
            new RateCardService(this.http).getRateCards().subscribe(res => {
                this.rateCards = res;
            });
        }
    }
    
    initRateList() {
        if ( this.rates == null ) {
            new RateService(this.http).getRates().subscribe(res => {
                this.rates = res;
            });
        }
    }
    
    initExperimentList() {
        if ( this.experiments == null ) {
            new ExperimentService(this.http).getExperiments().subscribe(res => {
                this.experiments = res;
            });
        }
    }
    
    initExperimentVariantList() {
        if ( this.experimentVariants == null ) {
            new ExperimentVariantService(this.http).getExperimentVariants().subscribe(res => {
                this.experimentVariants = res;
            });
        }
    }
    
    initGeoRegionList() {
        if ( this.geoRegions == null ) {
            new GeoRegionService(this.http).getGeoRegions().subscribe(res => {
                this.geoRegions = res;
            });
        }
    }
    
    
// comparison function for select controls  
    compareFn(user1: any, user2: any) {
        return user1 == user2
    }    
}
