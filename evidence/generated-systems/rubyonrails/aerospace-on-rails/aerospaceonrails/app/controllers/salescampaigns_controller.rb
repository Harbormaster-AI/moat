class SalesCampaignsController < ApplicationController
  def index
    @salesCampaigns = SalesCampaign.all
  end
 
  def show
    @salesCampaign = SalesCampaign.find(params[:id])
  end
 
  def new
    @salesCampaign = SalesCampaign.new
  end
 
  def edit
    @salesCampaign = SalesCampaign.find(params[:id])
  end
 
  def create
    @salesCampaign = SalesCampaign.new(salesCampaign_params)
 
    if @salesCampaign.save
      redirect_to salesCampaigns_path
    else
      render 'new'
    end
  end
 
  def update
    @salesCampaign = SalesCampaign.find(params[:id])
 
    if @salesCampaign.update(salesCampaign_params)
      redirect_to salesCampaigns_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @salesCampaign = SalesCampaign.find(params[:id])
    @salesCampaign.destroy
    redirect_to salesCampaigns_path
  end

 
  private
    def salesCampaign_params
      params.require(:salesCampaign).permit(:campaignCode, :Status)
    end
end