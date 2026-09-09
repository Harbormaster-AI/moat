import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import HomePageComponent from './components/HomePageComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListOrganizationComponent from './components/ListOrganizationComponent';
import CreateOrganizationComponent from './components/CreateOrganizationComponent';
import ViewOrganizationComponent from './components/ViewOrganizationComponent';
import ListUserComponent from './components/ListUserComponent';
import CreateUserComponent from './components/CreateUserComponent';
import ViewUserComponent from './components/ViewUserComponent';
import ListTeamComponent from './components/ListTeamComponent';
import CreateTeamComponent from './components/CreateTeamComponent';
import ViewTeamComponent from './components/ViewTeamComponent';
import ListTerritoryComponent from './components/ListTerritoryComponent';
import CreateTerritoryComponent from './components/CreateTerritoryComponent';
import ViewTerritoryComponent from './components/ViewTerritoryComponent';
import ListAccountComponent from './components/ListAccountComponent';
import CreateAccountComponent from './components/CreateAccountComponent';
import ViewAccountComponent from './components/ViewAccountComponent';
import ListContactComponent from './components/ListContactComponent';
import CreateContactComponent from './components/CreateContactComponent';
import ViewContactComponent from './components/ViewContactComponent';
import ListLeadComponent from './components/ListLeadComponent';
import CreateLeadComponent from './components/CreateLeadComponent';
import ViewLeadComponent from './components/ViewLeadComponent';
import ListOpportunityComponent from './components/ListOpportunityComponent';
import CreateOpportunityComponent from './components/CreateOpportunityComponent';
import ViewOpportunityComponent from './components/ViewOpportunityComponent';
import ListOpportunityLineItemComponent from './components/ListOpportunityLineItemComponent';
import CreateOpportunityLineItemComponent from './components/CreateOpportunityLineItemComponent';
import ViewOpportunityLineItemComponent from './components/ViewOpportunityLineItemComponent';
import ListOpportunityStageHistoryComponent from './components/ListOpportunityStageHistoryComponent';
import CreateOpportunityStageHistoryComponent from './components/CreateOpportunityStageHistoryComponent';
import ViewOpportunityStageHistoryComponent from './components/ViewOpportunityStageHistoryComponent';
import ListProductComponent from './components/ListProductComponent';
import CreateProductComponent from './components/CreateProductComponent';
import ViewProductComponent from './components/ViewProductComponent';
import ListPriceBookComponent from './components/ListPriceBookComponent';
import CreatePriceBookComponent from './components/CreatePriceBookComponent';
import ViewPriceBookComponent from './components/ViewPriceBookComponent';
import ListPriceBookEntryComponent from './components/ListPriceBookEntryComponent';
import CreatePriceBookEntryComponent from './components/CreatePriceBookEntryComponent';
import ViewPriceBookEntryComponent from './components/ViewPriceBookEntryComponent';
import ListQuoteComponent from './components/ListQuoteComponent';
import CreateQuoteComponent from './components/CreateQuoteComponent';
import ViewQuoteComponent from './components/ViewQuoteComponent';
import ListQuoteLineItemComponent from './components/ListQuoteLineItemComponent';
import CreateQuoteLineItemComponent from './components/CreateQuoteLineItemComponent';
import ViewQuoteLineItemComponent from './components/ViewQuoteLineItemComponent';
import ListOrderComponent from './components/ListOrderComponent';
import CreateOrderComponent from './components/CreateOrderComponent';
import ViewOrderComponent from './components/ViewOrderComponent';
import ListOrderItemComponent from './components/ListOrderItemComponent';
import CreateOrderItemComponent from './components/CreateOrderItemComponent';
import ViewOrderItemComponent from './components/ViewOrderItemComponent';
import ListContractComponent from './components/ListContractComponent';
import CreateContractComponent from './components/CreateContractComponent';
import ViewContractComponent from './components/ViewContractComponent';
import ListCase_Component from './components/ListCase_Component';
import CreateCase_Component from './components/CreateCase_Component';
import ViewCase_Component from './components/ViewCase_Component';
import ListActivityComponent from './components/ListActivityComponent';
import CreateActivityComponent from './components/CreateActivityComponent';
import ViewActivityComponent from './components/ViewActivityComponent';
import ListCampaignComponent from './components/ListCampaignComponent';
import CreateCampaignComponent from './components/CreateCampaignComponent';
import ViewCampaignComponent from './components/ViewCampaignComponent';
import ListCampaignMemberComponent from './components/ListCampaignMemberComponent';
import CreateCampaignMemberComponent from './components/CreateCampaignMemberComponent';
import ViewCampaignMemberComponent from './components/ViewCampaignMemberComponent';
import ListNoteComponent from './components/ListNoteComponent';
import CreateNoteComponent from './components/CreateNoteComponent';
import ViewNoteComponent from './components/ViewNoteComponent';
import ListEmailMessageComponent from './components/ListEmailMessageComponent';
import CreateEmailMessageComponent from './components/CreateEmailMessageComponent';
import ViewEmailMessageComponent from './components/ViewEmailMessageComponent';
function App() {
  return (
    <div>
        <Router>
                <HeaderComponent className="header"/>
                <div className="container">
                    <Switch>
                          <Route path = "/" exact component = {HomePageComponent}></Route>
                            <Route path = "/organizations" component = {ListOrganizationComponent}></Route>
                            <Route path = "/add-organization/:id" component = {CreateOrganizationComponent}></Route>
                            <Route path = "/view-organization/:id" component = {ViewOrganizationComponent}></Route>
                          {/* <Route path = "/update-organization/:id" component = {UpdateOrganizationComponent}></Route> */}
                            <Route path = "/users" component = {ListUserComponent}></Route>
                            <Route path = "/add-user/:id" component = {CreateUserComponent}></Route>
                            <Route path = "/view-user/:id" component = {ViewUserComponent}></Route>
                          {/* <Route path = "/update-user/:id" component = {UpdateUserComponent}></Route> */}
                            <Route path = "/teams" component = {ListTeamComponent}></Route>
                            <Route path = "/add-team/:id" component = {CreateTeamComponent}></Route>
                            <Route path = "/view-team/:id" component = {ViewTeamComponent}></Route>
                          {/* <Route path = "/update-team/:id" component = {UpdateTeamComponent}></Route> */}
                            <Route path = "/territorys" component = {ListTerritoryComponent}></Route>
                            <Route path = "/add-territory/:id" component = {CreateTerritoryComponent}></Route>
                            <Route path = "/view-territory/:id" component = {ViewTerritoryComponent}></Route>
                          {/* <Route path = "/update-territory/:id" component = {UpdateTerritoryComponent}></Route> */}
                            <Route path = "/accounts" component = {ListAccountComponent}></Route>
                            <Route path = "/add-account/:id" component = {CreateAccountComponent}></Route>
                            <Route path = "/view-account/:id" component = {ViewAccountComponent}></Route>
                          {/* <Route path = "/update-account/:id" component = {UpdateAccountComponent}></Route> */}
                            <Route path = "/contacts" component = {ListContactComponent}></Route>
                            <Route path = "/add-contact/:id" component = {CreateContactComponent}></Route>
                            <Route path = "/view-contact/:id" component = {ViewContactComponent}></Route>
                          {/* <Route path = "/update-contact/:id" component = {UpdateContactComponent}></Route> */}
                            <Route path = "/leads" component = {ListLeadComponent}></Route>
                            <Route path = "/add-lead/:id" component = {CreateLeadComponent}></Route>
                            <Route path = "/view-lead/:id" component = {ViewLeadComponent}></Route>
                          {/* <Route path = "/update-lead/:id" component = {UpdateLeadComponent}></Route> */}
                            <Route path = "/opportunitys" component = {ListOpportunityComponent}></Route>
                            <Route path = "/add-opportunity/:id" component = {CreateOpportunityComponent}></Route>
                            <Route path = "/view-opportunity/:id" component = {ViewOpportunityComponent}></Route>
                          {/* <Route path = "/update-opportunity/:id" component = {UpdateOpportunityComponent}></Route> */}
                            <Route path = "/opportunityLineItems" component = {ListOpportunityLineItemComponent}></Route>
                            <Route path = "/add-opportunityLineItem/:id" component = {CreateOpportunityLineItemComponent}></Route>
                            <Route path = "/view-opportunityLineItem/:id" component = {ViewOpportunityLineItemComponent}></Route>
                          {/* <Route path = "/update-opportunityLineItem/:id" component = {UpdateOpportunityLineItemComponent}></Route> */}
                            <Route path = "/opportunityStageHistorys" component = {ListOpportunityStageHistoryComponent}></Route>
                            <Route path = "/add-opportunityStageHistory/:id" component = {CreateOpportunityStageHistoryComponent}></Route>
                            <Route path = "/view-opportunityStageHistory/:id" component = {ViewOpportunityStageHistoryComponent}></Route>
                          {/* <Route path = "/update-opportunityStageHistory/:id" component = {UpdateOpportunityStageHistoryComponent}></Route> */}
                            <Route path = "/products" component = {ListProductComponent}></Route>
                            <Route path = "/add-product/:id" component = {CreateProductComponent}></Route>
                            <Route path = "/view-product/:id" component = {ViewProductComponent}></Route>
                          {/* <Route path = "/update-product/:id" component = {UpdateProductComponent}></Route> */}
                            <Route path = "/priceBooks" component = {ListPriceBookComponent}></Route>
                            <Route path = "/add-priceBook/:id" component = {CreatePriceBookComponent}></Route>
                            <Route path = "/view-priceBook/:id" component = {ViewPriceBookComponent}></Route>
                          {/* <Route path = "/update-priceBook/:id" component = {UpdatePriceBookComponent}></Route> */}
                            <Route path = "/priceBookEntrys" component = {ListPriceBookEntryComponent}></Route>
                            <Route path = "/add-priceBookEntry/:id" component = {CreatePriceBookEntryComponent}></Route>
                            <Route path = "/view-priceBookEntry/:id" component = {ViewPriceBookEntryComponent}></Route>
                          {/* <Route path = "/update-priceBookEntry/:id" component = {UpdatePriceBookEntryComponent}></Route> */}
                            <Route path = "/quotes" component = {ListQuoteComponent}></Route>
                            <Route path = "/add-quote/:id" component = {CreateQuoteComponent}></Route>
                            <Route path = "/view-quote/:id" component = {ViewQuoteComponent}></Route>
                          {/* <Route path = "/update-quote/:id" component = {UpdateQuoteComponent}></Route> */}
                            <Route path = "/quoteLineItems" component = {ListQuoteLineItemComponent}></Route>
                            <Route path = "/add-quoteLineItem/:id" component = {CreateQuoteLineItemComponent}></Route>
                            <Route path = "/view-quoteLineItem/:id" component = {ViewQuoteLineItemComponent}></Route>
                          {/* <Route path = "/update-quoteLineItem/:id" component = {UpdateQuoteLineItemComponent}></Route> */}
                            <Route path = "/orders" component = {ListOrderComponent}></Route>
                            <Route path = "/add-order/:id" component = {CreateOrderComponent}></Route>
                            <Route path = "/view-order/:id" component = {ViewOrderComponent}></Route>
                          {/* <Route path = "/update-order/:id" component = {UpdateOrderComponent}></Route> */}
                            <Route path = "/orderItems" component = {ListOrderItemComponent}></Route>
                            <Route path = "/add-orderItem/:id" component = {CreateOrderItemComponent}></Route>
                            <Route path = "/view-orderItem/:id" component = {ViewOrderItemComponent}></Route>
                          {/* <Route path = "/update-orderItem/:id" component = {UpdateOrderItemComponent}></Route> */}
                            <Route path = "/contracts" component = {ListContractComponent}></Route>
                            <Route path = "/add-contract/:id" component = {CreateContractComponent}></Route>
                            <Route path = "/view-contract/:id" component = {ViewContractComponent}></Route>
                          {/* <Route path = "/update-contract/:id" component = {UpdateContractComponent}></Route> */}
                            <Route path = "/case_s" component = {ListCase_Component}></Route>
                            <Route path = "/add-case_/:id" component = {CreateCase_Component}></Route>
                            <Route path = "/view-case_/:id" component = {ViewCase_Component}></Route>
                          {/* <Route path = "/update-case_/:id" component = {UpdateCase_Component}></Route> */}
                            <Route path = "/activitys" component = {ListActivityComponent}></Route>
                            <Route path = "/add-activity/:id" component = {CreateActivityComponent}></Route>
                            <Route path = "/view-activity/:id" component = {ViewActivityComponent}></Route>
                          {/* <Route path = "/update-activity/:id" component = {UpdateActivityComponent}></Route> */}
                            <Route path = "/campaigns" component = {ListCampaignComponent}></Route>
                            <Route path = "/add-campaign/:id" component = {CreateCampaignComponent}></Route>
                            <Route path = "/view-campaign/:id" component = {ViewCampaignComponent}></Route>
                          {/* <Route path = "/update-campaign/:id" component = {UpdateCampaignComponent}></Route> */}
                            <Route path = "/campaignMembers" component = {ListCampaignMemberComponent}></Route>
                            <Route path = "/add-campaignMember/:id" component = {CreateCampaignMemberComponent}></Route>
                            <Route path = "/view-campaignMember/:id" component = {ViewCampaignMemberComponent}></Route>
                          {/* <Route path = "/update-campaignMember/:id" component = {UpdateCampaignMemberComponent}></Route> */}
                            <Route path = "/notes" component = {ListNoteComponent}></Route>
                            <Route path = "/add-note/:id" component = {CreateNoteComponent}></Route>
                            <Route path = "/view-note/:id" component = {ViewNoteComponent}></Route>
                          {/* <Route path = "/update-note/:id" component = {UpdateNoteComponent}></Route> */}
                            <Route path = "/emailMessages" component = {ListEmailMessageComponent}></Route>
                            <Route path = "/add-emailMessage/:id" component = {CreateEmailMessageComponent}></Route>
                            <Route path = "/view-emailMessage/:id" component = {ViewEmailMessageComponent}></Route>
                          {/* <Route path = "/update-emailMessage/:id" component = {UpdateEmailMessageComponent}></Route> */}
                    </Switch>
                </div>
              <FooterComponent />
        </Router>
    </div>
    
  );
}

export default App;
