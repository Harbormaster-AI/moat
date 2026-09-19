class BonusPlansController < ApplicationController
  def index
    @bonusPlans = BonusPlan.all
  end
 
  def show
    @bonusPlan = BonusPlan.find(params[:id])
  end
 
  def new
    @bonusPlan = BonusPlan.new
  end
 
  def edit
    @bonusPlan = BonusPlan.find(params[:id])
  end
 
  def create
    @bonusPlan = BonusPlan.new(bonusPlan_params)
 
    if @bonusPlan.save
      redirect_to bonusPlans_path
    else
      render 'new'
    end
  end
 
  def update
    @bonusPlan = BonusPlan.find(params[:id])
 
    if @bonusPlan.update(bonusPlan_params)
      redirect_to bonusPlans_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @bonusPlan = BonusPlan.find(params[:id])
    @bonusPlan.destroy
    redirect_to bonusPlans_path
  end

 
  private
    def bonusPlan_params
      params.require(:bonusPlan).permit(:name, :targetPercentage)
    end
end