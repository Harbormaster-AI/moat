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

import {IndexOrganizationComponent} from './components/Organization/index/index.component';
import {CreateOrganizationComponent} from './components/Organization/create/create.component';
import {EditOrganizationComponent} from './components/Organization/edit/edit.component';
import {IndexUserComponent} from './components/User/index/index.component';
import {CreateUserComponent} from './components/User/create/create.component';
import {EditUserComponent} from './components/User/edit/edit.component';
import {IndexTeamComponent} from './components/Team/index/index.component';
import {CreateTeamComponent} from './components/Team/create/create.component';
import {EditTeamComponent} from './components/Team/edit/edit.component';
import {IndexTerritoryComponent} from './components/Territory/index/index.component';
import {CreateTerritoryComponent} from './components/Territory/create/create.component';
import {EditTerritoryComponent} from './components/Territory/edit/edit.component';
import {IndexAccountComponent} from './components/Account/index/index.component';
import {CreateAccountComponent} from './components/Account/create/create.component';
import {EditAccountComponent} from './components/Account/edit/edit.component';
import {IndexContactComponent} from './components/Contact/index/index.component';
import {CreateContactComponent} from './components/Contact/create/create.component';
import {EditContactComponent} from './components/Contact/edit/edit.component';
import {IndexLeadComponent} from './components/Lead/index/index.component';
import {CreateLeadComponent} from './components/Lead/create/create.component';
import {EditLeadComponent} from './components/Lead/edit/edit.component';
import {IndexOpportunityComponent} from './components/Opportunity/index/index.component';
import {CreateOpportunityComponent} from './components/Opportunity/create/create.component';
import {EditOpportunityComponent} from './components/Opportunity/edit/edit.component';
import {IndexOpportunityLineItemComponent} from './components/OpportunityLineItem/index/index.component';
import {CreateOpportunityLineItemComponent} from './components/OpportunityLineItem/create/create.component';
import {EditOpportunityLineItemComponent} from './components/OpportunityLineItem/edit/edit.component';
import {IndexOpportunityStageHistoryComponent} from './components/OpportunityStageHistory/index/index.component';
import {CreateOpportunityStageHistoryComponent} from './components/OpportunityStageHistory/create/create.component';
import {EditOpportunityStageHistoryComponent} from './components/OpportunityStageHistory/edit/edit.component';
import {IndexProductComponent} from './components/Product/index/index.component';
import {CreateProductComponent} from './components/Product/create/create.component';
import {EditProductComponent} from './components/Product/edit/edit.component';
import {IndexPriceBookComponent} from './components/PriceBook/index/index.component';
import {CreatePriceBookComponent} from './components/PriceBook/create/create.component';
import {EditPriceBookComponent} from './components/PriceBook/edit/edit.component';
import {IndexPriceBookEntryComponent} from './components/PriceBookEntry/index/index.component';
import {CreatePriceBookEntryComponent} from './components/PriceBookEntry/create/create.component';
import {EditPriceBookEntryComponent} from './components/PriceBookEntry/edit/edit.component';
import {IndexQuoteComponent} from './components/Quote/index/index.component';
import {CreateQuoteComponent} from './components/Quote/create/create.component';
import {EditQuoteComponent} from './components/Quote/edit/edit.component';
import {IndexQuoteLineItemComponent} from './components/QuoteLineItem/index/index.component';
import {CreateQuoteLineItemComponent} from './components/QuoteLineItem/create/create.component';
import {EditQuoteLineItemComponent} from './components/QuoteLineItem/edit/edit.component';
import {IndexOrderComponent} from './components/Order/index/index.component';
import {CreateOrderComponent} from './components/Order/create/create.component';
import {EditOrderComponent} from './components/Order/edit/edit.component';
import {IndexOrderItemComponent} from './components/OrderItem/index/index.component';
import {CreateOrderItemComponent} from './components/OrderItem/create/create.component';
import {EditOrderItemComponent} from './components/OrderItem/edit/edit.component';
import {IndexContractComponent} from './components/Contract/index/index.component';
import {CreateContractComponent} from './components/Contract/create/create.component';
import {EditContractComponent} from './components/Contract/edit/edit.component';
import {IndexCase_Component} from './components/Case_/index/index.component';
import {CreateCase_Component} from './components/Case_/create/create.component';
import {EditCase_Component} from './components/Case_/edit/edit.component';
import {IndexActivityComponent} from './components/Activity/index/index.component';
import {CreateActivityComponent} from './components/Activity/create/create.component';
import {EditActivityComponent} from './components/Activity/edit/edit.component';
import {IndexCampaignComponent} from './components/Campaign/index/index.component';
import {CreateCampaignComponent} from './components/Campaign/create/create.component';
import {EditCampaignComponent} from './components/Campaign/edit/edit.component';
import {IndexCampaignMemberComponent} from './components/CampaignMember/index/index.component';
import {CreateCampaignMemberComponent} from './components/CampaignMember/create/create.component';
import {EditCampaignMemberComponent} from './components/CampaignMember/edit/edit.component';
import {IndexNoteComponent} from './components/Note/index/index.component';
import {CreateNoteComponent} from './components/Note/create/create.component';
import {EditNoteComponent} from './components/Note/edit/edit.component';
import {IndexEmailMessageComponent} from './components/EmailMessage/index/index.component';
import {CreateEmailMessageComponent} from './components/EmailMessage/create/create.component';
import {EditEmailMessageComponent} from './components/EmailMessage/edit/edit.component';

import * as appRoutes from './routerConfig';

import {OrganizationService} from './services/Organization.service';
import {UserService} from './services/User.service';
import {TeamService} from './services/Team.service';
import {TerritoryService} from './services/Territory.service';
import {AccountService} from './services/Account.service';
import {ContactService} from './services/Contact.service';
import {LeadService} from './services/Lead.service';
import {OpportunityService} from './services/Opportunity.service';
import {OpportunityLineItemService} from './services/OpportunityLineItem.service';
import {OpportunityStageHistoryService} from './services/OpportunityStageHistory.service';
import {ProductService} from './services/Product.service';
import {PriceBookService} from './services/PriceBook.service';
import {PriceBookEntryService} from './services/PriceBookEntry.service';
import {QuoteService} from './services/Quote.service';
import {QuoteLineItemService} from './services/QuoteLineItem.service';
import {OrderService} from './services/Order.service';
import {OrderItemService} from './services/OrderItem.service';
import {ContractService} from './services/Contract.service';
import {Case_Service} from './services/Case_.service';
import {ActivityService} from './services/Activity.service';
import {CampaignService} from './services/Campaign.service';
import {CampaignMemberService} from './services/CampaignMember.service';
import {NoteService} from './services/Note.service';
import {EmailMessageService} from './services/EmailMessage.service';

@NgModule({
  declarations: [
    IndexOrganizationComponent,
    CreateOrganizationComponent,
    EditOrganizationComponent,
    IndexUserComponent,
    CreateUserComponent,
    EditUserComponent,
    IndexTeamComponent,
    CreateTeamComponent,
    EditTeamComponent,
    IndexTerritoryComponent,
    CreateTerritoryComponent,
    EditTerritoryComponent,
    IndexAccountComponent,
    CreateAccountComponent,
    EditAccountComponent,
    IndexContactComponent,
    CreateContactComponent,
    EditContactComponent,
    IndexLeadComponent,
    CreateLeadComponent,
    EditLeadComponent,
    IndexOpportunityComponent,
    CreateOpportunityComponent,
    EditOpportunityComponent,
    IndexOpportunityLineItemComponent,
    CreateOpportunityLineItemComponent,
    EditOpportunityLineItemComponent,
    IndexOpportunityStageHistoryComponent,
    CreateOpportunityStageHistoryComponent,
    EditOpportunityStageHistoryComponent,
    IndexProductComponent,
    CreateProductComponent,
    EditProductComponent,
    IndexPriceBookComponent,
    CreatePriceBookComponent,
    EditPriceBookComponent,
    IndexPriceBookEntryComponent,
    CreatePriceBookEntryComponent,
    EditPriceBookEntryComponent,
    IndexQuoteComponent,
    CreateQuoteComponent,
    EditQuoteComponent,
    IndexQuoteLineItemComponent,
    CreateQuoteLineItemComponent,
    EditQuoteLineItemComponent,
    IndexOrderComponent,
    CreateOrderComponent,
    EditOrderComponent,
    IndexOrderItemComponent,
    CreateOrderItemComponent,
    EditOrderItemComponent,
    IndexContractComponent,
    CreateContractComponent,
    EditContractComponent,
    IndexCase_Component,
    CreateCase_Component,
    EditCase_Component,
    IndexActivityComponent,
    CreateActivityComponent,
    EditActivityComponent,
    IndexCampaignComponent,
    CreateCampaignComponent,
    EditCampaignComponent,
    IndexCampaignMemberComponent,
    CreateCampaignMemberComponent,
    EditCampaignMemberComponent,
    IndexNoteComponent,
    CreateNoteComponent,
    EditNoteComponent,
    IndexEmailMessageComponent,
    CreateEmailMessageComponent,
    EditEmailMessageComponent,
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
    RouterModule.forRoot(appRoutes.OrganizationRoutes), 
    RouterModule.forRoot(appRoutes.UserRoutes), 
    RouterModule.forRoot(appRoutes.TeamRoutes), 
    RouterModule.forRoot(appRoutes.TerritoryRoutes), 
    RouterModule.forRoot(appRoutes.AccountRoutes), 
    RouterModule.forRoot(appRoutes.ContactRoutes), 
    RouterModule.forRoot(appRoutes.LeadRoutes), 
    RouterModule.forRoot(appRoutes.OpportunityRoutes), 
    RouterModule.forRoot(appRoutes.OpportunityLineItemRoutes), 
    RouterModule.forRoot(appRoutes.OpportunityStageHistoryRoutes), 
    RouterModule.forRoot(appRoutes.ProductRoutes), 
    RouterModule.forRoot(appRoutes.PriceBookRoutes), 
    RouterModule.forRoot(appRoutes.PriceBookEntryRoutes), 
    RouterModule.forRoot(appRoutes.QuoteRoutes), 
    RouterModule.forRoot(appRoutes.QuoteLineItemRoutes), 
    RouterModule.forRoot(appRoutes.OrderRoutes), 
    RouterModule.forRoot(appRoutes.OrderItemRoutes), 
    RouterModule.forRoot(appRoutes.ContractRoutes), 
    RouterModule.forRoot(appRoutes.Case_Routes), 
    RouterModule.forRoot(appRoutes.ActivityRoutes), 
    RouterModule.forRoot(appRoutes.CampaignRoutes), 
    RouterModule.forRoot(appRoutes.CampaignMemberRoutes), 
    RouterModule.forRoot(appRoutes.NoteRoutes), 
    RouterModule.forRoot(appRoutes.EmailMessageRoutes), 
  ],
  providers: [OrganizationService,UserService,TeamService,TerritoryService,AccountService,ContactService,LeadService,OpportunityService,OpportunityLineItemService,OpportunityStageHistoryService,ProductService,PriceBookService,PriceBookEntryService,QuoteService,QuoteLineItemService,OrderService,OrderItemService,ContractService,Case_Service,ActivityService,CampaignService,CampaignMemberService,NoteService,EmailMessageService],
  bootstrap: [AppComponent]
})
export class AppModule { }
