import { HttpClient } from '@angular/common/http';
import * as enumTypes from '../models/EnumTypes';

import {OrganizationService} from '../services/Organization.service';
import {UserService} from '../services/User.service';
import {TeamService} from '../services/Team.service';
import {TerritoryService} from '../services/Territory.service';
import {AccountService} from '../services/Account.service';
import {ContactService} from '../services/Contact.service';
import {LeadService} from '../services/Lead.service';
import {OpportunityService} from '../services/Opportunity.service';
import {OpportunityLineItemService} from '../services/OpportunityLineItem.service';
import {OpportunityStageHistoryService} from '../services/OpportunityStageHistory.service';
import {ProductService} from '../services/Product.service';
import {PriceBookService} from '../services/PriceBook.service';
import {PriceBookEntryService} from '../services/PriceBookEntry.service';
import {QuoteService} from '../services/Quote.service';
import {QuoteLineItemService} from '../services/QuoteLineItem.service';
import {OrderService} from '../services/Order.service';
import {OrderItemService} from '../services/OrderItem.service';
import {ContractService} from '../services/Contract.service';
import {Case_Service} from '../services/Case_.service';
import {ActivityService} from '../services/Activity.service';
import {CampaignService} from '../services/Campaign.service';
import {CampaignMemberService} from '../services/CampaignMember.service';
import {NoteService} from '../services/Note.service';
import {EmailMessageService} from '../services/EmailMessage.service';

import { Directive } from '@angular/core';

/**
 Base class of all Components.
 For convenience, contains all enums and entity lists
 **/

@Directive()
export class BaseComponent {

    constructor (private http: HttpClient) {}

// enum instances
    UserRoles = Object.keys(enumTypes.UserRole);
    UserStatuss = Object.keys(enumTypes.UserStatus);
    TeamTypes = Object.keys(enumTypes.TeamType);
    TerritoryTypes = Object.keys(enumTypes.TerritoryType);
    AccountTypes = Object.keys(enumTypes.AccountType);
    AccountLifecycleStages = Object.keys(enumTypes.AccountLifecycleStage);
    ContactMethods = Object.keys(enumTypes.ContactMethod);
    LeadStatuss = Object.keys(enumTypes.LeadStatus);
    LeadSources = Object.keys(enumTypes.LeadSource);
    LeadRatings = Object.keys(enumTypes.LeadRating);
    OpportunityStages = Object.keys(enumTypes.OpportunityStage);
    OpportunityTypes = Object.keys(enumTypes.OpportunityType);
    ForecastCategorys = Object.keys(enumTypes.ForecastCategory);
    ProductTypes = Object.keys(enumTypes.ProductType);
    UnitOfMeasures = Object.keys(enumTypes.UnitOfMeasure);
    QuoteStatuss = Object.keys(enumTypes.QuoteStatus);
    OrderStatuss = Object.keys(enumTypes.OrderStatus);
    ContractStatuss = Object.keys(enumTypes.ContractStatus);
    CaseStatuss = Object.keys(enumTypes.CaseStatus);
    CasePrioritys = Object.keys(enumTypes.CasePriority);
    CaseOrigins = Object.keys(enumTypes.CaseOrigin);
    CaseSeveritys = Object.keys(enumTypes.CaseSeverity);
    ActivityTypes = Object.keys(enumTypes.ActivityType);
    ActivityStatuss = Object.keys(enumTypes.ActivityStatus);
    ActivityPrioritys = Object.keys(enumTypes.ActivityPriority);
    CampaignStatuss = Object.keys(enumTypes.CampaignStatus);
    CampaignTypes = Object.keys(enumTypes.CampaignType);
    CampaignMemberStatuss = Object.keys(enumTypes.CampaignMemberStatus);
    CampaignMemberTypes = Object.keys(enumTypes.CampaignMemberType);
    EmailDirections = Object.keys(enumTypes.EmailDirection);
    EmailStatuss = Object.keys(enumTypes.EmailStatus);

// all collection instances
    organizations : any;
    users : any;
    teams : any;
    territorys : any;
    accounts : any;
    contacts : any;
    leads : any;
    opportunitys : any;
    opportunityLineItems : any;
    opportunityStageHistorys : any;
    products : any;
    priceBooks : any;
    priceBookEntrys : any;
    quotes : any;
    quoteLineItems : any;
    orders : any;
    orderItems : any;
    contracts : any;
    case_s : any;
    activitys : any;
    campaigns : any;
    campaignMembers : any;
    notes : any;
    emailMessages : any;
  
// initialization  
    ngOnInit() {
    }

    initOrganizationList() {
        if ( this.organizations == null ) {
            new OrganizationService(this.http).getOrganizations().subscribe(res => {
                this.organizations = res;
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
    
    initTeamList() {
        if ( this.teams == null ) {
            new TeamService(this.http).getTeams().subscribe(res => {
                this.teams = res;
            });
        }
    }
    
    initTerritoryList() {
        if ( this.territorys == null ) {
            new TerritoryService(this.http).getTerritorys().subscribe(res => {
                this.territorys = res;
            });
        }
    }
    
    initAccountList() {
        if ( this.accounts == null ) {
            new AccountService(this.http).getAccounts().subscribe(res => {
                this.accounts = res;
            });
        }
    }
    
    initContactList() {
        if ( this.contacts == null ) {
            new ContactService(this.http).getContacts().subscribe(res => {
                this.contacts = res;
            });
        }
    }
    
    initLeadList() {
        if ( this.leads == null ) {
            new LeadService(this.http).getLeads().subscribe(res => {
                this.leads = res;
            });
        }
    }
    
    initOpportunityList() {
        if ( this.opportunitys == null ) {
            new OpportunityService(this.http).getOpportunitys().subscribe(res => {
                this.opportunitys = res;
            });
        }
    }
    
    initOpportunityLineItemList() {
        if ( this.opportunityLineItems == null ) {
            new OpportunityLineItemService(this.http).getOpportunityLineItems().subscribe(res => {
                this.opportunityLineItems = res;
            });
        }
    }
    
    initOpportunityStageHistoryList() {
        if ( this.opportunityStageHistorys == null ) {
            new OpportunityStageHistoryService(this.http).getOpportunityStageHistorys().subscribe(res => {
                this.opportunityStageHistorys = res;
            });
        }
    }
    
    initProductList() {
        if ( this.products == null ) {
            new ProductService(this.http).getProducts().subscribe(res => {
                this.products = res;
            });
        }
    }
    
    initPriceBookList() {
        if ( this.priceBooks == null ) {
            new PriceBookService(this.http).getPriceBooks().subscribe(res => {
                this.priceBooks = res;
            });
        }
    }
    
    initPriceBookEntryList() {
        if ( this.priceBookEntrys == null ) {
            new PriceBookEntryService(this.http).getPriceBookEntrys().subscribe(res => {
                this.priceBookEntrys = res;
            });
        }
    }
    
    initQuoteList() {
        if ( this.quotes == null ) {
            new QuoteService(this.http).getQuotes().subscribe(res => {
                this.quotes = res;
            });
        }
    }
    
    initQuoteLineItemList() {
        if ( this.quoteLineItems == null ) {
            new QuoteLineItemService(this.http).getQuoteLineItems().subscribe(res => {
                this.quoteLineItems = res;
            });
        }
    }
    
    initOrderList() {
        if ( this.orders == null ) {
            new OrderService(this.http).getOrders().subscribe(res => {
                this.orders = res;
            });
        }
    }
    
    initOrderItemList() {
        if ( this.orderItems == null ) {
            new OrderItemService(this.http).getOrderItems().subscribe(res => {
                this.orderItems = res;
            });
        }
    }
    
    initContractList() {
        if ( this.contracts == null ) {
            new ContractService(this.http).getContracts().subscribe(res => {
                this.contracts = res;
            });
        }
    }
    
    initCase_List() {
        if ( this.case_s == null ) {
            new Case_Service(this.http).getCase_s().subscribe(res => {
                this.case_s = res;
            });
        }
    }
    
    initActivityList() {
        if ( this.activitys == null ) {
            new ActivityService(this.http).getActivitys().subscribe(res => {
                this.activitys = res;
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
    
    initCampaignMemberList() {
        if ( this.campaignMembers == null ) {
            new CampaignMemberService(this.http).getCampaignMembers().subscribe(res => {
                this.campaignMembers = res;
            });
        }
    }
    
    initNoteList() {
        if ( this.notes == null ) {
            new NoteService(this.http).getNotes().subscribe(res => {
                this.notes = res;
            });
        }
    }
    
    initEmailMessageList() {
        if ( this.emailMessages == null ) {
            new EmailMessageService(this.http).getEmailMessages().subscribe(res => {
                this.emailMessages = res;
            });
        }
    }
    
    
// comparison function for select controls  
    compareFn(user1: any, user2: any) {
        return user1 == user2
    }    
}
