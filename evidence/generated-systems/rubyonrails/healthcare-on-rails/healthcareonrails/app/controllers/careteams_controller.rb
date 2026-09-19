class CareTeamsController < ApplicationController
  def index
    @careTeams = CareTeam.all
  end
 
  def show
    @careTeam = CareTeam.find(params[:id])
  end
 
  def new
    @careTeam = CareTeam.new
  end
 
  def edit
    @careTeam = CareTeam.find(params[:id])
  end
 
  def create
    @careTeam = CareTeam.new(careTeam_params)
 
    if @careTeam.save
      redirect_to careTeams_path
    else
      render 'new'
    end
  end
 
  def update
    @careTeam = CareTeam.find(params[:id])
 
    if @careTeam.update(careTeam_params)
      redirect_to careTeams_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @careTeam = CareTeam.find(params[:id])
    @careTeam.destroy
    redirect_to careTeams_path
  end

 
  private
    def careTeam_params
      params.require(:careTeam).permit(:name, :CareSetting)
    end
end