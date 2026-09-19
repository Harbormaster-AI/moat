class CampaignMembersController < ApplicationController
  def index
    @campaignMembers = CampaignMember.all
  end
 
  def show
    @campaignMember = CampaignMember.find(params[:id])
  end
 
  def new
    @campaignMember = CampaignMember.new
  end
 
  def edit
    @campaignMember = CampaignMember.find(params[:id])
  end
 
  def create
    @campaignMember = CampaignMember.new(campaignMember_params)
 
    if @campaignMember.save
      redirect_to campaignMembers_path
    else
      render 'new'
    end
  end
 
  def update
    @campaignMember = CampaignMember.find(params[:id])
 
    if @campaignMember.update(campaignMember_params)
      redirect_to campaignMembers_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @campaignMember = CampaignMember.find(params[:id])
    @campaignMember.destroy
    redirect_to campaignMembers_path
  end

 
  private
    def campaignMember_params
      params.require(:campaignMember).permit(:responded, :Status, :MemberType)
    end
end