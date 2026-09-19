class SubscriptionsController < ApplicationController
  def index
    @subscriptions = Subscription.all
  end
 
  def show
    @subscription = Subscription.find(params[:id])
  end
 
  def new
    @subscription = Subscription.new
  end
 
  def edit
    @subscription = Subscription.find(params[:id])
  end
 
  def create
    @subscription = Subscription.new(subscription_params)
 
    if @subscription.save
      redirect_to subscriptions_path
    else
      render 'new'
    end
  end
 
  def update
    @subscription = Subscription.find(params[:id])
 
    if @subscription.update(subscription_params)
      redirect_to subscriptions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @subscription = Subscription.find(params[:id])
    @subscription.destroy
    redirect_to subscriptions_path
  end

 
  private
    def subscription_params
      params.require(:subscription).permit(:subscriptionNumber, :nextBillingDate, :startDate, :endDate, :Status, :Interval)
    end
end